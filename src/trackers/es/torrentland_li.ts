import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Torrentland_li extends WebChecker {
    public get url() { return "http://torrentland.li"; }
    public get name() { return "torrentland.li"; }
    public get language() { return "es"; }
    public get categories(): Array<String> { return ["movies"]; }
    public get description() { return ""; }
    public get registryUrl() { return "http://torrentland.li/sbg_login_classic.php"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.registryUrl);
        let exists = $(":contains('but registrations are closed')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}