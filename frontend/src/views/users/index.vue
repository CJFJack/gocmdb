<template>
    <d2-container>
        <template slot="header">用户管理</template>
        <el-button type="primary" @click="add('add')">新增</el-button>
        <template>
            <el-table
                    :data="tableData"
                    style="width: 100%">
                <el-table-column width="1">
                </el-table-column>
                <div v-for="col in tableColumns">
                    <el-table-column :prop="col.key" :label="col.title" v-if="col.key === 'Gender'">
                        <template slot-scope="scope">
                            <span>{{ genderTextMap[scope.row[col.key]] }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.key === 'Status'">
                        <template slot-scope="scope">
                            <span>{{ statusTextMap[scope.row[col.key]] }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else>
                    </el-table-column>
                </div>
                <el-table-column
                        fixed="right"
                        label="操作"
                        width="180">
                    <template slot-scope="scope">
                        <el-button @click="modify(scope.row)" type="success" size="mini">修改</el-button>
                        <el-button type="danger" @click="delConfirm(scope.row)" size="mini">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </template>

        <el-dialog title="新增 / 编辑" :visible.sync="dialogVisible">
            <el-form :model="form" label-width="120px">
                <el-form-item label="员工ID" prop="StaffID">
                    <el-input v-model="form.StaffID" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="用户名" prop="Name">
                    <el-input v-model="form.Name" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="昵称" prop="NickName">
                    <el-input v-model="form.NickName" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="登录密码" prop="Password">
                    <el-input type="password" v-model="form.Password" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="性别">
                    <el-select v-model="form.Gender" placeholder="请选择性别">
                        <el-option v-for="(text, key, index) in genderTextMap" :label="text" :value="index"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="电话" prop="Tel">
                    <el-input v-model="form.Tel" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="Email">
                    <el-input v-model="form.Email" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="地址" prop="Addr">
                    <el-input v-model="form.Addr" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="部门" prop="Department">
                    <el-input v-model="form.Department" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="form.Status" placeholder="请选择状态">
                        <el-option v-for="(text, key, index) in statusTextMap" :label="text" :value="index"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveForm">确 定</el-button>
            </div>
        </el-dialog>
    </d2-container>
</template>

<script>
    export default {
        name: "users",
        data() {
            return {
                tableData: [],
                tableColumns: [],
                dialogVisible: false,
                form: {
                    StaffID: '',
                    Name: '',
                    NickName: '',
                    Password: '',
                    Gender: 0,
                    Tel: "",
                    Email: "",
                    Addr: "",
                    Department: "",
                    Status: 0,
                },
                genderTextMap: {},
                statusTextMap: {},
                action: '',
            }
        },
        methods: {
            async listTable() {
                const res = await this.$api.LIST_USERS()
                this.tableData = res.tableData
                this.tableColumns = res.tableColumns
                this.genderTextMap = res.genderTextMap
                this.statusTextMap = res.statusTextMap
            },
            async saveForm() {
                if (this.action === 'add') {
                    await this.$api.USER_ADD(this.form)
                    this.listTable()
                } else if (this.action === 'modify') {
                    const res = await this.$api.USER_MODIFY(this.form)
                    this.listTable()
                }
                this.dialogVisible = false
            },

            add(action) {
                this.form = {
                    StaffID: '',
                        Name: '',
                        NickName: '',
                        Password: '',
                        Gender: 0,
                        Tel: "",
                        Email: "",
                        Addr: "",
                        Department: "",
                        Status: 0,
                }
                this.action = action
                this.dialogVisible = true
            },

            modify(row) {
                this.action = 'modify'
                this.form = row
                this.dialogVisible = true
            },

            async del (row){
                console.log(row)
                let data = {
                    id: row.ID
                }
                await this.$api.USER_DEL(data)
                this.$message({
                    type: 'success',
                    message: '删除成功!'
                });
                this.listTable()
            },

            delConfirm(row) {
                this.$confirm('此操作将删除该用户, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.del(row)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
            },

        },
        mounted() {
            this.listTable()
        }
    }
</script>

<style scoped>

</style>