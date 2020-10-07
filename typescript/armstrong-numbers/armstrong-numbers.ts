class ArmstrongNumbers {
    public static isArmstrongNumber(armstrong: number): boolean {
        return ArmstrongNumbers.armstrongAux(armstrong)[1] == armstrong
    }

    private static armstrongAux(num: number, counter: number = 0): [number, number] {
        if (!num)
            return [counter, 0]
        const sum = this.armstrongAux(Math.floor(num / 10), counter + 1)
        return [sum[0], Math.pow(num % 10, sum[0]) + sum[1]]
    }
}

export default ArmstrongNumbers
