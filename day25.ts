import fs from "node:fs"

function main() {
    const [lords, kings] = parseInput()

    var validLockAndKeys = 0
    for (let king of kings) {
        for (let lord of lords) {
            if (lord.intersection(king).size == 0) {
                validLockAndKeys++
            }
        }
    }
    console.log(`Part1: ${validLockAndKeys}`)
}

function parseInput(): [Set<number>[], Set<number>[]] {
    const pictures = fs.readFileSync("day25.input", "ascii").trim().split("\n\n").map((s) => s.trim())
    var locks: Set<number>[] = []
    var keys: Set<number>[] = []

    for (let picture of pictures) {
        var solidSet: Set<number> = new Set()
        const rows = picture.split("\n")
        for (let [i, row] of rows.entries()) {
            for (let [j, cell] of row.split("").entries()) {
                if (cell == "#") {
                    solidSet.add(i * 10 + j)
                }
            }
        }
        
        if (solidSet.has(0)) {
            locks.push(solidSet)
        } else {
            keys.push(solidSet)
        }
    }

    return [locks, keys]
}

if (require.main == module) {
    main()
}
