import {Sequelize} from 'sequelize-typescript';
import {CheckerManager} from "./checker/checkerManager";
import dbConfig from "./config/database.json";

import Language from "./models/language";
import User from "./models/user";
import Category from "./models/category";
import Tracker from "./models/tracker";
import TrackerCategories from "./models/trackerCategories";
import UserCategories from "./models/userCategories";
import * as Path from "path";

export class Server {
    private db: Sequelize;
    private checkerManager: CheckerManager;

    constructor() {
    }

    public async load(): Promise<any> {
        this.checkerManager = new CheckerManager();
        await this.loadDatabase();
        return null;
    }

    private async loadDatabase(): Promise<any> {
        this.db = new Sequelize(dbConfig);
        this.db.addModels([Language, User, Category, Tracker, TrackerCategories, UserCategories]);
        return this.db.sync({force: true})
    }

    public async init(): Promise<any> {
        return this.checkerManager.monitorize(Path.join(__dirname, "trackers"), 1000 * 60 * 10);
    }
}

const server = new Server();
server.load().then(() => {
    return server.init();
});