<template>
    <el-dialog
            :title="!dataForm.id ? '新增' : '修改'"
            :close-on-click-modal="false"
            :visible.sync="visible">
        <el-form :model="dataForm" :rules="dataRule" ref="dataForm" label-width="80px">
            <el-form-item label="参数名" prop="ParamKey">
                <el-input v-model="dataForm.ParamKey" placeholder="参数名"></el-input>
            </el-form-item>
            <el-form-item label="参数值" prop="ParamValue">
                <el-input v-model="dataForm.ParamValue" placeholder="参数值"></el-input>
            </el-form-item>
            <el-form-item label="备注" prop="Remark">
                <el-input v-model="dataForm.Remark" placeholder="备注"></el-input>
            </el-form-item>
            <el-form-item label="创建时间" prop="CreateTime">
                <el-input v-model="dataForm.CreateTime" placeholder="创建时间"></el-input>
            </el-form-item>
            <el-form-item label="更新时间" prop="UpdateTime">
                <el-input v-model="dataForm.UpdateTime" placeholder="更新时间"></el-input>
            </el-form-item>
            <el-form-item label="状态   0：隐藏   1：显示" prop="Status">
                <el-input v-model="dataForm.Status" placeholder="状态   0：隐藏   1：显示"></el-input>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" @click="dataFormSubmit()">确定</el-button>
    </span>
    </el-dialog>
</template>

<script>
    export default {
        data() {
            return {
                visible: false,
                dataForm: {
            Id
        :
            0,
            ParamKey
        :
            '',
            ParamValue
        :
            '',
            Remark
        :
            '',
            CreateTime
        :
            '',
            UpdateTime
        :
            '',
            Status
        :
            '',
        },
            dataRule: {
                ParamKey
            :
                [
                    {required: true, message: '参数名不能为空', trigger: 'blur'}
                ],
                ParamValue
            :
                [
                    {required: true, message: '参数值不能为空', trigger: 'blur'}
                ],
                Remark
            :
                [
                    {required: true, message: '备注不能为空', trigger: 'blur'}
                ],
                CreateTime
            :
                [
                    {required: true, message: '创建时间不能为空', trigger: 'blur'}
                ],
                UpdateTime
            :
                [
                    {required: true, message: '更新时间不能为空', trigger: 'blur'}
                ],
                Status
            :
                [
                    {required: true, message: '状态   0：隐藏   1：显示不能为空', trigger: 'blur'}
                ],
            }
        }
        },
        methods: {
            init(id) {
                this.dataForm.$
                {.
                    pk.AttrName
                }
                = id || 0
                this.visible = true
                this.$nextTick(() => {
                    this.$refs['dataForm'].resetFields()
                    if (this.dataForm.Id
                )
                    {
                        this.$http({
                            url: this.$http.adornUrl(`/sys/sysconfig/get/${this.dataForm.Id}`),
                            method: 'get',
                            params: this.$http.adornParams()
                        }).then(({data}) => {
                            if (data && data.code === 0) {
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                                this.dataForm.$
                                {
                                    $column.AttrName
                                }
                                = data.$
                                {
                                    $.classname
                                }
                            .
                                $
                                {
                                    $column.AttrName
                                }
                            }
                        })
                    }
                })
            },
            // 表单提交
            dataFormSubmit() {
                this.$refs['dataForm'].validate((valid) => {
                    if (valid) {
                        this.$http({
                            url: this.$http.adornUrl(`/sys/sysconfig/${!this.dataForm.Id ? 'save' : 'update'})`),
                            method: 'post',
                            data: this.$http.adornData({
                                'Id': this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ||
                        undefined,
                            'ParamKey'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'ParamValue'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Remark'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'CreateTime'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'UpdateTime'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Status'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                    })
                    }).
                        then(({data}) => {
                            if (data && data.code === 0) {
                                this.$message({
                                    message: '操作成功',
                                    type: 'success',
                                    duration: 1500,
                                    onClose: () => {
                                        this.visible = false
                                        this.$emit('refreshDataList')
                                    }
                                })
                            } else {
                                this.$message.error(data.msg)
                            }
                        })
                    }
                })
            }
        }
    }
</script>