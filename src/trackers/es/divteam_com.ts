import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Divteam_com extends WebChecker {
    public get url() { return "https://divteam.com"; }
    public get name() { return "hachede.me"; }
    public get language() { return "es"; }
    public get categories(): Array<String> { return ["movies"]; }
    public get description() { return ""; }
    public get registryUrl() { return "https://divteam.com/app/"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.registryUrl);
        let exists = $(":contains('REGISTRO TEMPORALMENTE DESHABILITADO')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}