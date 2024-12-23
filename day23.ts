import fs from "node:fs"

function main() {
    const input = fs.readFileSync("day23.input", "ascii").trim().split("\n").map((s) => s.trim())
    let network: Map<string, Set<string>> = new Map()

    for (let line of input) {
        let computerPair = line.split("-")
        for (let i = 0; i<computerPair.length; i++) {
            let computerConnections = network.get(computerPair[i % 2])
            if (computerConnections === undefined) {
                network.set(computerPair[i % 2], new Set([computerPair[0], computerPair[1]]))
            } else {
                computerConnections.add(computerPair[(i + 1) % 2])
            }
        }
    }

    let connectedTriples: string[][] = []
    for (let [computer1, connections] of network) {
        for (let computer2 of connections) {
            if (computer2 <= computer1) continue

            for (let computer3 of connections) {
                if (computer3 <= computer2) continue

                if ((computer1.startsWith("t") || computer2.startsWith("t") || computer3.startsWith("t")) && network.get(computer2)!!.has(computer3)) {
                    connectedTriples.push([computer1, computer2, computer3])
                }
            }
        }
    }

    console.log(`Part1: ${connectedTriples.length}`)

    var biggestClique: string[] = []
    for (let [computer, connections] of network) {
        let connectionsToCareAbout: string[] = []
        for (let connection of connections) {
            if (connection > computer) {
                connectionsToCareAbout.push(connection)
            }
        }

        var subsets: string[][] = [[computer]]
        for (let connection of connectionsToCareAbout) {
            let newSubsets: string[][] = []
            for (let subset of subsets) {
                newSubsets.push(subset)
                newSubsets.push(subset.concat([connection]))
            }
            subsets = newSubsets
        }

        for (let subset of subsets) {
            if (isMaximalClique(subset, network) && subset.length > biggestClique.length) {
                biggestClique = subset
            }
        }
    }

    biggestClique.sort()

    console.log(`Part2: ${biggestClique}`)
}

function isMaximalClique(maybeClique: string[], network: Map<string, Set<string>>): boolean {
    return maybeClique.map((computer) => network.get(computer)!!).reduce((acc, connections) => acc.intersection(connections)).size === maybeClique.length
}

if (require.main === module) {
    main()
}