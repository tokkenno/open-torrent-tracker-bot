import {Column, DataType, HasMany, Model, Table} from "sequelize-typescript";
import Language from "./language";
import Category from "./category";
import Tracker from "./tracker";

@Table({
    timestamps: true,
})
export default class User extends Model<User> {
    @Column({type: DataType.STRING})
    username: string;

    @HasMany(() => Language)
    subscribe_lang: Array<Language>;

    @HasMany(() => Category)
    subscribe_category: Array<Category>;

    @HasMany(() => Tracker)
    subscribe_tracker: Array<Tracker>;
}