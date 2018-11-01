import {BelongsToMany, Column, DataType, Model, Table} from "sequelize-typescript";
import Tracker from "./tracker";
import TrackerCategories from "./trackerCategories";
import UserCategories from "./userCategories";
import User from "./user";

@Table({
    timestamps: true,
})
export default class Category extends Model<Category> {
    @Column({type: DataType.STRING})
    name: string;

    @Column({type: DataType.STRING})
    description: string;

    @BelongsToMany(() => Tracker, () => TrackerCategories)
    trackers: Array<Tracker>;

    @BelongsToMany(() => User, () => UserCategories)
    users_subscribed: Array<User>;
}