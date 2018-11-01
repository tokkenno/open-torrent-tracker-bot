import {Column, DataType, HasMany, Model, Table} from "sequelize-typescript";
import User from "./user";

@Table
export default class Language extends Model<Language> {
    @Column({ primaryKey: true, type: DataType.STRING, validate: { is: ["^[a-z]{2}$",'i'] } })
    name: string;

    @HasMany(() => User)
    users: Array<User>
}