import fs from 'node:fs';

function main() {
    const input: String = fs.readFileSync('day19.input', 'utf8').trim().split("\n\n");
    const towels: Set<String> = new Set(input[0].split(",").map((s) => s.trim()))
    const maxTowelLength: number = Math.max(...[...towels].map((towel) => towel.length))
    const patterns: Array<String> = input[1].split("\n")

    // console.log(towels, towels.has('b'))
    console.log(`Part1: ${patterns.filter((pattern) => isPossiblePattern(pattern, towels, maxTowelLength, new Map([["", true]]))).length}`)

    const totalWays = patterns.map((pattern) => countPossibleWays(pattern, towels, maxTowelLength, new Map([["", 1]]))).reduce((acc, val) => acc + val)
    console.log(`Part2: ${totalWays}`)
}

function isPossiblePattern(pattern: String, towels: Set<String>, maxTowelLength: number, knownPatterns: Map<String, boolean>): boolean {
    // console.log(pattern)
    if (knownPatterns.has(pattern)) return knownPatterns.get(pattern)!!

    for (let i = 1; i <= maxTowelLength && i <= pattern.length; i++) {
        if (towels.has(pattern.slice(0, i)) &&
            isPossiblePattern(pattern.slice(i), towels, maxTowelLength, knownPatterns)) {
            knownPatterns.set(pattern, true)
            return true
        }
    }

    knownPatterns.set(pattern, false)
    return false
}

function countPossibleWays(pattern: String, towels: Set<String>, maxTowelLength: number, knownPatterns: Map<String, number>): number {
    // console.log(pattern)
    if (knownPatterns.has(pattern)) return knownPatterns.get(pattern)!!

    var possibleWays = 0
    for (let i = 1; i <= maxTowelLength && i <= pattern.length; i++) {
        if (towels.has(pattern.slice(0, i))) {
            possibleWays += countPossibleWays(pattern.slice(i), towels, maxTowelLength, knownPatterns)
        }
    }false

    knownPatterns.set(pattern, possibleWays)
    return possibleWays
}

if (require.main === module) {
    main()
}
