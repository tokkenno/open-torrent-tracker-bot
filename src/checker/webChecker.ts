import * as request from "request";
import {Checker} from "./checker";

export abstract class WebChecker extends Checker {
    protected static getUrl(url: string): Promise<string> {
        return new Promise((accept, reject) => {
            try {
                request.get(url, (err, res, body) => {
                    if (err) {
                        reject(err);
                    }
                    else accept(body);
                });
            } catch (err) {
                reject(err);
            }
        })
    }

    protected static getQueryUrl(url: string): Promise<CheerioStatic> {
        return new Promise((accept, reject) => {
            WebChecker.getUrl(url)
                .then((html) => {
                    accept(cheerio.load(html));
                })
                .catch(reject);
        });
    }
}