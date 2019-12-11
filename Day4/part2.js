function isValidPassword(password) {
    password = password.toString();
    let sameAdjacentPresent = false;
    let prevSame = false;
    let moreThan2 = false;
    let isValid = true;
    for (let i = 1; i < password.length; i++) {
        const prevChar = password.substr(i - 1, 1);
        const currChar = password.substr(i, 1);
        if (prevChar === currChar) {
            if (prevSame) {
                moreThan2 = true;
            }
            prevSame = true;

        } else {
            if (prevSame && !moreThan2) {
                sameAdjacentPresent = true;
            }
            prevSame = false;
            moreThan2 = false;

        }
        if (prevChar > currChar) {
            isValid = false;
            break;
        }

    }
    if (prevSame && !moreThan2) {
        sameAdjacentPresent = true;
    }
    if (!sameAdjacentPresent) {
        isValid = false;
    }
    return isValid;
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
