<template>
    <el-dialog
            :title="!dataForm.id ? '新增' : '修改'"
            :close-on-click-modal="false"
            :visible.sync="visible">
        <el-form :model="dataForm" :rules="dataRule" ref="dataForm" label-width="80px">
            <el-form-item label="鸟名" prop="Name">
                <el-input v-model="dataForm.Name" placeholder="鸟名"></el-input>
            </el-form-item>
            <el-form-item label="鸟别名" prop="Altername">
                <el-input v-model="dataForm.Altername" placeholder="鸟别名"></el-input>
            </el-form-item>
            <el-form-item label="拉丁名" prop="Latinname">
                <el-input v-model="dataForm.Latinname" placeholder="拉丁名"></el-input>
            </el-form-item>
            <el-form-item label="英文名" prop="Englishname">
                <el-input v-model="dataForm.Englishname" placeholder="英文名"></el-input>
            </el-form-item>
            <el-form-item label="目名" prop="Order">
                <el-input v-model="dataForm.Order" placeholder="目名"></el-input>
            </el-form-item>
            <el-form-item label="科名" prop="Family">
                <el-input v-model="dataForm.Family" placeholder="科名"></el-input>
            </el-form-item>
            <el-form-item label="属名" prop="Genus">
                <el-input v-model="dataForm.Genus" placeholder="属名"></el-input>
            </el-form-item>
            <el-form-item label="分布" prop="Home">
                <el-input v-model="dataForm.Home" placeholder="分布"></el-input>
            </el-form-item>
            <el-form-item label="特征细节" prop="Details">
                <el-input v-model="dataForm.Details" placeholder="特征细节"></el-input>
            </el-form-item>
            <el-form-item label="案值" prop="Price">
                <el-input v-model="dataForm.Price" placeholder="案值"></el-input>
            </el-form-item>
            <el-form-item label="法条" prop="Laws">
                <el-input v-model="dataForm.Laws" placeholder="法条"></el-input>
            </el-form-item>
            <el-form-item label="图片链接" prop="Images">
                <el-input v-model="dataForm.Images" placeholder="图片链接"></el-input>
            </el-form-item>
            <el-form-item label="保护等级" prop="ProtectLevel">
                <el-input v-model="dataForm.ProtectLevel" placeholder="保护等级"></el-input>
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
            Name
        :
            '',
            Altername
        :
            '',
            Latinname
        :
            '',
            Englishname
        :
            '',
            Order
        :
            '',
            Family
        :
            '',
            Genus
        :
            '',
            Home
        :
            '',
            Details
        :
            '',
            Price
        :
            '',
            Laws
        :
            '',
            Images
        :
            '',
            ProtectLevel
        :
            '',
            Id
        :
            0,
        },
            dataRule: {
                Name
            :
                [
                    {required: true, message: '鸟名不能为空', trigger: 'blur'}
                ],
                Altername
            :
                [
                    {required: true, message: '鸟别名不能为空', trigger: 'blur'}
                ],
                Latinname
            :
                [
                    {required: true, message: '拉丁名不能为空', trigger: 'blur'}
                ],
                Englishname
            :
                [
                    {required: true, message: '英文名不能为空', trigger: 'blur'}
                ],
                Order
            :
                [
                    {required: true, message: '目名不能为空', trigger: 'blur'}
                ],
                Family
            :
                [
                    {required: true, message: '科名不能为空', trigger: 'blur'}
                ],
                Genus
            :
                [
                    {required: true, message: '属名不能为空', trigger: 'blur'}
                ],
                Home
            :
                [
                    {required: true, message: '分布不能为空', trigger: 'blur'}
                ],
                Details
            :
                [
                    {required: true, message: '特征细节不能为空', trigger: 'blur'}
                ],
                Price
            :
                [
                    {required: true, message: '案值不能为空', trigger: 'blur'}
                ],
                Laws
            :
                [
                    {required: true, message: '法条不能为空', trigger: 'blur'}
                ],
                Images
            :
                [
                    {required: true, message: '图片链接不能为空', trigger: 'blur'}
                ],
                ProtectLevel
            :
                [
                    {required: true, message: '保护等级不能为空', trigger: 'blur'}
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
                            url: this.$http.adornUrl(`/ai/birds/get/${this.dataForm.Id}`),
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
                            url: this.$http.adornUrl(`/ai/birds/${!this.dataForm.Id ? 'save' : 'update'})`),
                            method: 'post',
                            data: this.$http.adornData({
                            'Name'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Altername'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Latinname'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Englishname'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Order'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Family'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Genus'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Home'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Details'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Price'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Laws'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'Images'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                            'ProtectLevel'
                    :
                        this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ,
                                'Id': this.dataForm.$
                        {
                            $column.AttrName
                        }
                    ||
                        undefined,
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