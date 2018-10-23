import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Tbplus_li extends WebChecker {
    public get url() { return "https://tbplus.li/index.php?page=signup"; }

    public get name() { return "tbplus.li"; }

    public get language() { return "es"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.url);
        let exists = $(":contains('pero los registros estan cerrados')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}