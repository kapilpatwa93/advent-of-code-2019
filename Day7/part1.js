
function intCode(input) {
     // arr = [3,8,1001,8,10,8,105,1,0,0,21,38,55,68,93,118,199,280,361,442,99999,3,9,1002,9,2,9,101,5,9,9,102,4,9,9,4,9,99,3,9,101,3,9,9,1002,9,3,9,1001,9,4,9,4,9,99,3,9,101,4,9,9,102,3,9,9,4,9,99,3,9,102,2,9,9,101,4,9,9,102,2,9,9,1001,9,4,9,102,4,9,9,4,9,99,3,9,1002,9,2,9,1001,9,2,9,1002,9,5,9,1001,9,2,9,1002,9,4,9,4,9,99,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,99];
    let arr  = [3,23,3,24,1002,24,10,24,1002,23,-1,23,
         101,5,23,23,1,24,23,23,4,23,99,0,0];
    let offset = 0;
    let i = 0;
    loop: while (true) {
        const process = getNextInstructionSet(offset,arr);
        offset = offset + process.length;
        if (!process.length) {
            break;
        }
        const inst = process[0].toString().split('');
        inst.splice(0, 0, ...(new Array(5 - inst.length).fill('0', 0)));
        const opCode = parseInt(inst[4]);
        if (opCode === 1 || opCode === 2) {
            const param1 = getParameter(inst[2], process[1],arr);
            const param2 = getParameter(inst[1], process[2],arr);
            const param3 = getParameter(1, process[3],arr);
            const res = opCode === 1 ? param1 + param2 : param1 * param2;
            setValue(res, param3,arr)
        } else if (opCode === 3) {
            const data = input[0];
            input = input.slice(1);
            // const inputData = await getUserInput();
            setValue(parseInt(data), process[1],arr);
        } else if (opCode === 4) {
            return getParameter(0, process[1],arr);
        } else if (opCode === 5) {
            const param1 = getParameter(inst[2], process[1],arr);
            const param2 = getParameter(inst[1], process[2],arr);
            if (param1 != 0) {
                offset = param2;
                continue;
            }
        } else if (opCode === 6) {
            const param1 = getParameter(inst[2], process[1],arr);
            const param2 = getParameter(inst[1], process[2],arr);
            if (param1 == 0) {
                offset = param2;
                continue;
            }
        } else if (opCode === 7) {
            const param1 = getParameter(inst[2], process[1],arr);
            const param2 = getParameter(inst[1], process[2],arr);
            const param3 = getParameter(1, process[3],arr);
            if (param1 < param2) {
                setValue(1, param3,arr)
            } else {
                setValue(0, param3,arr)
            }
        } else if (opCode === 8) {
            const param1 = getParameter(inst[2], process[1],arr);
            const param2 = getParameter(inst[1], process[2],arr);
            const param3 = getParameter(1, process[3],arr);
            if (param1 === param2) {
                setValue(1, param3,arr)
            } else {
                setValue(0, param3,arr)
            }
        }
    }
}

function getUserInput(input) {

}

function getNextInstructionSet(offset,arr) {
    const instruction = arr[offset];
    const opCode = getOpCode(instruction);
    const instSet = arr.slice(offset, offset + getTotalOperandsFromOpCode(opCode));
    offset += instSet.length;
    return instSet

}

function getParameter(mode, value,arr) {
    mode = parseInt(mode);
    switch (mode) {
        case 0:
            return arr[value];
        default:
            return value;
    }
}

function setValue(val, pos,arr) {
    arr[pos] = val
}

function getOpCode(instruction) {
    instruction = instruction.toString();
    return instruction.substring(instruction.length - 1, instruction.length)
}

function getTotalOperandsFromOpCode(opCode) {
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
            return 0
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
    const elems = [0,1,2,3,4];
    let output = 0;
    let maxOutput = Number.NEGATIVE_INFINITY
    const permutations = getPermutations(elems);
    permutations.forEach((p) =>{
        p.forEach((phase) => {
            console.log(phase);
            output = intCode([phase,output]);
            // console.log("output", output)
        })
        maxOutput = output < maxOutput ? output : maxOutput
        console.log("----")
    })
    console.log(maxOutput)
    // intCode([3,0]);
}
main()


