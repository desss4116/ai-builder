/*========================================
WCKD OMEGA COGNITIVE CORE
========================================*/

"use client"

import React, {
  useMemo,
  useState,
  useEffect,
  useRef,
  Suspense,
} from "react"

import {
  Card,
  CardContent,
} from "@/components/ui/card"

import {
  Table,
  TableHeader,
  TableBody,
  TableHead,
  TableRow,
  TableCell,
} from "@/components/ui/table"

import { motion } from "framer-motion"

/*========================================
TOKENIZER
========================================*/

const STOP_WORDS = new Set([
  "the","a","an","and","or","of","to","in","on","at","is",
  "are","was","were","be","been","this","that","with",
  "for","from","as","by","it","its","into","about",
  "what","which","who","whom","their","them","they",
  "you","your","our","ours","i","me","my","mine"
])

export function normalizeText(input = "") {
  return input
    .toLowerCase()
    .replace(/[^\p{L}\p{N}\s]/gu, " ")
    .replace(/\s+/g, " ")
    .trim()
}

export function tokenize(text = "") {
  return normalizeText(text)
    .split(" ")
    .filter(Boolean)
    .filter(word => !STOP_WORDS.has(word))
}

/*========================================
BM25 ENGINE
========================================*/

export class BM25Engine {

  constructor(documents = [], k1 = 1.5, b = 0.75) {
    this.documents = documents
    this.k1 = k1
    this.b = b

    this.docLengths = []
    this.avgDocLength = 0

    this.termFreqs = []
    this.documentFreqs = new Map()

    this.initialize()
  }

  initialize() {

    let totalLength = 0

    for (const doc of this.documents) {

      const tokens = tokenize(doc)

      totalLength += tokens.length

      this.docLengths.push(tokens.length)

      const frequencies = new Map()

      for (const token of tokens) {
        frequencies.set(
          token,
          (frequencies.get(token) || 0) + 1
        )
      }

      this.termFreqs.push(frequencies)

      const unique = new Set(tokens)

      for (const token of unique) {
        this.documentFreqs.set(
          token,
          (this.documentFreqs.get(token) || 0) + 1
        )
      }
    }

    this.avgDocLength =
      totalLength / Math.max(1, this.documents.length)
  }

  idf(term) {

    const df =
      this.documentFreqs.get(term) || 0

    const N = this.documents.length

    return Math.log(
      1 +
      (
        (N - df + 0.5) /
        (df + 0.5)
      )
    )
  }

  score(query, documentIndex) {

    const queryTokens = tokenize(query)

    const frequencies =
      this.termFreqs[documentIndex]

    const docLength =
      this.docLengths[documentIndex]

    let score = 0

    for (const term of queryTokens) {

      const tf =
        frequencies.get(term) || 0

      const idf = this.idf(term)

      const numerator =
        tf * (this.k1 + 1)

      const denominator =
        tf +
        this.k1 *
        (
          1 -
          this.b +
          this.b *
          (docLength / this.avgDocLength)
        )

      score +=
        idf *
        (numerator / Math.max(1, denominator))
    }

    return score
  }

  search(query, limit = 5) {

    const results = this.documents.map(
      (doc, index) => ({
        index,
        score: this.score(query, index),
        document: doc,
      })
    )

    return results
      .sort((a, b) => b.score - a.score)
      .slice(0, limit)
  }
}

/*========================================
DOM DENSITY EXTRACTION
========================================*/

export class DOMDensityExtractor {

  constructor(html = "") {
    this.html = html
  }

  extractBlocks() {

    const regex =
      /<([a-zA-Z0-9]+)(.*?)>(.*?)<\/\1>/gs

    const blocks = []

    let match

    while ((match = regex.exec(this.html)) !== null) {

      const tag = match[1]
      const content = match[3]

      const clean =
        content
          .replace(/<[^>]*>/g, "")
          .replace(/\s+/g, " ")
          .trim()

      const tagCount =
        (content.match(/<[^>]+>/g) || []).length

      const charCount =
        clean.length

      const density =
        charCount / Math.max(1, tagCount)

      blocks.push({
        tag,
        density,
        text: clean,
      })
    }

    return blocks
  }

