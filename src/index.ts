import {CheckerManager} from "./checker/checkerManager";
import * as Path from "path";

export class Server {
    constructor() {
        let x = CheckerManager.loadInstances(Path.join(__dirname, "trackers"))
        console.log(x)
    }
}

new Server();