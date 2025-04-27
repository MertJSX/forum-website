function isValidEmail(testString: string) {
    const regex: RegExp = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return regex.test(testString);
}

export default isValidEmail