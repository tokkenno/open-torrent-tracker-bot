import {CheckerResult} from "./checkerResult";

export abstract class Checker {
    abstract readonly name: string;
    abstract readonly language: string;
    abstract readonly url: string;
    abstract readonly registryUrl: string;
    abstract readonly description: string;
    abstract readonly categories: Array<String>;

    abstract isOpen(): Promise<CheckerResult>;
}