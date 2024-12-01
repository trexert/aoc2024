const fs = require('node:fs');

const input: String = fs.readFileSync('day1.input', 'utf8');

console.log(input)

let list1: Array<number> = [];
let list2: Array<number> = [];
for (let line of input.split("\n")) {
    if (line == "") continue
    let vals: Array<number> = line.split("   ").map((num) => +num);
    list1.push(vals[0]);
    list2.push(vals[1]);
}

list1.sort();
list2.sort();

let distance: number = 0
for (let i=0; i<list1.length; i++) {
    distance += Math.abs(list1[i] - list2[i]);
}

console.log("Part 1: " + distance);

let counts: Object = {}

for (let val of list2) {
    counts[val] = counts[val] ? counts[val] + 1 : 1;
}

// console.log(counts)

let similarity: number = 0
for (let val of list1) {
    similarity += counts[val] ? val * counts[val] : 0
}

console.log(`Part 2: ${similarity}`)
