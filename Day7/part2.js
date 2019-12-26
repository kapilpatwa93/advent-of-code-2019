class intCode {
    constructor(arr) {
        this.arr = arr;
        this.offset = 0;
        this.complete = false;

    }

    run(input) {
        let output;
        loop: while (true) {
            const process = this.getNextInstructionSet();
            this.offset = this.offset + process.length;
            if (!process.length) {
                break;
            }
            const inst = process[0].toString().split('');
            inst.splice(0, 0, ...(new Array(5 - inst.length).fill('0', 0)));
            const opCode = parseInt(inst[4]);
            if (opCode === 1 || opCode === 2) {
                const param1 = this.getParameter(inst[2], process[1]);
                const param2 = this.getParameter(inst[1], process[2]);
                const param3 = this.getParameter(1, process[3]);
                const res = opCode === 1 ? param1 + param2 : param1 * param2;
                this.setValue(res, param3)
            } else if (opCode === 3) {
                const data = input[0];
                input = input.slice(1);
                this.setValue(parseInt(data), process[1]);
            } else if (opCode === 4) {
                output = this.getParameter(0, process[1]);
                return output
            } else if (opCode === 5) {
                const param1 = this.getParameter(inst[2], process[1]);
                const param2 = this.getParameter(inst[1], process[2]);
                if (param1 != 0) {
                    this.offset = param2;
                    continue;
                }
            } else if (opCode === 6) {
                const param1 = this.getParameter(inst[2], process[1]);
                const param2 = this.getParameter(inst[1], process[2]);
                if (param1 == 0) {
                    this.offset = param2;
                    continue;
                }
            } else if (opCode === 7) {
                const param1 = this.getParameter(inst[2], process[1]);
                const param2 = this.getParameter(inst[1], process[2]);
                const param3 = this.getParameter(1, process[3]);
                if (param1 < param2) {
                    this.setValue(1, param3)
                } else {
                    this.setValue(0, param3)
                }
            } else if (opCode === 8) {
                const param1 = this.getParameter(inst[2], process[1]);
                const param2 = this.getParameter(inst[1], process[2]);
                const param3 = this.getParameter(1, process[3]);
                if (param1 === param2) {
                    this.setValue(1, param3)
                } else {
                    this.setValue(0, param3)
                }
            } else if (opCode === 9) {
                this.complete = true;
                return null
            }
        }
    }

    getNextInstructionSet() {
        const instruction = this.arr[this.offset];
        const opCode = this.getOpCode(instruction);
        const instSet = this.arr.slice(this.offset, this.offset + this.getTotalOperandsFromOpCode(opCode));
        return instSet
    }

    getParameter(mode, value) {
        mode = parseInt(mode);
        switch (mode) {
            case 0:
                return this.arr[value];
            default:
                return value;
        }
    }

    setValue(val, pos) {
        this.arr[pos] = val
    }

    getOpCode(instruction) {
        instruction = instruction.toString();
        return instruction.substring(instruction.length - 1, instruction.length)
    }

    getTotalOperandsFromOpCode(opCode) {
        opCode = parseInt(opCode);
        switch (opCode) {
            case 1:
            case 2:
            case 7:
            case 8:
                return 4;
            case 3:
            case 4:
                return 2;
            case 5:
            case 6:
                return 3;
            default:
                return 1
        }
    }
}

function getPermutations(elems) {
    const perms = [];
    function swap(pos, arr) {
        if (pos == arr.length - 1) {
            perms.push(arr);
        } else {
            for (let i = pos; i < arr.length; i++) {
                const newArr = arr.slice();
                let temp = newArr[pos];
                newArr[pos] = newArr[i];
                newArr[i] = temp;
                swap(pos + 1, newArr)
            }
        }
    }
    swap(0, elems);
    return perms
}

function main() {
    const elems = [5, 6, 7, 8, 9];
    let maxOutput = Number.NEGATIVE_INFINITY;
    const permutations = getPermutations(elems);
    const arr = [3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 55, 68, 93, 118, 199, 280, 361, 442, 99999, 3, 9, 1002, 9, 2, 9, 101, 5, 9, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 3, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 101, 4, 9, 9, 102, 2, 9, 9, 1001, 9, 4, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99];

    permutations.forEach((p) => {
        const ampArr = [];
        const inputArr = [];
        for (let i = 0; i < 5; i++) {
            if (!ampArr[i]) {
                ampArr.push(new intCode(arr.slice()));
            }
            inputArr.push([p[i]])
        }
        let i = 0;
        let output = 0;
        while (!(ampArr[4].complete)) {
            const input = inputArr[i];
            input.push(output);
            inputArr[i] = [];
            const op = ampArr[i].run(input);
            if (op) {
                output = op
            }
            if (i == 4) {
                i = 0;
            } else {
                i++
            }
        }
        maxOutput = output > maxOutput ? output : maxOutput
    });
    console.log(maxOutput)
}

main();


