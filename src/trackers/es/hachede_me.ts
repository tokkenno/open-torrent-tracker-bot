import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Hachede_me extends WebChecker {
    public get url() { return "https://hachede.me/"; }
    public get name() { return "hachede.me"; }
    public get language() { return "es"; }
    public get categories(): Array<String> { return ["movies"]; }
    public get description() { return ""; }
    public get registryUrl() { return "https://hachede.me/?p=signup&pid=16"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.registryUrl);
        let exists = $(":contains('Lo sentimos, pero en estos momentos los registros')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}