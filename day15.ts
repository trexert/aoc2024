import { assert } from 'node:console';
import fs from 'node:fs';

class Point {
    row: number
    col: number
    constructor(row: number, col: number) {
        this.row = row
        this.col = col
    }

    public add(other: Point): Point {
        return new Point(this.row + other.row, this.col + other.col)
    }

    public subtract(other: Point): Point {
        return new Point(this.row - other.row, this.col - other.col)
    }
}

function main() {
    const input: Array<String> = fs.readFileSync('day15.input', 'utf8').trim().split('\n\n');
    const warehouse: Array<Array<String>> = input[0].split('\n').map((s) => s.split(""))
    const instructions: String = input[1].replaceAll("\n", "")

    let robotStart: Point
    for (let row = 0; row < warehouse.length; row++) {
        outerLoop:
        for (let col = 0; col < warehouse[row].length; col++) {
            if (warehouse[row][col] == '@') {
                robotStart = new Point(row, col)
                break outerLoop
            }
        }
    }

    var robotLoc = robotStart
    for (let instruction of instructions) {
        robotLoc = moveRobot(warehouse, robotLoc, instruction)
    }

    console.log(`Part1: ${getBoxGpss(warehouse)}`)
}

function moveRobot(warehouse: Array<Array<String>>, robotLoc: Point, instruction: String): Point {
    let direction: Point
    switch (instruction) {
        case "^":
            direction = new Point(-1, 0)
            break
        case ">":
            direction = new Point(0, 1)
            break
        case "v":
            direction = new Point(1, 0)
            break
        case "<":
            direction = new Point(0, -1)
            break
        default:
            throw new Error(`Unexpected instruction ${instruction}`)
    }

    assert(warehouse[robotLoc.row][robotLoc.col] == "@")
    const firstStep = robotLoc.add(direction)
    var pointToCheck = firstStep
    loop:
    while (true) {
        switch (warehouse[pointToCheck.row][pointToCheck.col]) {
            case "#":
                break loop
            case "O":
                pointToCheck = pointToCheck.add(direction)
                break
            case ".":
                warehouse[pointToCheck.row][pointToCheck.col] = "O"
                warehouse[firstStep.row][firstStep.col] = "@"
                warehouse[robotLoc.row][robotLoc.col] = "."
                robotLoc = firstStep
                break loop
            default:
                throw new Error(`Unexpected warehouse character ${warehouse[pointToCheck.row][pointToCheck.col]}`)
        }
    }
    return robotLoc
}

function getBoxGpss(warehouse: Array<Array<String>>): number {
    var totalGps = 0
    for (let row = 0; row < warehouse.length; row++) {
        for (let col = 0; col < warehouse[row].length; col++) {
            if (warehouse[row][col] == "O") {
                totalGps += row * 100 + col
            }
        }
    }
    return totalGps
}

function printWarehouse(warehouse: Array<Array<String>>) {
    console.log(warehouse.map((line) => line.join("")).join("\n"))
}

if (require.main === module) {
    main();
}