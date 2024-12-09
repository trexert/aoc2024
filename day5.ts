const fs = require('node:fs');

const input: String = fs.readFileSync('day5.input', 'utf8').trim().split("\n\n");

const rules: String = input[0];
const prints: String = input[1];

let before: Map<String, Set<String>> = new Map();
let after: Map<String, Set<String>> = new Map()

for (let line of rules.split('\n')) {
    let rule = line.split('|');
    if (before.get(rule[0]) == undefined) {
        before.set(rule[0], new Set());
    }
    if (before.get(rule[1]) == undefined) {
        before.set(rule[1], new Set());
    }
    if (after.get(rule[0]) == undefined) {
        after.set(rule[0], new Set());
    }
    if (after.get(rule[1]) == undefined) {
        after.set(rule[1], new Set())
    }
    before.get(rule[0]).add(rule[1]);
    after.get(rule[1]).add(rule[0])
}

let validMiddles = 0
let postSortingMiddles = 0
for (let line of prints.split('\n')) {
    let printLine = line.split(',');
    if (testPrints(printLine)) {
        validMiddles += +printLine[Math.floor(printLine.length / 2)]
    } else {
        let sortedLine = sortPages(printLine)
        postSortingMiddles += +sortedLine[Math.floor(sortedLine.length / 2)]
    }
}

console.log(`Part1: ${validMiddles}`);
console.log(`Part2: ${postSortingMiddles}`);

function testPrints(printLine: Array<String>): boolean {
    let seenPages: Set<String> = new Set();
    for (let page of printLine) {
        let mustBeBefores = before.get(page)
        for (let mustBeBefore of mustBeBefores) {
            if (seenPages.has(mustBeBefore)) {
                return false
            }
        }
        seenPages.add(page)
    }
    return true
}

function sortPages(printLine: Array<String>): Array<String> {
    let remainingPages: Set<String> = new Set(printLine);
    let sortedPages: Array<String> = [];
    while (remainingPages.size > 0) {
        for (let page of remainingPages) {
            if (after.get(page).intersection(remainingPages).size == 0) {
                sortedPages.push(page);
                remainingPages.delete(page)
                break
            }
        }
    }

    return sortedPages
}
