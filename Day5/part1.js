// main()
const arr = [111,22,34,5,113,7,4,3,122,3,4,5,99, 10];
let start = 0;
function intCode() {
    let offset = 0;
    i = 0;
    loop: while (true) {
        const process = getNextInstructionSet();
        if (!process.length) {
            break;
        }
        console.log(process)
    }
}

function getNextInstructionSet() {
    const instruction = arr[start];
    const inst = instruction.toString().split('');
    inst.splice(0, 0, ...(new Array(5 - inst.length).fill('0', 0)));
    const instSet=  arr.slice(start, start + getTotalOperandsFromOpCode(getOpCode(instruction)));
    start += instSet.length;
    return instSet

}
function getOpCode(instruction) {
    instruction = instruction.toString()
    return instruction.substring(instruction.length-1,instruction.length)
}
function getTotalOperandsFromOpCode(opCode) {
    opCode = parseInt(opCode);
    switch (opCode) {
        case 1:
        case 2:
            return 4;
        case 3:
        case 4:
            return 2;
        default:
            return 0
    }
}

intCode();
