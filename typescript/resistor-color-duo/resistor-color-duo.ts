export class ResistorColor {

    private static colors: Map<string, string> = new Map<string, string>(
        [
            ['black', '0'],
            ['brown', '1'],
            ['red', '2'],
            ['orange', '3'],
            ['yellow', '4'],
            ['green', '5'],
            ['blue', '6'],
            ['violet', '7'],
            ['grey', '8'],
            ['white', '9']
        ]
    );

    private resistance: number = 0;

    constructor(colorsIn: string[]) {
        try {
            const firstValue = this.pickValueFromColor(colorsIn[0])
            const secondValue = this.pickValueFromColor(colorsIn[1])
            this.resistance = parseInt(firstValue + secondValue)
        } catch {
            throw "At least two colors need to be present";
            
        }
    }

    private pickValueFromColor(color: string): string {
        return ResistorColor.colors.get(color)!
    }

    value = (): number => this.resistance;
}
