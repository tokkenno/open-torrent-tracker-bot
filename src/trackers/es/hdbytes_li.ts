import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Hdbytes_li extends WebChecker {
    public get url() { return "http://www.hdbytes.li"; }
    public get name() { return "hdbytes.li"; }
    public get language() { return "es"; }
    public get categories(): Array<String> { return ["movies"]; }
    public get description() { return ""; }
    public get registryUrl() { return "http://www.hdbytes.li/index.php?page=signup"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.registryUrl);
        let exists = $(":contains('para registrarte')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}