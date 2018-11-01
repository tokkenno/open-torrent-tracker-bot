import {Column, DataType, ForeignKey, Model, Table} from "sequelize-typescript";
import User from "./user";
import Category from "./category";

@Table({
    timestamps: false,
})
export default class UserCategories extends Model<UserCategories> {
    @Column({type: DataType.INTEGER, primaryKey: true})
    @ForeignKey(() => Category)
    idCategory: number;

    @Column({type: DataType.INTEGER, primaryKey: true})
    @ForeignKey(() => User)
    idUser: number;
}