import {Column, DataType, ForeignKey, Model, Table} from "sequelize-typescript";
import Tracker from "./tracker";
import Category from "./category";

@Table({
    timestamps: false,
})
export default class TrackerCategories extends Model<TrackerCategories> {
    @Column({type: DataType.INTEGER, primaryKey: true})
    @ForeignKey(() => Tracker)
    idTracker: number;

    @Column({type: DataType.INTEGER, primaryKey: true})
    @ForeignKey(() => Category)
    idCategory: number;
}