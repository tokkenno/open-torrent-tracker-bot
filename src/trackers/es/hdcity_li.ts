import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Hdcity_li extends WebChecker {
    public get url() { return "https://hdcity.li/index.php?page=account"; }

    public get name() { return "hdcity.li"; }

    public get language() { return "es"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.url);
        let exists = $(":contains('but registrations are closed')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}