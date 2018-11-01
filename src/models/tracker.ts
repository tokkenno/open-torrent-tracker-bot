import {BelongsToMany, Column, DataType, Model, Table} from "sequelize-typescript";
import Category from "./category";
import TrackerCategories from "./trackerCategories";

@Table({
    timestamps: true,
})
export default class Tracker extends Model<Tracker> {
    @Column({type: DataType.STRING, unique: true})
    name: string;

    @Column({type: DataType.STRING})
    language: string;

    @Column({type: DataType.STRING})
    description: string;

    @BelongsToMany(() => Category, () => TrackerCategories)
    categories: Array<Category>;

    @Column({type: DataType.DATE})
    lastSeen: Date;

    @Column({type: DataType.DATE})
    lastOpen: Date;

    @Column({type: DataType.BOOLEAN})
    opened: boolean;
}