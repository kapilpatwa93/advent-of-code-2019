function isValidPassword(password) {
    password = password.toString();
    let sameAdjacentPresent = false;
    let isValid = true;
    for (let i = 1; i < password.length; i++) {
        const prevChar = password.substr(i - 1, 1);
        const currChar = password.substr(i, 1);
        if (prevChar === currChar) {
            sameAdjacentPresent = true
        }
        if (prevChar > currChar) {
            isValid = false;
            break;
        }

    }
    if (!sameAdjacentPresent) {
        isValid = false;
    }
    return isValid
}

function main() {
    const start = 271973;
    const end = 785961;
    let totalIsValid = 0;
    for (let i = start; i <= end; i++) {
        if (isValidPassword(i)) {
            totalIsValid++;
        }
    }
    console.log(totalIsValid)
}

main();
