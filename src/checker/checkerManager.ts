import {Checker} from "./checker";
import * as FS from "fs";
import * as Path from "path";

export class CheckerManager {
    private static loadInstance(path: string): Checker {
        try {
            let classDef = require(path).default;
            return (new classDef()) as Checker;
        } catch (e) {
            return null
        }
    }

    private static loadInstances(path: string): Array<Checker> {
        let instances = [];
        for (let file of FS.readdirSync(path)) {
            let filePath = Path.join(path, file);
            let fileInfo = FS.lstatSync(filePath);

            if (fileInfo.isFile()) {
                instances.push(CheckerManager.loadInstance(filePath));
            } else if (fileInfo.isDirectory() && file.indexOf("..") == -1) {
                instances = instances.concat(this.loadInstances(filePath));
            }
        }
        return instances.filter((instance) => instance != null);
    }
}