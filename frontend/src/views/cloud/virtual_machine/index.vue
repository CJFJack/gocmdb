<template>
    <d2-container>
        <template slot="header">云主机管理</template>
        <template>
            <el-table
                    size="mini"
                    :data="tableData"
                    style="width: 100%">
                <el-table-column width="1">
                </el-table-column>
                <div v-for="col in tableColumns">
                    <el-table-column :prop="col.key" :label="col.title" v-if="col.title==='IP地址'">
                        <template slot-scope="scope">
                            <span v-if="scope.row['PublicAddrs']">公网地址：{{scope.row["PublicAddrs"] }}</span><br>
                            <span v-if="scope.row['PrivateAddrs']">私网地址：{{scope.row["PrivateAddrs"] }}</span><br>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.title==='时间'" width="230">
                        <template slot-scope="scope">
                            <span>创建时间：{{scope.row["VmCreatedTime"] }}</span><br>
                            <span>过期时间：{{scope.row["VmExpiredTime"] }}</span><br>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.title==='平台'">
                        <template slot-scope="scope">
                            <div v-html="scope.row[col.key].Name"></div>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else-if="col.title==='配置'">
                        <template slot-scope="scope">
                            <span>CPU：{{scope.row["CPU"] }} 核</span><br>
                            <span>内存：{{scope.row["Mem"] }} M</span><br>
                        </template>
                    </el-table-column>
                    <el-table-column :prop="col.key" :label="col.title" v-else>
                    </el-table-column>
                </div>
                <el-table-column
                        fixed="right"
                        label="操作"
                        width="230">
                    <template slot-scope="scope">
                        <el-button type="danger" @click="rebootConfirm(scope.row)" size="mini">重启</el-button>
                        <el-button type="success" @click="startConfirm(scope.row)" size="mini">开机</el-button>
                        <el-button type="warning" @click="stopConfirm(scope.row)" size="mini">关机</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <br>
            <div class="block">
                <el-pagination
                        @size-change="paginationSizeChange"
                        @current-change="paginationCurrentChange"
                        :current-page="pagination.currentPage"
                        :page-sizes="[10, 20, 50, 100, 200]"
                        :page-size="pagination.pageSize"
                        layout="total, sizes, prev, pager, next, jumper"
                        :total="pagination.total"
                        :prev-text="pagination.prevText"
                        :next-text="pagination.nextText">
                </el-pagination>
            </div>
        </template>

        <el-dialog title="新增 / 编辑" :visible.sync="dialogVisible" :destroy-on-close="true">
            <el-form ref="form" :model="form" :rules="rules" label-width="120px">
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
        name: "cloud_virtual_machine",
        data() {
            return {
                tableData: [],
                tableColumns: [],
                tableTotal: 0,
                // 分页设置
                pagination: {
                    currentPage: 1,
                    pageSize: 10,
                    total: 100,
                    prevText: '上一页',
                    nextText: '下一页',
                    pageSizes: [5, 10, 20, 30, 40, 50, 100],
                    layout: 'sizes, prev, pager, next, jumper, ->, total, slot'
                },
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
                rules: {
                    StaffID: [{required: true, message: '请输入员工ID', trigger: 'blur'}],
                    Name: [{required: true, message: '请输入用户名', trigger: 'blur'}],
                    NickName: [{required: true, message: '请输入昵称', trigger: 'blur'}],
                    Password: [{required: true, message: '请输入密码', trigger: 'blur'}],
                },
                genderTextMap: {},
                statusTextMap: {},
                action: '',
            }
        },
        methods: {
            notice(title, message) {
                this.$alert(message, title, {
                    confirmButtonText: '确定',
                });
            },

            async listTable() {
                const data = {
                    pagination: {
                        currentPage: this.pagination.currentPage,
                        pageSize: this.pagination.pageSize,
                    }
                }
                const res = await this.$api.LIST_VIRTUAL_MACHINE(data)
                this.tableData = res.tableData
                this.tableColumns = res.tableColumns
                this.pagination.total = res.tableTotal
            },

            saveForm() {
                if (this.action === 'add') {
                    this.rules.Password = [{required: true, message: '请输入密码', trigger: 'blur'}]
                } else if (this.action === 'modify') {
                    this.rules.Password = [{required: false}]
                }
                this.$refs["form"].validate(async (valid) => {
                    if (valid) {
                        if (this.action === 'add') {
                            await this.$api.USER_ADD(this.form)
                            this.listTable()
                        } else if (this.action === 'modify') {
                            await this.$api.USER_MODIFY(this.form)
                            this.notice("恭喜", "修改成功")
                            this.listTable()
                        }
                        this.dialogVisible = false
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
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

            async del(row) {
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

            async start(row) {
                await this.$api.START_VIRTUAL_MACHINE(row)
                this.listTable()
                this.$message({
                    type: 'success',
                    message: '启动成功'
                });
            },

            async stop(row) {
                await this.$api.STOP_VIRTUAL_MACHINE(row)
                this.listTable()
                this.$message({
                    type: 'success',
                    message: '关机成功'
                });
            },

            async reboot(row) {
                await this.$api.REBOOT_VIRTUAL_MACHINE(row)
                this.listTable()
                this.$message({
                    type: 'success',
                    message: '重启成功'
                });
            },

            rebootConfirm(row) {
                this.$confirm('此操作将重启云主机, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.reboot(row)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消重启'
                    });
                });
            },

            startConfirm(row) {
                this.$confirm('此操作将启动云主机, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.start(row)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消启动'
                    });
                });
            },

            stopConfirm(row) {
                this.$confirm('此操作将关闭云主机, 是否继续?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    this.stop(row)
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消关闭'
                    });
                });
            },

            // 当分页改变时触发
            paginationCurrentChange(currentPage) {
                this.pagination.currentPage = currentPage
                this.listTable()
            },

            // 分页选择器改变后触发的函数
            paginationSizeChange(pageSize) {
                this.pagination.pageSize = pageSize
                this.listTable()
            },

        },
        mounted() {
            this.listTable()
            this.timer = setInterval(() => {
                //获取数据
                this.listTable()
            }, 10 * 1000)
        },
        destroyed() {
            clearInterval(this.timer)
        }
    }
</script>

<style scoped>

</style>
