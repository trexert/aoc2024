const fs = require('node:fs');

const input: String = fs.readFileSync('day7.input', 'utf8').trim().split('\n');

class OperationRecord {
    total: number
    lastIndex: number
    constructor(total: number, lastIndex: number) {
        this.total = total;
        this.lastIndex = lastIndex;
    }
}

let calibrationResultPart1 = 0
let calibrationResultPart2 = 0
for (let line of input) {
    let splitLine = line.split(' ');
    let target = +splitLine[0].slice(0, -1);
    let values = splitLine.slice(1).map((s: string) => +s);
    if (verifyValuesPart1(target, values)) {
        calibrationResultPart1 += target;
        calibrationResultPart2 += target;
    } else if (verifyValuesPart2(target, values)) {
        calibrationResultPart2 += target;
    }
}

console.log("Part1: ", calibrationResultPart1);
console.log("Part2: ", calibrationResultPart2);

function verifyValuesPart1(target: number, values: Array<number>): boolean {
    let operationHistory = [new OperationRecord(values[0], 0)];
    let possible = false;
    while (operationHistory.length > 0) {
        let currentRecord = operationHistory.pop()!!;
        if (currentRecord.lastIndex >= values.length - 1) {
            // Check total once we've used all the values
            if (currentRecord.total == target) {
                possible = true;
                break;
            }
            continue;
        }

        if (currentRecord.total > target) {
            // Early exit if impossible from this point
            continue;
        }

        let nextValue = values[currentRecord.lastIndex + 1];
        operationHistory.push(new OperationRecord(currentRecord.total + nextValue, currentRecord.lastIndex + 1));
        operationHistory.push(new OperationRecord(currentRecord.total * nextValue, currentRecord.lastIndex + 1));
    }
    return possible
}

function verifyValuesPart2(target: number, values: Array<number>): boolean {
    let operationHistory = [new OperationRecord(values[0], 0)];
    let possible = false;
    while (operationHistory.length > 0) {
        let currentRecord = operationHistory.pop()!!;
        if (currentRecord.lastIndex >= values.length - 1) {
            // Check total once we've used all the values
            if (currentRecord.total == target) {
                possible = true;
                break;
            }
            continue;
        }

        if (currentRecord.total > target) {
            // Early exit if impossible from this point
            continue;
        }

        let nextValue = values[currentRecord.lastIndex + 1];
        operationHistory.push(new OperationRecord(currentRecord.total + nextValue, currentRecord.lastIndex + 1));
        operationHistory.push(new OperationRecord(currentRecord.total * nextValue, currentRecord.lastIndex + 1));
        operationHistory.push(new OperationRecord(+("" + currentRecord.total + nextValue), currentRecord.lastIndex + 1));
    }
    return possible
}

