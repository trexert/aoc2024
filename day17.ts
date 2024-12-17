import fs from 'node:fs';

function main() {
    const input: String = fs.readFileSync('day17.input', 'utf8').trim();
//     const input = `Register A: 729
// Register B: 0
// Register C: 0

// Program: 0,1,5,4,3,0`
    const re = /Register A: (\d+)\nRegister B: (\d+)\nRegister C: (\d+)\n\nProgram: ([\d,]+)/m
    const match = input.match(re)!!
    const registers: Array<bigint> = match.slice(1,4).map((s) => BigInt(s))
    const program: Array<number> = match[4].split(",").map((s) => +s);

    const computer = new Computer(registers.slice(), program.slice())
    computer.runProgram()
    console.log("Part1: ", computer.output.join(","))
    
    var possibleAs: Array<bigint> = [0n]
    for (const opNumber of program.reverse()) {
        // console.log(opNumber)
        const nextPossibleAs: Array<bigint> = []
        for (const possibleA of possibleAs) {
            for (let i = 0n; i<8n; i++) {
                const a = (possibleA << 3n) + i
                computer.reset(a)
                if (computer.singlePass() == opNumber) {
                    nextPossibleAs.push(a)
                }
            }
        }
        possibleAs = nextPossibleAs
        // console.log(possibleAs)
    }

    possibleAs.sort()

    console.log("Part2: ", possibleAs[0])
}

class Computer {
    registers: Array<bigint>
    program: Array<number>
    instructionPointer: number
    output: Array<number>

    constructor(registers: Array<bigint>, program: Array<number>) {
        this.registers = registers
        this.program = program
        this.instructionPointer = 0
        this.output = []
    }

    public reset(registerA: bigint) {
        this.registers = [registerA, 0n, 0n]
        this.instructionPointer = 0
        this.output = []
    }

    public runProgram() {
        while (this.instructionPointer + 1 < this.program.length) {
            this.performOp()
            // console.log(this)
        }
    }

    public singlePass(): number {
        do {
            this.performOp()
        } while (this.instructionPointer > 0 && this.instructionPointer < this.program.length)
        return this.output[0]
    }

    private performOp() {
        switch (this.program[this.instructionPointer]) {
            case Operation.adv: {
                const operand = this.comboOp()
                this.registers[0] = this.registers[0] >> operand
                break
            }
            case Operation.bxl: {
                const operand = this.literOp()
                this.registers[1] = this.registers[1] ^ BigInt(operand)
                break
            }
            case Operation.bst: {
                const operand = this.comboOp()
                this.registers[1] = operand % 8n
                break
            }
            case Operation.jnz: {
                const operand = this.literOp()
                if (this.registers[0] != 0n) {
                    this.instructionPointer = operand - 2
                }
                break
            }
            case Operation.bxc: {
                this.registers[1] = this.registers[1] ^ this.registers[2]
                break
            }
            case Operation.out: {
                const operand = this.comboOp()
                this.output.push(Number(operand % 8n))
                break
            }
            case Operation.bdv: {
                const operand = this.comboOp()
                this.registers[1] = this.registers[0] >> operand
                break
            }
            case Operation.cdv: {
                const operand = this.comboOp()
                this.registers[2] = this.registers[0] >> operand
                break
            }
            default:
                throw new Error(`Invalid operator ${this.program[this.instructionPointer]}`)
        }
        this.instructionPointer += 2
    }

    private literOp(): number {
        return this.program[this.instructionPointer + 1]
    }

    private comboOp(): bigint {
        const operand = this.program[this.instructionPointer + 1]
        switch(operand) {
            case 0:
            case 1:
            case 2:
            case 3:
                return BigInt(operand)
            case 4:
            case 5:
            case 6:
                return this.registers[operand - 4]
            default:
                throw new Error(`Invalid combo operand, ${this.instructionPointer} ${operand}`)
        }
    } 
}

enum Operation {
    adv = 0,
    bxl = 1,
    bst = 2,
    jnz = 3,
    bxc = 4,
    out = 5,
    bdv = 6,
    cdv = 7,
}

if (require.main === module) {
    main();
}
