import {BelongsTo, BelongsToMany, Column, DataType, ForeignKey, HasMany, Model, Table} from "sequelize-typescript";
import Language from "./language";
import Category from "./category";
import UserCategories from "./userCategories";

@Table({
    timestamps: true,
})
export default class User extends Model<User> {
    @Column({type: DataType.STRING})
    username: string;

    @BelongsTo(() => Language)
    subscribe_lang: Array<Language>;

    @Column({type: DataType.INTEGER})
    @ForeignKey(() => Language)
    subscribe_lang_id: number;

    @BelongsToMany(() => Category, () => UserCategories)
    subscribe_category: Array<Category>;
/*
    @BelongsTo(() => Tracker)
    subscribe_tracker: Array<Tracker>;*/
}