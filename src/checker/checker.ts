import {CheckerResult} from "./checkerResult";

export abstract class Checker {
    abstract readonly name: string;
    abstract readonly language: string;
    abstract readonly url: string;

    abstract isOpen(): Promise<CheckerResult>;
}