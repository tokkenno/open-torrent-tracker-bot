import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Tbplus_li extends WebChecker {
    public get url() { return "https://tbplus.li"; }
    public get name() { return "tbplus.li"; }
    public get language() { return "es"; }
    public get categories(): Array<String> { return ["movies"]; }
    public get description() { return ""; }
    public get registryUrl() { return "https://tbplus.li/index.php?page=signup"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.registryUrl);
        let exists = $(":contains('pero los registros estan cerrados')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}