  extractMainContent(threshold = 120) {

    const blocks = this.extractBlocks()

    const useful =
      blocks.filter(
        block =>
          block.density >= threshold &&
          block.text.length > 80
      )

    return useful
      .map(x => x.text)
      .join("\n\n")
  }
}

/*========================================
KNOWLEDGE GRAPH
========================================*/

export class KnowledgeGraph {

  constructor() {
    this.nodes = new Map()
    this.edges = new Map()
  }

  ensureNode(node) {

    if (!this.nodes.has(node)) {
      this.nodes.set(node, {
        id: node,
        createdAt: Date.now(),
      })
    }

    if (!this.edges.has(node)) {
      this.edges.set(node, [])
    }
  }

  addTriple(subject, predicate, object) {

    this.ensureNode(subject)
    this.ensureNode(object)

    this.edges.get(subject).push({
      predicate,
      object,
      timestamp: Date.now(),
    })
  }

  getRelations(node) {

    return this.edges.get(node) || []
  }

  search(query) {

    const results = []

    for (const [subject, relations] of this.edges.entries()) {

      for (const relation of relations) {

        const blob =
          `${subject} ${relation.predicate} ${relation.object}`

        if (
          normalizeText(blob)
            .includes(normalizeText(query))
        ) {
          results.push({
            subject,
            predicate: relation.predicate,
            object: relation.object,
          })
        }
      }
    }

    return results
  }

  shortestPath(start, end) {

    const queue = [[start]]
    const visited = new Set()

    while (queue.length > 0) {

      const path = queue.shift()

      const node =
        path[path.length - 1]

      if (node === end) {
        return path
      }

      if (!visited.has(node)) {

        visited.add(node)

        const relations =
          this.getRelations(node)

        for (const relation of relations) {

          queue.push([
            ...path,
            relation.object,
          ])
        }
      }
    }

    return null
  }
}

/*========================================
ATOMIC UI DISPATCHER
========================================*/

function TimelineComponent({ data }) {

  return (
    <Card className="bg-black text-cyan-400 border-cyan-800">
      <CardContent className="p-6">
        <div className="space-y-4">
          {data.map((item, i) => (
            <motion.div
              key={i}
              initial={{ opacity: 0, x: -20 }}
              animate={{ opacity: 1, x: 0 }}
              className="border-l border-cyan-600 pl-4"
            >
              <div className="text-sm opacity-60">
                {item.date}
              </div>
              <div className="font-bold">
                {item.label}
              </div>
            </motion.div>
          ))}
        </div>
      </CardContent>
    </Card>
  )
}

function MapComponent({ data }) {

  return (
    <Card className="bg-black text-green-400 border-green-700">
      <CardContent className="p-4">
        <div className="grid grid-cols-10 gap-1">
          {data.map((point, i) => (
            <div
              key={i}
              className="w-4 h-4 rounded-full bg-green-500 animate-pulse"
              style={{
                transform:
                  `translate(${point.x}px, ${point.y}px)`
              }}
            />
          ))}
        </div>
      </CardContent>
    </Card>
  )
}

