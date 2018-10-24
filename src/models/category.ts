import {Column, DataType, Model, Table} from "sequelize-typescript";

@Table({
    timestamps: true,
})
export default class Category extends Model<Category> {
    @Column({type: DataType.STRING})
    name: string;

    @Column({type: DataType.STRING})
    description: string;
}