import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Divteam_com extends WebChecker {
    public get url() { return "https://divteam.com/app/"; }

    public get name() { return "hachede.me"; }

    public get language() { return "es"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.url);
        let exists = $(":contains('REGISTRO TEMPORALMENTE DESHABILITADO')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}