const fs = require('node:fs');

const input: String = fs.readFileSync('day9.input', 'utf8').trim();

let disk: Array<number> = [];
let files: Array<[number, number, number]> = [];
let gaps: Array<[number, number]> = [];
let k: number = 0;
for (let [i, c] of [...input].entries()) {
    let count = +c
    for (let j = 0; j < count; j++) {
        if (i % 2 == 0) {
            disk.push(i / 2);
        } else {
            disk.push(-1);
        }
    }
    if (i % 2 == 0) {
        files.push([k, k + count, i / 2]);
    } else {
        gaps.push([k, k+count])
    }
    k += count;
}

let result = 0
let endPointer = disk.length
for (let [i, val] of disk.entries()) {
    if (i >= endPointer) {
        break;
    }
    if (val >= 0) {
        result += i * val;
    } else {
        do {
            endPointer--
        } while (disk[endPointer] < 0)
        if (i >= endPointer) {
            break
        }
        result += i * disk[endPointer]
    }
}

console.log("Part1: ", result)

files.reverse()
for (let fileId in files) {
    let fileSize = files[fileId][1] - files[fileId][0];
    for (let gapId in gaps) {
        if (gaps[gapId][0] > files[fileId][0]) {
            // Haven't found a big enough gap to the left of the file
            break
        }
        let gapSize = gaps[gapId][1] - gaps[gapId][0]
        if (gapSize >= fileSize) {
            files[fileId][0] = gaps[gapId][0];
            files[fileId][1] = files[fileId][0] + fileSize;
            gaps[gapId][0] += fileSize;
            break;
        }
    }
}

result = 0
for (let file of files) {
    let fileSize = file[1] - file[0];
    result += (fileSize * file[0] + fileSize * (fileSize - 1) / 2) * file[2];
}

console.log("Part2: ", result)
