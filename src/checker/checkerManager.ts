import {Checker} from "./checker";
import * as FS from "fs";
import * as Path from "path";
import Timeout = NodeJS.Timeout;

export class ActiveCheckerManager {
    private readonly parent: CheckerManager;
    private readonly path: string;
    private readonly timeoutFunc: Function;
    private readonly timeoutTime: number;
    private readonly timeoutId: Timeout;

    constructor(manager: CheckerManager, path: string, timeout: number) {
        this.parent = manager;
        this.path = path;
        this.timeoutTime = timeout;
        this.timeoutFunc = (context: ActiveCheckerManager) => {
            return () => {
                context.check(context.path).catch((err) => {
                    console.error(err);
                });
            };
        };
        this.timeoutFunc(this)();
    }

    public stop() {
        clearTimeout(this.timeoutId);
    }

    private async check(path: string) {
        console.log("llega")
        let instances = CheckerManager.loadInstances(path);

        if (instances.length > 0) {
            console.log("Updating", instances.length, "trackers...");

            for (let instance of instances) {
                try {
                    console.log("Updating the tracker '" + instance.name + "' (" + instance.language + ")...");
                    let result = await instance.isOpen()
console.log(result)
                } catch (err) {
                    console.error("The check of the tracker '" + instance.name + "' (" + instance.language + ") has failed.\n", err);
                }
            }

            console.log("Updated finished...");
        } else {
            console.log("Don't exists definitions for any tracker. Maybe has defined the incorrect directory? (Directory:" + path + ")");
        }
    }
}

export class CheckerManager {
    public monitorize(path: string, timeInterval: number) {
        if (!timeInterval || timeInterval < 1000) timeInterval = 1000;
        return new ActiveCheckerManager(this, path, timeInterval)
    }

    public static loadInstance(path: string): Checker {
        try {
            let classDef = require(path).default;
            return (new classDef()) as Checker;
        } catch (e) {
            return null
        }
    }

    public static loadInstances(path: string): Array<Checker> {
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