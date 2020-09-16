class ReverseString {
    static reverse(text: string): string {
        return ReverseString._reverse(text)
    }

    private static _reverse(text: string, acc: string = '', index: number = text.length): string {
        if (!index) return acc

        const accAdded = acc + text[index]
        const _index = index - 1
        return this._reverse(text, accAdded, _index)
    }
}

export default ReverseString
