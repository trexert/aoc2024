const fs = require('node:fs');

const input: String = fs.readFileSync('day11.input', 'utf8').trim();

let stonesArray = input.split(' ')
// let stonesArray = ["125", "17"]
let stones: Map<String, number> = new Map(stonesArray.map((s) => [s, 1]))

for (let i = 0; i < 25; i++) {
    // console.log(i)
    // console.log(stones)
    stones = stoneUpdate(stones)
}
// console.log(stones)

let totalStones = 0
stones.forEach((count) => totalStones += count)

console.log(`Part1: ${totalStones}`)

for (let i = 0; i < 50; i++) {
    // console.log(i)
    stones = stoneUpdate(stones)
}

totalStones = 0
stones.forEach((count) => totalStones += count)

console.log(`Part2: ${totalStones}`)

function stoneUpdate(stones: Map<String, number>): Map<String, number> {
    let newStones = new Map<string, number>()
    for (let [oldStone, count] of stones) {
        if (oldStone.length % 2 == 0) {
            const stone1 = oldStone.slice(0, oldStone.length / 2)
            const stone2 = "" + +oldStone.slice(oldStone.length / 2)
            newStones.set(stone1, (newStones.get(stone1) ?? 0) + count)
            newStones.set(stone2, (newStones.get(stone2) ?? 0) + count)
        } else if (oldStone == "0") {
            newStones.set("1", (newStones.get("1") ?? 0) + count)
        } else {
            const stone = "" + (+oldStone * 2024)
            newStones.set(stone, (newStones.get(stone) ?? 0) + count)
        }
    }
    return newStones
}