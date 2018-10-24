import {WebChecker} from "../../checker/webChecker";
import {CheckerResult} from "../../checker/checkerResult";

export default class Puntotorrent_ch extends WebChecker {
    public get url() { return "https://xbt.puntotorrent.ch"; }
    public get name() { return "puntotorrent.ch"; }
    public get language() { return "es"; }
    public get categories(): Array<String> { return ["movies"]; }
    public get description() { return ""; }
    public get registryUrl() { return "https://xbt.puntotorrent.ch/index.php?page=signup"; }

    async isOpen(): Promise<CheckerResult> {
        let $ = await WebChecker.getQueryUrl(this.registryUrl);
        let exists = $(":contains('Lo sentimos pero los registros')").length > 0;

        return {
            online: true,
            regClosed: !exists,
            regOpened: false
        };
    }
}