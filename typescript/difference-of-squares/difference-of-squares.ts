class Squares {
    public squareOfSum: number = 0
    public sumOfSquares: number = 0
    public difference: number = 0

    constructor(n: number) {
        this.setSquareOfSum(n)
        this.setSumOfSquare(n)
        this.setDifference()
    }

    private setSquareOfSum(n: number): void {
        for (let index = 1; index <= n; index++) {
            this.squareOfSum += index
        }
        this.squareOfSum *= this.squareOfSum
    }

    private setSumOfSquare(n: number): void {
        for (let index = 1; index <= n; index++) {
            this.sumOfSquares += index * index
        }
    }

    private setDifference(): void {
        this.difference = this.squareOfSum - this.sumOfSquares
    }
}

export default Squares
