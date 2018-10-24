import {Column, DataType, Model, Table} from "sequelize-typescript";

@Table
export default class Language extends Model<Language> {
    @Column({ primaryKey: true, type: DataType.STRING, validate: { is: ["^[a-z]{2}$",'i'] } })
    name: string;
}