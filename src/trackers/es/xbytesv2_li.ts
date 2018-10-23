import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Xbytesv2_li extends WebChecker {
    public get url() { return "http://xbytesv2.li/sbg_login_classic.php"; }

    public get name() { return "xbytesv2.li"; }

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