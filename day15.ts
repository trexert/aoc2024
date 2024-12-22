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
    const input: string[] = fs.readFileSync('day15.input', 'utf8').trim().split('\n\n');
    const warehouse: string[][] = input[0].split('\n').map((s) => s.split(""))
    const instructions: string = input[1].replaceAll("\n", "")
    const p2Warehouse: string[][] = warehouse.map((warehouseLine: string[]) => warehouseLine.flatMap((warehouseCell) => {
        var newCell: string[]
        switch (warehouseCell) {
            case "#":
                newCell = ["#", "#"]
                break
            case "O":
                newCell = ["[", "]"]
                break
            case ".":
                newCell = [".", "."]
                break
            case "@":
                newCell = ["@", "."]
                break
            default:
                throw new Error(`Unexpected warehouseCharacter ${warehouseCell}`)
        }
        return newCell
    }))

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
    const robotStartP2: Point = new Point(robotStart.row, robotStart.col * 2)

    var robotLoc = robotStart
    for (let instruction of instructions) {
        robotLoc = moveRobot(warehouse, robotLoc, instruction)
    }

    console.log(`Part1: ${getBoxGpss(warehouse)}`)

    robotLoc = robotStartP2
    // printWarehouse(p2Warehouse)
    for (let instruction of instructions) {
        robotLoc = moveRobotP2(p2Warehouse, robotLoc, instruction)
        // printWarehouse(p2Warehouse)
    }

    console.log(`Part2: ${getBoxGpss(p2Warehouse)}`)
}

function moveRobot(warehouse: string[][], robotLoc: Point, instruction: string): Point {
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

function moveRobotP2(warehouse: string[][], robotLoc: Point, instruction: string): Point {
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
    var pointsToCheck = new Set([firstStep])
    var canMove = true
    loop:
    while (pointsToCheck.size > 0) {
        let nextPointsToCheck: Set<Point> = new Set()
        for (let pointToCheck of pointsToCheck) {
            switch (warehouse[pointToCheck.row][pointToCheck.col]) {
                case "#":
                    canMove = false
                    break loop
                case "]":
                case "[":
                    switch (instruction) {
                        case "<":
                        case ">":
                            nextPointsToCheck.add(pointToCheck.add(direction).add(direction))
                            break
                        case "^":
                        case "v":
                            nextPointsToCheck.add(pointToCheck.add(direction))
                            if (warehouse[pointToCheck.row][pointToCheck.col] == "[") {
                                nextPointsToCheck.add(pointToCheck.add(direction).add(new Point(0, 1)))
                            } else {
                                nextPointsToCheck.add(pointToCheck.add(direction).add(new Point(0, -1)))
                            }
                            break
                    }
                    break
                case ".":
                    break
                default:
                    throw new Error(`Unexpected warehouse character ${warehouse[pointToCheck.row][pointToCheck.col]}`)
            }
        }
        pointsToCheck = nextPointsToCheck
    }

    if (canMove) {
        var pointsToMove = new Map([[firstStep, "@"]])
        warehouse[robotLoc.row][robotLoc.col] = "."
        while (pointsToMove.size > 0) {
            let nextPointsToMove: Map<Point, string> = new Map()
            for (let [pointToMove, newValue] of pointsToMove) {
                switch (warehouse[pointToMove.row][pointToMove.col]) {
                    case "]":
                    case "[":
                        switch (instruction) {
                            case "<":
                            case ">":
                                nextPointsToMove.set(pointToMove.add(direction), warehouse[pointToMove.row][pointToMove.col])
                                break
                            case "^":
                            case "v":
                                nextPointsToMove.set(pointToMove.add(direction), warehouse[pointToMove.row][pointToMove.col])
                                let otherSideOfBox: Point
                                if (warehouse[pointToMove.row][pointToMove.col] == "[") {
                                    nextPointsToMove.set(pointToMove.add(direction).add(new Point(0, 1)), "]")
                                    otherSideOfBox = pointToMove.add(new Point(0, 1))
                                } else {
                                    nextPointsToMove.set(pointToMove.add(direction).add(new Point(0, -1)), "[")
                                    otherSideOfBox = pointToMove.add(new Point(0, -1))
                                }
                                if (!pointsToMove.has(otherSideOfBox)) {
                                    warehouse[otherSideOfBox.row][otherSideOfBox.col] = "."
                                }
                                break
                        }
                        break
                    case ".":
                        break
                    default:
                        throw new Error(`Unexpected warehouse character ${warehouse[pointToMove.row][pointToMove.col]}`)
                }
                warehouse[pointToMove.row][pointToMove.col] = newValue
            }
            pointsToMove = nextPointsToMove
        }
        robotLoc = robotLoc.add(direction)
    }

    return robotLoc
}

function getBoxGpss(warehouse: string[][]): number {
    var totalGps = 0
    for (let row = 0; row < warehouse.length; row++) {
        for (let col = 0; col < warehouse[row].length; col++) {
            if (warehouse[row][col] == "O" || warehouse[row][col] == "[") {
                totalGps += row * 100 + col
            }
        }
    }
    return totalGps
}

function printWarehouse(warehouse: string[][]) {
    console.log(warehouse.map((line) => line.join("")).join("\n"))
}

if (require.main === module) {
    main();
}