function DataTable({ data }) {

  const headers =
    Object.keys(data[0] || {})

  return (
    <Table>
      <TableHeader>
        <TableRow>
          {headers.map(header => (
            <TableHead key={header}>
              {header}
            </TableHead>
          ))}
        </TableRow>
      </TableHeader>

      <TableBody>
        {data.map((row, i) => (
          <TableRow key={i}>
            {headers.map(header => (
              <TableCell key={header}>
                {String(row[header])}
              </TableCell>
            ))}
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}

export function AtomicUIDispatcher({ payload }) {

  const type =
    useMemo(() => {

      if (
        Array.isArray(payload) &&
        payload.length > 0
      ) {

        const sample = payload[0]

        if (
          typeof sample.x === "number" &&
          typeof sample.y === "number"
        ) {
          return "map"
        }

        if (
          sample.date &&
          sample.label
        ) {
          return "timeline"
        }

        if (
          typeof sample === "object"
        ) {
          return "table"
        }
      }

      return "unknown"

    }, [payload])

  if (type === "map") {
    return <MapComponent data={payload} />
  }

  if (type === "timeline") {
    return <TimelineComponent data={payload} />
  }

  if (type === "table") {
    return <DataTable data={payload} />
  }

  return (
    <Card className="bg-zinc-950 text-zinc-200">
      <CardContent className="p-6">
        Unsupported Payload
      </CardContent>
    </Card>
  )
}

/*========================================
WCKD SEARCH CORE
========================================*/

export class WCKDBrain {

  constructor() {

    this.documents = []

    this.graph =
      new KnowledgeGraph()
  }

  ingestDocument(text) {

    this.documents.push(text)
  }

  buildSemanticEngine() {

    return new BM25Engine(
      this.documents
    )
  }

  learnTriple(s, p, o) {

    this.graph.addTriple(s, p, o)
  }

  ask(query) {

    const semantic =
      this.buildSemanticEngine()

    const ranked =
      semantic.search(query)

    const graph =
      this.graph.search(query)

    return {
      ranked,
      graph,
    }
  }
}

/*========================================
LIVE DEMO COMPONENT
========================================*/

export default function WCKDOmegaDemo() {

  const [query, setQuery] =
    useState("maze runner")

  const brainRef =
    useRef(null)

  useEffect(() => {

    const brain =
      new WCKDBrain()

    brain.ingestDocument(
      "Maze Runner contains procedural labyrinth systems and WCKD experiments."
    )

    brain.ingestDocument(
      "Three.js powers cinematic WebGL experiences."
    )

    brain.ingestDocument(
      "Framer Motion creates immersive UI transitions."
    )

    brain.learnTriple(
      "WCKD",
      "controls",
      "The Maze"
    )

    brain.learnTriple(
      "Thomas",
      "escaped",
      "The Glade"
    )

    brain.learnTriple(
      "Minho",
      "mapped",
      "Maze Corridors"
    )

    brainRef.current = brain

  }, [])

  const results =
    useMemo(() => {

      if (!brainRef.current) {
        return null
      }

      return brainRef.current.ask(query)

    }, [query])

  const timelineData = [
    {
      date: "2026",
      label: "WCKD System Activated",
    },
    {
      date: "2027",
      label: "Maze Protocol Initialized",
    },
  ]

  const mapData = [
    { x: 10, y: 20 },
    { x: 30, y: 50 },
    { x: 90, y: 80 },
  ]

  return (
    <div className="min-h-screen bg-black text-cyan-400 p-10 space-y-10">

      <motion.h1
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        className="text-5xl font-black tracking-widest"
      >
        WCKD OMEGA BRAIN
      </motion.h1>

      <input
        value={query}
        onChange={e => setQuery(e.target.value)}
        className="
          bg-zinc-900
          border
          border-cyan-700
          px-4
          py-3
          w-full
          rounded-xl
        "
      />

      <div className="grid grid-cols-2 gap-8">

        <Card className="bg-zinc-950 border-cyan-900">
          <CardContent className="p-6">
            <div className="text-2xl mb-4">
              Semantic Results
            </div>

            <div className="space-y-4">

              {results?.ranked.map((r, i) => (
                <div
                  key={i}
                  className="border border-cyan-800 p-4 rounded-xl"
                >
                  <div className="text-sm opacity-60">
                    SCORE:
                    {" "}
                    {r.score.toFixed(4)}
                  </div>

                  <div>
                    {r.document}
                  </div>
                </div>
              ))}

            </div>
          </CardContent>
        </Card>

        <Card className="bg-zinc-950 border-green-900">
          <CardContent className="p-6">
            <div className="text-2xl mb-4">
              Knowledge Graph
            </div>

            <div className="space-y-3">

              {results?.graph.map((g, i) => (
                <div
                  key={i}
                  className="
                    border
                    border-green-800
                    rounded-xl
                    p-4
                  "
                >
                  <span className="text-cyan-400">
                    {g.subject}
                  </span>

                  {" "}
                  →
                  {" "}

                  <span className="text-pink-400">
                    {g.predicate}
                  </span>

                  {" "}
                  →
                  {" "}

                  <span className="text-green-400">
                    {g.object}
                  </span>
                </div>
              ))}

            </div>
          </CardContent>
        </Card>

      </div>

      <div className="grid grid-cols-2 gap-8">

        <AtomicUIDispatcher
          payload={timelineData}
        />

        <AtomicUIDispatcher
          payload={mapData}
        />

      </div>

    </div>
  )
        }
