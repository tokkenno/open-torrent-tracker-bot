import {Column, DataType, HasMany, Model, Table} from "sequelize-typescript";
import Category from "./category";

@Table({
    timestamps: true,
})
export default class Tracker extends Model<Tracker> {
    @Column({type: DataType.STRING})
    name: string;

    @Column({type: DataType.STRING})
    language: string;

    @Column({type: DataType.STRING})
    description: string;

    @HasMany(() => Category)
    categories: Array<Category>;

    @Column({type: DataType.DATE})
    lastSeen: Date;

    @Column({type: DataType.DATE})
    lastOpen: Date;

    @Column({type: DataType.BOOLEAN})
    opened: boolean;
}