import fs from 'node:fs';

function main() {
    const input: String = fs.readFileSync('day13.input', 'utf8').trim();
    const parsedInput: Array<Machine> = parseInput(input)
    
    const part1Tokens = parsedInput.map((machine) => machine.tokensRequired(0))
        .filter((tokens) => tokens != null)
        .reduce((sum, tokens) => sum + tokens, 0)

    console.log("Part1: ", part1Tokens)
    
    const part2Tokens = parsedInput.map((machine) => machine.tokensRequired(10000000000000))
        .filter((tokens) => tokens != null)
        .reduce((sum, tokens) => sum + tokens, 0)

    console.log("Part2: ", part2Tokens)
}

const re = /Button A: X([\+-]\d+), Y([\+-]\d+)\nButton B: X([\+-]\d+), Y([\+-]\d+)\nPrize: X=(\d+), Y=(\d+)/m
class Machine {
    buttonA: [number, number]
    buttonB: [number, number]
    prize: [number, number]

    constructor(machineDescription: String) {
        var regexMatch = machineDescription.match(re)!!
        // console.log(regexMatch)
        this.buttonA = [parseInt(regexMatch[1]), parseInt(regexMatch[2])]
        this.buttonB = [parseInt(regexMatch[3]), parseInt(regexMatch[4])]
        this.prize = [parseInt(regexMatch[5]), parseInt(regexMatch[6])]
    }

    public tokensRequired(extra: number): number | null {
        const a = this.buttonA[0]
        const b = this.buttonB[0]
        const c = this.buttonA[1]
        const d = this.buttonB[1]
        const px = this.prize[0] + extra
        const py = this.prize[1] + extra
        const determinant = a * d - b * c
        const buttonAs = (px * d + py * (-b)) / determinant
        const buttonBs = (px * (-c) + py * a) / determinant
        return Number.isInteger(buttonAs) && Number.isInteger(buttonBs) ?
            buttonAs * 3 + buttonBs : null
    }
}

function parseInput(input: String): Array<Machine> {
    return input.split("\n\n")
        .map((machineDescription) => new Machine(machineDescription))
}

if (require.main === module) {
    main();
  }
