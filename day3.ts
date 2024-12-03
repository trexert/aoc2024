const fs = require('node:fs');

const input: String = fs.readFileSync('day3.input', 'utf8');

const regex = /mul\((\d{1,3}),(\d{1,3})\)/g;

let totalMuls = 0
for (let match of input.matchAll(regex)) {
    totalMuls += +match[1] * +match[2]
}

console.log(`Part1: ${totalMuls}`)

const regex2 = /mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)/g;

totalMuls = 0
let doMuls = true
for (let match of input.matchAll(regex2)) {
    switch (match[0].substring(0, 3)) {
        case "mul":
            if (doMuls) {
                totalMuls += +match[1] * +match[2]
            }
            break
        case "do(":
            doMuls = true
            break
        case "don":
            doMuls = false
            break
        default:
            console.error(`Unexpected match sequence ${match[0]}`)
    }
}

console.log(`Part2: ${totalMuls}`)
