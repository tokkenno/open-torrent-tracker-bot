import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Hachede_me extends WebChecker {
    public get url() { return "https://hachede.me/?p=signup"; }

    public get name() { return "hachede.me"; }

    public get language() { return "es"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.url);
        let exists = $(":contains('Lo sentimos, pero en estos momentos los registros')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}