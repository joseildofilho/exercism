class Gigasecond {
    private static gigasecond: number = 1_000_000_000

    private bufferDate: Date;

    constructor(_date: Date) {
        this.bufferDate = _date
        this.addSeconds(Gigasecond.gigasecond)
    }

    date(): Date {
        return this.bufferDate
    }

    private addSeconds(seconds: number) {
        const secondsBufferDate = this.bufferDate.getTime()
        const milliseconds: number = seconds * 1000
        this.bufferDate = new Date(secondsBufferDate + milliseconds)
    }
}

export default Gigasecond
