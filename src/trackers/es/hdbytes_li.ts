import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Hdbytes_li extends WebChecker {
    public get url() { return "http://www.hdbytes.li/index.php?page=signup"; }

    public get name() { return "hdbytes.li"; }

    public get language() { return "es"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.url);
        let exists = $(":contains('para registrarte')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}