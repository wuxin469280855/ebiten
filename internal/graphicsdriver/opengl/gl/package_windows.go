// SPDX-License-Identifier: MIT

package gl

import (
	"errors"
	"math"
	"syscall"
	"unsafe"
)

var (
	gpAccum                                                  uintptr
	gpAccumxOES                                              uintptr
	gpAcquireKeyedMutexWin32EXT                              uintptr
	gpActiveProgramEXT                                       uintptr
	gpActiveShaderProgram                                    uintptr
	gpActiveShaderProgramEXT                                 uintptr
	gpActiveStencilFaceEXT                                   uintptr
	gpActiveTexture                                          uintptr
	gpActiveTextureARB                                       uintptr
	gpActiveVaryingNV                                        uintptr
	gpAlphaFragmentOp1ATI                                    uintptr
	gpAlphaFragmentOp2ATI                                    uintptr
	gpAlphaFragmentOp3ATI                                    uintptr
	gpAlphaFunc                                              uintptr
	gpAlphaFuncxOES                                          uintptr
	gpAlphaToCoverageDitherControlNV                         uintptr
	gpApplyFramebufferAttachmentCMAAINTEL                    uintptr
	gpApplyTextureEXT                                        uintptr
	gpAreProgramsResidentNV                                  uintptr
	gpAreTexturesResident                                    uintptr
	gpAreTexturesResidentEXT                                 uintptr
	gpArrayElement                                           uintptr
	gpArrayElementEXT                                        uintptr
	gpArrayObjectATI                                         uintptr
	gpAsyncMarkerSGIX                                        uintptr
	gpAttachObjectARB                                        uintptr
	gpAttachShader                                           uintptr
	gpBegin                                                  uintptr
	gpBeginConditionalRenderNV                               uintptr
	gpBeginConditionalRenderNVX                              uintptr
	gpBeginFragmentShaderATI                                 uintptr
	gpBeginOcclusionQueryNV                                  uintptr
	gpBeginPerfMonitorAMD                                    uintptr
	gpBeginPerfQueryINTEL                                    uintptr
	gpBeginQuery                                             uintptr
	gpBeginQueryARB                                          uintptr
	gpBeginQueryIndexed                                      uintptr
	gpBeginTransformFeedbackEXT                              uintptr
	gpBeginTransformFeedbackNV                               uintptr
	gpBeginVertexShaderEXT                                   uintptr
	gpBeginVideoCaptureNV                                    uintptr
	gpBindAttribLocation                                     uintptr
	gpBindAttribLocationARB                                  uintptr
	gpBindBuffer                                             uintptr
	gpBindBufferARB                                          uintptr
	gpBindBufferBase                                         uintptr
	gpBindBufferBaseEXT                                      uintptr
	gpBindBufferBaseNV                                       uintptr
	gpBindBufferOffsetEXT                                    uintptr
	gpBindBufferOffsetNV                                     uintptr
	gpBindBufferRange                                        uintptr
	gpBindBufferRangeEXT                                     uintptr
	gpBindBufferRangeNV                                      uintptr
	gpBindBuffersBase                                        uintptr
	gpBindBuffersRange                                       uintptr
	gpBindFragDataLocationEXT                                uintptr
	gpBindFragDataLocationIndexed                            uintptr
	gpBindFragmentShaderATI                                  uintptr
	gpBindFramebuffer                                        uintptr
	gpBindFramebufferEXT                                     uintptr
	gpBindImageTexture                                       uintptr
	gpBindImageTextureEXT                                    uintptr
	gpBindImageTextures                                      uintptr
	gpBindLightParameterEXT                                  uintptr
	gpBindMaterialParameterEXT                               uintptr
	gpBindMultiTextureEXT                                    uintptr
	gpBindParameterEXT                                       uintptr
	gpBindProgramARB                                         uintptr
	gpBindProgramNV                                          uintptr
	gpBindProgramPipeline                                    uintptr
	gpBindProgramPipelineEXT                                 uintptr
	gpBindRenderbuffer                                       uintptr
	gpBindRenderbufferEXT                                    uintptr
	gpBindSampler                                            uintptr
	gpBindSamplers                                           uintptr
	gpBindTexGenParameterEXT                                 uintptr
	gpBindTexture                                            uintptr
	gpBindTextureEXT                                         uintptr
	gpBindTextureUnit                                        uintptr
	gpBindTextureUnitParameterEXT                            uintptr
	gpBindTextures                                           uintptr
	gpBindTransformFeedback                                  uintptr
	gpBindTransformFeedbackNV                                uintptr
	gpBindVertexArray                                        uintptr
	gpBindVertexArrayAPPLE                                   uintptr
	gpBindVertexBuffer                                       uintptr
	gpBindVertexBuffers                                      uintptr
	gpBindVertexShaderEXT                                    uintptr
	gpBindVideoCaptureStreamBufferNV                         uintptr
	gpBindVideoCaptureStreamTextureNV                        uintptr
	gpBinormal3bEXT                                          uintptr
	gpBinormal3bvEXT                                         uintptr
	gpBinormal3dEXT                                          uintptr
	gpBinormal3dvEXT                                         uintptr
	gpBinormal3fEXT                                          uintptr
	gpBinormal3fvEXT                                         uintptr
	gpBinormal3iEXT                                          uintptr
	gpBinormal3ivEXT                                         uintptr
	gpBinormal3sEXT                                          uintptr
	gpBinormal3svEXT                                         uintptr
	gpBinormalPointerEXT                                     uintptr
	gpBitmap                                                 uintptr
	gpBitmapxOES                                             uintptr
	gpBlendBarrierKHR                                        uintptr
	gpBlendBarrierNV                                         uintptr
	gpBlendColor                                             uintptr
	gpBlendColorEXT                                          uintptr
	gpBlendColorxOES                                         uintptr
	gpBlendEquation                                          uintptr
	gpBlendEquationEXT                                       uintptr
	gpBlendEquationIndexedAMD                                uintptr
	gpBlendEquationSeparate                                  uintptr
	gpBlendEquationSeparateEXT                               uintptr
	gpBlendEquationSeparateIndexedAMD                        uintptr
	gpBlendEquationSeparateiARB                              uintptr
	gpBlendEquationiARB                                      uintptr
	gpBlendFunc                                              uintptr
	gpBlendFuncIndexedAMD                                    uintptr
	gpBlendFuncSeparate                                      uintptr
	gpBlendFuncSeparateEXT                                   uintptr
	gpBlendFuncSeparateINGR                                  uintptr
	gpBlendFuncSeparateIndexedAMD                            uintptr
	gpBlendFuncSeparateiARB                                  uintptr
	gpBlendFunciARB                                          uintptr
	gpBlendParameteriNV                                      uintptr
	gpBlitFramebuffer                                        uintptr
	gpBlitFramebufferEXT                                     uintptr
	gpBlitNamedFramebuffer                                   uintptr
	gpBufferAddressRangeNV                                   uintptr
	gpBufferData                                             uintptr
	gpBufferDataARB                                          uintptr
	gpBufferPageCommitmentARB                                uintptr
	gpBufferParameteriAPPLE                                  uintptr
	gpBufferStorage                                          uintptr
	gpBufferStorageExternalEXT                               uintptr
	gpBufferStorageMemEXT                                    uintptr
	gpBufferSubData                                          uintptr
	gpBufferSubDataARB                                       uintptr
	gpCallCommandListNV                                      uintptr
	gpCallList                                               uintptr
	gpCallLists                                              uintptr
	gpCheckFramebufferStatus                                 uintptr
	gpCheckFramebufferStatusEXT                              uintptr
	gpCheckNamedFramebufferStatus                            uintptr
	gpCheckNamedFramebufferStatusEXT                         uintptr
	gpClampColorARB                                          uintptr
	gpClear                                                  uintptr
	gpClearAccum                                             uintptr
	gpClearAccumxOES                                         uintptr
	gpClearBufferData                                        uintptr
	gpClearBufferSubData                                     uintptr
	gpClearColor                                             uintptr
	gpClearColorIiEXT                                        uintptr
	gpClearColorIuiEXT                                       uintptr
	gpClearColorxOES                                         uintptr
	gpClearDepth                                             uintptr
	gpClearDepthdNV                                          uintptr
	gpClearDepthf                                            uintptr
	gpClearDepthfOES                                         uintptr
	gpClearDepthxOES                                         uintptr
	gpClearIndex                                             uintptr
	gpClearNamedBufferData                                   uintptr
	gpClearNamedBufferDataEXT                                uintptr
	gpClearNamedBufferSubData                                uintptr
	gpClearNamedBufferSubDataEXT                             uintptr
	gpClearNamedFramebufferfi                                uintptr
	gpClearNamedFramebufferfv                                uintptr
	gpClearNamedFramebufferiv                                uintptr
	gpClearNamedFramebufferuiv                               uintptr
	gpClearStencil                                           uintptr
	gpClearTexImage                                          uintptr
	gpClearTexSubImage                                       uintptr
	gpClientActiveTexture                                    uintptr
	gpClientActiveTextureARB                                 uintptr
	gpClientActiveVertexStreamATI                            uintptr
	gpClientAttribDefaultEXT                                 uintptr
	gpClientWaitSync                                         uintptr
	gpClipControl                                            uintptr
	gpClipPlane                                              uintptr
	gpClipPlanefOES                                          uintptr
	gpClipPlanexOES                                          uintptr
	gpColor3b                                                uintptr
	gpColor3bv                                               uintptr
	gpColor3d                                                uintptr
	gpColor3dv                                               uintptr
	gpColor3f                                                uintptr
	gpColor3fVertex3fSUN                                     uintptr
	gpColor3fVertex3fvSUN                                    uintptr
	gpColor3fv                                               uintptr
	gpColor3hNV                                              uintptr
	gpColor3hvNV                                             uintptr
	gpColor3i                                                uintptr
	gpColor3iv                                               uintptr
	gpColor3s                                                uintptr
	gpColor3sv                                               uintptr
	gpColor3ub                                               uintptr
	gpColor3ubv                                              uintptr
	gpColor3ui                                               uintptr
	gpColor3uiv                                              uintptr
	gpColor3us                                               uintptr
	gpColor3usv                                              uintptr
	gpColor3xOES                                             uintptr
	gpColor3xvOES                                            uintptr
	gpColor4b                                                uintptr
	gpColor4bv                                               uintptr
	gpColor4d                                                uintptr
	gpColor4dv                                               uintptr
	gpColor4f                                                uintptr
	gpColor4fNormal3fVertex3fSUN                             uintptr
	gpColor4fNormal3fVertex3fvSUN                            uintptr
	gpColor4fv                                               uintptr
	gpColor4hNV                                              uintptr
	gpColor4hvNV                                             uintptr
	gpColor4i                                                uintptr
	gpColor4iv                                               uintptr
	gpColor4s                                                uintptr
	gpColor4sv                                               uintptr
	gpColor4ub                                               uintptr
	gpColor4ubVertex2fSUN                                    uintptr
	gpColor4ubVertex2fvSUN                                   uintptr
	gpColor4ubVertex3fSUN                                    uintptr
	gpColor4ubVertex3fvSUN                                   uintptr
	gpColor4ubv                                              uintptr
	gpColor4ui                                               uintptr
	gpColor4uiv                                              uintptr
	gpColor4us                                               uintptr
	gpColor4usv                                              uintptr
	gpColor4xOES                                             uintptr
	gpColor4xvOES                                            uintptr
	gpColorFormatNV                                          uintptr
	gpColorFragmentOp1ATI                                    uintptr
	gpColorFragmentOp2ATI                                    uintptr
	gpColorFragmentOp3ATI                                    uintptr
	gpColorMask                                              uintptr
	gpColorMaskIndexedEXT                                    uintptr
	gpColorMaterial                                          uintptr
	gpColorPointer                                           uintptr
	gpColorPointerEXT                                        uintptr
	gpColorPointerListIBM                                    uintptr
	gpColorPointervINTEL                                     uintptr
	gpColorSubTableEXT                                       uintptr
	gpColorTableEXT                                          uintptr
	gpColorTableParameterfvSGI                               uintptr
	gpColorTableParameterivSGI                               uintptr
	gpColorTableSGI                                          uintptr
	gpCombinerInputNV                                        uintptr
	gpCombinerOutputNV                                       uintptr
	gpCombinerParameterfNV                                   uintptr
	gpCombinerParameterfvNV                                  uintptr
	gpCombinerParameteriNV                                   uintptr
	gpCombinerParameterivNV                                  uintptr
	gpCombinerStageParameterfvNV                             uintptr
	gpCommandListSegmentsNV                                  uintptr
	gpCompileCommandListNV                                   uintptr
	gpCompileShader                                          uintptr
	gpCompileShaderARB                                       uintptr
	gpCompileShaderIncludeARB                                uintptr
	gpCompressedMultiTexImage1DEXT                           uintptr
	gpCompressedMultiTexImage2DEXT                           uintptr
	gpCompressedMultiTexImage3DEXT                           uintptr
	gpCompressedMultiTexSubImage1DEXT                        uintptr
	gpCompressedMultiTexSubImage2DEXT                        uintptr
	gpCompressedMultiTexSubImage3DEXT                        uintptr
	gpCompressedTexImage1D                                   uintptr
	gpCompressedTexImage1DARB                                uintptr
	gpCompressedTexImage2D                                   uintptr
	gpCompressedTexImage2DARB                                uintptr
	gpCompressedTexImage3D                                   uintptr
	gpCompressedTexImage3DARB                                uintptr
	gpCompressedTexSubImage1D                                uintptr
	gpCompressedTexSubImage1DARB                             uintptr
	gpCompressedTexSubImage2D                                uintptr
	gpCompressedTexSubImage2DARB                             uintptr
	gpCompressedTexSubImage3D                                uintptr
	gpCompressedTexSubImage3DARB                             uintptr
	gpCompressedTextureImage1DEXT                            uintptr
	gpCompressedTextureImage2DEXT                            uintptr
	gpCompressedTextureImage3DEXT                            uintptr
	gpCompressedTextureSubImage1D                            uintptr
	gpCompressedTextureSubImage1DEXT                         uintptr
	gpCompressedTextureSubImage2D                            uintptr
	gpCompressedTextureSubImage2DEXT                         uintptr
	gpCompressedTextureSubImage3D                            uintptr
	gpCompressedTextureSubImage3DEXT                         uintptr
	gpConservativeRasterParameterfNV                         uintptr
	gpConservativeRasterParameteriNV                         uintptr
	gpConvolutionFilter1DEXT                                 uintptr
	gpConvolutionFilter2DEXT                                 uintptr
	gpConvolutionParameterfEXT                               uintptr
	gpConvolutionParameterfvEXT                              uintptr
	gpConvolutionParameteriEXT                               uintptr
	gpConvolutionParameterivEXT                              uintptr
	gpConvolutionParameterxOES                               uintptr
	gpConvolutionParameterxvOES                              uintptr
	gpCopyBufferSubData                                      uintptr
	gpCopyColorSubTableEXT                                   uintptr
	gpCopyColorTableSGI                                      uintptr
	gpCopyConvolutionFilter1DEXT                             uintptr
	gpCopyConvolutionFilter2DEXT                             uintptr
	gpCopyImageSubData                                       uintptr
	gpCopyImageSubDataNV                                     uintptr
	gpCopyMultiTexImage1DEXT                                 uintptr
	gpCopyMultiTexImage2DEXT                                 uintptr
	gpCopyMultiTexSubImage1DEXT                              uintptr
	gpCopyMultiTexSubImage2DEXT                              uintptr
	gpCopyMultiTexSubImage3DEXT                              uintptr
	gpCopyNamedBufferSubData                                 uintptr
	gpCopyPathNV                                             uintptr
	gpCopyPixels                                             uintptr
	gpCopyTexImage1D                                         uintptr
	gpCopyTexImage1DEXT                                      uintptr
	gpCopyTexImage2D                                         uintptr
	gpCopyTexImage2DEXT                                      uintptr
	gpCopyTexSubImage1D                                      uintptr
	gpCopyTexSubImage1DEXT                                   uintptr
	gpCopyTexSubImage2D                                      uintptr
	gpCopyTexSubImage2DEXT                                   uintptr
	gpCopyTexSubImage3D                                      uintptr
	gpCopyTexSubImage3DEXT                                   uintptr
	gpCopyTextureImage1DEXT                                  uintptr
	gpCopyTextureImage2DEXT                                  uintptr
	gpCopyTextureSubImage1D                                  uintptr
	gpCopyTextureSubImage1DEXT                               uintptr
	gpCopyTextureSubImage2D                                  uintptr
	gpCopyTextureSubImage2DEXT                               uintptr
	gpCopyTextureSubImage3D                                  uintptr
	gpCopyTextureSubImage3DEXT                               uintptr
	gpCoverFillPathInstancedNV                               uintptr
	gpCoverFillPathNV                                        uintptr
	gpCoverStrokePathInstancedNV                             uintptr
	gpCoverStrokePathNV                                      uintptr
	gpCoverageModulationNV                                   uintptr
	gpCoverageModulationTableNV                              uintptr
	gpCreateBuffers                                          uintptr
	gpCreateCommandListsNV                                   uintptr
	gpCreateFramebuffers                                     uintptr
	gpCreateMemoryObjectsEXT                                 uintptr
	gpCreatePerfQueryINTEL                                   uintptr
	gpCreateProgram                                          uintptr
	gpCreateProgramObjectARB                                 uintptr
	gpCreateProgramPipelines                                 uintptr
	gpCreateQueries                                          uintptr
	gpCreateRenderbuffers                                    uintptr
	gpCreateSamplers                                         uintptr
	gpCreateShader                                           uintptr
	gpCreateShaderObjectARB                                  uintptr
	gpCreateShaderProgramEXT                                 uintptr
	gpCreateShaderProgramv                                   uintptr
	gpCreateShaderProgramvEXT                                uintptr
	gpCreateStatesNV                                         uintptr
	gpCreateSyncFromCLeventARB                               uintptr
	gpCreateTextures                                         uintptr
	gpCreateTransformFeedbacks                               uintptr
	gpCreateVertexArrays                                     uintptr
	gpCullFace                                               uintptr
	gpCullParameterdvEXT                                     uintptr
	gpCullParameterfvEXT                                     uintptr
	gpCurrentPaletteMatrixARB                                uintptr
	gpDebugMessageCallback                                   uintptr
	gpDebugMessageCallbackAMD                                uintptr
	gpDebugMessageCallbackARB                                uintptr
	gpDebugMessageCallbackKHR                                uintptr
	gpDebugMessageControl                                    uintptr
	gpDebugMessageControlARB                                 uintptr
	gpDebugMessageControlKHR                                 uintptr
	gpDebugMessageEnableAMD                                  uintptr
	gpDebugMessageInsert                                     uintptr
	gpDebugMessageInsertAMD                                  uintptr
	gpDebugMessageInsertARB                                  uintptr
	gpDebugMessageInsertKHR                                  uintptr
	gpDeformSGIX                                             uintptr
	gpDeformationMap3dSGIX                                   uintptr
	gpDeformationMap3fSGIX                                   uintptr
	gpDeleteAsyncMarkersSGIX                                 uintptr
	gpDeleteBuffers                                          uintptr
	gpDeleteBuffersARB                                       uintptr
	gpDeleteCommandListsNV                                   uintptr
	gpDeleteFencesAPPLE                                      uintptr
	gpDeleteFencesNV                                         uintptr
	gpDeleteFragmentShaderATI                                uintptr
	gpDeleteFramebuffers                                     uintptr
	gpDeleteFramebuffersEXT                                  uintptr
	gpDeleteLists                                            uintptr
	gpDeleteMemoryObjectsEXT                                 uintptr
	gpDeleteNamedStringARB                                   uintptr
	gpDeleteNamesAMD                                         uintptr
	gpDeleteObjectARB                                        uintptr
	gpDeleteOcclusionQueriesNV                               uintptr
	gpDeletePathsNV                                          uintptr
	gpDeletePerfMonitorsAMD                                  uintptr
	gpDeletePerfQueryINTEL                                   uintptr
	gpDeleteProgram                                          uintptr
	gpDeleteProgramPipelines                                 uintptr
	gpDeleteProgramPipelinesEXT                              uintptr
	gpDeleteProgramsARB                                      uintptr
	gpDeleteProgramsNV                                       uintptr
	gpDeleteQueries                                          uintptr
	gpDeleteQueriesARB                                       uintptr
	gpDeleteQueryResourceTagNV                               uintptr
	gpDeleteRenderbuffers                                    uintptr
	gpDeleteRenderbuffersEXT                                 uintptr
	gpDeleteSamplers                                         uintptr
	gpDeleteSemaphoresEXT                                    uintptr
	gpDeleteShader                                           uintptr
	gpDeleteStatesNV                                         uintptr
	gpDeleteSync                                             uintptr
	gpDeleteTextures                                         uintptr
	gpDeleteTexturesEXT                                      uintptr
	gpDeleteTransformFeedbacks                               uintptr
	gpDeleteTransformFeedbacksNV                             uintptr
	gpDeleteVertexArrays                                     uintptr
	gpDeleteVertexArraysAPPLE                                uintptr
	gpDeleteVertexShaderEXT                                  uintptr
	gpDepthBoundsEXT                                         uintptr
	gpDepthBoundsdNV                                         uintptr
	gpDepthFunc                                              uintptr
	gpDepthMask                                              uintptr
	gpDepthRange                                             uintptr
	gpDepthRangeArrayv                                       uintptr
	gpDepthRangeIndexed                                      uintptr
	gpDepthRangedNV                                          uintptr
	gpDepthRangef                                            uintptr
	gpDepthRangefOES                                         uintptr
	gpDepthRangexOES                                         uintptr
	gpDetachObjectARB                                        uintptr
	gpDetachShader                                           uintptr
	gpDetailTexFuncSGIS                                      uintptr
	gpDisable                                                uintptr
	gpDisableClientState                                     uintptr
	gpDisableClientStateIndexedEXT                           uintptr
	gpDisableClientStateiEXT                                 uintptr
	gpDisableIndexedEXT                                      uintptr
	gpDisableVariantClientStateEXT                           uintptr
	gpDisableVertexArrayAttrib                               uintptr
	gpDisableVertexArrayAttribEXT                            uintptr
	gpDisableVertexArrayEXT                                  uintptr
	gpDisableVertexAttribAPPLE                               uintptr
	gpDisableVertexAttribArray                               uintptr
	gpDisableVertexAttribArrayARB                            uintptr
	gpDispatchCompute                                        uintptr
	gpDispatchComputeGroupSizeARB                            uintptr
	gpDispatchComputeIndirect                                uintptr
	gpDrawArrays                                             uintptr
	gpDrawArraysEXT                                          uintptr
	gpDrawArraysIndirect                                     uintptr
	gpDrawArraysInstancedARB                                 uintptr
	gpDrawArraysInstancedBaseInstance                        uintptr
	gpDrawArraysInstancedEXT                                 uintptr
	gpDrawBuffer                                             uintptr
	gpDrawBuffers                                            uintptr
	gpDrawBuffersARB                                         uintptr
	gpDrawBuffersATI                                         uintptr
	gpDrawCommandsAddressNV                                  uintptr
	gpDrawCommandsNV                                         uintptr
	gpDrawCommandsStatesAddressNV                            uintptr
	gpDrawCommandsStatesNV                                   uintptr
	gpDrawElementArrayAPPLE                                  uintptr
	gpDrawElementArrayATI                                    uintptr
	gpDrawElements                                           uintptr
	gpDrawElementsBaseVertex                                 uintptr
	gpDrawElementsIndirect                                   uintptr
	gpDrawElementsInstancedARB                               uintptr
	gpDrawElementsInstancedBaseInstance                      uintptr
	gpDrawElementsInstancedBaseVertex                        uintptr
	gpDrawElementsInstancedBaseVertexBaseInstance            uintptr
	gpDrawElementsInstancedEXT                               uintptr
	gpDrawMeshArraysSUN                                      uintptr
	gpDrawPixels                                             uintptr
	gpDrawRangeElementArrayAPPLE                             uintptr
	gpDrawRangeElementArrayATI                               uintptr
	gpDrawRangeElements                                      uintptr
	gpDrawRangeElementsBaseVertex                            uintptr
	gpDrawRangeElementsEXT                                   uintptr
	gpDrawTextureNV                                          uintptr
	gpDrawTransformFeedback                                  uintptr
	gpDrawTransformFeedbackInstanced                         uintptr
	gpDrawTransformFeedbackNV                                uintptr
	gpDrawTransformFeedbackStream                            uintptr
	gpDrawTransformFeedbackStreamInstanced                   uintptr
	gpDrawVkImageNV                                          uintptr
	gpEGLImageTargetTexStorageEXT                            uintptr
	gpEGLImageTargetTextureStorageEXT                        uintptr
	gpEdgeFlag                                               uintptr
	gpEdgeFlagFormatNV                                       uintptr
	gpEdgeFlagPointer                                        uintptr
	gpEdgeFlagPointerEXT                                     uintptr
	gpEdgeFlagPointerListIBM                                 uintptr
	gpEdgeFlagv                                              uintptr
	gpElementPointerAPPLE                                    uintptr
	gpElementPointerATI                                      uintptr
	gpEnable                                                 uintptr
	gpEnableClientState                                      uintptr
	gpEnableClientStateIndexedEXT                            uintptr
	gpEnableClientStateiEXT                                  uintptr
	gpEnableIndexedEXT                                       uintptr
	gpEnableVariantClientStateEXT                            uintptr
	gpEnableVertexArrayAttrib                                uintptr
	gpEnableVertexArrayAttribEXT                             uintptr
	gpEnableVertexArrayEXT                                   uintptr
	gpEnableVertexAttribAPPLE                                uintptr
	gpEnableVertexAttribArray                                uintptr
	gpEnableVertexAttribArrayARB                             uintptr
	gpEnd                                                    uintptr
	gpEndConditionalRenderNV                                 uintptr
	gpEndConditionalRenderNVX                                uintptr
	gpEndFragmentShaderATI                                   uintptr
	gpEndList                                                uintptr
	gpEndOcclusionQueryNV                                    uintptr
	gpEndPerfMonitorAMD                                      uintptr
	gpEndPerfQueryINTEL                                      uintptr
	gpEndQuery                                               uintptr
	gpEndQueryARB                                            uintptr
	gpEndQueryIndexed                                        uintptr
	gpEndTransformFeedbackEXT                                uintptr
	gpEndTransformFeedbackNV                                 uintptr
	gpEndVertexShaderEXT                                     uintptr
	gpEndVideoCaptureNV                                      uintptr
	gpEvalCoord1d                                            uintptr
	gpEvalCoord1dv                                           uintptr
	gpEvalCoord1f                                            uintptr
	gpEvalCoord1fv                                           uintptr
	gpEvalCoord1xOES                                         uintptr
	gpEvalCoord1xvOES                                        uintptr
	gpEvalCoord2d                                            uintptr
	gpEvalCoord2dv                                           uintptr
	gpEvalCoord2f                                            uintptr
	gpEvalCoord2fv                                           uintptr
	gpEvalCoord2xOES                                         uintptr
	gpEvalCoord2xvOES                                        uintptr
	gpEvalMapsNV                                             uintptr
	gpEvalMesh1                                              uintptr
	gpEvalMesh2                                              uintptr
	gpEvalPoint1                                             uintptr
	gpEvalPoint2                                             uintptr
	gpEvaluateDepthValuesARB                                 uintptr
	gpExecuteProgramNV                                       uintptr
	gpExtractComponentEXT                                    uintptr
	gpFeedbackBuffer                                         uintptr
	gpFeedbackBufferxOES                                     uintptr
	gpFenceSync                                              uintptr
	gpFinalCombinerInputNV                                   uintptr
	gpFinish                                                 uintptr
	gpFinishAsyncSGIX                                        uintptr
	gpFinishFenceAPPLE                                       uintptr
	gpFinishFenceNV                                          uintptr
	gpFinishObjectAPPLE                                      uintptr
	gpFinishTextureSUNX                                      uintptr
	gpFlush                                                  uintptr
	gpFlushMappedBufferRange                                 uintptr
	gpFlushMappedBufferRangeAPPLE                            uintptr
	gpFlushMappedNamedBufferRange                            uintptr
	gpFlushMappedNamedBufferRangeEXT                         uintptr
	gpFlushPixelDataRangeNV                                  uintptr
	gpFlushRasterSGIX                                        uintptr
	gpFlushStaticDataIBM                                     uintptr
	gpFlushVertexArrayRangeAPPLE                             uintptr
	gpFlushVertexArrayRangeNV                                uintptr
	gpFogCoordFormatNV                                       uintptr
	gpFogCoordPointer                                        uintptr
	gpFogCoordPointerEXT                                     uintptr
	gpFogCoordPointerListIBM                                 uintptr
	gpFogCoordd                                              uintptr
	gpFogCoorddEXT                                           uintptr
	gpFogCoorddv                                             uintptr
	gpFogCoorddvEXT                                          uintptr
	gpFogCoordf                                              uintptr
	gpFogCoordfEXT                                           uintptr
	gpFogCoordfv                                             uintptr
	gpFogCoordfvEXT                                          uintptr
	gpFogCoordhNV                                            uintptr
	gpFogCoordhvNV                                           uintptr
	gpFogFuncSGIS                                            uintptr
	gpFogf                                                   uintptr
	gpFogfv                                                  uintptr
	gpFogi                                                   uintptr
	gpFogiv                                                  uintptr
	gpFogxOES                                                uintptr
	gpFogxvOES                                               uintptr
	gpFragmentColorMaterialSGIX                              uintptr
	gpFragmentCoverageColorNV                                uintptr
	gpFragmentLightModelfSGIX                                uintptr
	gpFragmentLightModelfvSGIX                               uintptr
	gpFragmentLightModeliSGIX                                uintptr
	gpFragmentLightModelivSGIX                               uintptr
	gpFragmentLightfSGIX                                     uintptr
	gpFragmentLightfvSGIX                                    uintptr
	gpFragmentLightiSGIX                                     uintptr
	gpFragmentLightivSGIX                                    uintptr
	gpFragmentMaterialfSGIX                                  uintptr
	gpFragmentMaterialfvSGIX                                 uintptr
	gpFragmentMaterialiSGIX                                  uintptr
	gpFragmentMaterialivSGIX                                 uintptr
	gpFrameTerminatorGREMEDY                                 uintptr
	gpFrameZoomSGIX                                          uintptr
	gpFramebufferDrawBufferEXT                               uintptr
	gpFramebufferDrawBuffersEXT                              uintptr
	gpFramebufferFetchBarrierEXT                             uintptr
	gpFramebufferParameteri                                  uintptr
	gpFramebufferReadBufferEXT                               uintptr
	gpFramebufferRenderbuffer                                uintptr
	gpFramebufferRenderbufferEXT                             uintptr
	gpFramebufferSampleLocationsfvARB                        uintptr
	gpFramebufferSampleLocationsfvNV                         uintptr
	gpFramebufferSamplePositionsfvAMD                        uintptr
	gpFramebufferTexture1D                                   uintptr
	gpFramebufferTexture1DEXT                                uintptr
	gpFramebufferTexture2D                                   uintptr
	gpFramebufferTexture2DEXT                                uintptr
	gpFramebufferTexture3D                                   uintptr
	gpFramebufferTexture3DEXT                                uintptr
	gpFramebufferTextureARB                                  uintptr
	gpFramebufferTextureEXT                                  uintptr
	gpFramebufferTextureFaceARB                              uintptr
	gpFramebufferTextureFaceEXT                              uintptr
	gpFramebufferTextureLayer                                uintptr
	gpFramebufferTextureLayerARB                             uintptr
	gpFramebufferTextureLayerEXT                             uintptr
	gpFramebufferTextureMultiviewOVR                         uintptr
	gpFreeObjectBufferATI                                    uintptr
	gpFrontFace                                              uintptr
	gpFrustum                                                uintptr
	gpFrustumfOES                                            uintptr
	gpFrustumxOES                                            uintptr
	gpGenAsyncMarkersSGIX                                    uintptr
	gpGenBuffers                                             uintptr
	gpGenBuffersARB                                          uintptr
	gpGenFencesAPPLE                                         uintptr
	gpGenFencesNV                                            uintptr
	gpGenFragmentShadersATI                                  uintptr
	gpGenFramebuffers                                        uintptr
	gpGenFramebuffersEXT                                     uintptr
	gpGenLists                                               uintptr
	gpGenNamesAMD                                            uintptr
	gpGenOcclusionQueriesNV                                  uintptr
	gpGenPathsNV                                             uintptr
	gpGenPerfMonitorsAMD                                     uintptr
	gpGenProgramPipelines                                    uintptr
	gpGenProgramPipelinesEXT                                 uintptr
	gpGenProgramsARB                                         uintptr
	gpGenProgramsNV                                          uintptr
	gpGenQueries                                             uintptr
	gpGenQueriesARB                                          uintptr
	gpGenQueryResourceTagNV                                  uintptr
	gpGenRenderbuffers                                       uintptr
	gpGenRenderbuffersEXT                                    uintptr
	gpGenSamplers                                            uintptr
	gpGenSemaphoresEXT                                       uintptr
	gpGenSymbolsEXT                                          uintptr
	gpGenTextures                                            uintptr
	gpGenTexturesEXT                                         uintptr
	gpGenTransformFeedbacks                                  uintptr
	gpGenTransformFeedbacksNV                                uintptr
	gpGenVertexArrays                                        uintptr
	gpGenVertexArraysAPPLE                                   uintptr
	gpGenVertexShadersEXT                                    uintptr
	gpGenerateMipmap                                         uintptr
	gpGenerateMipmapEXT                                      uintptr
	gpGenerateMultiTexMipmapEXT                              uintptr
	gpGenerateTextureMipmap                                  uintptr
	gpGenerateTextureMipmapEXT                               uintptr
	gpGetActiveAtomicCounterBufferiv                         uintptr
	gpGetActiveAttrib                                        uintptr
	gpGetActiveAttribARB                                     uintptr
	gpGetActiveSubroutineName                                uintptr
	gpGetActiveSubroutineUniformName                         uintptr
	gpGetActiveSubroutineUniformiv                           uintptr
	gpGetActiveUniform                                       uintptr
	gpGetActiveUniformARB                                    uintptr
	gpGetActiveUniformBlockName                              uintptr
	gpGetActiveUniformBlockiv                                uintptr
	gpGetActiveUniformName                                   uintptr
	gpGetActiveUniformsiv                                    uintptr
	gpGetActiveVaryingNV                                     uintptr
	gpGetArrayObjectfvATI                                    uintptr
	gpGetArrayObjectivATI                                    uintptr
	gpGetAttachedObjectsARB                                  uintptr
	gpGetAttachedShaders                                     uintptr
	gpGetAttribLocation                                      uintptr
	gpGetAttribLocationARB                                   uintptr
	gpGetBooleanIndexedvEXT                                  uintptr
	gpGetBooleanv                                            uintptr
	gpGetBufferParameteriv                                   uintptr
	gpGetBufferParameterivARB                                uintptr
	gpGetBufferParameterui64vNV                              uintptr
	gpGetBufferPointerv                                      uintptr
	gpGetBufferPointervARB                                   uintptr
	gpGetBufferSubData                                       uintptr
	gpGetBufferSubDataARB                                    uintptr
	gpGetClipPlane                                           uintptr
	gpGetClipPlanefOES                                       uintptr
	gpGetClipPlanexOES                                       uintptr
	gpGetColorTableEXT                                       uintptr
	gpGetColorTableParameterfvEXT                            uintptr
	gpGetColorTableParameterfvSGI                            uintptr
	gpGetColorTableParameterivEXT                            uintptr
	gpGetColorTableParameterivSGI                            uintptr
	gpGetColorTableSGI                                       uintptr
	gpGetCombinerInputParameterfvNV                          uintptr
	gpGetCombinerInputParameterivNV                          uintptr
	gpGetCombinerOutputParameterfvNV                         uintptr
	gpGetCombinerOutputParameterivNV                         uintptr
	gpGetCombinerStageParameterfvNV                          uintptr
	gpGetCommandHeaderNV                                     uintptr
	gpGetCompressedMultiTexImageEXT                          uintptr
	gpGetCompressedTexImage                                  uintptr
	gpGetCompressedTexImageARB                               uintptr
	gpGetCompressedTextureImage                              uintptr
	gpGetCompressedTextureImageEXT                           uintptr
	gpGetCompressedTextureSubImage                           uintptr
	gpGetConvolutionFilterEXT                                uintptr
	gpGetConvolutionParameterfvEXT                           uintptr
	gpGetConvolutionParameterivEXT                           uintptr
	gpGetConvolutionParameterxvOES                           uintptr
	gpGetCoverageModulationTableNV                           uintptr
	gpGetDebugMessageLog                                     uintptr
	gpGetDebugMessageLogAMD                                  uintptr
	gpGetDebugMessageLogARB                                  uintptr
	gpGetDebugMessageLogKHR                                  uintptr
	gpGetDetailTexFuncSGIS                                   uintptr
	gpGetDoubleIndexedvEXT                                   uintptr
	gpGetDoublei_v                                           uintptr
	gpGetDoublei_vEXT                                        uintptr
	gpGetDoublev                                             uintptr
	gpGetError                                               uintptr
	gpGetFenceivNV                                           uintptr
	gpGetFinalCombinerInputParameterfvNV                     uintptr
	gpGetFinalCombinerInputParameterivNV                     uintptr
	gpGetFirstPerfQueryIdINTEL                               uintptr
	gpGetFixedvOES                                           uintptr
	gpGetFloatIndexedvEXT                                    uintptr
	gpGetFloati_v                                            uintptr
	gpGetFloati_vEXT                                         uintptr
	gpGetFloatv                                              uintptr
	gpGetFogFuncSGIS                                         uintptr
	gpGetFragDataIndex                                       uintptr
	gpGetFragDataLocationEXT                                 uintptr
	gpGetFragmentLightfvSGIX                                 uintptr
	gpGetFragmentLightivSGIX                                 uintptr
	gpGetFragmentMaterialfvSGIX                              uintptr
	gpGetFragmentMaterialivSGIX                              uintptr
	gpGetFramebufferAttachmentParameteriv                    uintptr
	gpGetFramebufferAttachmentParameterivEXT                 uintptr
	gpGetFramebufferParameterfvAMD                           uintptr
	gpGetFramebufferParameteriv                              uintptr
	gpGetFramebufferParameterivEXT                           uintptr
	gpGetGraphicsResetStatus                                 uintptr
	gpGetGraphicsResetStatusARB                              uintptr
	gpGetGraphicsResetStatusKHR                              uintptr
	gpGetHandleARB                                           uintptr
	gpGetHistogramEXT                                        uintptr
	gpGetHistogramParameterfvEXT                             uintptr
	gpGetHistogramParameterivEXT                             uintptr
	gpGetHistogramParameterxvOES                             uintptr
	gpGetImageHandleARB                                      uintptr
	gpGetImageHandleNV                                       uintptr
	gpGetImageTransformParameterfvHP                         uintptr
	gpGetImageTransformParameterivHP                         uintptr
	gpGetInfoLogARB                                          uintptr
	gpGetInstrumentsSGIX                                     uintptr
	gpGetInteger64v                                          uintptr
	gpGetIntegerIndexedvEXT                                  uintptr
	gpGetIntegeri_v                                          uintptr
	gpGetIntegerui64i_vNV                                    uintptr
	gpGetIntegerui64vNV                                      uintptr
	gpGetIntegerv                                            uintptr
	gpGetInternalformatSampleivNV                            uintptr
	gpGetInternalformati64v                                  uintptr
	gpGetInternalformativ                                    uintptr
	gpGetInvariantBooleanvEXT                                uintptr
	gpGetInvariantFloatvEXT                                  uintptr
	gpGetInvariantIntegervEXT                                uintptr
	gpGetLightfv                                             uintptr
	gpGetLightiv                                             uintptr
	gpGetLightxOES                                           uintptr
	gpGetLightxvOES                                          uintptr
	gpGetListParameterfvSGIX                                 uintptr
	gpGetListParameterivSGIX                                 uintptr
	gpGetLocalConstantBooleanvEXT                            uintptr
	gpGetLocalConstantFloatvEXT                              uintptr
	gpGetLocalConstantIntegervEXT                            uintptr
	gpGetMapAttribParameterfvNV                              uintptr
	gpGetMapAttribParameterivNV                              uintptr
	gpGetMapControlPointsNV                                  uintptr
	gpGetMapParameterfvNV                                    uintptr
	gpGetMapParameterivNV                                    uintptr
	gpGetMapdv                                               uintptr
	gpGetMapfv                                               uintptr
	gpGetMapiv                                               uintptr
	gpGetMapxvOES                                            uintptr
	gpGetMaterialfv                                          uintptr
	gpGetMaterialiv                                          uintptr
	gpGetMaterialxOES                                        uintptr
	gpGetMaterialxvOES                                       uintptr
	gpGetMemoryObjectParameterivEXT                          uintptr
	gpGetMinmaxEXT                                           uintptr
	gpGetMinmaxParameterfvEXT                                uintptr
	gpGetMinmaxParameterivEXT                                uintptr
	gpGetMultiTexEnvfvEXT                                    uintptr
	gpGetMultiTexEnvivEXT                                    uintptr
	gpGetMultiTexGendvEXT                                    uintptr
	gpGetMultiTexGenfvEXT                                    uintptr
	gpGetMultiTexGenivEXT                                    uintptr
	gpGetMultiTexImageEXT                                    uintptr
	gpGetMultiTexLevelParameterfvEXT                         uintptr
	gpGetMultiTexLevelParameterivEXT                         uintptr
	gpGetMultiTexParameterIivEXT                             uintptr
	gpGetMultiTexParameterIuivEXT                            uintptr
	gpGetMultiTexParameterfvEXT                              uintptr
	gpGetMultiTexParameterivEXT                              uintptr
	gpGetMultisamplefv                                       uintptr
	gpGetMultisamplefvNV                                     uintptr
	gpGetNamedBufferParameteri64v                            uintptr
	gpGetNamedBufferParameteriv                              uintptr
	gpGetNamedBufferParameterivEXT                           uintptr
	gpGetNamedBufferParameterui64vNV                         uintptr
	gpGetNamedBufferPointerv                                 uintptr
	gpGetNamedBufferPointervEXT                              uintptr
	gpGetNamedBufferSubData                                  uintptr
	gpGetNamedBufferSubDataEXT                               uintptr
	gpGetNamedFramebufferAttachmentParameteriv               uintptr
	gpGetNamedFramebufferAttachmentParameterivEXT            uintptr
	gpGetNamedFramebufferParameterfvAMD                      uintptr
	gpGetNamedFramebufferParameteriv                         uintptr
	gpGetNamedFramebufferParameterivEXT                      uintptr
	gpGetNamedProgramLocalParameterIivEXT                    uintptr
	gpGetNamedProgramLocalParameterIuivEXT                   uintptr
	gpGetNamedProgramLocalParameterdvEXT                     uintptr
	gpGetNamedProgramLocalParameterfvEXT                     uintptr
	gpGetNamedProgramStringEXT                               uintptr
	gpGetNamedProgramivEXT                                   uintptr
	gpGetNamedRenderbufferParameteriv                        uintptr
	gpGetNamedRenderbufferParameterivEXT                     uintptr
	gpGetNamedStringARB                                      uintptr
	gpGetNamedStringivARB                                    uintptr
	gpGetNextPerfQueryIdINTEL                                uintptr
	gpGetObjectBufferfvATI                                   uintptr
	gpGetObjectBufferivATI                                   uintptr
	gpGetObjectLabel                                         uintptr
	gpGetObjectLabelEXT                                      uintptr
	gpGetObjectLabelKHR                                      uintptr
	gpGetObjectParameterfvARB                                uintptr
	gpGetObjectParameterivAPPLE                              uintptr
	gpGetObjectParameterivARB                                uintptr
	gpGetObjectPtrLabel                                      uintptr
	gpGetObjectPtrLabelKHR                                   uintptr
	gpGetOcclusionQueryivNV                                  uintptr
	gpGetOcclusionQueryuivNV                                 uintptr
	gpGetPathCommandsNV                                      uintptr
	gpGetPathCoordsNV                                        uintptr
	gpGetPathDashArrayNV                                     uintptr
	gpGetPathLengthNV                                        uintptr
	gpGetPathMetricRangeNV                                   uintptr
	gpGetPathMetricsNV                                       uintptr
	gpGetPathParameterfvNV                                   uintptr
	gpGetPathParameterivNV                                   uintptr
	gpGetPathSpacingNV                                       uintptr
	gpGetPerfCounterInfoINTEL                                uintptr
	gpGetPerfMonitorCounterDataAMD                           uintptr
	gpGetPerfMonitorCounterInfoAMD                           uintptr
	gpGetPerfMonitorCounterStringAMD                         uintptr
	gpGetPerfMonitorCountersAMD                              uintptr
	gpGetPerfMonitorGroupStringAMD                           uintptr
	gpGetPerfMonitorGroupsAMD                                uintptr
	gpGetPerfQueryDataINTEL                                  uintptr
	gpGetPerfQueryIdByNameINTEL                              uintptr
	gpGetPerfQueryInfoINTEL                                  uintptr
	gpGetPixelMapfv                                          uintptr
	gpGetPixelMapuiv                                         uintptr
	gpGetPixelMapusv                                         uintptr
	gpGetPixelMapxv                                          uintptr
	gpGetPixelTexGenParameterfvSGIS                          uintptr
	gpGetPixelTexGenParameterivSGIS                          uintptr
	gpGetPixelTransformParameterfvEXT                        uintptr
	gpGetPixelTransformParameterivEXT                        uintptr
	gpGetPointerIndexedvEXT                                  uintptr
	gpGetPointeri_vEXT                                       uintptr
	gpGetPointerv                                            uintptr
	gpGetPointervEXT                                         uintptr
	gpGetPointervKHR                                         uintptr
	gpGetPolygonStipple                                      uintptr
	gpGetProgramBinary                                       uintptr
	gpGetProgramEnvParameterIivNV                            uintptr
	gpGetProgramEnvParameterIuivNV                           uintptr
	gpGetProgramEnvParameterdvARB                            uintptr
	gpGetProgramEnvParameterfvARB                            uintptr
	gpGetProgramInfoLog                                      uintptr
	gpGetProgramInterfaceiv                                  uintptr
	gpGetProgramLocalParameterIivNV                          uintptr
	gpGetProgramLocalParameterIuivNV                         uintptr
	gpGetProgramLocalParameterdvARB                          uintptr
	gpGetProgramLocalParameterfvARB                          uintptr
	gpGetProgramNamedParameterdvNV                           uintptr
	gpGetProgramNamedParameterfvNV                           uintptr
	gpGetProgramParameterdvNV                                uintptr
	gpGetProgramParameterfvNV                                uintptr
	gpGetProgramPipelineInfoLog                              uintptr
	gpGetProgramPipelineInfoLogEXT                           uintptr
	gpGetProgramPipelineiv                                   uintptr
	gpGetProgramPipelineivEXT                                uintptr
	gpGetProgramResourceIndex                                uintptr
	gpGetProgramResourceLocation                             uintptr
	gpGetProgramResourceLocationIndex                        uintptr
	gpGetProgramResourceName                                 uintptr
	gpGetProgramResourcefvNV                                 uintptr
	gpGetProgramResourceiv                                   uintptr
	gpGetProgramStageiv                                      uintptr
	gpGetProgramStringARB                                    uintptr
	gpGetProgramStringNV                                     uintptr
	gpGetProgramSubroutineParameteruivNV                     uintptr
	gpGetProgramiv                                           uintptr
	gpGetProgramivARB                                        uintptr
	gpGetProgramivNV                                         uintptr
	gpGetQueryBufferObjecti64v                               uintptr
	gpGetQueryBufferObjectiv                                 uintptr
	gpGetQueryBufferObjectui64v                              uintptr
	gpGetQueryBufferObjectuiv                                uintptr
	gpGetQueryIndexediv                                      uintptr
	gpGetQueryObjecti64v                                     uintptr
	gpGetQueryObjecti64vEXT                                  uintptr
	gpGetQueryObjectiv                                       uintptr
	gpGetQueryObjectivARB                                    uintptr
	gpGetQueryObjectui64v                                    uintptr
	gpGetQueryObjectui64vEXT                                 uintptr
	gpGetQueryObjectuiv                                      uintptr
	gpGetQueryObjectuivARB                                   uintptr
	gpGetQueryiv                                             uintptr
	gpGetQueryivARB                                          uintptr
	gpGetRenderbufferParameteriv                             uintptr
	gpGetRenderbufferParameterivEXT                          uintptr
	gpGetSamplerParameterIiv                                 uintptr
	gpGetSamplerParameterIuiv                                uintptr
	gpGetSamplerParameterfv                                  uintptr
	gpGetSamplerParameteriv                                  uintptr
	gpGetSemaphoreParameterui64vEXT                          uintptr
	gpGetSeparableFilterEXT                                  uintptr
	gpGetShaderInfoLog                                       uintptr
	gpGetShaderPrecisionFormat                               uintptr
	gpGetShaderSource                                        uintptr
	gpGetShaderSourceARB                                     uintptr
	gpGetShaderiv                                            uintptr
	gpGetSharpenTexFuncSGIS                                  uintptr
	gpGetStageIndexNV                                        uintptr
	gpGetString                                              uintptr
	gpGetSubroutineIndex                                     uintptr
	gpGetSubroutineUniformLocation                           uintptr
	gpGetSynciv                                              uintptr
	gpGetTexBumpParameterfvATI                               uintptr
	gpGetTexBumpParameterivATI                               uintptr
	gpGetTexEnvfv                                            uintptr
	gpGetTexEnviv                                            uintptr
	gpGetTexEnvxvOES                                         uintptr
	gpGetTexFilterFuncSGIS                                   uintptr
	gpGetTexGendv                                            uintptr
	gpGetTexGenfv                                            uintptr
	gpGetTexGeniv                                            uintptr
	gpGetTexGenxvOES                                         uintptr
	gpGetTexImage                                            uintptr
	gpGetTexLevelParameterfv                                 uintptr
	gpGetTexLevelParameteriv                                 uintptr
	gpGetTexLevelParameterxvOES                              uintptr
	gpGetTexParameterIivEXT                                  uintptr
	gpGetTexParameterIuivEXT                                 uintptr
	gpGetTexParameterPointervAPPLE                           uintptr
	gpGetTexParameterfv                                      uintptr
	gpGetTexParameteriv                                      uintptr
	gpGetTexParameterxvOES                                   uintptr
	gpGetTextureHandleARB                                    uintptr
	gpGetTextureHandleNV                                     uintptr
	gpGetTextureImage                                        uintptr
	gpGetTextureImageEXT                                     uintptr
	gpGetTextureLevelParameterfv                             uintptr
	gpGetTextureLevelParameterfvEXT                          uintptr
	gpGetTextureLevelParameteriv                             uintptr
	gpGetTextureLevelParameterivEXT                          uintptr
	gpGetTextureParameterIiv                                 uintptr
	gpGetTextureParameterIivEXT                              uintptr
	gpGetTextureParameterIuiv                                uintptr
	gpGetTextureParameterIuivEXT                             uintptr
	gpGetTextureParameterfv                                  uintptr
	gpGetTextureParameterfvEXT                               uintptr
	gpGetTextureParameteriv                                  uintptr
	gpGetTextureParameterivEXT                               uintptr
	gpGetTextureSamplerHandleARB                             uintptr
	gpGetTextureSamplerHandleNV                              uintptr
	gpGetTextureSubImage                                     uintptr
	gpGetTrackMatrixivNV                                     uintptr
	gpGetTransformFeedbackVaryingEXT                         uintptr
	gpGetTransformFeedbackVaryingNV                          uintptr
	gpGetTransformFeedbacki64_v                              uintptr
	gpGetTransformFeedbacki_v                                uintptr
	gpGetTransformFeedbackiv                                 uintptr
	gpGetUniformBlockIndex                                   uintptr
	gpGetUniformBufferSizeEXT                                uintptr
	gpGetUniformIndices                                      uintptr
	gpGetUniformLocation                                     uintptr
	gpGetUniformLocationARB                                  uintptr
	gpGetUniformOffsetEXT                                    uintptr
	gpGetUniformSubroutineuiv                                uintptr
	gpGetUniformdv                                           uintptr
	gpGetUniformfv                                           uintptr
	gpGetUniformfvARB                                        uintptr
	gpGetUniformi64vARB                                      uintptr
	gpGetUniformi64vNV                                       uintptr
	gpGetUniformiv                                           uintptr
	gpGetUniformivARB                                        uintptr
	gpGetUniformui64vARB                                     uintptr
	gpGetUniformui64vNV                                      uintptr
	gpGetUniformuivEXT                                       uintptr
	gpGetUnsignedBytei_vEXT                                  uintptr
	gpGetUnsignedBytevEXT                                    uintptr
	gpGetVariantArrayObjectfvATI                             uintptr
	gpGetVariantArrayObjectivATI                             uintptr
	gpGetVariantBooleanvEXT                                  uintptr
	gpGetVariantFloatvEXT                                    uintptr
	gpGetVariantIntegervEXT                                  uintptr
	gpGetVariantPointervEXT                                  uintptr
	gpGetVaryingLocationNV                                   uintptr
	gpGetVertexArrayIndexed64iv                              uintptr
	gpGetVertexArrayIndexediv                                uintptr
	gpGetVertexArrayIntegeri_vEXT                            uintptr
	gpGetVertexArrayIntegervEXT                              uintptr
	gpGetVertexArrayPointeri_vEXT                            uintptr
	gpGetVertexArrayPointervEXT                              uintptr
	gpGetVertexArrayiv                                       uintptr
	gpGetVertexAttribArrayObjectfvATI                        uintptr
	gpGetVertexAttribArrayObjectivATI                        uintptr
	gpGetVertexAttribIivEXT                                  uintptr
	gpGetVertexAttribIuivEXT                                 uintptr
	gpGetVertexAttribLdv                                     uintptr
	gpGetVertexAttribLdvEXT                                  uintptr
	gpGetVertexAttribLi64vNV                                 uintptr
	gpGetVertexAttribLui64vARB                               uintptr
	gpGetVertexAttribLui64vNV                                uintptr
	gpGetVertexAttribPointerv                                uintptr
	gpGetVertexAttribPointervARB                             uintptr
	gpGetVertexAttribPointervNV                              uintptr
	gpGetVertexAttribdv                                      uintptr
	gpGetVertexAttribdvARB                                   uintptr
	gpGetVertexAttribdvNV                                    uintptr
	gpGetVertexAttribfv                                      uintptr
	gpGetVertexAttribfvARB                                   uintptr
	gpGetVertexAttribfvNV                                    uintptr
	gpGetVertexAttribiv                                      uintptr
	gpGetVertexAttribivARB                                   uintptr
	gpGetVertexAttribivNV                                    uintptr
	gpGetVideoCaptureStreamdvNV                              uintptr
	gpGetVideoCaptureStreamfvNV                              uintptr
	gpGetVideoCaptureStreamivNV                              uintptr
	gpGetVideoCaptureivNV                                    uintptr
	gpGetVideoi64vNV                                         uintptr
	gpGetVideoivNV                                           uintptr
	gpGetVideoui64vNV                                        uintptr
	gpGetVideouivNV                                          uintptr
	gpGetVkProcAddrNV                                        uintptr
	gpGetnCompressedTexImageARB                              uintptr
	gpGetnTexImageARB                                        uintptr
	gpGetnUniformdvARB                                       uintptr
	gpGetnUniformfv                                          uintptr
	gpGetnUniformfvARB                                       uintptr
	gpGetnUniformfvKHR                                       uintptr
	gpGetnUniformi64vARB                                     uintptr
	gpGetnUniformiv                                          uintptr
	gpGetnUniformivARB                                       uintptr
	gpGetnUniformivKHR                                       uintptr
	gpGetnUniformui64vARB                                    uintptr
	gpGetnUniformuiv                                         uintptr
	gpGetnUniformuivARB                                      uintptr
	gpGetnUniformuivKHR                                      uintptr
	gpGlobalAlphaFactorbSUN                                  uintptr
	gpGlobalAlphaFactordSUN                                  uintptr
	gpGlobalAlphaFactorfSUN                                  uintptr
	gpGlobalAlphaFactoriSUN                                  uintptr
	gpGlobalAlphaFactorsSUN                                  uintptr
	gpGlobalAlphaFactorubSUN                                 uintptr
	gpGlobalAlphaFactoruiSUN                                 uintptr
	gpGlobalAlphaFactorusSUN                                 uintptr
	gpHint                                                   uintptr
	gpHintPGI                                                uintptr
	gpHistogramEXT                                           uintptr
	gpIglooInterfaceSGIX                                     uintptr
	gpImageTransformParameterfHP                             uintptr
	gpImageTransformParameterfvHP                            uintptr
	gpImageTransformParameteriHP                             uintptr
	gpImageTransformParameterivHP                            uintptr
	gpImportMemoryFdEXT                                      uintptr
	gpImportMemoryWin32HandleEXT                             uintptr
	gpImportMemoryWin32NameEXT                               uintptr
	gpImportSemaphoreFdEXT                                   uintptr
	gpImportSemaphoreWin32HandleEXT                          uintptr
	gpImportSemaphoreWin32NameEXT                            uintptr
	gpImportSyncEXT                                          uintptr
	gpIndexFormatNV                                          uintptr
	gpIndexFuncEXT                                           uintptr
	gpIndexMask                                              uintptr
	gpIndexMaterialEXT                                       uintptr
	gpIndexPointer                                           uintptr
	gpIndexPointerEXT                                        uintptr
	gpIndexPointerListIBM                                    uintptr
	gpIndexd                                                 uintptr
	gpIndexdv                                                uintptr
	gpIndexf                                                 uintptr
	gpIndexfv                                                uintptr
	gpIndexi                                                 uintptr
	gpIndexiv                                                uintptr
	gpIndexs                                                 uintptr
	gpIndexsv                                                uintptr
	gpIndexub                                                uintptr
	gpIndexubv                                               uintptr
	gpIndexxOES                                              uintptr
	gpIndexxvOES                                             uintptr
	gpInitNames                                              uintptr
	gpInsertComponentEXT                                     uintptr
	gpInsertEventMarkerEXT                                   uintptr
	gpInstrumentsBufferSGIX                                  uintptr
	gpInterleavedArrays                                      uintptr
	gpInterpolatePathsNV                                     uintptr
	gpInvalidateBufferData                                   uintptr
	gpInvalidateBufferSubData                                uintptr
	gpInvalidateFramebuffer                                  uintptr
	gpInvalidateNamedFramebufferData                         uintptr
	gpInvalidateNamedFramebufferSubData                      uintptr
	gpInvalidateSubFramebuffer                               uintptr
	gpInvalidateTexImage                                     uintptr
	gpInvalidateTexSubImage                                  uintptr
	gpIsAsyncMarkerSGIX                                      uintptr
	gpIsBuffer                                               uintptr
	gpIsBufferARB                                            uintptr
	gpIsBufferResidentNV                                     uintptr
	gpIsCommandListNV                                        uintptr
	gpIsEnabled                                              uintptr
	gpIsEnabledIndexedEXT                                    uintptr
	gpIsFenceAPPLE                                           uintptr
	gpIsFenceNV                                              uintptr
	gpIsFramebuffer                                          uintptr
	gpIsFramebufferEXT                                       uintptr
	gpIsImageHandleResidentARB                               uintptr
	gpIsImageHandleResidentNV                                uintptr
	gpIsList                                                 uintptr
	gpIsMemoryObjectEXT                                      uintptr
	gpIsNameAMD                                              uintptr
	gpIsNamedBufferResidentNV                                uintptr
	gpIsNamedStringARB                                       uintptr
	gpIsObjectBufferATI                                      uintptr
	gpIsOcclusionQueryNV                                     uintptr
	gpIsPathNV                                               uintptr
	gpIsPointInFillPathNV                                    uintptr
	gpIsPointInStrokePathNV                                  uintptr
	gpIsProgram                                              uintptr
	gpIsProgramARB                                           uintptr
	gpIsProgramNV                                            uintptr
	gpIsProgramPipeline                                      uintptr
	gpIsProgramPipelineEXT                                   uintptr
	gpIsQuery                                                uintptr
	gpIsQueryARB                                             uintptr
	gpIsRenderbuffer                                         uintptr
	gpIsRenderbufferEXT                                      uintptr
	gpIsSampler                                              uintptr
	gpIsSemaphoreEXT                                         uintptr
	gpIsShader                                               uintptr
	gpIsStateNV                                              uintptr
	gpIsSync                                                 uintptr
	gpIsTexture                                              uintptr
	gpIsTextureEXT                                           uintptr
	gpIsTextureHandleResidentARB                             uintptr
	gpIsTextureHandleResidentNV                              uintptr
	gpIsTransformFeedback                                    uintptr
	gpIsTransformFeedbackNV                                  uintptr
	gpIsVariantEnabledEXT                                    uintptr
	gpIsVertexArray                                          uintptr
	gpIsVertexArrayAPPLE                                     uintptr
	gpIsVertexAttribEnabledAPPLE                             uintptr
	gpLGPUCopyImageSubDataNVX                                uintptr
	gpLGPUInterlockNVX                                       uintptr
	gpLGPUNamedBufferSubDataNVX                              uintptr
	gpLabelObjectEXT                                         uintptr
	gpLightEnviSGIX                                          uintptr
	gpLightModelf                                            uintptr
	gpLightModelfv                                           uintptr
	gpLightModeli                                            uintptr
	gpLightModeliv                                           uintptr
	gpLightModelxOES                                         uintptr
	gpLightModelxvOES                                        uintptr
	gpLightf                                                 uintptr
	gpLightfv                                                uintptr
	gpLighti                                                 uintptr
	gpLightiv                                                uintptr
	gpLightxOES                                              uintptr
	gpLightxvOES                                             uintptr
	gpLineStipple                                            uintptr
	gpLineWidth                                              uintptr
	gpLineWidthxOES                                          uintptr
	gpLinkProgram                                            uintptr
	gpLinkProgramARB                                         uintptr
	gpListBase                                               uintptr
	gpListDrawCommandsStatesClientNV                         uintptr
	gpListParameterfSGIX                                     uintptr
	gpListParameterfvSGIX                                    uintptr
	gpListParameteriSGIX                                     uintptr
	gpListParameterivSGIX                                    uintptr
	gpLoadIdentity                                           uintptr
	gpLoadIdentityDeformationMapSGIX                         uintptr
	gpLoadMatrixd                                            uintptr
	gpLoadMatrixf                                            uintptr
	gpLoadMatrixxOES                                         uintptr
	gpLoadName                                               uintptr
	gpLoadProgramNV                                          uintptr
	gpLoadTransposeMatrixd                                   uintptr
	gpLoadTransposeMatrixdARB                                uintptr
	gpLoadTransposeMatrixf                                   uintptr
	gpLoadTransposeMatrixfARB                                uintptr
	gpLoadTransposeMatrixxOES                                uintptr
	gpLockArraysEXT                                          uintptr
	gpLogicOp                                                uintptr
	gpMakeBufferNonResidentNV                                uintptr
	gpMakeBufferResidentNV                                   uintptr
	gpMakeImageHandleNonResidentARB                          uintptr
	gpMakeImageHandleNonResidentNV                           uintptr
	gpMakeImageHandleResidentARB                             uintptr
	gpMakeImageHandleResidentNV                              uintptr
	gpMakeNamedBufferNonResidentNV                           uintptr
	gpMakeNamedBufferResidentNV                              uintptr
	gpMakeTextureHandleNonResidentARB                        uintptr
	gpMakeTextureHandleNonResidentNV                         uintptr
	gpMakeTextureHandleResidentARB                           uintptr
	gpMakeTextureHandleResidentNV                            uintptr
	gpMap1d                                                  uintptr
	gpMap1f                                                  uintptr
	gpMap1xOES                                               uintptr
	gpMap2d                                                  uintptr
	gpMap2f                                                  uintptr
	gpMap2xOES                                               uintptr
	gpMapBuffer                                              uintptr
	gpMapBufferARB                                           uintptr
	gpMapBufferRange                                         uintptr
	gpMapControlPointsNV                                     uintptr
	gpMapGrid1d                                              uintptr
	gpMapGrid1f                                              uintptr
	gpMapGrid1xOES                                           uintptr
	gpMapGrid2d                                              uintptr
	gpMapGrid2f                                              uintptr
	gpMapGrid2xOES                                           uintptr
	gpMapNamedBuffer                                         uintptr
	gpMapNamedBufferEXT                                      uintptr
	gpMapNamedBufferRange                                    uintptr
	gpMapNamedBufferRangeEXT                                 uintptr
	gpMapObjectBufferATI                                     uintptr
	gpMapParameterfvNV                                       uintptr
	gpMapParameterivNV                                       uintptr
	gpMapTexture2DINTEL                                      uintptr
	gpMapVertexAttrib1dAPPLE                                 uintptr
	gpMapVertexAttrib1fAPPLE                                 uintptr
	gpMapVertexAttrib2dAPPLE                                 uintptr
	gpMapVertexAttrib2fAPPLE                                 uintptr
	gpMaterialf                                              uintptr
	gpMaterialfv                                             uintptr
	gpMateriali                                              uintptr
	gpMaterialiv                                             uintptr
	gpMaterialxOES                                           uintptr
	gpMaterialxvOES                                          uintptr
	gpMatrixFrustumEXT                                       uintptr
	gpMatrixIndexPointerARB                                  uintptr
	gpMatrixIndexubvARB                                      uintptr
	gpMatrixIndexuivARB                                      uintptr
	gpMatrixIndexusvARB                                      uintptr
	gpMatrixLoad3x2fNV                                       uintptr
	gpMatrixLoad3x3fNV                                       uintptr
	gpMatrixLoadIdentityEXT                                  uintptr
	gpMatrixLoadTranspose3x3fNV                              uintptr
	gpMatrixLoadTransposedEXT                                uintptr
	gpMatrixLoadTransposefEXT                                uintptr
	gpMatrixLoaddEXT                                         uintptr
	gpMatrixLoadfEXT                                         uintptr
	gpMatrixMode                                             uintptr
	gpMatrixMult3x2fNV                                       uintptr
	gpMatrixMult3x3fNV                                       uintptr
	gpMatrixMultTranspose3x3fNV                              uintptr
	gpMatrixMultTransposedEXT                                uintptr
	gpMatrixMultTransposefEXT                                uintptr
	gpMatrixMultdEXT                                         uintptr
	gpMatrixMultfEXT                                         uintptr
	gpMatrixOrthoEXT                                         uintptr
	gpMatrixPopEXT                                           uintptr
	gpMatrixPushEXT                                          uintptr
	gpMatrixRotatedEXT                                       uintptr
	gpMatrixRotatefEXT                                       uintptr
	gpMatrixScaledEXT                                        uintptr
	gpMatrixScalefEXT                                        uintptr
	gpMatrixTranslatedEXT                                    uintptr
	gpMatrixTranslatefEXT                                    uintptr
	gpMaxShaderCompilerThreadsARB                            uintptr
	gpMaxShaderCompilerThreadsKHR                            uintptr
	gpMemoryBarrier                                          uintptr
	gpMemoryBarrierByRegion                                  uintptr
	gpMemoryBarrierEXT                                       uintptr
	gpMemoryObjectParameterivEXT                             uintptr
	gpMinSampleShadingARB                                    uintptr
	gpMinmaxEXT                                              uintptr
	gpMultMatrixd                                            uintptr
	gpMultMatrixf                                            uintptr
	gpMultMatrixxOES                                         uintptr
	gpMultTransposeMatrixd                                   uintptr
	gpMultTransposeMatrixdARB                                uintptr
	gpMultTransposeMatrixf                                   uintptr
	gpMultTransposeMatrixfARB                                uintptr
	gpMultTransposeMatrixxOES                                uintptr
	gpMultiDrawArrays                                        uintptr
	gpMultiDrawArraysEXT                                     uintptr
	gpMultiDrawArraysIndirect                                uintptr
	gpMultiDrawArraysIndirectAMD                             uintptr
	gpMultiDrawArraysIndirectBindlessCountNV                 uintptr
	gpMultiDrawArraysIndirectBindlessNV                      uintptr
	gpMultiDrawArraysIndirectCountARB                        uintptr
	gpMultiDrawElementArrayAPPLE                             uintptr
	gpMultiDrawElements                                      uintptr
	gpMultiDrawElementsBaseVertex                            uintptr
	gpMultiDrawElementsEXT                                   uintptr
	gpMultiDrawElementsIndirect                              uintptr
	gpMultiDrawElementsIndirectAMD                           uintptr
	gpMultiDrawElementsIndirectBindlessCountNV               uintptr
	gpMultiDrawElementsIndirectBindlessNV                    uintptr
	gpMultiDrawElementsIndirectCountARB                      uintptr
	gpMultiDrawRangeElementArrayAPPLE                        uintptr
	gpMultiModeDrawArraysIBM                                 uintptr
	gpMultiModeDrawElementsIBM                               uintptr
	gpMultiTexBufferEXT                                      uintptr
	gpMultiTexCoord1bOES                                     uintptr
	gpMultiTexCoord1bvOES                                    uintptr
	gpMultiTexCoord1d                                        uintptr
	gpMultiTexCoord1dARB                                     uintptr
	gpMultiTexCoord1dv                                       uintptr
	gpMultiTexCoord1dvARB                                    uintptr
	gpMultiTexCoord1f                                        uintptr
	gpMultiTexCoord1fARB                                     uintptr
	gpMultiTexCoord1fv                                       uintptr
	gpMultiTexCoord1fvARB                                    uintptr
	gpMultiTexCoord1hNV                                      uintptr
	gpMultiTexCoord1hvNV                                     uintptr
	gpMultiTexCoord1i                                        uintptr
	gpMultiTexCoord1iARB                                     uintptr
	gpMultiTexCoord1iv                                       uintptr
	gpMultiTexCoord1ivARB                                    uintptr
	gpMultiTexCoord1s                                        uintptr
	gpMultiTexCoord1sARB                                     uintptr
	gpMultiTexCoord1sv                                       uintptr
	gpMultiTexCoord1svARB                                    uintptr
	gpMultiTexCoord1xOES                                     uintptr
	gpMultiTexCoord1xvOES                                    uintptr
	gpMultiTexCoord2bOES                                     uintptr
	gpMultiTexCoord2bvOES                                    uintptr
	gpMultiTexCoord2d                                        uintptr
	gpMultiTexCoord2dARB                                     uintptr
	gpMultiTexCoord2dv                                       uintptr
	gpMultiTexCoord2dvARB                                    uintptr
	gpMultiTexCoord2f                                        uintptr
	gpMultiTexCoord2fARB                                     uintptr
	gpMultiTexCoord2fv                                       uintptr
	gpMultiTexCoord2fvARB                                    uintptr
	gpMultiTexCoord2hNV                                      uintptr
	gpMultiTexCoord2hvNV                                     uintptr
	gpMultiTexCoord2i                                        uintptr
	gpMultiTexCoord2iARB                                     uintptr
	gpMultiTexCoord2iv                                       uintptr
	gpMultiTexCoord2ivARB                                    uintptr
	gpMultiTexCoord2s                                        uintptr
	gpMultiTexCoord2sARB                                     uintptr
	gpMultiTexCoord2sv                                       uintptr
	gpMultiTexCoord2svARB                                    uintptr
	gpMultiTexCoord2xOES                                     uintptr
	gpMultiTexCoord2xvOES                                    uintptr
	gpMultiTexCoord3bOES                                     uintptr
	gpMultiTexCoord3bvOES                                    uintptr
	gpMultiTexCoord3d                                        uintptr
	gpMultiTexCoord3dARB                                     uintptr
	gpMultiTexCoord3dv                                       uintptr
	gpMultiTexCoord3dvARB                                    uintptr
	gpMultiTexCoord3f                                        uintptr
	gpMultiTexCoord3fARB                                     uintptr
	gpMultiTexCoord3fv                                       uintptr
	gpMultiTexCoord3fvARB                                    uintptr
	gpMultiTexCoord3hNV                                      uintptr
	gpMultiTexCoord3hvNV                                     uintptr
	gpMultiTexCoord3i                                        uintptr
	gpMultiTexCoord3iARB                                     uintptr
	gpMultiTexCoord3iv                                       uintptr
	gpMultiTexCoord3ivARB                                    uintptr
	gpMultiTexCoord3s                                        uintptr
	gpMultiTexCoord3sARB                                     uintptr
	gpMultiTexCoord3sv                                       uintptr
	gpMultiTexCoord3svARB                                    uintptr
	gpMultiTexCoord3xOES                                     uintptr
	gpMultiTexCoord3xvOES                                    uintptr
	gpMultiTexCoord4bOES                                     uintptr
	gpMultiTexCoord4bvOES                                    uintptr
	gpMultiTexCoord4d                                        uintptr
	gpMultiTexCoord4dARB                                     uintptr
	gpMultiTexCoord4dv                                       uintptr
	gpMultiTexCoord4dvARB                                    uintptr
	gpMultiTexCoord4f                                        uintptr
	gpMultiTexCoord4fARB                                     uintptr
	gpMultiTexCoord4fv                                       uintptr
	gpMultiTexCoord4fvARB                                    uintptr
	gpMultiTexCoord4hNV                                      uintptr
	gpMultiTexCoord4hvNV                                     uintptr
	gpMultiTexCoord4i                                        uintptr
	gpMultiTexCoord4iARB                                     uintptr
	gpMultiTexCoord4iv                                       uintptr
	gpMultiTexCoord4ivARB                                    uintptr
	gpMultiTexCoord4s                                        uintptr
	gpMultiTexCoord4sARB                                     uintptr
	gpMultiTexCoord4sv                                       uintptr
	gpMultiTexCoord4svARB                                    uintptr
	gpMultiTexCoord4xOES                                     uintptr
	gpMultiTexCoord4xvOES                                    uintptr
	gpMultiTexCoordPointerEXT                                uintptr
	gpMultiTexEnvfEXT                                        uintptr
	gpMultiTexEnvfvEXT                                       uintptr
	gpMultiTexEnviEXT                                        uintptr
	gpMultiTexEnvivEXT                                       uintptr
	gpMultiTexGendEXT                                        uintptr
	gpMultiTexGendvEXT                                       uintptr
	gpMultiTexGenfEXT                                        uintptr
	gpMultiTexGenfvEXT                                       uintptr
	gpMultiTexGeniEXT                                        uintptr
	gpMultiTexGenivEXT                                       uintptr
	gpMultiTexImage1DEXT                                     uintptr
	gpMultiTexImage2DEXT                                     uintptr
	gpMultiTexImage3DEXT                                     uintptr
	gpMultiTexParameterIivEXT                                uintptr
	gpMultiTexParameterIuivEXT                               uintptr
	gpMultiTexParameterfEXT                                  uintptr
	gpMultiTexParameterfvEXT                                 uintptr
	gpMultiTexParameteriEXT                                  uintptr
	gpMultiTexParameterivEXT                                 uintptr
	gpMultiTexRenderbufferEXT                                uintptr
	gpMultiTexSubImage1DEXT                                  uintptr
	gpMultiTexSubImage2DEXT                                  uintptr
	gpMultiTexSubImage3DEXT                                  uintptr
	gpMulticastBarrierNV                                     uintptr
	gpMulticastBlitFramebufferNV                             uintptr
	gpMulticastBufferSubDataNV                               uintptr
	gpMulticastCopyBufferSubDataNV                           uintptr
	gpMulticastCopyImageSubDataNV                            uintptr
	gpMulticastFramebufferSampleLocationsfvNV                uintptr
	gpMulticastGetQueryObjecti64vNV                          uintptr
	gpMulticastGetQueryObjectivNV                            uintptr
	gpMulticastGetQueryObjectui64vNV                         uintptr
	gpMulticastGetQueryObjectuivNV                           uintptr
	gpMulticastWaitSyncNV                                    uintptr
	gpNamedBufferData                                        uintptr
	gpNamedBufferDataEXT                                     uintptr
	gpNamedBufferPageCommitmentARB                           uintptr
	gpNamedBufferPageCommitmentEXT                           uintptr
	gpNamedBufferStorage                                     uintptr
	gpNamedBufferStorageEXT                                  uintptr
	gpNamedBufferStorageExternalEXT                          uintptr
	gpNamedBufferStorageMemEXT                               uintptr
	gpNamedBufferSubData                                     uintptr
	gpNamedBufferSubDataEXT                                  uintptr
	gpNamedCopyBufferSubDataEXT                              uintptr
	gpNamedFramebufferDrawBuffer                             uintptr
	gpNamedFramebufferDrawBuffers                            uintptr
	gpNamedFramebufferParameteri                             uintptr
	gpNamedFramebufferParameteriEXT                          uintptr
	gpNamedFramebufferReadBuffer                             uintptr
	gpNamedFramebufferRenderbuffer                           uintptr
	gpNamedFramebufferRenderbufferEXT                        uintptr
	gpNamedFramebufferSampleLocationsfvARB                   uintptr
	gpNamedFramebufferSampleLocationsfvNV                    uintptr
	gpNamedFramebufferSamplePositionsfvAMD                   uintptr
	gpNamedFramebufferTexture                                uintptr
	gpNamedFramebufferTexture1DEXT                           uintptr
	gpNamedFramebufferTexture2DEXT                           uintptr
	gpNamedFramebufferTexture3DEXT                           uintptr
	gpNamedFramebufferTextureEXT                             uintptr
	gpNamedFramebufferTextureFaceEXT                         uintptr
	gpNamedFramebufferTextureLayer                           uintptr
	gpNamedFramebufferTextureLayerEXT                        uintptr
	gpNamedProgramLocalParameter4dEXT                        uintptr
	gpNamedProgramLocalParameter4dvEXT                       uintptr
	gpNamedProgramLocalParameter4fEXT                        uintptr
	gpNamedProgramLocalParameter4fvEXT                       uintptr
	gpNamedProgramLocalParameterI4iEXT                       uintptr
	gpNamedProgramLocalParameterI4ivEXT                      uintptr
	gpNamedProgramLocalParameterI4uiEXT                      uintptr
	gpNamedProgramLocalParameterI4uivEXT                     uintptr
	gpNamedProgramLocalParameters4fvEXT                      uintptr
	gpNamedProgramLocalParametersI4ivEXT                     uintptr
	gpNamedProgramLocalParametersI4uivEXT                    uintptr
	gpNamedProgramStringEXT                                  uintptr
	gpNamedRenderbufferStorage                               uintptr
	gpNamedRenderbufferStorageEXT                            uintptr
	gpNamedRenderbufferStorageMultisample                    uintptr
	gpNamedRenderbufferStorageMultisampleCoverageEXT         uintptr
	gpNamedRenderbufferStorageMultisampleEXT                 uintptr
	gpNamedStringARB                                         uintptr
	gpNewList                                                uintptr
	gpNewObjectBufferATI                                     uintptr
	gpNormal3b                                               uintptr
	gpNormal3bv                                              uintptr
	gpNormal3d                                               uintptr
	gpNormal3dv                                              uintptr
	gpNormal3f                                               uintptr
	gpNormal3fVertex3fSUN                                    uintptr
	gpNormal3fVertex3fvSUN                                   uintptr
	gpNormal3fv                                              uintptr
	gpNormal3hNV                                             uintptr
	gpNormal3hvNV                                            uintptr
	gpNormal3i                                               uintptr
	gpNormal3iv                                              uintptr
	gpNormal3s                                               uintptr
	gpNormal3sv                                              uintptr
	gpNormal3xOES                                            uintptr
	gpNormal3xvOES                                           uintptr
	gpNormalFormatNV                                         uintptr
	gpNormalPointer                                          uintptr
	gpNormalPointerEXT                                       uintptr
	gpNormalPointerListIBM                                   uintptr
	gpNormalPointervINTEL                                    uintptr
	gpNormalStream3bATI                                      uintptr
	gpNormalStream3bvATI                                     uintptr
	gpNormalStream3dATI                                      uintptr
	gpNormalStream3dvATI                                     uintptr
	gpNormalStream3fATI                                      uintptr
	gpNormalStream3fvATI                                     uintptr
	gpNormalStream3iATI                                      uintptr
	gpNormalStream3ivATI                                     uintptr
	gpNormalStream3sATI                                      uintptr
	gpNormalStream3svATI                                     uintptr
	gpObjectLabel                                            uintptr
	gpObjectLabelKHR                                         uintptr
	gpObjectPtrLabel                                         uintptr
	gpObjectPtrLabelKHR                                      uintptr
	gpObjectPurgeableAPPLE                                   uintptr
	gpObjectUnpurgeableAPPLE                                 uintptr
	gpOrtho                                                  uintptr
	gpOrthofOES                                              uintptr
	gpOrthoxOES                                              uintptr
	gpPNTrianglesfATI                                        uintptr
	gpPNTrianglesiATI                                        uintptr
	gpPassTexCoordATI                                        uintptr
	gpPassThrough                                            uintptr
	gpPassThroughxOES                                        uintptr
	gpPatchParameterfv                                       uintptr
	gpPatchParameteri                                        uintptr
	gpPathCommandsNV                                         uintptr
	gpPathCoordsNV                                           uintptr
	gpPathCoverDepthFuncNV                                   uintptr
	gpPathDashArrayNV                                        uintptr
	gpPathGlyphIndexArrayNV                                  uintptr
	gpPathGlyphIndexRangeNV                                  uintptr
	gpPathGlyphRangeNV                                       uintptr
	gpPathGlyphsNV                                           uintptr
	gpPathMemoryGlyphIndexArrayNV                            uintptr
	gpPathParameterfNV                                       uintptr
	gpPathParameterfvNV                                      uintptr
	gpPathParameteriNV                                       uintptr
	gpPathParameterivNV                                      uintptr
	gpPathStencilDepthOffsetNV                               uintptr
	gpPathStencilFuncNV                                      uintptr
	gpPathStringNV                                           uintptr
	gpPathSubCommandsNV                                      uintptr
	gpPathSubCoordsNV                                        uintptr
	gpPauseTransformFeedback                                 uintptr
	gpPauseTransformFeedbackNV                               uintptr
	gpPixelDataRangeNV                                       uintptr
	gpPixelMapfv                                             uintptr
	gpPixelMapuiv                                            uintptr
	gpPixelMapusv                                            uintptr
	gpPixelMapx                                              uintptr
	gpPixelStoref                                            uintptr
	gpPixelStorei                                            uintptr
	gpPixelStorex                                            uintptr
	gpPixelTexGenParameterfSGIS                              uintptr
	gpPixelTexGenParameterfvSGIS                             uintptr
	gpPixelTexGenParameteriSGIS                              uintptr
	gpPixelTexGenParameterivSGIS                             uintptr
	gpPixelTexGenSGIX                                        uintptr
	gpPixelTransferf                                         uintptr
	gpPixelTransferi                                         uintptr
	gpPixelTransferxOES                                      uintptr
	gpPixelTransformParameterfEXT                            uintptr
	gpPixelTransformParameterfvEXT                           uintptr
	gpPixelTransformParameteriEXT                            uintptr
	gpPixelTransformParameterivEXT                           uintptr
	gpPixelZoom                                              uintptr
	gpPixelZoomxOES                                          uintptr
	gpPointAlongPathNV                                       uintptr
	gpPointParameterf                                        uintptr
	gpPointParameterfARB                                     uintptr
	gpPointParameterfEXT                                     uintptr
	gpPointParameterfSGIS                                    uintptr
	gpPointParameterfv                                       uintptr
	gpPointParameterfvARB                                    uintptr
	gpPointParameterfvEXT                                    uintptr
	gpPointParameterfvSGIS                                   uintptr
	gpPointParameteri                                        uintptr
	gpPointParameteriNV                                      uintptr
	gpPointParameteriv                                       uintptr
	gpPointParameterivNV                                     uintptr
	gpPointParameterxOES                                     uintptr
	gpPointParameterxvOES                                    uintptr
	gpPointSize                                              uintptr
	gpPointSizexOES                                          uintptr
	gpPollAsyncSGIX                                          uintptr
	gpPollInstrumentsSGIX                                    uintptr
	gpPolygonMode                                            uintptr
	gpPolygonOffset                                          uintptr
	gpPolygonOffsetClamp                                     uintptr
	gpPolygonOffsetClampEXT                                  uintptr
	gpPolygonOffsetEXT                                       uintptr
	gpPolygonOffsetxOES                                      uintptr
	gpPolygonStipple                                         uintptr
	gpPopAttrib                                              uintptr
	gpPopClientAttrib                                        uintptr
	gpPopDebugGroup                                          uintptr
	gpPopDebugGroupKHR                                       uintptr
	gpPopGroupMarkerEXT                                      uintptr
	gpPopMatrix                                              uintptr
	gpPopName                                                uintptr
	gpPresentFrameDualFillNV                                 uintptr
	gpPresentFrameKeyedNV                                    uintptr
	gpPrimitiveBoundingBoxARB                                uintptr
	gpPrimitiveRestartIndexNV                                uintptr
	gpPrimitiveRestartNV                                     uintptr
	gpPrioritizeTextures                                     uintptr
	gpPrioritizeTexturesEXT                                  uintptr
	gpPrioritizeTexturesxOES                                 uintptr
	gpProgramBinary                                          uintptr
	gpProgramBufferParametersIivNV                           uintptr
	gpProgramBufferParametersIuivNV                          uintptr
	gpProgramBufferParametersfvNV                            uintptr
	gpProgramEnvParameter4dARB                               uintptr
	gpProgramEnvParameter4dvARB                              uintptr
	gpProgramEnvParameter4fARB                               uintptr
	gpProgramEnvParameter4fvARB                              uintptr
	gpProgramEnvParameterI4iNV                               uintptr
	gpProgramEnvParameterI4ivNV                              uintptr
	gpProgramEnvParameterI4uiNV                              uintptr
	gpProgramEnvParameterI4uivNV                             uintptr
	gpProgramEnvParameters4fvEXT                             uintptr
	gpProgramEnvParametersI4ivNV                             uintptr
	gpProgramEnvParametersI4uivNV                            uintptr
	gpProgramLocalParameter4dARB                             uintptr
	gpProgramLocalParameter4dvARB                            uintptr
	gpProgramLocalParameter4fARB                             uintptr
	gpProgramLocalParameter4fvARB                            uintptr
	gpProgramLocalParameterI4iNV                             uintptr
	gpProgramLocalParameterI4ivNV                            uintptr
	gpProgramLocalParameterI4uiNV                            uintptr
	gpProgramLocalParameterI4uivNV                           uintptr
	gpProgramLocalParameters4fvEXT                           uintptr
	gpProgramLocalParametersI4ivNV                           uintptr
	gpProgramLocalParametersI4uivNV                          uintptr
	gpProgramNamedParameter4dNV                              uintptr
	gpProgramNamedParameter4dvNV                             uintptr
	gpProgramNamedParameter4fNV                              uintptr
	gpProgramNamedParameter4fvNV                             uintptr
	gpProgramParameter4dNV                                   uintptr
	gpProgramParameter4dvNV                                  uintptr
	gpProgramParameter4fNV                                   uintptr
	gpProgramParameter4fvNV                                  uintptr
	gpProgramParameteri                                      uintptr
	gpProgramParameteriARB                                   uintptr
	gpProgramParameteriEXT                                   uintptr
	gpProgramParameters4dvNV                                 uintptr
	gpProgramParameters4fvNV                                 uintptr
	gpProgramPathFragmentInputGenNV                          uintptr
	gpProgramStringARB                                       uintptr
	gpProgramSubroutineParametersuivNV                       uintptr
	gpProgramUniform1d                                       uintptr
	gpProgramUniform1dEXT                                    uintptr
	gpProgramUniform1dv                                      uintptr
	gpProgramUniform1dvEXT                                   uintptr
	gpProgramUniform1f                                       uintptr
	gpProgramUniform1fEXT                                    uintptr
	gpProgramUniform1fv                                      uintptr
	gpProgramUniform1fvEXT                                   uintptr
	gpProgramUniform1i                                       uintptr
	gpProgramUniform1i64ARB                                  uintptr
	gpProgramUniform1i64NV                                   uintptr
	gpProgramUniform1i64vARB                                 uintptr
	gpProgramUniform1i64vNV                                  uintptr
	gpProgramUniform1iEXT                                    uintptr
	gpProgramUniform1iv                                      uintptr
	gpProgramUniform1ivEXT                                   uintptr
	gpProgramUniform1ui                                      uintptr
	gpProgramUniform1ui64ARB                                 uintptr
	gpProgramUniform1ui64NV                                  uintptr
	gpProgramUniform1ui64vARB                                uintptr
	gpProgramUniform1ui64vNV                                 uintptr
	gpProgramUniform1uiEXT                                   uintptr
	gpProgramUniform1uiv                                     uintptr
	gpProgramUniform1uivEXT                                  uintptr
	gpProgramUniform2d                                       uintptr
	gpProgramUniform2dEXT                                    uintptr
	gpProgramUniform2dv                                      uintptr
	gpProgramUniform2dvEXT                                   uintptr
	gpProgramUniform2f                                       uintptr
	gpProgramUniform2fEXT                                    uintptr
	gpProgramUniform2fv                                      uintptr
	gpProgramUniform2fvEXT                                   uintptr
	gpProgramUniform2i                                       uintptr
	gpProgramUniform2i64ARB                                  uintptr
	gpProgramUniform2i64NV                                   uintptr
	gpProgramUniform2i64vARB                                 uintptr
	gpProgramUniform2i64vNV                                  uintptr
	gpProgramUniform2iEXT                                    uintptr
	gpProgramUniform2iv                                      uintptr
	gpProgramUniform2ivEXT                                   uintptr
	gpProgramUniform2ui                                      uintptr
	gpProgramUniform2ui64ARB                                 uintptr
	gpProgramUniform2ui64NV                                  uintptr
	gpProgramUniform2ui64vARB                                uintptr
	gpProgramUniform2ui64vNV                                 uintptr
	gpProgramUniform2uiEXT                                   uintptr
	gpProgramUniform2uiv                                     uintptr
	gpProgramUniform2uivEXT                                  uintptr
	gpProgramUniform3d                                       uintptr
	gpProgramUniform3dEXT                                    uintptr
	gpProgramUniform3dv                                      uintptr
	gpProgramUniform3dvEXT                                   uintptr
	gpProgramUniform3f                                       uintptr
	gpProgramUniform3fEXT                                    uintptr
	gpProgramUniform3fv                                      uintptr
	gpProgramUniform3fvEXT                                   uintptr
	gpProgramUniform3i                                       uintptr
	gpProgramUniform3i64ARB                                  uintptr
	gpProgramUniform3i64NV                                   uintptr
	gpProgramUniform3i64vARB                                 uintptr
	gpProgramUniform3i64vNV                                  uintptr
	gpProgramUniform3iEXT                                    uintptr
	gpProgramUniform3iv                                      uintptr
	gpProgramUniform3ivEXT                                   uintptr
	gpProgramUniform3ui                                      uintptr
	gpProgramUniform3ui64ARB                                 uintptr
	gpProgramUniform3ui64NV                                  uintptr
	gpProgramUniform3ui64vARB                                uintptr
	gpProgramUniform3ui64vNV                                 uintptr
	gpProgramUniform3uiEXT                                   uintptr
	gpProgramUniform3uiv                                     uintptr
	gpProgramUniform3uivEXT                                  uintptr
	gpProgramUniform4d                                       uintptr
	gpProgramUniform4dEXT                                    uintptr
	gpProgramUniform4dv                                      uintptr
	gpProgramUniform4dvEXT                                   uintptr
	gpProgramUniform4f                                       uintptr
	gpProgramUniform4fEXT                                    uintptr
	gpProgramUniform4fv                                      uintptr
	gpProgramUniform4fvEXT                                   uintptr
	gpProgramUniform4i                                       uintptr
	gpProgramUniform4i64ARB                                  uintptr
	gpProgramUniform4i64NV                                   uintptr
	gpProgramUniform4i64vARB                                 uintptr
	gpProgramUniform4i64vNV                                  uintptr
	gpProgramUniform4iEXT                                    uintptr
	gpProgramUniform4iv                                      uintptr
	gpProgramUniform4ivEXT                                   uintptr
	gpProgramUniform4ui                                      uintptr
	gpProgramUniform4ui64ARB                                 uintptr
	gpProgramUniform4ui64NV                                  uintptr
	gpProgramUniform4ui64vARB                                uintptr
	gpProgramUniform4ui64vNV                                 uintptr
	gpProgramUniform4uiEXT                                   uintptr
	gpProgramUniform4uiv                                     uintptr
	gpProgramUniform4uivEXT                                  uintptr
	gpProgramUniformHandleui64ARB                            uintptr
	gpProgramUniformHandleui64NV                             uintptr
	gpProgramUniformHandleui64vARB                           uintptr
	gpProgramUniformHandleui64vNV                            uintptr
	gpProgramUniformMatrix2dv                                uintptr
	gpProgramUniformMatrix2dvEXT                             uintptr
	gpProgramUniformMatrix2fv                                uintptr
	gpProgramUniformMatrix2fvEXT                             uintptr
	gpProgramUniformMatrix2x3dv                              uintptr
	gpProgramUniformMatrix2x3dvEXT                           uintptr
	gpProgramUniformMatrix2x3fv                              uintptr
	gpProgramUniformMatrix2x3fvEXT                           uintptr
	gpProgramUniformMatrix2x4dv                              uintptr
	gpProgramUniformMatrix2x4dvEXT                           uintptr
	gpProgramUniformMatrix2x4fv                              uintptr
	gpProgramUniformMatrix2x4fvEXT                           uintptr
	gpProgramUniformMatrix3dv                                uintptr
	gpProgramUniformMatrix3dvEXT                             uintptr
	gpProgramUniformMatrix3fv                                uintptr
	gpProgramUniformMatrix3fvEXT                             uintptr
	gpProgramUniformMatrix3x2dv                              uintptr
	gpProgramUniformMatrix3x2dvEXT                           uintptr
	gpProgramUniformMatrix3x2fv                              uintptr
	gpProgramUniformMatrix3x2fvEXT                           uintptr
	gpProgramUniformMatrix3x4dv                              uintptr
	gpProgramUniformMatrix3x4dvEXT                           uintptr
	gpProgramUniformMatrix3x4fv                              uintptr
	gpProgramUniformMatrix3x4fvEXT                           uintptr
	gpProgramUniformMatrix4dv                                uintptr
	gpProgramUniformMatrix4dvEXT                             uintptr
	gpProgramUniformMatrix4fv                                uintptr
	gpProgramUniformMatrix4fvEXT                             uintptr
	gpProgramUniformMatrix4x2dv                              uintptr
	gpProgramUniformMatrix4x2dvEXT                           uintptr
	gpProgramUniformMatrix4x2fv                              uintptr
	gpProgramUniformMatrix4x2fvEXT                           uintptr
	gpProgramUniformMatrix4x3dv                              uintptr
	gpProgramUniformMatrix4x3dvEXT                           uintptr
	gpProgramUniformMatrix4x3fv                              uintptr
	gpProgramUniformMatrix4x3fvEXT                           uintptr
	gpProgramUniformui64NV                                   uintptr
	gpProgramUniformui64vNV                                  uintptr
	gpProgramVertexLimitNV                                   uintptr
	gpProvokingVertex                                        uintptr
	gpProvokingVertexEXT                                     uintptr
	gpPushAttrib                                             uintptr
	gpPushClientAttrib                                       uintptr
	gpPushClientAttribDefaultEXT                             uintptr
	gpPushDebugGroup                                         uintptr
	gpPushDebugGroupKHR                                      uintptr
	gpPushGroupMarkerEXT                                     uintptr
	gpPushMatrix                                             uintptr
	gpPushName                                               uintptr
	gpQueryCounter                                           uintptr
	gpQueryMatrixxOES                                        uintptr
	gpQueryObjectParameteruiAMD                              uintptr
	gpQueryResourceNV                                        uintptr
	gpQueryResourceTagNV                                     uintptr
	gpRasterPos2d                                            uintptr
	gpRasterPos2dv                                           uintptr
	gpRasterPos2f                                            uintptr
	gpRasterPos2fv                                           uintptr
	gpRasterPos2i                                            uintptr
	gpRasterPos2iv                                           uintptr
	gpRasterPos2s                                            uintptr
	gpRasterPos2sv                                           uintptr
	gpRasterPos2xOES                                         uintptr
	gpRasterPos2xvOES                                        uintptr
	gpRasterPos3d                                            uintptr
	gpRasterPos3dv                                           uintptr
	gpRasterPos3f                                            uintptr
	gpRasterPos3fv                                           uintptr
	gpRasterPos3i                                            uintptr
	gpRasterPos3iv                                           uintptr
	gpRasterPos3s                                            uintptr
	gpRasterPos3sv                                           uintptr
	gpRasterPos3xOES                                         uintptr
	gpRasterPos3xvOES                                        uintptr
	gpRasterPos4d                                            uintptr
	gpRasterPos4dv                                           uintptr
	gpRasterPos4f                                            uintptr
	gpRasterPos4fv                                           uintptr
	gpRasterPos4i                                            uintptr
	gpRasterPos4iv                                           uintptr
	gpRasterPos4s                                            uintptr
	gpRasterPos4sv                                           uintptr
	gpRasterPos4xOES                                         uintptr
	gpRasterPos4xvOES                                        uintptr
	gpRasterSamplesEXT                                       uintptr
	gpReadBuffer                                             uintptr
	gpReadInstrumentsSGIX                                    uintptr
	gpReadPixels                                             uintptr
	gpReadnPixels                                            uintptr
	gpReadnPixelsARB                                         uintptr
	gpReadnPixelsKHR                                         uintptr
	gpRectd                                                  uintptr
	gpRectdv                                                 uintptr
	gpRectf                                                  uintptr
	gpRectfv                                                 uintptr
	gpRecti                                                  uintptr
	gpRectiv                                                 uintptr
	gpRects                                                  uintptr
	gpRectsv                                                 uintptr
	gpRectxOES                                               uintptr
	gpRectxvOES                                              uintptr
	gpReferencePlaneSGIX                                     uintptr
	gpReleaseKeyedMutexWin32EXT                              uintptr
	gpReleaseShaderCompiler                                  uintptr
	gpRenderGpuMaskNV                                        uintptr
	gpRenderMode                                             uintptr
	gpRenderbufferStorage                                    uintptr
	gpRenderbufferStorageEXT                                 uintptr
	gpRenderbufferStorageMultisample                         uintptr
	gpRenderbufferStorageMultisampleCoverageNV               uintptr
	gpRenderbufferStorageMultisampleEXT                      uintptr
	gpReplacementCodePointerSUN                              uintptr
	gpReplacementCodeubSUN                                   uintptr
	gpReplacementCodeubvSUN                                  uintptr
	gpReplacementCodeuiColor3fVertex3fSUN                    uintptr
	gpReplacementCodeuiColor3fVertex3fvSUN                   uintptr
	gpReplacementCodeuiColor4fNormal3fVertex3fSUN            uintptr
	gpReplacementCodeuiColor4fNormal3fVertex3fvSUN           uintptr
	gpReplacementCodeuiColor4ubVertex3fSUN                   uintptr
	gpReplacementCodeuiColor4ubVertex3fvSUN                  uintptr
	gpReplacementCodeuiNormal3fVertex3fSUN                   uintptr
	gpReplacementCodeuiNormal3fVertex3fvSUN                  uintptr
	gpReplacementCodeuiSUN                                   uintptr
	gpReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN  uintptr
	gpReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN uintptr
	gpReplacementCodeuiTexCoord2fNormal3fVertex3fSUN         uintptr
	gpReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN        uintptr
	gpReplacementCodeuiTexCoord2fVertex3fSUN                 uintptr
	gpReplacementCodeuiTexCoord2fVertex3fvSUN                uintptr
	gpReplacementCodeuiVertex3fSUN                           uintptr
	gpReplacementCodeuiVertex3fvSUN                          uintptr
	gpReplacementCodeuivSUN                                  uintptr
	gpReplacementCodeusSUN                                   uintptr
	gpReplacementCodeusvSUN                                  uintptr
	gpRequestResidentProgramsNV                              uintptr
	gpResetHistogramEXT                                      uintptr
	gpResetMinmaxEXT                                         uintptr
	gpResizeBuffersMESA                                      uintptr
	gpResolveDepthValuesNV                                   uintptr
	gpResumeTransformFeedback                                uintptr
	gpResumeTransformFeedbackNV                              uintptr
	gpRotated                                                uintptr
	gpRotatef                                                uintptr
	gpRotatexOES                                             uintptr
	gpSampleCoverage                                         uintptr
	gpSampleCoverageARB                                      uintptr
	gpSampleCoveragexOES                                     uintptr
	gpSampleMapATI                                           uintptr
	gpSampleMaskEXT                                          uintptr
	gpSampleMaskIndexedNV                                    uintptr
	gpSampleMaskSGIS                                         uintptr
	gpSampleMaski                                            uintptr
	gpSamplePatternEXT                                       uintptr
	gpSamplePatternSGIS                                      uintptr
	gpSamplerParameterIiv                                    uintptr
	gpSamplerParameterIuiv                                   uintptr
	gpSamplerParameterf                                      uintptr
	gpSamplerParameterfv                                     uintptr
	gpSamplerParameteri                                      uintptr
	gpSamplerParameteriv                                     uintptr
	gpScaled                                                 uintptr
	gpScalef                                                 uintptr
	gpScalexOES                                              uintptr
	gpScissor                                                uintptr
	gpScissorArrayv                                          uintptr
	gpScissorIndexed                                         uintptr
	gpScissorIndexedv                                        uintptr
	gpSecondaryColor3b                                       uintptr
	gpSecondaryColor3bEXT                                    uintptr
	gpSecondaryColor3bv                                      uintptr
	gpSecondaryColor3bvEXT                                   uintptr
	gpSecondaryColor3d                                       uintptr
	gpSecondaryColor3dEXT                                    uintptr
	gpSecondaryColor3dv                                      uintptr
	gpSecondaryColor3dvEXT                                   uintptr
	gpSecondaryColor3f                                       uintptr
	gpSecondaryColor3fEXT                                    uintptr
	gpSecondaryColor3fv                                      uintptr
	gpSecondaryColor3fvEXT                                   uintptr
	gpSecondaryColor3hNV                                     uintptr
	gpSecondaryColor3hvNV                                    uintptr
	gpSecondaryColor3i                                       uintptr
	gpSecondaryColor3iEXT                                    uintptr
	gpSecondaryColor3iv                                      uintptr
	gpSecondaryColor3ivEXT                                   uintptr
	gpSecondaryColor3s                                       uintptr
	gpSecondaryColor3sEXT                                    uintptr
	gpSecondaryColor3sv                                      uintptr
	gpSecondaryColor3svEXT                                   uintptr
	gpSecondaryColor3ub                                      uintptr
	gpSecondaryColor3ubEXT                                   uintptr
	gpSecondaryColor3ubv                                     uintptr
	gpSecondaryColor3ubvEXT                                  uintptr
	gpSecondaryColor3ui                                      uintptr
	gpSecondaryColor3uiEXT                                   uintptr
	gpSecondaryColor3uiv                                     uintptr
	gpSecondaryColor3uivEXT                                  uintptr
	gpSecondaryColor3us                                      uintptr
	gpSecondaryColor3usEXT                                   uintptr
	gpSecondaryColor3usv                                     uintptr
	gpSecondaryColor3usvEXT                                  uintptr
	gpSecondaryColorFormatNV                                 uintptr
	gpSecondaryColorPointer                                  uintptr
	gpSecondaryColorPointerEXT                               uintptr
	gpSecondaryColorPointerListIBM                           uintptr
	gpSelectBuffer                                           uintptr
	gpSelectPerfMonitorCountersAMD                           uintptr
	gpSemaphoreParameterui64vEXT                             uintptr
	gpSeparableFilter2DEXT                                   uintptr
	gpSetFenceAPPLE                                          uintptr
	gpSetFenceNV                                             uintptr
	gpSetFragmentShaderConstantATI                           uintptr
	gpSetInvariantEXT                                        uintptr
	gpSetLocalConstantEXT                                    uintptr
	gpSetMultisamplefvAMD                                    uintptr
	gpShadeModel                                             uintptr
	gpShaderBinary                                           uintptr
	gpShaderOp1EXT                                           uintptr
	gpShaderOp2EXT                                           uintptr
	gpShaderOp3EXT                                           uintptr
	gpShaderSource                                           uintptr
	gpShaderSourceARB                                        uintptr
	gpShaderStorageBlockBinding                              uintptr
	gpSharpenTexFuncSGIS                                     uintptr
	gpSignalSemaphoreEXT                                     uintptr
	gpSignalVkFenceNV                                        uintptr
	gpSignalVkSemaphoreNV                                    uintptr
	gpSpecializeShaderARB                                    uintptr
	gpSpriteParameterfSGIX                                   uintptr
	gpSpriteParameterfvSGIX                                  uintptr
	gpSpriteParameteriSGIX                                   uintptr
	gpSpriteParameterivSGIX                                  uintptr
	gpStartInstrumentsSGIX                                   uintptr
	gpStateCaptureNV                                         uintptr
	gpStencilClearTagEXT                                     uintptr
	gpStencilFillPathInstancedNV                             uintptr
	gpStencilFillPathNV                                      uintptr
	gpStencilFunc                                            uintptr
	gpStencilFuncSeparate                                    uintptr
	gpStencilFuncSeparateATI                                 uintptr
	gpStencilMask                                            uintptr
	gpStencilMaskSeparate                                    uintptr
	gpStencilOp                                              uintptr
	gpStencilOpSeparate                                      uintptr
	gpStencilOpSeparateATI                                   uintptr
	gpStencilOpValueAMD                                      uintptr
	gpStencilStrokePathInstancedNV                           uintptr
	gpStencilStrokePathNV                                    uintptr
	gpStencilThenCoverFillPathInstancedNV                    uintptr
	gpStencilThenCoverFillPathNV                             uintptr
	gpStencilThenCoverStrokePathInstancedNV                  uintptr
	gpStencilThenCoverStrokePathNV                           uintptr
	gpStopInstrumentsSGIX                                    uintptr
	gpStringMarkerGREMEDY                                    uintptr
	gpSubpixelPrecisionBiasNV                                uintptr
	gpSwizzleEXT                                             uintptr
	gpSyncTextureINTEL                                       uintptr
	gpTagSampleBufferSGIX                                    uintptr
	gpTangent3bEXT                                           uintptr
	gpTangent3bvEXT                                          uintptr
	gpTangent3dEXT                                           uintptr
	gpTangent3dvEXT                                          uintptr
	gpTangent3fEXT                                           uintptr
	gpTangent3fvEXT                                          uintptr
	gpTangent3iEXT                                           uintptr
	gpTangent3ivEXT                                          uintptr
	gpTangent3sEXT                                           uintptr
	gpTangent3svEXT                                          uintptr
	gpTangentPointerEXT                                      uintptr
	gpTbufferMask3DFX                                        uintptr
	gpTessellationFactorAMD                                  uintptr
	gpTessellationModeAMD                                    uintptr
	gpTestFenceAPPLE                                         uintptr
	gpTestFenceNV                                            uintptr
	gpTestObjectAPPLE                                        uintptr
	gpTexBufferARB                                           uintptr
	gpTexBufferEXT                                           uintptr
	gpTexBufferRange                                         uintptr
	gpTexBumpParameterfvATI                                  uintptr
	gpTexBumpParameterivATI                                  uintptr
	gpTexCoord1bOES                                          uintptr
	gpTexCoord1bvOES                                         uintptr
	gpTexCoord1d                                             uintptr
	gpTexCoord1dv                                            uintptr
	gpTexCoord1f                                             uintptr
	gpTexCoord1fv                                            uintptr
	gpTexCoord1hNV                                           uintptr
	gpTexCoord1hvNV                                          uintptr
	gpTexCoord1i                                             uintptr
	gpTexCoord1iv                                            uintptr
	gpTexCoord1s                                             uintptr
	gpTexCoord1sv                                            uintptr
	gpTexCoord1xOES                                          uintptr
	gpTexCoord1xvOES                                         uintptr
	gpTexCoord2bOES                                          uintptr
	gpTexCoord2bvOES                                         uintptr
	gpTexCoord2d                                             uintptr
	gpTexCoord2dv                                            uintptr
	gpTexCoord2f                                             uintptr
	gpTexCoord2fColor3fVertex3fSUN                           uintptr
	gpTexCoord2fColor3fVertex3fvSUN                          uintptr
	gpTexCoord2fColor4fNormal3fVertex3fSUN                   uintptr
	gpTexCoord2fColor4fNormal3fVertex3fvSUN                  uintptr
	gpTexCoord2fColor4ubVertex3fSUN                          uintptr
	gpTexCoord2fColor4ubVertex3fvSUN                         uintptr
	gpTexCoord2fNormal3fVertex3fSUN                          uintptr
	gpTexCoord2fNormal3fVertex3fvSUN                         uintptr
	gpTexCoord2fVertex3fSUN                                  uintptr
	gpTexCoord2fVertex3fvSUN                                 uintptr
	gpTexCoord2fv                                            uintptr
	gpTexCoord2hNV                                           uintptr
	gpTexCoord2hvNV                                          uintptr
	gpTexCoord2i                                             uintptr
	gpTexCoord2iv                                            uintptr
	gpTexCoord2s                                             uintptr
	gpTexCoord2sv                                            uintptr
	gpTexCoord2xOES                                          uintptr
	gpTexCoord2xvOES                                         uintptr
	gpTexCoord3bOES                                          uintptr
	gpTexCoord3bvOES                                         uintptr
	gpTexCoord3d                                             uintptr
	gpTexCoord3dv                                            uintptr
	gpTexCoord3f                                             uintptr
	gpTexCoord3fv                                            uintptr
	gpTexCoord3hNV                                           uintptr
	gpTexCoord3hvNV                                          uintptr
	gpTexCoord3i                                             uintptr
	gpTexCoord3iv                                            uintptr
	gpTexCoord3s                                             uintptr
	gpTexCoord3sv                                            uintptr
	gpTexCoord3xOES                                          uintptr
	gpTexCoord3xvOES                                         uintptr
	gpTexCoord4bOES                                          uintptr
	gpTexCoord4bvOES                                         uintptr
	gpTexCoord4d                                             uintptr
	gpTexCoord4dv                                            uintptr
	gpTexCoord4f                                             uintptr
	gpTexCoord4fColor4fNormal3fVertex4fSUN                   uintptr
	gpTexCoord4fColor4fNormal3fVertex4fvSUN                  uintptr
	gpTexCoord4fVertex4fSUN                                  uintptr
	gpTexCoord4fVertex4fvSUN                                 uintptr
	gpTexCoord4fv                                            uintptr
	gpTexCoord4hNV                                           uintptr
	gpTexCoord4hvNV                                          uintptr
	gpTexCoord4i                                             uintptr
	gpTexCoord4iv                                            uintptr
	gpTexCoord4s                                             uintptr
	gpTexCoord4sv                                            uintptr
	gpTexCoord4xOES                                          uintptr
	gpTexCoord4xvOES                                         uintptr
	gpTexCoordFormatNV                                       uintptr
	gpTexCoordPointer                                        uintptr
	gpTexCoordPointerEXT                                     uintptr
	gpTexCoordPointerListIBM                                 uintptr
	gpTexCoordPointervINTEL                                  uintptr
	gpTexEnvf                                                uintptr
	gpTexEnvfv                                               uintptr
	gpTexEnvi                                                uintptr
	gpTexEnviv                                               uintptr
	gpTexEnvxOES                                             uintptr
	gpTexEnvxvOES                                            uintptr
	gpTexFilterFuncSGIS                                      uintptr
	gpTexGend                                                uintptr
	gpTexGendv                                               uintptr
	gpTexGenf                                                uintptr
	gpTexGenfv                                               uintptr
	gpTexGeni                                                uintptr
	gpTexGeniv                                               uintptr
	gpTexGenxOES                                             uintptr
	gpTexGenxvOES                                            uintptr
	gpTexImage1D                                             uintptr
	gpTexImage2D                                             uintptr
	gpTexImage2DMultisample                                  uintptr
	gpTexImage2DMultisampleCoverageNV                        uintptr
	gpTexImage3D                                             uintptr
	gpTexImage3DEXT                                          uintptr
	gpTexImage3DMultisample                                  uintptr
	gpTexImage3DMultisampleCoverageNV                        uintptr
	gpTexImage4DSGIS                                         uintptr
	gpTexPageCommitmentARB                                   uintptr
	gpTexParameterIivEXT                                     uintptr
	gpTexParameterIuivEXT                                    uintptr
	gpTexParameterf                                          uintptr
	gpTexParameterfv                                         uintptr
	gpTexParameteri                                          uintptr
	gpTexParameteriv                                         uintptr
	gpTexParameterxOES                                       uintptr
	gpTexParameterxvOES                                      uintptr
	gpTexRenderbufferNV                                      uintptr
	gpTexStorage1D                                           uintptr
	gpTexStorage2D                                           uintptr
	gpTexStorage2DMultisample                                uintptr
	gpTexStorage3D                                           uintptr
	gpTexStorage3DMultisample                                uintptr
	gpTexStorageMem1DEXT                                     uintptr
	gpTexStorageMem2DEXT                                     uintptr
	gpTexStorageMem2DMultisampleEXT                          uintptr
	gpTexStorageMem3DEXT                                     uintptr
	gpTexStorageMem3DMultisampleEXT                          uintptr
	gpTexStorageSparseAMD                                    uintptr
	gpTexSubImage1D                                          uintptr
	gpTexSubImage1DEXT                                       uintptr
	gpTexSubImage2D                                          uintptr
	gpTexSubImage2DEXT                                       uintptr
	gpTexSubImage3D                                          uintptr
	gpTexSubImage3DEXT                                       uintptr
	gpTexSubImage4DSGIS                                      uintptr
	gpTextureBarrier                                         uintptr
	gpTextureBarrierNV                                       uintptr
	gpTextureBuffer                                          uintptr
	gpTextureBufferEXT                                       uintptr
	gpTextureBufferRange                                     uintptr
	gpTextureBufferRangeEXT                                  uintptr
	gpTextureColorMaskSGIS                                   uintptr
	gpTextureImage1DEXT                                      uintptr
	gpTextureImage2DEXT                                      uintptr
	gpTextureImage2DMultisampleCoverageNV                    uintptr
	gpTextureImage2DMultisampleNV                            uintptr
	gpTextureImage3DEXT                                      uintptr
	gpTextureImage3DMultisampleCoverageNV                    uintptr
	gpTextureImage3DMultisampleNV                            uintptr
	gpTextureLightEXT                                        uintptr
	gpTextureMaterialEXT                                     uintptr
	gpTextureNormalEXT                                       uintptr
	gpTexturePageCommitmentEXT                               uintptr
	gpTextureParameterIiv                                    uintptr
	gpTextureParameterIivEXT                                 uintptr
	gpTextureParameterIuiv                                   uintptr
	gpTextureParameterIuivEXT                                uintptr
	gpTextureParameterf                                      uintptr
	gpTextureParameterfEXT                                   uintptr
	gpTextureParameterfv                                     uintptr
	gpTextureParameterfvEXT                                  uintptr
	gpTextureParameteri                                      uintptr
	gpTextureParameteriEXT                                   uintptr
	gpTextureParameteriv                                     uintptr
	gpTextureParameterivEXT                                  uintptr
	gpTextureRangeAPPLE                                      uintptr
	gpTextureRenderbufferEXT                                 uintptr
	gpTextureStorage1D                                       uintptr
	gpTextureStorage1DEXT                                    uintptr
	gpTextureStorage2D                                       uintptr
	gpTextureStorage2DEXT                                    uintptr
	gpTextureStorage2DMultisample                            uintptr
	gpTextureStorage2DMultisampleEXT                         uintptr
	gpTextureStorage3D                                       uintptr
	gpTextureStorage3DEXT                                    uintptr
	gpTextureStorage3DMultisample                            uintptr
	gpTextureStorage3DMultisampleEXT                         uintptr
	gpTextureStorageMem1DEXT                                 uintptr
	gpTextureStorageMem2DEXT                                 uintptr
	gpTextureStorageMem2DMultisampleEXT                      uintptr
	gpTextureStorageMem3DEXT                                 uintptr
	gpTextureStorageMem3DMultisampleEXT                      uintptr
	gpTextureStorageSparseAMD                                uintptr
	gpTextureSubImage1D                                      uintptr
	gpTextureSubImage1DEXT                                   uintptr
	gpTextureSubImage2D                                      uintptr
	gpTextureSubImage2DEXT                                   uintptr
	gpTextureSubImage3D                                      uintptr
	gpTextureSubImage3DEXT                                   uintptr
	gpTextureView                                            uintptr
	gpTrackMatrixNV                                          uintptr
	gpTransformFeedbackAttribsNV                             uintptr
	gpTransformFeedbackBufferBase                            uintptr
	gpTransformFeedbackBufferRange                           uintptr
	gpTransformFeedbackStreamAttribsNV                       uintptr
	gpTransformFeedbackVaryingsEXT                           uintptr
	gpTransformFeedbackVaryingsNV                            uintptr
	gpTransformPathNV                                        uintptr
	gpTranslated                                             uintptr
	gpTranslatef                                             uintptr
	gpTranslatexOES                                          uintptr
	gpUniform1d                                              uintptr
	gpUniform1dv                                             uintptr
	gpUniform1f                                              uintptr
	gpUniform1fARB                                           uintptr
	gpUniform1fv                                             uintptr
	gpUniform1fvARB                                          uintptr
	gpUniform1i                                              uintptr
	gpUniform1i64ARB                                         uintptr
	gpUniform1i64NV                                          uintptr
	gpUniform1i64vARB                                        uintptr
	gpUniform1i64vNV                                         uintptr
	gpUniform1iARB                                           uintptr
	gpUniform1iv                                             uintptr
	gpUniform1ivARB                                          uintptr
	gpUniform1ui64ARB                                        uintptr
	gpUniform1ui64NV                                         uintptr
	gpUniform1ui64vARB                                       uintptr
	gpUniform1ui64vNV                                        uintptr
	gpUniform1uiEXT                                          uintptr
	gpUniform1uivEXT                                         uintptr
	gpUniform2d                                              uintptr
	gpUniform2dv                                             uintptr
	gpUniform2f                                              uintptr
	gpUniform2fARB                                           uintptr
	gpUniform2fv                                             uintptr
	gpUniform2fvARB                                          uintptr
	gpUniform2i                                              uintptr
	gpUniform2i64ARB                                         uintptr
	gpUniform2i64NV                                          uintptr
	gpUniform2i64vARB                                        uintptr
	gpUniform2i64vNV                                         uintptr
	gpUniform2iARB                                           uintptr
	gpUniform2iv                                             uintptr
	gpUniform2ivARB                                          uintptr
	gpUniform2ui64ARB                                        uintptr
	gpUniform2ui64NV                                         uintptr
	gpUniform2ui64vARB                                       uintptr
	gpUniform2ui64vNV                                        uintptr
	gpUniform2uiEXT                                          uintptr
	gpUniform2uivEXT                                         uintptr
	gpUniform3d                                              uintptr
	gpUniform3dv                                             uintptr
	gpUniform3f                                              uintptr
	gpUniform3fARB                                           uintptr
	gpUniform3fv                                             uintptr
	gpUniform3fvARB                                          uintptr
	gpUniform3i                                              uintptr
	gpUniform3i64ARB                                         uintptr
	gpUniform3i64NV                                          uintptr
	gpUniform3i64vARB                                        uintptr
	gpUniform3i64vNV                                         uintptr
	gpUniform3iARB                                           uintptr
	gpUniform3iv                                             uintptr
	gpUniform3ivARB                                          uintptr
	gpUniform3ui64ARB                                        uintptr
	gpUniform3ui64NV                                         uintptr
	gpUniform3ui64vARB                                       uintptr
	gpUniform3ui64vNV                                        uintptr
	gpUniform3uiEXT                                          uintptr
	gpUniform3uivEXT                                         uintptr
	gpUniform4d                                              uintptr
	gpUniform4dv                                             uintptr
	gpUniform4f                                              uintptr
	gpUniform4fARB                                           uintptr
	gpUniform4fv                                             uintptr
	gpUniform4fvARB                                          uintptr
	gpUniform4i                                              uintptr
	gpUniform4i64ARB                                         uintptr
	gpUniform4i64NV                                          uintptr
	gpUniform4i64vARB                                        uintptr
	gpUniform4i64vNV                                         uintptr
	gpUniform4iARB                                           uintptr
	gpUniform4iv                                             uintptr
	gpUniform4ivARB                                          uintptr
	gpUniform4ui64ARB                                        uintptr
	gpUniform4ui64NV                                         uintptr
	gpUniform4ui64vARB                                       uintptr
	gpUniform4ui64vNV                                        uintptr
	gpUniform4uiEXT                                          uintptr
	gpUniform4uivEXT                                         uintptr
	gpUniformBlockBinding                                    uintptr
	gpUniformBufferEXT                                       uintptr
	gpUniformHandleui64ARB                                   uintptr
	gpUniformHandleui64NV                                    uintptr
	gpUniformHandleui64vARB                                  uintptr
	gpUniformHandleui64vNV                                   uintptr
	gpUniformMatrix2dv                                       uintptr
	gpUniformMatrix2fv                                       uintptr
	gpUniformMatrix2fvARB                                    uintptr
	gpUniformMatrix2x3dv                                     uintptr
	gpUniformMatrix2x3fv                                     uintptr
	gpUniformMatrix2x4dv                                     uintptr
	gpUniformMatrix2x4fv                                     uintptr
	gpUniformMatrix3dv                                       uintptr
	gpUniformMatrix3fv                                       uintptr
	gpUniformMatrix3fvARB                                    uintptr
	gpUniformMatrix3x2dv                                     uintptr
	gpUniformMatrix3x2fv                                     uintptr
	gpUniformMatrix3x4dv                                     uintptr
	gpUniformMatrix3x4fv                                     uintptr
	gpUniformMatrix4dv                                       uintptr
	gpUniformMatrix4fv                                       uintptr
	gpUniformMatrix4fvARB                                    uintptr
	gpUniformMatrix4x2dv                                     uintptr
	gpUniformMatrix4x2fv                                     uintptr
	gpUniformMatrix4x3dv                                     uintptr
	gpUniformMatrix4x3fv                                     uintptr
	gpUniformSubroutinesuiv                                  uintptr
	gpUniformui64NV                                          uintptr
	gpUniformui64vNV                                         uintptr
	gpUnlockArraysEXT                                        uintptr
	gpUnmapBuffer                                            uintptr
	gpUnmapBufferARB                                         uintptr
	gpUnmapNamedBuffer                                       uintptr
	gpUnmapNamedBufferEXT                                    uintptr
	gpUnmapObjectBufferATI                                   uintptr
	gpUnmapTexture2DINTEL                                    uintptr
	gpUpdateObjectBufferATI                                  uintptr
	gpUseProgram                                             uintptr
	gpUseProgramObjectARB                                    uintptr
	gpUseProgramStages                                       uintptr
	gpUseProgramStagesEXT                                    uintptr
	gpUseShaderProgramEXT                                    uintptr
	gpVDPAUFiniNV                                            uintptr
	gpVDPAUGetSurfaceivNV                                    uintptr
	gpVDPAUInitNV                                            uintptr
	gpVDPAUIsSurfaceNV                                       uintptr
	gpVDPAUMapSurfacesNV                                     uintptr
	gpVDPAURegisterOutputSurfaceNV                           uintptr
	gpVDPAURegisterVideoSurfaceNV                            uintptr
	gpVDPAUSurfaceAccessNV                                   uintptr
	gpVDPAUUnmapSurfacesNV                                   uintptr
	gpVDPAUUnregisterSurfaceNV                               uintptr
	gpValidateProgram                                        uintptr
	gpValidateProgramARB                                     uintptr
	gpValidateProgramPipeline                                uintptr
	gpValidateProgramPipelineEXT                             uintptr
	gpVariantArrayObjectATI                                  uintptr
	gpVariantPointerEXT                                      uintptr
	gpVariantbvEXT                                           uintptr
	gpVariantdvEXT                                           uintptr
	gpVariantfvEXT                                           uintptr
	gpVariantivEXT                                           uintptr
	gpVariantsvEXT                                           uintptr
	gpVariantubvEXT                                          uintptr
	gpVariantuivEXT                                          uintptr
	gpVariantusvEXT                                          uintptr
	gpVertex2bOES                                            uintptr
	gpVertex2bvOES                                           uintptr
	gpVertex2d                                               uintptr
	gpVertex2dv                                              uintptr
	gpVertex2f                                               uintptr
	gpVertex2fv                                              uintptr
	gpVertex2hNV                                             uintptr
	gpVertex2hvNV                                            uintptr
	gpVertex2i                                               uintptr
	gpVertex2iv                                              uintptr
	gpVertex2s                                               uintptr
	gpVertex2sv                                              uintptr
	gpVertex2xOES                                            uintptr
	gpVertex2xvOES                                           uintptr
	gpVertex3bOES                                            uintptr
	gpVertex3bvOES                                           uintptr
	gpVertex3d                                               uintptr
	gpVertex3dv                                              uintptr
	gpVertex3f                                               uintptr
	gpVertex3fv                                              uintptr
	gpVertex3hNV                                             uintptr
	gpVertex3hvNV                                            uintptr
	gpVertex3i                                               uintptr
	gpVertex3iv                                              uintptr
	gpVertex3s                                               uintptr
	gpVertex3sv                                              uintptr
	gpVertex3xOES                                            uintptr
	gpVertex3xvOES                                           uintptr
	gpVertex4bOES                                            uintptr
	gpVertex4bvOES                                           uintptr
	gpVertex4d                                               uintptr
	gpVertex4dv                                              uintptr
	gpVertex4f                                               uintptr
	gpVertex4fv                                              uintptr
	gpVertex4hNV                                             uintptr
	gpVertex4hvNV                                            uintptr
	gpVertex4i                                               uintptr
	gpVertex4iv                                              uintptr
	gpVertex4s                                               uintptr
	gpVertex4sv                                              uintptr
	gpVertex4xOES                                            uintptr
	gpVertex4xvOES                                           uintptr
	gpVertexArrayAttribBinding                               uintptr
	gpVertexArrayAttribFormat                                uintptr
	gpVertexArrayAttribIFormat                               uintptr
	gpVertexArrayAttribLFormat                               uintptr
	gpVertexArrayBindVertexBufferEXT                         uintptr
	gpVertexArrayBindingDivisor                              uintptr
	gpVertexArrayColorOffsetEXT                              uintptr
	gpVertexArrayEdgeFlagOffsetEXT                           uintptr
	gpVertexArrayElementBuffer                               uintptr
	gpVertexArrayFogCoordOffsetEXT                           uintptr
	gpVertexArrayIndexOffsetEXT                              uintptr
	gpVertexArrayMultiTexCoordOffsetEXT                      uintptr
	gpVertexArrayNormalOffsetEXT                             uintptr
	gpVertexArrayParameteriAPPLE                             uintptr
	gpVertexArrayRangeAPPLE                                  uintptr
	gpVertexArrayRangeNV                                     uintptr
	gpVertexArraySecondaryColorOffsetEXT                     uintptr
	gpVertexArrayTexCoordOffsetEXT                           uintptr
	gpVertexArrayVertexAttribBindingEXT                      uintptr
	gpVertexArrayVertexAttribDivisorEXT                      uintptr
	gpVertexArrayVertexAttribFormatEXT                       uintptr
	gpVertexArrayVertexAttribIFormatEXT                      uintptr
	gpVertexArrayVertexAttribIOffsetEXT                      uintptr
	gpVertexArrayVertexAttribLFormatEXT                      uintptr
	gpVertexArrayVertexAttribLOffsetEXT                      uintptr
	gpVertexArrayVertexAttribOffsetEXT                       uintptr
	gpVertexArrayVertexBindingDivisorEXT                     uintptr
	gpVertexArrayVertexBuffer                                uintptr
	gpVertexArrayVertexBuffers                               uintptr
	gpVertexArrayVertexOffsetEXT                             uintptr
	gpVertexAttrib1d                                         uintptr
	gpVertexAttrib1dARB                                      uintptr
	gpVertexAttrib1dNV                                       uintptr
	gpVertexAttrib1dv                                        uintptr
	gpVertexAttrib1dvARB                                     uintptr
	gpVertexAttrib1dvNV                                      uintptr
	gpVertexAttrib1f                                         uintptr
	gpVertexAttrib1fARB                                      uintptr
	gpVertexAttrib1fNV                                       uintptr
	gpVertexAttrib1fv                                        uintptr
	gpVertexAttrib1fvARB                                     uintptr
	gpVertexAttrib1fvNV                                      uintptr
	gpVertexAttrib1hNV                                       uintptr
	gpVertexAttrib1hvNV                                      uintptr
	gpVertexAttrib1s                                         uintptr
	gpVertexAttrib1sARB                                      uintptr
	gpVertexAttrib1sNV                                       uintptr
	gpVertexAttrib1sv                                        uintptr
	gpVertexAttrib1svARB                                     uintptr
	gpVertexAttrib1svNV                                      uintptr
	gpVertexAttrib2d                                         uintptr
	gpVertexAttrib2dARB                                      uintptr
	gpVertexAttrib2dNV                                       uintptr
	gpVertexAttrib2dv                                        uintptr
	gpVertexAttrib2dvARB                                     uintptr
	gpVertexAttrib2dvNV                                      uintptr
	gpVertexAttrib2f                                         uintptr
	gpVertexAttrib2fARB                                      uintptr
	gpVertexAttrib2fNV                                       uintptr
	gpVertexAttrib2fv                                        uintptr
	gpVertexAttrib2fvARB                                     uintptr
	gpVertexAttrib2fvNV                                      uintptr
	gpVertexAttrib2hNV                                       uintptr
	gpVertexAttrib2hvNV                                      uintptr
	gpVertexAttrib2s                                         uintptr
	gpVertexAttrib2sARB                                      uintptr
	gpVertexAttrib2sNV                                       uintptr
	gpVertexAttrib2sv                                        uintptr
	gpVertexAttrib2svARB                                     uintptr
	gpVertexAttrib2svNV                                      uintptr
	gpVertexAttrib3d                                         uintptr
	gpVertexAttrib3dARB                                      uintptr
	gpVertexAttrib3dNV                                       uintptr
	gpVertexAttrib3dv                                        uintptr
	gpVertexAttrib3dvARB                                     uintptr
	gpVertexAttrib3dvNV                                      uintptr
	gpVertexAttrib3f                                         uintptr
	gpVertexAttrib3fARB                                      uintptr
	gpVertexAttrib3fNV                                       uintptr
	gpVertexAttrib3fv                                        uintptr
	gpVertexAttrib3fvARB                                     uintptr
	gpVertexAttrib3fvNV                                      uintptr
	gpVertexAttrib3hNV                                       uintptr
	gpVertexAttrib3hvNV                                      uintptr
	gpVertexAttrib3s                                         uintptr
	gpVertexAttrib3sARB                                      uintptr
	gpVertexAttrib3sNV                                       uintptr
	gpVertexAttrib3sv                                        uintptr
	gpVertexAttrib3svARB                                     uintptr
	gpVertexAttrib3svNV                                      uintptr
	gpVertexAttrib4Nbv                                       uintptr
	gpVertexAttrib4NbvARB                                    uintptr
	gpVertexAttrib4Niv                                       uintptr
	gpVertexAttrib4NivARB                                    uintptr
	gpVertexAttrib4Nsv                                       uintptr
	gpVertexAttrib4NsvARB                                    uintptr
	gpVertexAttrib4Nub                                       uintptr
	gpVertexAttrib4NubARB                                    uintptr
	gpVertexAttrib4Nubv                                      uintptr
	gpVertexAttrib4NubvARB                                   uintptr
	gpVertexAttrib4Nuiv                                      uintptr
	gpVertexAttrib4NuivARB                                   uintptr
	gpVertexAttrib4Nusv                                      uintptr
	gpVertexAttrib4NusvARB                                   uintptr
	gpVertexAttrib4bv                                        uintptr
	gpVertexAttrib4bvARB                                     uintptr
	gpVertexAttrib4d                                         uintptr
	gpVertexAttrib4dARB                                      uintptr
	gpVertexAttrib4dNV                                       uintptr
	gpVertexAttrib4dv                                        uintptr
	gpVertexAttrib4dvARB                                     uintptr
	gpVertexAttrib4dvNV                                      uintptr
	gpVertexAttrib4f                                         uintptr
	gpVertexAttrib4fARB                                      uintptr
	gpVertexAttrib4fNV                                       uintptr
	gpVertexAttrib4fv                                        uintptr
	gpVertexAttrib4fvARB                                     uintptr
	gpVertexAttrib4fvNV                                      uintptr
	gpVertexAttrib4hNV                                       uintptr
	gpVertexAttrib4hvNV                                      uintptr
	gpVertexAttrib4iv                                        uintptr
	gpVertexAttrib4ivARB                                     uintptr
	gpVertexAttrib4s                                         uintptr
	gpVertexAttrib4sARB                                      uintptr
	gpVertexAttrib4sNV                                       uintptr
	gpVertexAttrib4sv                                        uintptr
	gpVertexAttrib4svARB                                     uintptr
	gpVertexAttrib4svNV                                      uintptr
	gpVertexAttrib4ubNV                                      uintptr
	gpVertexAttrib4ubv                                       uintptr
	gpVertexAttrib4ubvARB                                    uintptr
	gpVertexAttrib4ubvNV                                     uintptr
	gpVertexAttrib4uiv                                       uintptr
	gpVertexAttrib4uivARB                                    uintptr
	gpVertexAttrib4usv                                       uintptr
	gpVertexAttrib4usvARB                                    uintptr
	gpVertexAttribArrayObjectATI                             uintptr
	gpVertexAttribBinding                                    uintptr
	gpVertexAttribDivisorARB                                 uintptr
	gpVertexAttribFormat                                     uintptr
	gpVertexAttribFormatNV                                   uintptr
	gpVertexAttribI1iEXT                                     uintptr
	gpVertexAttribI1ivEXT                                    uintptr
	gpVertexAttribI1uiEXT                                    uintptr
	gpVertexAttribI1uivEXT                                   uintptr
	gpVertexAttribI2iEXT                                     uintptr
	gpVertexAttribI2ivEXT                                    uintptr
	gpVertexAttribI2uiEXT                                    uintptr
	gpVertexAttribI2uivEXT                                   uintptr
	gpVertexAttribI3iEXT                                     uintptr
	gpVertexAttribI3ivEXT                                    uintptr
	gpVertexAttribI3uiEXT                                    uintptr
	gpVertexAttribI3uivEXT                                   uintptr
	gpVertexAttribI4bvEXT                                    uintptr
	gpVertexAttribI4iEXT                                     uintptr
	gpVertexAttribI4ivEXT                                    uintptr
	gpVertexAttribI4svEXT                                    uintptr
	gpVertexAttribI4ubvEXT                                   uintptr
	gpVertexAttribI4uiEXT                                    uintptr
	gpVertexAttribI4uivEXT                                   uintptr
	gpVertexAttribI4usvEXT                                   uintptr
	gpVertexAttribIFormat                                    uintptr
	gpVertexAttribIFormatNV                                  uintptr
	gpVertexAttribIPointerEXT                                uintptr
	gpVertexAttribL1d                                        uintptr
	gpVertexAttribL1dEXT                                     uintptr
	gpVertexAttribL1dv                                       uintptr
	gpVertexAttribL1dvEXT                                    uintptr
	gpVertexAttribL1i64NV                                    uintptr
	gpVertexAttribL1i64vNV                                   uintptr
	gpVertexAttribL1ui64ARB                                  uintptr
	gpVertexAttribL1ui64NV                                   uintptr
	gpVertexAttribL1ui64vARB                                 uintptr
	gpVertexAttribL1ui64vNV                                  uintptr
	gpVertexAttribL2d                                        uintptr
	gpVertexAttribL2dEXT                                     uintptr
	gpVertexAttribL2dv                                       uintptr
	gpVertexAttribL2dvEXT                                    uintptr
	gpVertexAttribL2i64NV                                    uintptr
	gpVertexAttribL2i64vNV                                   uintptr
	gpVertexAttribL2ui64NV                                   uintptr
	gpVertexAttribL2ui64vNV                                  uintptr
	gpVertexAttribL3d                                        uintptr
	gpVertexAttribL3dEXT                                     uintptr
	gpVertexAttribL3dv                                       uintptr
	gpVertexAttribL3dvEXT                                    uintptr
	gpVertexAttribL3i64NV                                    uintptr
	gpVertexAttribL3i64vNV                                   uintptr
	gpVertexAttribL3ui64NV                                   uintptr
	gpVertexAttribL3ui64vNV                                  uintptr
	gpVertexAttribL4d                                        uintptr
	gpVertexAttribL4dEXT                                     uintptr
	gpVertexAttribL4dv                                       uintptr
	gpVertexAttribL4dvEXT                                    uintptr
	gpVertexAttribL4i64NV                                    uintptr
	gpVertexAttribL4i64vNV                                   uintptr
	gpVertexAttribL4ui64NV                                   uintptr
	gpVertexAttribL4ui64vNV                                  uintptr
	gpVertexAttribLFormat                                    uintptr
	gpVertexAttribLFormatNV                                  uintptr
	gpVertexAttribLPointer                                   uintptr
	gpVertexAttribLPointerEXT                                uintptr
	gpVertexAttribP1ui                                       uintptr
	gpVertexAttribP1uiv                                      uintptr
	gpVertexAttribP2ui                                       uintptr
	gpVertexAttribP2uiv                                      uintptr
	gpVertexAttribP3ui                                       uintptr
	gpVertexAttribP3uiv                                      uintptr
	gpVertexAttribP4ui                                       uintptr
	gpVertexAttribP4uiv                                      uintptr
	gpVertexAttribParameteriAMD                              uintptr
	gpVertexAttribPointer                                    uintptr
	gpVertexAttribPointerARB                                 uintptr
	gpVertexAttribPointerNV                                  uintptr
	gpVertexAttribs1dvNV                                     uintptr
	gpVertexAttribs1fvNV                                     uintptr
	gpVertexAttribs1hvNV                                     uintptr
	gpVertexAttribs1svNV                                     uintptr
	gpVertexAttribs2dvNV                                     uintptr
	gpVertexAttribs2fvNV                                     uintptr
	gpVertexAttribs2hvNV                                     uintptr
	gpVertexAttribs2svNV                                     uintptr
	gpVertexAttribs3dvNV                                     uintptr
	gpVertexAttribs3fvNV                                     uintptr
	gpVertexAttribs3hvNV                                     uintptr
	gpVertexAttribs3svNV                                     uintptr
	gpVertexAttribs4dvNV                                     uintptr
	gpVertexAttribs4fvNV                                     uintptr
	gpVertexAttribs4hvNV                                     uintptr
	gpVertexAttribs4svNV                                     uintptr
	gpVertexAttribs4ubvNV                                    uintptr
	gpVertexBindingDivisor                                   uintptr
	gpVertexBlendARB                                         uintptr
	gpVertexBlendEnvfATI                                     uintptr
	gpVertexBlendEnviATI                                     uintptr
	gpVertexFormatNV                                         uintptr
	gpVertexPointer                                          uintptr
	gpVertexPointerEXT                                       uintptr
	gpVertexPointerListIBM                                   uintptr
	gpVertexPointervINTEL                                    uintptr
	gpVertexStream1dATI                                      uintptr
	gpVertexStream1dvATI                                     uintptr
	gpVertexStream1fATI                                      uintptr
	gpVertexStream1fvATI                                     uintptr
	gpVertexStream1iATI                                      uintptr
	gpVertexStream1ivATI                                     uintptr
	gpVertexStream1sATI                                      uintptr
	gpVertexStream1svATI                                     uintptr
	gpVertexStream2dATI                                      uintptr
	gpVertexStream2dvATI                                     uintptr
	gpVertexStream2fATI                                      uintptr
	gpVertexStream2fvATI                                     uintptr
	gpVertexStream2iATI                                      uintptr
	gpVertexStream2ivATI                                     uintptr
	gpVertexStream2sATI                                      uintptr
	gpVertexStream2svATI                                     uintptr
	gpVertexStream3dATI                                      uintptr
	gpVertexStream3dvATI                                     uintptr
	gpVertexStream3fATI                                      uintptr
	gpVertexStream3fvATI                                     uintptr
	gpVertexStream3iATI                                      uintptr
	gpVertexStream3ivATI                                     uintptr
	gpVertexStream3sATI                                      uintptr
	gpVertexStream3svATI                                     uintptr
	gpVertexStream4dATI                                      uintptr
	gpVertexStream4dvATI                                     uintptr
	gpVertexStream4fATI                                      uintptr
	gpVertexStream4fvATI                                     uintptr
	gpVertexStream4iATI                                      uintptr
	gpVertexStream4ivATI                                     uintptr
	gpVertexStream4sATI                                      uintptr
	gpVertexStream4svATI                                     uintptr
	gpVertexWeightPointerEXT                                 uintptr
	gpVertexWeightfEXT                                       uintptr
	gpVertexWeightfvEXT                                      uintptr
	gpVertexWeighthNV                                        uintptr
	gpVertexWeighthvNV                                       uintptr
	gpVideoCaptureNV                                         uintptr
	gpVideoCaptureStreamParameterdvNV                        uintptr
	gpVideoCaptureStreamParameterfvNV                        uintptr
	gpVideoCaptureStreamParameterivNV                        uintptr
	gpViewport                                               uintptr
	gpViewportArrayv                                         uintptr
	gpViewportIndexedf                                       uintptr
	gpViewportIndexedfv                                      uintptr
	gpViewportPositionWScaleNV                               uintptr
	gpViewportSwizzleNV                                      uintptr
	gpWaitSemaphoreEXT                                       uintptr
	gpWaitSync                                               uintptr
	gpWaitVkSemaphoreNV                                      uintptr
	gpWeightPathsNV                                          uintptr
	gpWeightPointerARB                                       uintptr
	gpWeightbvARB                                            uintptr
	gpWeightdvARB                                            uintptr
	gpWeightfvARB                                            uintptr
	gpWeightivARB                                            uintptr
	gpWeightsvARB                                            uintptr
	gpWeightubvARB                                           uintptr
	gpWeightuivARB                                           uintptr
	gpWeightusvARB                                           uintptr
	gpWindowPos2d                                            uintptr
	gpWindowPos2dARB                                         uintptr
	gpWindowPos2dMESA                                        uintptr
	gpWindowPos2dv                                           uintptr
	gpWindowPos2dvARB                                        uintptr
	gpWindowPos2dvMESA                                       uintptr
	gpWindowPos2f                                            uintptr
	gpWindowPos2fARB                                         uintptr
	gpWindowPos2fMESA                                        uintptr
	gpWindowPos2fv                                           uintptr
	gpWindowPos2fvARB                                        uintptr
	gpWindowPos2fvMESA                                       uintptr
	gpWindowPos2i                                            uintptr
	gpWindowPos2iARB                                         uintptr
	gpWindowPos2iMESA                                        uintptr
	gpWindowPos2iv                                           uintptr
	gpWindowPos2ivARB                                        uintptr
	gpWindowPos2ivMESA                                       uintptr
	gpWindowPos2s                                            uintptr
	gpWindowPos2sARB                                         uintptr
	gpWindowPos2sMESA                                        uintptr
	gpWindowPos2sv                                           uintptr
	gpWindowPos2svARB                                        uintptr
	gpWindowPos2svMESA                                       uintptr
	gpWindowPos3d                                            uintptr
	gpWindowPos3dARB                                         uintptr
	gpWindowPos3dMESA                                        uintptr
	gpWindowPos3dv                                           uintptr
	gpWindowPos3dvARB                                        uintptr
	gpWindowPos3dvMESA                                       uintptr
	gpWindowPos3f                                            uintptr
	gpWindowPos3fARB                                         uintptr
	gpWindowPos3fMESA                                        uintptr
	gpWindowPos3fv                                           uintptr
	gpWindowPos3fvARB                                        uintptr
	gpWindowPos3fvMESA                                       uintptr
	gpWindowPos3i                                            uintptr
	gpWindowPos3iARB                                         uintptr
	gpWindowPos3iMESA                                        uintptr
	gpWindowPos3iv                                           uintptr
	gpWindowPos3ivARB                                        uintptr
	gpWindowPos3ivMESA                                       uintptr
	gpWindowPos3s                                            uintptr
	gpWindowPos3sARB                                         uintptr
	gpWindowPos3sMESA                                        uintptr
	gpWindowPos3sv                                           uintptr
	gpWindowPos3svARB                                        uintptr
	gpWindowPos3svMESA                                       uintptr
	gpWindowPos4dMESA                                        uintptr
	gpWindowPos4dvMESA                                       uintptr
	gpWindowPos4fMESA                                        uintptr
	gpWindowPos4fvMESA                                       uintptr
	gpWindowPos4iMESA                                        uintptr
	gpWindowPos4ivMESA                                       uintptr
	gpWindowPos4sMESA                                        uintptr
	gpWindowPos4svMESA                                       uintptr
	gpWindowRectanglesEXT                                    uintptr
	gpWriteMaskEXT                                           uintptr
)

func boolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// operate on the accumulation buffer
func Accum(op uint32, value float32) {
	syscall.Syscall(gpAccum, 2, uintptr(op), uintptr(math.Float32bits(value)), 0)
}
func AccumxOES(op uint32, value int32) {
	syscall.Syscall(gpAccumxOES, 2, uintptr(op), uintptr(value), 0)
}
func AcquireKeyedMutexWin32EXT(memory uint32, key uint64, timeout uint32) bool {
	ret, _, _ := syscall.Syscall(gpAcquireKeyedMutexWin32EXT, 3, uintptr(memory), uintptr(key), uintptr(timeout))
	return ret != 0
}
func ActiveProgramEXT(program uint32) {
	syscall.Syscall(gpActiveProgramEXT, 1, uintptr(program), 0, 0)
}

// set the active program object for a program pipeline object
func ActiveShaderProgram(pipeline uint32, program uint32) {
	syscall.Syscall(gpActiveShaderProgram, 2, uintptr(pipeline), uintptr(program), 0)
}
func ActiveShaderProgramEXT(pipeline uint32, program uint32) {
	syscall.Syscall(gpActiveShaderProgramEXT, 2, uintptr(pipeline), uintptr(program), 0)
}
func ActiveStencilFaceEXT(face uint32) {
	syscall.Syscall(gpActiveStencilFaceEXT, 1, uintptr(face), 0, 0)
}

// select active texture unit
func ActiveTexture(texture uint32) {
	syscall.Syscall(gpActiveTexture, 1, uintptr(texture), 0, 0)
}
func ActiveTextureARB(texture uint32) {
	syscall.Syscall(gpActiveTextureARB, 1, uintptr(texture), 0, 0)
}
func ActiveVaryingNV(program uint32, name *uint8) {
	syscall.Syscall(gpActiveVaryingNV, 2, uintptr(program), uintptr(unsafe.Pointer(name)), 0)
}
func AlphaFragmentOp1ATI(op uint32, dst uint32, dstMod uint32, arg1 uint32, arg1Rep uint32, arg1Mod uint32) {
	syscall.Syscall6(gpAlphaFragmentOp1ATI, 6, uintptr(op), uintptr(dst), uintptr(dstMod), uintptr(arg1), uintptr(arg1Rep), uintptr(arg1Mod))
}
func AlphaFragmentOp2ATI(op uint32, dst uint32, dstMod uint32, arg1 uint32, arg1Rep uint32, arg1Mod uint32, arg2 uint32, arg2Rep uint32, arg2Mod uint32) {
	syscall.Syscall9(gpAlphaFragmentOp2ATI, 9, uintptr(op), uintptr(dst), uintptr(dstMod), uintptr(arg1), uintptr(arg1Rep), uintptr(arg1Mod), uintptr(arg2), uintptr(arg2Rep), uintptr(arg2Mod))
}
func AlphaFragmentOp3ATI(op uint32, dst uint32, dstMod uint32, arg1 uint32, arg1Rep uint32, arg1Mod uint32, arg2 uint32, arg2Rep uint32, arg2Mod uint32, arg3 uint32, arg3Rep uint32, arg3Mod uint32) {
	syscall.Syscall12(gpAlphaFragmentOp3ATI, 12, uintptr(op), uintptr(dst), uintptr(dstMod), uintptr(arg1), uintptr(arg1Rep), uintptr(arg1Mod), uintptr(arg2), uintptr(arg2Rep), uintptr(arg2Mod), uintptr(arg3), uintptr(arg3Rep), uintptr(arg3Mod))
}

// specify the alpha test function
func AlphaFunc(xfunc uint32, ref float32) {
	syscall.Syscall(gpAlphaFunc, 2, uintptr(xfunc), uintptr(math.Float32bits(ref)), 0)
}
func AlphaFuncxOES(xfunc uint32, ref int32) {
	syscall.Syscall(gpAlphaFuncxOES, 2, uintptr(xfunc), uintptr(ref), 0)
}
func AlphaToCoverageDitherControlNV(mode uint32) {
	syscall.Syscall(gpAlphaToCoverageDitherControlNV, 1, uintptr(mode), 0, 0)
}
func ApplyFramebufferAttachmentCMAAINTEL() {
	syscall.Syscall(gpApplyFramebufferAttachmentCMAAINTEL, 0, 0, 0, 0)
}
func ApplyTextureEXT(mode uint32) {
	syscall.Syscall(gpApplyTextureEXT, 1, uintptr(mode), 0, 0)
}
func AreProgramsResidentNV(n int32, programs *uint32, residences *bool) bool {
	ret, _, _ := syscall.Syscall(gpAreProgramsResidentNV, 3, uintptr(n), uintptr(unsafe.Pointer(programs)), uintptr(unsafe.Pointer(residences)))
	return ret != 0
}

// determine if textures are loaded in texture memory
func AreTexturesResident(n int32, textures *uint32, residences *bool) bool {
	ret, _, _ := syscall.Syscall(gpAreTexturesResident, 3, uintptr(n), uintptr(unsafe.Pointer(textures)), uintptr(unsafe.Pointer(residences)))
	return ret != 0
}
func AreTexturesResidentEXT(n int32, textures *uint32, residences *bool) bool {
	ret, _, _ := syscall.Syscall(gpAreTexturesResidentEXT, 3, uintptr(n), uintptr(unsafe.Pointer(textures)), uintptr(unsafe.Pointer(residences)))
	return ret != 0
}

// render a vertex using the specified vertex array element
func ArrayElement(i int32) {
	syscall.Syscall(gpArrayElement, 1, uintptr(i), 0, 0)
}
func ArrayElementEXT(i int32) {
	syscall.Syscall(gpArrayElementEXT, 1, uintptr(i), 0, 0)
}
func ArrayObjectATI(array uint32, size int32, xtype uint32, stride int32, buffer uint32, offset uint32) {
	syscall.Syscall6(gpArrayObjectATI, 6, uintptr(array), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(buffer), uintptr(offset))
}
func AsyncMarkerSGIX(marker uint32) {
	syscall.Syscall(gpAsyncMarkerSGIX, 1, uintptr(marker), 0, 0)
}
func AttachObjectARB(containerObj uintptr, obj uintptr) {
	syscall.Syscall(gpAttachObjectARB, 2, uintptr(containerObj), uintptr(obj), 0)
}

// Attaches a shader object to a program object
func AttachShader(program uint32, shader uint32) {
	syscall.Syscall(gpAttachShader, 2, uintptr(program), uintptr(shader), 0)
}

// delimit the vertices of a primitive or a group of like primitives
func Begin(mode uint32) {
	syscall.Syscall(gpBegin, 1, uintptr(mode), 0, 0)
}
func BeginConditionalRenderNV(id uint32, mode uint32) {
	syscall.Syscall(gpBeginConditionalRenderNV, 2, uintptr(id), uintptr(mode), 0)
}
func BeginConditionalRenderNVX(id uint32) {
	syscall.Syscall(gpBeginConditionalRenderNVX, 1, uintptr(id), 0, 0)
}
func BeginFragmentShaderATI() {
	syscall.Syscall(gpBeginFragmentShaderATI, 0, 0, 0, 0)
}
func BeginOcclusionQueryNV(id uint32) {
	syscall.Syscall(gpBeginOcclusionQueryNV, 1, uintptr(id), 0, 0)
}
func BeginPerfMonitorAMD(monitor uint32) {
	syscall.Syscall(gpBeginPerfMonitorAMD, 1, uintptr(monitor), 0, 0)
}
func BeginPerfQueryINTEL(queryHandle uint32) {
	syscall.Syscall(gpBeginPerfQueryINTEL, 1, uintptr(queryHandle), 0, 0)
}

// delimit the boundaries of a query object
func BeginQuery(target uint32, id uint32) {
	syscall.Syscall(gpBeginQuery, 2, uintptr(target), uintptr(id), 0)
}
func BeginQueryARB(target uint32, id uint32) {
	syscall.Syscall(gpBeginQueryARB, 2, uintptr(target), uintptr(id), 0)
}
func BeginQueryIndexed(target uint32, index uint32, id uint32) {
	syscall.Syscall(gpBeginQueryIndexed, 3, uintptr(target), uintptr(index), uintptr(id))
}
func BeginTransformFeedbackEXT(primitiveMode uint32) {
	syscall.Syscall(gpBeginTransformFeedbackEXT, 1, uintptr(primitiveMode), 0, 0)
}
func BeginTransformFeedbackNV(primitiveMode uint32) {
	syscall.Syscall(gpBeginTransformFeedbackNV, 1, uintptr(primitiveMode), 0, 0)
}
func BeginVertexShaderEXT() {
	syscall.Syscall(gpBeginVertexShaderEXT, 0, 0, 0, 0)
}
func BeginVideoCaptureNV(video_capture_slot uint32) {
	syscall.Syscall(gpBeginVideoCaptureNV, 1, uintptr(video_capture_slot), 0, 0)
}

// Associates a generic vertex attribute index with a named attribute variable
func BindAttribLocation(program uint32, index uint32, name *uint8) {
	syscall.Syscall(gpBindAttribLocation, 3, uintptr(program), uintptr(index), uintptr(unsafe.Pointer(name)))
}
func BindAttribLocationARB(programObj uintptr, index uint32, name *uint8) {
	syscall.Syscall(gpBindAttribLocationARB, 3, uintptr(programObj), uintptr(index), uintptr(unsafe.Pointer(name)))
}

// bind a named buffer object
func BindBuffer(target uint32, buffer uint32) {
	syscall.Syscall(gpBindBuffer, 2, uintptr(target), uintptr(buffer), 0)
}
func BindBufferARB(target uint32, buffer uint32) {
	syscall.Syscall(gpBindBufferARB, 2, uintptr(target), uintptr(buffer), 0)
}

// bind a buffer object to an indexed buffer target
func BindBufferBase(target uint32, index uint32, buffer uint32) {
	syscall.Syscall(gpBindBufferBase, 3, uintptr(target), uintptr(index), uintptr(buffer))
}
func BindBufferBaseEXT(target uint32, index uint32, buffer uint32) {
	syscall.Syscall(gpBindBufferBaseEXT, 3, uintptr(target), uintptr(index), uintptr(buffer))
}
func BindBufferBaseNV(target uint32, index uint32, buffer uint32) {
	syscall.Syscall(gpBindBufferBaseNV, 3, uintptr(target), uintptr(index), uintptr(buffer))
}
func BindBufferOffsetEXT(target uint32, index uint32, buffer uint32, offset int) {
	syscall.Syscall6(gpBindBufferOffsetEXT, 4, uintptr(target), uintptr(index), uintptr(buffer), uintptr(offset), 0, 0)
}
func BindBufferOffsetNV(target uint32, index uint32, buffer uint32, offset int) {
	syscall.Syscall6(gpBindBufferOffsetNV, 4, uintptr(target), uintptr(index), uintptr(buffer), uintptr(offset), 0, 0)
}

// bind a range within a buffer object to an indexed buffer target
func BindBufferRange(target uint32, index uint32, buffer uint32, offset int, size int) {
	syscall.Syscall6(gpBindBufferRange, 5, uintptr(target), uintptr(index), uintptr(buffer), uintptr(offset), uintptr(size), 0)
}
func BindBufferRangeEXT(target uint32, index uint32, buffer uint32, offset int, size int) {
	syscall.Syscall6(gpBindBufferRangeEXT, 5, uintptr(target), uintptr(index), uintptr(buffer), uintptr(offset), uintptr(size), 0)
}
func BindBufferRangeNV(target uint32, index uint32, buffer uint32, offset int, size int) {
	syscall.Syscall6(gpBindBufferRangeNV, 5, uintptr(target), uintptr(index), uintptr(buffer), uintptr(offset), uintptr(size), 0)
}

// bind one or more buffer objects to a sequence of indexed buffer targets
func BindBuffersBase(target uint32, first uint32, count int32, buffers *uint32) {
	syscall.Syscall6(gpBindBuffersBase, 4, uintptr(target), uintptr(first), uintptr(count), uintptr(unsafe.Pointer(buffers)), 0, 0)
}

// bind ranges of one or more buffer objects to a sequence of indexed buffer targets
func BindBuffersRange(target uint32, first uint32, count int32, buffers *uint32, offsets *int, sizes *int) {
	syscall.Syscall6(gpBindBuffersRange, 6, uintptr(target), uintptr(first), uintptr(count), uintptr(unsafe.Pointer(buffers)), uintptr(unsafe.Pointer(offsets)), uintptr(unsafe.Pointer(sizes)))
}
func BindFragDataLocationEXT(program uint32, color uint32, name *uint8) {
	syscall.Syscall(gpBindFragDataLocationEXT, 3, uintptr(program), uintptr(color), uintptr(unsafe.Pointer(name)))
}

// bind a user-defined varying out variable to a fragment shader color number and index
func BindFragDataLocationIndexed(program uint32, colorNumber uint32, index uint32, name *uint8) {
	syscall.Syscall6(gpBindFragDataLocationIndexed, 4, uintptr(program), uintptr(colorNumber), uintptr(index), uintptr(unsafe.Pointer(name)), 0, 0)
}
func BindFragmentShaderATI(id uint32) {
	syscall.Syscall(gpBindFragmentShaderATI, 1, uintptr(id), 0, 0)
}

// bind a framebuffer to a framebuffer target
func BindFramebuffer(target uint32, framebuffer uint32) {
	syscall.Syscall(gpBindFramebuffer, 2, uintptr(target), uintptr(framebuffer), 0)
}
func BindFramebufferEXT(target uint32, framebuffer uint32) {
	syscall.Syscall(gpBindFramebufferEXT, 2, uintptr(target), uintptr(framebuffer), 0)
}

// bind a level of a texture to an image unit
func BindImageTexture(unit uint32, texture uint32, level int32, layered bool, layer int32, access uint32, format uint32) {
	syscall.Syscall9(gpBindImageTexture, 7, uintptr(unit), uintptr(texture), uintptr(level), boolToUintptr(layered), uintptr(layer), uintptr(access), uintptr(format), 0, 0)
}
func BindImageTextureEXT(index uint32, texture uint32, level int32, layered bool, layer int32, access uint32, format int32) {
	syscall.Syscall9(gpBindImageTextureEXT, 7, uintptr(index), uintptr(texture), uintptr(level), boolToUintptr(layered), uintptr(layer), uintptr(access), uintptr(format), 0, 0)
}

// bind one or more named texture images to a sequence of consecutive image units
func BindImageTextures(first uint32, count int32, textures *uint32) {
	syscall.Syscall(gpBindImageTextures, 3, uintptr(first), uintptr(count), uintptr(unsafe.Pointer(textures)))
}
func BindLightParameterEXT(light uint32, value uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpBindLightParameterEXT, 2, uintptr(light), uintptr(value), 0)
	return (uint32)(ret)
}
func BindMaterialParameterEXT(face uint32, value uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpBindMaterialParameterEXT, 2, uintptr(face), uintptr(value), 0)
	return (uint32)(ret)
}
func BindMultiTextureEXT(texunit uint32, target uint32, texture uint32) {
	syscall.Syscall(gpBindMultiTextureEXT, 3, uintptr(texunit), uintptr(target), uintptr(texture))
}
func BindParameterEXT(value uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpBindParameterEXT, 1, uintptr(value), 0, 0)
	return (uint32)(ret)
}
func BindProgramARB(target uint32, program uint32) {
	syscall.Syscall(gpBindProgramARB, 2, uintptr(target), uintptr(program), 0)
}
func BindProgramNV(target uint32, id uint32) {
	syscall.Syscall(gpBindProgramNV, 2, uintptr(target), uintptr(id), 0)
}

// bind a program pipeline to the current context
func BindProgramPipeline(pipeline uint32) {
	syscall.Syscall(gpBindProgramPipeline, 1, uintptr(pipeline), 0, 0)
}
func BindProgramPipelineEXT(pipeline uint32) {
	syscall.Syscall(gpBindProgramPipelineEXT, 1, uintptr(pipeline), 0, 0)
}

// bind a renderbuffer to a renderbuffer target
func BindRenderbuffer(target uint32, renderbuffer uint32) {
	syscall.Syscall(gpBindRenderbuffer, 2, uintptr(target), uintptr(renderbuffer), 0)
}
func BindRenderbufferEXT(target uint32, renderbuffer uint32) {
	syscall.Syscall(gpBindRenderbufferEXT, 2, uintptr(target), uintptr(renderbuffer), 0)
}

// bind a named sampler to a texturing target
func BindSampler(unit uint32, sampler uint32) {
	syscall.Syscall(gpBindSampler, 2, uintptr(unit), uintptr(sampler), 0)
}

// bind one or more named sampler objects to a sequence of consecutive sampler units
func BindSamplers(first uint32, count int32, samplers *uint32) {
	syscall.Syscall(gpBindSamplers, 3, uintptr(first), uintptr(count), uintptr(unsafe.Pointer(samplers)))
}
func BindTexGenParameterEXT(unit uint32, coord uint32, value uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpBindTexGenParameterEXT, 3, uintptr(unit), uintptr(coord), uintptr(value))
	return (uint32)(ret)
}

// bind a named texture to a texturing target
func BindTexture(target uint32, texture uint32) {
	syscall.Syscall(gpBindTexture, 2, uintptr(target), uintptr(texture), 0)
}
func BindTextureEXT(target uint32, texture uint32) {
	syscall.Syscall(gpBindTextureEXT, 2, uintptr(target), uintptr(texture), 0)
}

// bind an existing texture object to the specified texture unit
func BindTextureUnit(unit uint32, texture uint32) {
	syscall.Syscall(gpBindTextureUnit, 2, uintptr(unit), uintptr(texture), 0)
}
func BindTextureUnitParameterEXT(unit uint32, value uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpBindTextureUnitParameterEXT, 2, uintptr(unit), uintptr(value), 0)
	return (uint32)(ret)
}

// bind one or more named textures to a sequence of consecutive texture units
func BindTextures(first uint32, count int32, textures *uint32) {
	syscall.Syscall(gpBindTextures, 3, uintptr(first), uintptr(count), uintptr(unsafe.Pointer(textures)))
}

// bind a transform feedback object
func BindTransformFeedback(target uint32, id uint32) {
	syscall.Syscall(gpBindTransformFeedback, 2, uintptr(target), uintptr(id), 0)
}
func BindTransformFeedbackNV(target uint32, id uint32) {
	syscall.Syscall(gpBindTransformFeedbackNV, 2, uintptr(target), uintptr(id), 0)
}

// bind a vertex array object
func BindVertexArray(array uint32) {
	syscall.Syscall(gpBindVertexArray, 1, uintptr(array), 0, 0)
}
func BindVertexArrayAPPLE(array uint32) {
	syscall.Syscall(gpBindVertexArrayAPPLE, 1, uintptr(array), 0, 0)
}

// bind a buffer to a vertex buffer bind point
func BindVertexBuffer(bindingindex uint32, buffer uint32, offset int, stride int32) {
	syscall.Syscall6(gpBindVertexBuffer, 4, uintptr(bindingindex), uintptr(buffer), uintptr(offset), uintptr(stride), 0, 0)
}

// attach multiple buffer objects to a vertex array object
func BindVertexBuffers(first uint32, count int32, buffers *uint32, offsets *int, strides *int32) {
	syscall.Syscall6(gpBindVertexBuffers, 5, uintptr(first), uintptr(count), uintptr(unsafe.Pointer(buffers)), uintptr(unsafe.Pointer(offsets)), uintptr(unsafe.Pointer(strides)), 0)
}
func BindVertexShaderEXT(id uint32) {
	syscall.Syscall(gpBindVertexShaderEXT, 1, uintptr(id), 0, 0)
}
func BindVideoCaptureStreamBufferNV(video_capture_slot uint32, stream uint32, frame_region uint32, offset int) {
	syscall.Syscall6(gpBindVideoCaptureStreamBufferNV, 4, uintptr(video_capture_slot), uintptr(stream), uintptr(frame_region), uintptr(offset), 0, 0)
}
func BindVideoCaptureStreamTextureNV(video_capture_slot uint32, stream uint32, frame_region uint32, target uint32, texture uint32) {
	syscall.Syscall6(gpBindVideoCaptureStreamTextureNV, 5, uintptr(video_capture_slot), uintptr(stream), uintptr(frame_region), uintptr(target), uintptr(texture), 0)
}
func Binormal3bEXT(bx int8, by int8, bz int8) {
	syscall.Syscall(gpBinormal3bEXT, 3, uintptr(bx), uintptr(by), uintptr(bz))
}
func Binormal3bvEXT(v *int8) {
	syscall.Syscall(gpBinormal3bvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Binormal3dEXT(bx float64, by float64, bz float64) {
	syscall.Syscall(gpBinormal3dEXT, 3, uintptr(math.Float64bits(bx)), uintptr(math.Float64bits(by)), uintptr(math.Float64bits(bz)))
}
func Binormal3dvEXT(v *float64) {
	syscall.Syscall(gpBinormal3dvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Binormal3fEXT(bx float32, by float32, bz float32) {
	syscall.Syscall(gpBinormal3fEXT, 3, uintptr(math.Float32bits(bx)), uintptr(math.Float32bits(by)), uintptr(math.Float32bits(bz)))
}
func Binormal3fvEXT(v *float32) {
	syscall.Syscall(gpBinormal3fvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Binormal3iEXT(bx int32, by int32, bz int32) {
	syscall.Syscall(gpBinormal3iEXT, 3, uintptr(bx), uintptr(by), uintptr(bz))
}
func Binormal3ivEXT(v *int32) {
	syscall.Syscall(gpBinormal3ivEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Binormal3sEXT(bx int16, by int16, bz int16) {
	syscall.Syscall(gpBinormal3sEXT, 3, uintptr(bx), uintptr(by), uintptr(bz))
}
func Binormal3svEXT(v *int16) {
	syscall.Syscall(gpBinormal3svEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func BinormalPointerEXT(xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpBinormalPointerEXT, 3, uintptr(xtype), uintptr(stride), uintptr(pointer))
}

// draw a bitmap
func Bitmap(width int32, height int32, xorig float32, yorig float32, xmove float32, ymove float32, bitmap *uint8) {
	syscall.Syscall9(gpBitmap, 7, uintptr(width), uintptr(height), uintptr(math.Float32bits(xorig)), uintptr(math.Float32bits(yorig)), uintptr(math.Float32bits(xmove)), uintptr(math.Float32bits(ymove)), uintptr(unsafe.Pointer(bitmap)), 0, 0)
}
func BitmapxOES(width int32, height int32, xorig int32, yorig int32, xmove int32, ymove int32, bitmap *uint8) {
	syscall.Syscall9(gpBitmapxOES, 7, uintptr(width), uintptr(height), uintptr(xorig), uintptr(yorig), uintptr(xmove), uintptr(ymove), uintptr(unsafe.Pointer(bitmap)), 0, 0)
}
func BlendBarrierKHR() {
	syscall.Syscall(gpBlendBarrierKHR, 0, 0, 0, 0)
}
func BlendBarrierNV() {
	syscall.Syscall(gpBlendBarrierNV, 0, 0, 0, 0)
}

// set the blend color
func BlendColor(red float32, green float32, blue float32, alpha float32) {
	syscall.Syscall6(gpBlendColor, 4, uintptr(math.Float32bits(red)), uintptr(math.Float32bits(green)), uintptr(math.Float32bits(blue)), uintptr(math.Float32bits(alpha)), 0, 0)
}
func BlendColorEXT(red float32, green float32, blue float32, alpha float32) {
	syscall.Syscall6(gpBlendColorEXT, 4, uintptr(math.Float32bits(red)), uintptr(math.Float32bits(green)), uintptr(math.Float32bits(blue)), uintptr(math.Float32bits(alpha)), 0, 0)
}
func BlendColorxOES(red int32, green int32, blue int32, alpha int32) {
	syscall.Syscall6(gpBlendColorxOES, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}

// specify the equation used for both the RGB blend equation and the Alpha blend equation
func BlendEquation(mode uint32) {
	syscall.Syscall(gpBlendEquation, 1, uintptr(mode), 0, 0)
}
func BlendEquationEXT(mode uint32) {
	syscall.Syscall(gpBlendEquationEXT, 1, uintptr(mode), 0, 0)
}
func BlendEquationIndexedAMD(buf uint32, mode uint32) {
	syscall.Syscall(gpBlendEquationIndexedAMD, 2, uintptr(buf), uintptr(mode), 0)
}

// set the RGB blend equation and the alpha blend equation separately
func BlendEquationSeparate(modeRGB uint32, modeAlpha uint32) {
	syscall.Syscall(gpBlendEquationSeparate, 2, uintptr(modeRGB), uintptr(modeAlpha), 0)
}
func BlendEquationSeparateEXT(modeRGB uint32, modeAlpha uint32) {
	syscall.Syscall(gpBlendEquationSeparateEXT, 2, uintptr(modeRGB), uintptr(modeAlpha), 0)
}
func BlendEquationSeparateIndexedAMD(buf uint32, modeRGB uint32, modeAlpha uint32) {
	syscall.Syscall(gpBlendEquationSeparateIndexedAMD, 3, uintptr(buf), uintptr(modeRGB), uintptr(modeAlpha))
}
func BlendEquationSeparateiARB(buf uint32, modeRGB uint32, modeAlpha uint32) {
	syscall.Syscall(gpBlendEquationSeparateiARB, 3, uintptr(buf), uintptr(modeRGB), uintptr(modeAlpha))
}
func BlendEquationiARB(buf uint32, mode uint32) {
	syscall.Syscall(gpBlendEquationiARB, 2, uintptr(buf), uintptr(mode), 0)
}

// specify pixel arithmetic
func BlendFunc(sfactor uint32, dfactor uint32) {
	syscall.Syscall(gpBlendFunc, 2, uintptr(sfactor), uintptr(dfactor), 0)
}
func BlendFuncIndexedAMD(buf uint32, src uint32, dst uint32) {
	syscall.Syscall(gpBlendFuncIndexedAMD, 3, uintptr(buf), uintptr(src), uintptr(dst))
}

// specify pixel arithmetic for RGB and alpha components separately
func BlendFuncSeparate(sfactorRGB uint32, dfactorRGB uint32, sfactorAlpha uint32, dfactorAlpha uint32) {
	syscall.Syscall6(gpBlendFuncSeparate, 4, uintptr(sfactorRGB), uintptr(dfactorRGB), uintptr(sfactorAlpha), uintptr(dfactorAlpha), 0, 0)
}
func BlendFuncSeparateEXT(sfactorRGB uint32, dfactorRGB uint32, sfactorAlpha uint32, dfactorAlpha uint32) {
	syscall.Syscall6(gpBlendFuncSeparateEXT, 4, uintptr(sfactorRGB), uintptr(dfactorRGB), uintptr(sfactorAlpha), uintptr(dfactorAlpha), 0, 0)
}
func BlendFuncSeparateINGR(sfactorRGB uint32, dfactorRGB uint32, sfactorAlpha uint32, dfactorAlpha uint32) {
	syscall.Syscall6(gpBlendFuncSeparateINGR, 4, uintptr(sfactorRGB), uintptr(dfactorRGB), uintptr(sfactorAlpha), uintptr(dfactorAlpha), 0, 0)
}
func BlendFuncSeparateIndexedAMD(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	syscall.Syscall6(gpBlendFuncSeparateIndexedAMD, 5, uintptr(buf), uintptr(srcRGB), uintptr(dstRGB), uintptr(srcAlpha), uintptr(dstAlpha), 0)
}
func BlendFuncSeparateiARB(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	syscall.Syscall6(gpBlendFuncSeparateiARB, 5, uintptr(buf), uintptr(srcRGB), uintptr(dstRGB), uintptr(srcAlpha), uintptr(dstAlpha), 0)
}
func BlendFunciARB(buf uint32, src uint32, dst uint32) {
	syscall.Syscall(gpBlendFunciARB, 3, uintptr(buf), uintptr(src), uintptr(dst))
}
func BlendParameteriNV(pname uint32, value int32) {
	syscall.Syscall(gpBlendParameteriNV, 2, uintptr(pname), uintptr(value), 0)
}

// copy a block of pixels from one framebuffer object to another
func BlitFramebuffer(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	syscall.Syscall12(gpBlitFramebuffer, 10, uintptr(srcX0), uintptr(srcY0), uintptr(srcX1), uintptr(srcY1), uintptr(dstX0), uintptr(dstY0), uintptr(dstX1), uintptr(dstY1), uintptr(mask), uintptr(filter), 0, 0)
}
func BlitFramebufferEXT(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	syscall.Syscall12(gpBlitFramebufferEXT, 10, uintptr(srcX0), uintptr(srcY0), uintptr(srcX1), uintptr(srcY1), uintptr(dstX0), uintptr(dstY0), uintptr(dstX1), uintptr(dstY1), uintptr(mask), uintptr(filter), 0, 0)
}

// copy a block of pixels from one framebuffer object to another
func BlitNamedFramebuffer(readFramebuffer uint32, drawFramebuffer uint32, srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	syscall.Syscall12(gpBlitNamedFramebuffer, 12, uintptr(readFramebuffer), uintptr(drawFramebuffer), uintptr(srcX0), uintptr(srcY0), uintptr(srcX1), uintptr(srcY1), uintptr(dstX0), uintptr(dstY0), uintptr(dstX1), uintptr(dstY1), uintptr(mask), uintptr(filter))
}
func BufferAddressRangeNV(pname uint32, index uint32, address uint64, length int) {
	syscall.Syscall6(gpBufferAddressRangeNV, 4, uintptr(pname), uintptr(index), uintptr(address), uintptr(length), 0, 0)
}

// creates and initializes a buffer object's data     store
func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	syscall.Syscall6(gpBufferData, 4, uintptr(target), uintptr(size), uintptr(data), uintptr(usage), 0, 0)
}
func BufferDataARB(target uint32, size int, data unsafe.Pointer, usage uint32) {
	syscall.Syscall6(gpBufferDataARB, 4, uintptr(target), uintptr(size), uintptr(data), uintptr(usage), 0, 0)
}
func BufferPageCommitmentARB(target uint32, offset int, size int, commit bool) {
	syscall.Syscall6(gpBufferPageCommitmentARB, 4, uintptr(target), uintptr(offset), uintptr(size), boolToUintptr(commit), 0, 0)
}
func BufferParameteriAPPLE(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpBufferParameteriAPPLE, 3, uintptr(target), uintptr(pname), uintptr(param))
}

// creates and initializes a buffer object's immutable data     store
func BufferStorage(target uint32, size int, data unsafe.Pointer, flags uint32) {
	syscall.Syscall6(gpBufferStorage, 4, uintptr(target), uintptr(size), uintptr(data), uintptr(flags), 0, 0)
}

// Parameter clientBuffer has type C.GLeglClientBufferEXT.
func BufferStorageExternalEXT(target uint32, offset int, size int, clientBuffer unsafe.Pointer, flags uint32) {
	syscall.Syscall6(gpBufferStorageExternalEXT, 5, uintptr(target), uintptr(offset), uintptr(size), uintptr(clientBuffer), uintptr(flags), 0)
}
func BufferStorageMemEXT(target uint32, size int, memory uint32, offset uint64) {
	syscall.Syscall6(gpBufferStorageMemEXT, 4, uintptr(target), uintptr(size), uintptr(memory), uintptr(offset), 0, 0)
}

// updates a subset of a buffer object's data store
func BufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpBufferSubData, 4, uintptr(target), uintptr(offset), uintptr(size), uintptr(data), 0, 0)
}
func BufferSubDataARB(target uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpBufferSubDataARB, 4, uintptr(target), uintptr(offset), uintptr(size), uintptr(data), 0, 0)
}
func CallCommandListNV(list uint32) {
	syscall.Syscall(gpCallCommandListNV, 1, uintptr(list), 0, 0)
}

// execute a display list
func CallList(list uint32) {
	syscall.Syscall(gpCallList, 1, uintptr(list), 0, 0)
}

// execute a list of display lists
func CallLists(n int32, xtype uint32, lists unsafe.Pointer) {
	syscall.Syscall(gpCallLists, 3, uintptr(n), uintptr(xtype), uintptr(lists))
}

// check the completeness status of a framebuffer
func CheckFramebufferStatus(target uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpCheckFramebufferStatus, 1, uintptr(target), 0, 0)
	return (uint32)(ret)
}
func CheckFramebufferStatusEXT(target uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpCheckFramebufferStatusEXT, 1, uintptr(target), 0, 0)
	return (uint32)(ret)
}

// check the completeness status of a framebuffer
func CheckNamedFramebufferStatus(framebuffer uint32, target uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpCheckNamedFramebufferStatus, 2, uintptr(framebuffer), uintptr(target), 0)
	return (uint32)(ret)
}
func CheckNamedFramebufferStatusEXT(framebuffer uint32, target uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpCheckNamedFramebufferStatusEXT, 2, uintptr(framebuffer), uintptr(target), 0)
	return (uint32)(ret)
}
func ClampColorARB(target uint32, clamp uint32) {
	syscall.Syscall(gpClampColorARB, 2, uintptr(target), uintptr(clamp), 0)
}

// clear buffers to preset values
func Clear(mask uint32) {
	syscall.Syscall(gpClear, 1, uintptr(mask), 0, 0)
}

// specify clear values for the accumulation buffer
func ClearAccum(red float32, green float32, blue float32, alpha float32) {
	syscall.Syscall6(gpClearAccum, 4, uintptr(math.Float32bits(red)), uintptr(math.Float32bits(green)), uintptr(math.Float32bits(blue)), uintptr(math.Float32bits(alpha)), 0, 0)
}
func ClearAccumxOES(red int32, green int32, blue int32, alpha int32) {
	syscall.Syscall6(gpClearAccumxOES, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}

// fill a buffer object's data store with a fixed value
func ClearBufferData(target uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall6(gpClearBufferData, 5, uintptr(target), uintptr(internalformat), uintptr(format), uintptr(xtype), uintptr(data), 0)
}

// fill all or part of buffer object's data store with a fixed value
func ClearBufferSubData(target uint32, internalformat uint32, offset int, size int, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall9(gpClearBufferSubData, 7, uintptr(target), uintptr(internalformat), uintptr(offset), uintptr(size), uintptr(format), uintptr(xtype), uintptr(data), 0, 0)
}

// specify clear values for the color buffers
func ClearColor(red float32, green float32, blue float32, alpha float32) {
	syscall.Syscall6(gpClearColor, 4, uintptr(math.Float32bits(red)), uintptr(math.Float32bits(green)), uintptr(math.Float32bits(blue)), uintptr(math.Float32bits(alpha)), 0, 0)
}
func ClearColorIiEXT(red int32, green int32, blue int32, alpha int32) {
	syscall.Syscall6(gpClearColorIiEXT, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func ClearColorIuiEXT(red uint32, green uint32, blue uint32, alpha uint32) {
	syscall.Syscall6(gpClearColorIuiEXT, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func ClearColorxOES(red int32, green int32, blue int32, alpha int32) {
	syscall.Syscall6(gpClearColorxOES, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}

// specify the clear value for the depth buffer
func ClearDepth(depth float64) {
	syscall.Syscall(gpClearDepth, 1, uintptr(math.Float64bits(depth)), 0, 0)
}
func ClearDepthdNV(depth float64) {
	syscall.Syscall(gpClearDepthdNV, 1, uintptr(math.Float64bits(depth)), 0, 0)
}

// specify the clear value for the depth buffer
func ClearDepthf(d float32) {
	syscall.Syscall(gpClearDepthf, 1, uintptr(math.Float32bits(d)), 0, 0)
}
func ClearDepthfOES(depth float32) {
	syscall.Syscall(gpClearDepthfOES, 1, uintptr(math.Float32bits(depth)), 0, 0)
}
func ClearDepthxOES(depth int32) {
	syscall.Syscall(gpClearDepthxOES, 1, uintptr(depth), 0, 0)
}

// specify the clear value for the color index buffers
func ClearIndex(c float32) {
	syscall.Syscall(gpClearIndex, 1, uintptr(math.Float32bits(c)), 0, 0)
}

// fill a buffer object's data store with a fixed value
func ClearNamedBufferData(buffer uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall6(gpClearNamedBufferData, 5, uintptr(buffer), uintptr(internalformat), uintptr(format), uintptr(xtype), uintptr(data), 0)
}
func ClearNamedBufferDataEXT(buffer uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall6(gpClearNamedBufferDataEXT, 5, uintptr(buffer), uintptr(internalformat), uintptr(format), uintptr(xtype), uintptr(data), 0)
}

// fill all or part of buffer object's data store with a fixed value
func ClearNamedBufferSubData(buffer uint32, internalformat uint32, offset int, size int, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall9(gpClearNamedBufferSubData, 7, uintptr(buffer), uintptr(internalformat), uintptr(offset), uintptr(size), uintptr(format), uintptr(xtype), uintptr(data), 0, 0)
}
func ClearNamedBufferSubDataEXT(buffer uint32, internalformat uint32, offset int, size int, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall9(gpClearNamedBufferSubDataEXT, 7, uintptr(buffer), uintptr(internalformat), uintptr(offset), uintptr(size), uintptr(format), uintptr(xtype), uintptr(data), 0, 0)
}
func ClearNamedFramebufferfi(framebuffer uint32, buffer uint32, drawbuffer int32, depth float32, stencil int32) {
	syscall.Syscall6(gpClearNamedFramebufferfi, 5, uintptr(framebuffer), uintptr(buffer), uintptr(drawbuffer), uintptr(math.Float32bits(depth)), uintptr(stencil), 0)
}
func ClearNamedFramebufferfv(framebuffer uint32, buffer uint32, drawbuffer int32, value *float32) {
	syscall.Syscall6(gpClearNamedFramebufferfv, 4, uintptr(framebuffer), uintptr(buffer), uintptr(drawbuffer), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ClearNamedFramebufferiv(framebuffer uint32, buffer uint32, drawbuffer int32, value *int32) {
	syscall.Syscall6(gpClearNamedFramebufferiv, 4, uintptr(framebuffer), uintptr(buffer), uintptr(drawbuffer), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ClearNamedFramebufferuiv(framebuffer uint32, buffer uint32, drawbuffer int32, value *uint32) {
	syscall.Syscall6(gpClearNamedFramebufferuiv, 4, uintptr(framebuffer), uintptr(buffer), uintptr(drawbuffer), uintptr(unsafe.Pointer(value)), 0, 0)
}

// specify the clear value for the stencil buffer
func ClearStencil(s int32) {
	syscall.Syscall(gpClearStencil, 1, uintptr(s), 0, 0)
}

// fills all a texture image with a constant value
func ClearTexImage(texture uint32, level int32, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall6(gpClearTexImage, 5, uintptr(texture), uintptr(level), uintptr(format), uintptr(xtype), uintptr(data), 0)
}

// fills all or part of a texture image with a constant value
func ClearTexSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall12(gpClearTexSubImage, 11, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(xtype), uintptr(data), 0)
}

// select active texture unit
func ClientActiveTexture(texture uint32) {
	syscall.Syscall(gpClientActiveTexture, 1, uintptr(texture), 0, 0)
}
func ClientActiveTextureARB(texture uint32) {
	syscall.Syscall(gpClientActiveTextureARB, 1, uintptr(texture), 0, 0)
}
func ClientActiveVertexStreamATI(stream uint32) {
	syscall.Syscall(gpClientActiveVertexStreamATI, 1, uintptr(stream), 0, 0)
}
func ClientAttribDefaultEXT(mask uint32) {
	syscall.Syscall(gpClientAttribDefaultEXT, 1, uintptr(mask), 0, 0)
}

// block and wait for a sync object to become signaled
func ClientWaitSync(sync uintptr, flags uint32, timeout uint64) uint32 {
	ret, _, _ := syscall.Syscall(gpClientWaitSync, 3, uintptr(sync), uintptr(flags), uintptr(timeout))
	return (uint32)(ret)
}

// control clip coordinate to window coordinate behavior
func ClipControl(origin uint32, depth uint32) {
	syscall.Syscall(gpClipControl, 2, uintptr(origin), uintptr(depth), 0)
}

// specify a plane against which all geometry is clipped
func ClipPlane(plane uint32, equation *float64) {
	syscall.Syscall(gpClipPlane, 2, uintptr(plane), uintptr(unsafe.Pointer(equation)), 0)
}
func ClipPlanefOES(plane uint32, equation *float32) {
	syscall.Syscall(gpClipPlanefOES, 2, uintptr(plane), uintptr(unsafe.Pointer(equation)), 0)
}
func ClipPlanexOES(plane uint32, equation *int32) {
	syscall.Syscall(gpClipPlanexOES, 2, uintptr(plane), uintptr(unsafe.Pointer(equation)), 0)
}
func Color3b(red int8, green int8, blue int8) {
	syscall.Syscall(gpColor3b, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func Color3bv(v *int8) {
	syscall.Syscall(gpColor3bv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3d(red float64, green float64, blue float64) {
	syscall.Syscall(gpColor3d, 3, uintptr(math.Float64bits(red)), uintptr(math.Float64bits(green)), uintptr(math.Float64bits(blue)))
}
func Color3dv(v *float64) {
	syscall.Syscall(gpColor3dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3f(red float32, green float32, blue float32) {
	syscall.Syscall(gpColor3f, 3, uintptr(math.Float32bits(red)), uintptr(math.Float32bits(green)), uintptr(math.Float32bits(blue)))
}
func Color3fVertex3fSUN(r float32, g float32, b float32, x float32, y float32, z float32) {
	syscall.Syscall6(gpColor3fVertex3fSUN, 6, uintptr(math.Float32bits(r)), uintptr(math.Float32bits(g)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func Color3fVertex3fvSUN(c *float32, v *float32) {
	syscall.Syscall(gpColor3fVertex3fvSUN, 2, uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(v)), 0)
}
func Color3fv(v *float32) {
	syscall.Syscall(gpColor3fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3hNV(red uint16, green uint16, blue uint16) {
	syscall.Syscall(gpColor3hNV, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func Color3hvNV(v *uint16) {
	syscall.Syscall(gpColor3hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3i(red int32, green int32, blue int32) {
	syscall.Syscall(gpColor3i, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func Color3iv(v *int32) {
	syscall.Syscall(gpColor3iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3s(red int16, green int16, blue int16) {
	syscall.Syscall(gpColor3s, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func Color3sv(v *int16) {
	syscall.Syscall(gpColor3sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3ub(red uint8, green uint8, blue uint8) {
	syscall.Syscall(gpColor3ub, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func Color3ubv(v *uint8) {
	syscall.Syscall(gpColor3ubv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3ui(red uint32, green uint32, blue uint32) {
	syscall.Syscall(gpColor3ui, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func Color3uiv(v *uint32) {
	syscall.Syscall(gpColor3uiv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3us(red uint16, green uint16, blue uint16) {
	syscall.Syscall(gpColor3us, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func Color3usv(v *uint16) {
	syscall.Syscall(gpColor3usv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color3xOES(red int32, green int32, blue int32) {
	syscall.Syscall(gpColor3xOES, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func Color3xvOES(components *int32) {
	syscall.Syscall(gpColor3xvOES, 1, uintptr(unsafe.Pointer(components)), 0, 0)
}
func Color4b(red int8, green int8, blue int8, alpha int8) {
	syscall.Syscall6(gpColor4b, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func Color4bv(v *int8) {
	syscall.Syscall(gpColor4bv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4d(red float64, green float64, blue float64, alpha float64) {
	syscall.Syscall6(gpColor4d, 4, uintptr(math.Float64bits(red)), uintptr(math.Float64bits(green)), uintptr(math.Float64bits(blue)), uintptr(math.Float64bits(alpha)), 0, 0)
}
func Color4dv(v *float64) {
	syscall.Syscall(gpColor4dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4f(red float32, green float32, blue float32, alpha float32) {
	syscall.Syscall6(gpColor4f, 4, uintptr(math.Float32bits(red)), uintptr(math.Float32bits(green)), uintptr(math.Float32bits(blue)), uintptr(math.Float32bits(alpha)), 0, 0)
}
func Color4fNormal3fVertex3fSUN(r float32, g float32, b float32, a float32, nx float32, ny float32, nz float32, x float32, y float32, z float32) {
	syscall.Syscall12(gpColor4fNormal3fVertex3fSUN, 10, uintptr(math.Float32bits(r)), uintptr(math.Float32bits(g)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(a)), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func Color4fNormal3fVertex3fvSUN(c *float32, n *float32, v *float32) {
	syscall.Syscall(gpColor4fNormal3fVertex3fvSUN, 3, uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)))
}
func Color4fv(v *float32) {
	syscall.Syscall(gpColor4fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4hNV(red uint16, green uint16, blue uint16, alpha uint16) {
	syscall.Syscall6(gpColor4hNV, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func Color4hvNV(v *uint16) {
	syscall.Syscall(gpColor4hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4i(red int32, green int32, blue int32, alpha int32) {
	syscall.Syscall6(gpColor4i, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func Color4iv(v *int32) {
	syscall.Syscall(gpColor4iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4s(red int16, green int16, blue int16, alpha int16) {
	syscall.Syscall6(gpColor4s, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func Color4sv(v *int16) {
	syscall.Syscall(gpColor4sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4ub(red uint8, green uint8, blue uint8, alpha uint8) {
	syscall.Syscall6(gpColor4ub, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func Color4ubVertex2fSUN(r uint8, g uint8, b uint8, a uint8, x float32, y float32) {
	syscall.Syscall6(gpColor4ubVertex2fSUN, 6, uintptr(r), uintptr(g), uintptr(b), uintptr(a), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)))
}
func Color4ubVertex2fvSUN(c *uint8, v *float32) {
	syscall.Syscall(gpColor4ubVertex2fvSUN, 2, uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(v)), 0)
}
func Color4ubVertex3fSUN(r uint8, g uint8, b uint8, a uint8, x float32, y float32, z float32) {
	syscall.Syscall9(gpColor4ubVertex3fSUN, 7, uintptr(r), uintptr(g), uintptr(b), uintptr(a), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func Color4ubVertex3fvSUN(c *uint8, v *float32) {
	syscall.Syscall(gpColor4ubVertex3fvSUN, 2, uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(v)), 0)
}
func Color4ubv(v *uint8) {
	syscall.Syscall(gpColor4ubv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4ui(red uint32, green uint32, blue uint32, alpha uint32) {
	syscall.Syscall6(gpColor4ui, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func Color4uiv(v *uint32) {
	syscall.Syscall(gpColor4uiv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4us(red uint16, green uint16, blue uint16, alpha uint16) {
	syscall.Syscall6(gpColor4us, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func Color4usv(v *uint16) {
	syscall.Syscall(gpColor4usv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Color4xOES(red int32, green int32, blue int32, alpha int32) {
	syscall.Syscall6(gpColor4xOES, 4, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha), 0, 0)
}
func Color4xvOES(components *int32) {
	syscall.Syscall(gpColor4xvOES, 1, uintptr(unsafe.Pointer(components)), 0, 0)
}
func ColorFormatNV(size int32, xtype uint32, stride int32) {
	syscall.Syscall(gpColorFormatNV, 3, uintptr(size), uintptr(xtype), uintptr(stride))
}
func ColorFragmentOp1ATI(op uint32, dst uint32, dstMask uint32, dstMod uint32, arg1 uint32, arg1Rep uint32, arg1Mod uint32) {
	syscall.Syscall9(gpColorFragmentOp1ATI, 7, uintptr(op), uintptr(dst), uintptr(dstMask), uintptr(dstMod), uintptr(arg1), uintptr(arg1Rep), uintptr(arg1Mod), 0, 0)
}
func ColorFragmentOp2ATI(op uint32, dst uint32, dstMask uint32, dstMod uint32, arg1 uint32, arg1Rep uint32, arg1Mod uint32, arg2 uint32, arg2Rep uint32, arg2Mod uint32) {
	syscall.Syscall12(gpColorFragmentOp2ATI, 10, uintptr(op), uintptr(dst), uintptr(dstMask), uintptr(dstMod), uintptr(arg1), uintptr(arg1Rep), uintptr(arg1Mod), uintptr(arg2), uintptr(arg2Rep), uintptr(arg2Mod), 0, 0)
}
func ColorFragmentOp3ATI(op uint32, dst uint32, dstMask uint32, dstMod uint32, arg1 uint32, arg1Rep uint32, arg1Mod uint32, arg2 uint32, arg2Rep uint32, arg2Mod uint32, arg3 uint32, arg3Rep uint32, arg3Mod uint32) {
	syscall.Syscall15(gpColorFragmentOp3ATI, 13, uintptr(op), uintptr(dst), uintptr(dstMask), uintptr(dstMod), uintptr(arg1), uintptr(arg1Rep), uintptr(arg1Mod), uintptr(arg2), uintptr(arg2Rep), uintptr(arg2Mod), uintptr(arg3), uintptr(arg3Rep), uintptr(arg3Mod), 0, 0)
}
func ColorMask(red bool, green bool, blue bool, alpha bool) {
	syscall.Syscall6(gpColorMask, 4, boolToUintptr(red), boolToUintptr(green), boolToUintptr(blue), boolToUintptr(alpha), 0, 0)
}
func ColorMaskIndexedEXT(index uint32, r bool, g bool, b bool, a bool) {
	syscall.Syscall6(gpColorMaskIndexedEXT, 5, uintptr(index), boolToUintptr(r), boolToUintptr(g), boolToUintptr(b), boolToUintptr(a), 0)
}

// cause a material color to track the current color
func ColorMaterial(face uint32, mode uint32) {
	syscall.Syscall(gpColorMaterial, 2, uintptr(face), uintptr(mode), 0)
}

// define an array of colors
func ColorPointer(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpColorPointer, 4, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0, 0)
}
func ColorPointerEXT(size int32, xtype uint32, stride int32, count int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpColorPointerEXT, 5, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(count), uintptr(pointer), 0)
}
func ColorPointerListIBM(size int32, xtype uint32, stride int32, pointer *unsafe.Pointer, ptrstride int32) {
	syscall.Syscall6(gpColorPointerListIBM, 5, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(unsafe.Pointer(pointer)), uintptr(ptrstride), 0)
}
func ColorPointervINTEL(size int32, xtype uint32, pointer *unsafe.Pointer) {
	syscall.Syscall(gpColorPointervINTEL, 3, uintptr(size), uintptr(xtype), uintptr(unsafe.Pointer(pointer)))
}
func ColorSubTableEXT(target uint32, start int32, count int32, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall6(gpColorSubTableEXT, 6, uintptr(target), uintptr(start), uintptr(count), uintptr(format), uintptr(xtype), uintptr(data))
}
func ColorTableEXT(target uint32, internalFormat uint32, width int32, format uint32, xtype uint32, table unsafe.Pointer) {
	syscall.Syscall6(gpColorTableEXT, 6, uintptr(target), uintptr(internalFormat), uintptr(width), uintptr(format), uintptr(xtype), uintptr(table))
}
func ColorTableParameterfvSGI(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpColorTableParameterfvSGI, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func ColorTableParameterivSGI(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpColorTableParameterivSGI, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func ColorTableSGI(target uint32, internalformat uint32, width int32, format uint32, xtype uint32, table unsafe.Pointer) {
	syscall.Syscall6(gpColorTableSGI, 6, uintptr(target), uintptr(internalformat), uintptr(width), uintptr(format), uintptr(xtype), uintptr(table))
}
func CombinerInputNV(stage uint32, portion uint32, variable uint32, input uint32, mapping uint32, componentUsage uint32) {
	syscall.Syscall6(gpCombinerInputNV, 6, uintptr(stage), uintptr(portion), uintptr(variable), uintptr(input), uintptr(mapping), uintptr(componentUsage))
}
func CombinerOutputNV(stage uint32, portion uint32, abOutput uint32, cdOutput uint32, sumOutput uint32, scale uint32, bias uint32, abDotProduct bool, cdDotProduct bool, muxSum bool) {
	syscall.Syscall12(gpCombinerOutputNV, 10, uintptr(stage), uintptr(portion), uintptr(abOutput), uintptr(cdOutput), uintptr(sumOutput), uintptr(scale), uintptr(bias), boolToUintptr(abDotProduct), boolToUintptr(cdDotProduct), boolToUintptr(muxSum), 0, 0)
}
func CombinerParameterfNV(pname uint32, param float32) {
	syscall.Syscall(gpCombinerParameterfNV, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func CombinerParameterfvNV(pname uint32, params *float32) {
	syscall.Syscall(gpCombinerParameterfvNV, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func CombinerParameteriNV(pname uint32, param int32) {
	syscall.Syscall(gpCombinerParameteriNV, 2, uintptr(pname), uintptr(param), 0)
}
func CombinerParameterivNV(pname uint32, params *int32) {
	syscall.Syscall(gpCombinerParameterivNV, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func CombinerStageParameterfvNV(stage uint32, pname uint32, params *float32) {
	syscall.Syscall(gpCombinerStageParameterfvNV, 3, uintptr(stage), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func CommandListSegmentsNV(list uint32, segments uint32) {
	syscall.Syscall(gpCommandListSegmentsNV, 2, uintptr(list), uintptr(segments), 0)
}
func CompileCommandListNV(list uint32) {
	syscall.Syscall(gpCompileCommandListNV, 1, uintptr(list), 0, 0)
}

// Compiles a shader object
func CompileShader(shader uint32) {
	syscall.Syscall(gpCompileShader, 1, uintptr(shader), 0, 0)
}
func CompileShaderARB(shaderObj uintptr) {
	syscall.Syscall(gpCompileShaderARB, 1, uintptr(shaderObj), 0, 0)
}
func CompileShaderIncludeARB(shader uint32, count int32, path **uint8, length *int32) {
	syscall.Syscall6(gpCompileShaderIncludeARB, 4, uintptr(shader), uintptr(count), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(length)), 0, 0)
}
func CompressedMultiTexImage1DEXT(texunit uint32, target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall9(gpCompressedMultiTexImage1DEXT, 8, uintptr(texunit), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(border), uintptr(imageSize), uintptr(bits), 0)
}
func CompressedMultiTexImage2DEXT(texunit uint32, target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall9(gpCompressedMultiTexImage2DEXT, 9, uintptr(texunit), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(border), uintptr(imageSize), uintptr(bits))
}
func CompressedMultiTexImage3DEXT(texunit uint32, target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall12(gpCompressedMultiTexImage3DEXT, 10, uintptr(texunit), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(border), uintptr(imageSize), uintptr(bits), 0, 0)
}
func CompressedMultiTexSubImage1DEXT(texunit uint32, target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall9(gpCompressedMultiTexSubImage1DEXT, 8, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(imageSize), uintptr(bits), 0)
}
func CompressedMultiTexSubImage2DEXT(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall12(gpCompressedMultiTexSubImage2DEXT, 10, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(imageSize), uintptr(bits), 0, 0)
}
func CompressedMultiTexSubImage3DEXT(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall12(gpCompressedMultiTexSubImage3DEXT, 12, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(imageSize), uintptr(bits))
}

// specify a one-dimensional texture image in a compressed format
func CompressedTexImage1D(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexImage1D, 7, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(border), uintptr(imageSize), uintptr(data), 0, 0)
}
func CompressedTexImage1DARB(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexImage1DARB, 7, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(border), uintptr(imageSize), uintptr(data), 0, 0)
}

// specify a two-dimensional texture image in a compressed format
func CompressedTexImage2D(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexImage2D, 8, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(border), uintptr(imageSize), uintptr(data), 0)
}
func CompressedTexImage2DARB(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexImage2DARB, 8, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(border), uintptr(imageSize), uintptr(data), 0)
}

// specify a three-dimensional texture image in a compressed format
func CompressedTexImage3D(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexImage3D, 9, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(border), uintptr(imageSize), uintptr(data))
}
func CompressedTexImage3DARB(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexImage3DARB, 9, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(border), uintptr(imageSize), uintptr(data))
}

// specify a one-dimensional texture subimage in a compressed     format
func CompressedTexSubImage1D(target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexSubImage1D, 7, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(imageSize), uintptr(data), 0, 0)
}
func CompressedTexSubImage1DARB(target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexSubImage1DARB, 7, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(imageSize), uintptr(data), 0, 0)
}

// specify a two-dimensional texture subimage in a compressed format
func CompressedTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexSubImage2D, 9, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(imageSize), uintptr(data))
}
func CompressedTexSubImage2DARB(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTexSubImage2DARB, 9, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(imageSize), uintptr(data))
}

// specify a three-dimensional texture subimage in a compressed format
func CompressedTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall12(gpCompressedTexSubImage3D, 11, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(imageSize), uintptr(data), 0)
}
func CompressedTexSubImage3DARB(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall12(gpCompressedTexSubImage3DARB, 11, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(imageSize), uintptr(data), 0)
}
func CompressedTextureImage1DEXT(texture uint32, target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTextureImage1DEXT, 8, uintptr(texture), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(border), uintptr(imageSize), uintptr(bits), 0)
}
func CompressedTextureImage2DEXT(texture uint32, target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTextureImage2DEXT, 9, uintptr(texture), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(border), uintptr(imageSize), uintptr(bits))
}
func CompressedTextureImage3DEXT(texture uint32, target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall12(gpCompressedTextureImage3DEXT, 10, uintptr(texture), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(border), uintptr(imageSize), uintptr(bits), 0, 0)
}

// specify a one-dimensional texture subimage in a compressed     format
func CompressedTextureSubImage1D(texture uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTextureSubImage1D, 7, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(imageSize), uintptr(data), 0, 0)
}
func CompressedTextureSubImage1DEXT(texture uint32, target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTextureSubImage1DEXT, 8, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(imageSize), uintptr(bits), 0)
}

// specify a two-dimensional texture subimage in a compressed format
func CompressedTextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpCompressedTextureSubImage2D, 9, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(imageSize), uintptr(data))
}
func CompressedTextureSubImage2DEXT(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall12(gpCompressedTextureSubImage2DEXT, 10, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(imageSize), uintptr(bits), 0, 0)
}

// specify a three-dimensional texture subimage in a compressed format
func CompressedTextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	syscall.Syscall12(gpCompressedTextureSubImage3D, 11, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(imageSize), uintptr(data), 0)
}
func CompressedTextureSubImage3DEXT(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, bits unsafe.Pointer) {
	syscall.Syscall12(gpCompressedTextureSubImage3DEXT, 12, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(imageSize), uintptr(bits))
}
func ConservativeRasterParameterfNV(pname uint32, value float32) {
	syscall.Syscall(gpConservativeRasterParameterfNV, 2, uintptr(pname), uintptr(math.Float32bits(value)), 0)
}
func ConservativeRasterParameteriNV(pname uint32, param int32) {
	syscall.Syscall(gpConservativeRasterParameteriNV, 2, uintptr(pname), uintptr(param), 0)
}
func ConvolutionFilter1DEXT(target uint32, internalformat uint32, width int32, format uint32, xtype uint32, image unsafe.Pointer) {
	syscall.Syscall6(gpConvolutionFilter1DEXT, 6, uintptr(target), uintptr(internalformat), uintptr(width), uintptr(format), uintptr(xtype), uintptr(image))
}
func ConvolutionFilter2DEXT(target uint32, internalformat uint32, width int32, height int32, format uint32, xtype uint32, image unsafe.Pointer) {
	syscall.Syscall9(gpConvolutionFilter2DEXT, 7, uintptr(target), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(image), 0, 0)
}
func ConvolutionParameterfEXT(target uint32, pname uint32, params float32) {
	syscall.Syscall(gpConvolutionParameterfEXT, 3, uintptr(target), uintptr(pname), uintptr(math.Float32bits(params)))
}
func ConvolutionParameterfvEXT(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpConvolutionParameterfvEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func ConvolutionParameteriEXT(target uint32, pname uint32, params int32) {
	syscall.Syscall(gpConvolutionParameteriEXT, 3, uintptr(target), uintptr(pname), uintptr(params))
}
func ConvolutionParameterivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpConvolutionParameterivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func ConvolutionParameterxOES(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpConvolutionParameterxOES, 3, uintptr(target), uintptr(pname), uintptr(param))
}
func ConvolutionParameterxvOES(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpConvolutionParameterxvOES, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// copy all or part of the data store of a buffer object to the data store of another buffer object
func CopyBufferSubData(readTarget uint32, writeTarget uint32, readOffset int, writeOffset int, size int) {
	syscall.Syscall6(gpCopyBufferSubData, 5, uintptr(readTarget), uintptr(writeTarget), uintptr(readOffset), uintptr(writeOffset), uintptr(size), 0)
}
func CopyColorSubTableEXT(target uint32, start int32, x int32, y int32, width int32) {
	syscall.Syscall6(gpCopyColorSubTableEXT, 5, uintptr(target), uintptr(start), uintptr(x), uintptr(y), uintptr(width), 0)
}
func CopyColorTableSGI(target uint32, internalformat uint32, x int32, y int32, width int32) {
	syscall.Syscall6(gpCopyColorTableSGI, 5, uintptr(target), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), 0)
}
func CopyConvolutionFilter1DEXT(target uint32, internalformat uint32, x int32, y int32, width int32) {
	syscall.Syscall6(gpCopyConvolutionFilter1DEXT, 5, uintptr(target), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), 0)
}
func CopyConvolutionFilter2DEXT(target uint32, internalformat uint32, x int32, y int32, width int32, height int32) {
	syscall.Syscall6(gpCopyConvolutionFilter2DEXT, 6, uintptr(target), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

// perform a raw data copy between two images
func CopyImageSubData(srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, srcWidth int32, srcHeight int32, srcDepth int32) {
	syscall.Syscall15(gpCopyImageSubData, 15, uintptr(srcName), uintptr(srcTarget), uintptr(srcLevel), uintptr(srcX), uintptr(srcY), uintptr(srcZ), uintptr(dstName), uintptr(dstTarget), uintptr(dstLevel), uintptr(dstX), uintptr(dstY), uintptr(dstZ), uintptr(srcWidth), uintptr(srcHeight), uintptr(srcDepth))
}
func CopyImageSubDataNV(srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, width int32, height int32, depth int32) {
	syscall.Syscall15(gpCopyImageSubDataNV, 15, uintptr(srcName), uintptr(srcTarget), uintptr(srcLevel), uintptr(srcX), uintptr(srcY), uintptr(srcZ), uintptr(dstName), uintptr(dstTarget), uintptr(dstLevel), uintptr(dstX), uintptr(dstY), uintptr(dstZ), uintptr(width), uintptr(height), uintptr(depth))
}
func CopyMultiTexImage1DEXT(texunit uint32, target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32) {
	syscall.Syscall9(gpCopyMultiTexImage1DEXT, 8, uintptr(texunit), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(border), 0)
}
func CopyMultiTexImage2DEXT(texunit uint32, target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32) {
	syscall.Syscall9(gpCopyMultiTexImage2DEXT, 9, uintptr(texunit), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(border))
}
func CopyMultiTexSubImage1DEXT(texunit uint32, target uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	syscall.Syscall9(gpCopyMultiTexSubImage1DEXT, 7, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(x), uintptr(y), uintptr(width), 0, 0)
}
func CopyMultiTexSubImage2DEXT(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpCopyMultiTexSubImage2DEXT, 9, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}
func CopyMultiTexSubImage3DEXT(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall12(gpCopyMultiTexSubImage3DEXT, 10, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0, 0)
}

// copy all or part of the data store of a buffer object to the data store of another buffer object
func CopyNamedBufferSubData(readBuffer uint32, writeBuffer uint32, readOffset int, writeOffset int, size int) {
	syscall.Syscall6(gpCopyNamedBufferSubData, 5, uintptr(readBuffer), uintptr(writeBuffer), uintptr(readOffset), uintptr(writeOffset), uintptr(size), 0)
}
func CopyPathNV(resultPath uint32, srcPath uint32) {
	syscall.Syscall(gpCopyPathNV, 2, uintptr(resultPath), uintptr(srcPath), 0)
}

// copy pixels in the frame buffer
func CopyPixels(x int32, y int32, width int32, height int32, xtype uint32) {
	syscall.Syscall6(gpCopyPixels, 5, uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(xtype), 0)
}

// copy pixels into a 1D texture image
func CopyTexImage1D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32) {
	syscall.Syscall9(gpCopyTexImage1D, 7, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(border), 0, 0)
}
func CopyTexImage1DEXT(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32) {
	syscall.Syscall9(gpCopyTexImage1DEXT, 7, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(border), 0, 0)
}

// copy pixels into a 2D texture image
func CopyTexImage2D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32) {
	syscall.Syscall9(gpCopyTexImage2D, 8, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(border), 0)
}
func CopyTexImage2DEXT(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32) {
	syscall.Syscall9(gpCopyTexImage2DEXT, 8, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(border), 0)
}

// copy a one-dimensional texture subimage
func CopyTexSubImage1D(target uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	syscall.Syscall6(gpCopyTexSubImage1D, 6, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(x), uintptr(y), uintptr(width))
}
func CopyTexSubImage1DEXT(target uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	syscall.Syscall6(gpCopyTexSubImage1DEXT, 6, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(x), uintptr(y), uintptr(width))
}

// copy a two-dimensional texture subimage
func CopyTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpCopyTexSubImage2D, 8, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0)
}
func CopyTexSubImage2DEXT(target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpCopyTexSubImage2DEXT, 8, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0)
}

// copy a three-dimensional texture subimage
func CopyTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpCopyTexSubImage3D, 9, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}
func CopyTexSubImage3DEXT(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpCopyTexSubImage3DEXT, 9, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}
func CopyTextureImage1DEXT(texture uint32, target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32) {
	syscall.Syscall9(gpCopyTextureImage1DEXT, 8, uintptr(texture), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(border), 0)
}
func CopyTextureImage2DEXT(texture uint32, target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32) {
	syscall.Syscall9(gpCopyTextureImage2DEXT, 9, uintptr(texture), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(border))
}

// copy a one-dimensional texture subimage
func CopyTextureSubImage1D(texture uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	syscall.Syscall6(gpCopyTextureSubImage1D, 6, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(x), uintptr(y), uintptr(width))
}
func CopyTextureSubImage1DEXT(texture uint32, target uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	syscall.Syscall9(gpCopyTextureSubImage1DEXT, 7, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(x), uintptr(y), uintptr(width), 0, 0)
}

// copy a two-dimensional texture subimage
func CopyTextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpCopyTextureSubImage2D, 8, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0)
}
func CopyTextureSubImage2DEXT(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpCopyTextureSubImage2DEXT, 9, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

// copy a three-dimensional texture subimage
func CopyTextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpCopyTextureSubImage3D, 9, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}
func CopyTextureSubImage3DEXT(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	syscall.Syscall12(gpCopyTextureSubImage3DEXT, 10, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0, 0)
}
func CoverFillPathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, coverMode uint32, transformType uint32, transformValues *float32) {
	syscall.Syscall9(gpCoverFillPathInstancedNV, 7, uintptr(numPaths), uintptr(pathNameType), uintptr(paths), uintptr(pathBase), uintptr(coverMode), uintptr(transformType), uintptr(unsafe.Pointer(transformValues)), 0, 0)
}
func CoverFillPathNV(path uint32, coverMode uint32) {
	syscall.Syscall(gpCoverFillPathNV, 2, uintptr(path), uintptr(coverMode), 0)
}
func CoverStrokePathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, coverMode uint32, transformType uint32, transformValues *float32) {
	syscall.Syscall9(gpCoverStrokePathInstancedNV, 7, uintptr(numPaths), uintptr(pathNameType), uintptr(paths), uintptr(pathBase), uintptr(coverMode), uintptr(transformType), uintptr(unsafe.Pointer(transformValues)), 0, 0)
}
func CoverStrokePathNV(path uint32, coverMode uint32) {
	syscall.Syscall(gpCoverStrokePathNV, 2, uintptr(path), uintptr(coverMode), 0)
}
func CoverageModulationNV(components uint32) {
	syscall.Syscall(gpCoverageModulationNV, 1, uintptr(components), 0, 0)
}
func CoverageModulationTableNV(n int32, v *float32) {
	syscall.Syscall(gpCoverageModulationTableNV, 2, uintptr(n), uintptr(unsafe.Pointer(v)), 0)
}

// create buffer objects
func CreateBuffers(n int32, buffers *uint32) {
	syscall.Syscall(gpCreateBuffers, 2, uintptr(n), uintptr(unsafe.Pointer(buffers)), 0)
}
func CreateCommandListsNV(n int32, lists *uint32) {
	syscall.Syscall(gpCreateCommandListsNV, 2, uintptr(n), uintptr(unsafe.Pointer(lists)), 0)
}

// create framebuffer objects
func CreateFramebuffers(n int32, framebuffers *uint32) {
	syscall.Syscall(gpCreateFramebuffers, 2, uintptr(n), uintptr(unsafe.Pointer(framebuffers)), 0)
}
func CreateMemoryObjectsEXT(n int32, memoryObjects *uint32) {
	syscall.Syscall(gpCreateMemoryObjectsEXT, 2, uintptr(n), uintptr(unsafe.Pointer(memoryObjects)), 0)
}
func CreatePerfQueryINTEL(queryId uint32, queryHandle *uint32) {
	syscall.Syscall(gpCreatePerfQueryINTEL, 2, uintptr(queryId), uintptr(unsafe.Pointer(queryHandle)), 0)
}

// Creates a program object
func CreateProgram() uint32 {
	ret, _, _ := syscall.Syscall(gpCreateProgram, 0, 0, 0, 0)
	return (uint32)(ret)
}
func CreateProgramObjectARB() uintptr {
	ret, _, _ := syscall.Syscall(gpCreateProgramObjectARB, 0, 0, 0, 0)
	return (uintptr)(ret)
}

// create program pipeline objects
func CreateProgramPipelines(n int32, pipelines *uint32) {
	syscall.Syscall(gpCreateProgramPipelines, 2, uintptr(n), uintptr(unsafe.Pointer(pipelines)), 0)
}

// create query objects
func CreateQueries(target uint32, n int32, ids *uint32) {
	syscall.Syscall(gpCreateQueries, 3, uintptr(target), uintptr(n), uintptr(unsafe.Pointer(ids)))
}

// create renderbuffer objects
func CreateRenderbuffers(n int32, renderbuffers *uint32) {
	syscall.Syscall(gpCreateRenderbuffers, 2, uintptr(n), uintptr(unsafe.Pointer(renderbuffers)), 0)
}

// create sampler objects
func CreateSamplers(n int32, samplers *uint32) {
	syscall.Syscall(gpCreateSamplers, 2, uintptr(n), uintptr(unsafe.Pointer(samplers)), 0)
}

// Creates a shader object
func CreateShader(xtype uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpCreateShader, 1, uintptr(xtype), 0, 0)
	return (uint32)(ret)
}
func CreateShaderObjectARB(shaderType uint32) uintptr {
	ret, _, _ := syscall.Syscall(gpCreateShaderObjectARB, 1, uintptr(shaderType), 0, 0)
	return (uintptr)(ret)
}
func CreateShaderProgramEXT(xtype uint32, xstring *uint8) uint32 {
	ret, _, _ := syscall.Syscall(gpCreateShaderProgramEXT, 2, uintptr(xtype), uintptr(unsafe.Pointer(xstring)), 0)
	return (uint32)(ret)
}

// create a stand-alone program from an array of null-terminated source code strings
func CreateShaderProgramv(xtype uint32, count int32, strings **uint8) uint32 {
	ret, _, _ := syscall.Syscall(gpCreateShaderProgramv, 3, uintptr(xtype), uintptr(count), uintptr(unsafe.Pointer(strings)))
	return (uint32)(ret)
}
func CreateShaderProgramvEXT(xtype uint32, count int32, strings **uint8) uint32 {
	ret, _, _ := syscall.Syscall(gpCreateShaderProgramvEXT, 3, uintptr(xtype), uintptr(count), uintptr(unsafe.Pointer(strings)))
	return (uint32)(ret)
}
func CreateStatesNV(n int32, states *uint32) {
	syscall.Syscall(gpCreateStatesNV, 2, uintptr(n), uintptr(unsafe.Pointer(states)), 0)
}

// Parameter context has type *C.struct__cl_context.
// Parameter event has type *C.struct__cl_event.
func CreateSyncFromCLeventARB(context unsafe.Pointer, event unsafe.Pointer, flags uint32) uintptr {
	ret, _, _ := syscall.Syscall(gpCreateSyncFromCLeventARB, 3, uintptr(context), uintptr(event), uintptr(flags))
	return (uintptr)(ret)
}

// create texture objects
func CreateTextures(target uint32, n int32, textures *uint32) {
	syscall.Syscall(gpCreateTextures, 3, uintptr(target), uintptr(n), uintptr(unsafe.Pointer(textures)))
}

// create transform feedback objects
func CreateTransformFeedbacks(n int32, ids *uint32) {
	syscall.Syscall(gpCreateTransformFeedbacks, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}

// create vertex array objects
func CreateVertexArrays(n int32, arrays *uint32) {
	syscall.Syscall(gpCreateVertexArrays, 2, uintptr(n), uintptr(unsafe.Pointer(arrays)), 0)
}

// specify whether front- or back-facing facets can be culled
func CullFace(mode uint32) {
	syscall.Syscall(gpCullFace, 1, uintptr(mode), 0, 0)
}
func CullParameterdvEXT(pname uint32, params *float64) {
	syscall.Syscall(gpCullParameterdvEXT, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func CullParameterfvEXT(pname uint32, params *float32) {
	syscall.Syscall(gpCullParameterfvEXT, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func CurrentPaletteMatrixARB(index int32) {
	syscall.Syscall(gpCurrentPaletteMatrixARB, 1, uintptr(index), 0, 0)
}

// specify a callback to receive debugging messages from the GL
// Parameter callback has type C.GLDEBUGPROC.
func DebugMessageCallback(callback unsafe.Pointer, userParam unsafe.Pointer) {
	syscall.Syscall(gpDebugMessageCallback, 2, syscall.NewCallback(callback), uintptr(userParam), 0)
}

// Parameter callback has type C.GLDEBUGPROCAMD.
func DebugMessageCallbackAMD(callback unsafe.Pointer, userParam unsafe.Pointer) {
	syscall.Syscall(gpDebugMessageCallbackAMD, 2, uintptr(callback), uintptr(userParam), 0)
}

// Parameter callback has type C.GLDEBUGPROCARB.
func DebugMessageCallbackARB(callback unsafe.Pointer, userParam unsafe.Pointer) {
	syscall.Syscall(gpDebugMessageCallbackARB, 2, syscall.NewCallback(callback), uintptr(userParam), 0)
}

// Parameter callback has type C.GLDEBUGPROCKHR.
func DebugMessageCallbackKHR(callback unsafe.Pointer, userParam unsafe.Pointer) {
	syscall.Syscall(gpDebugMessageCallbackKHR, 2, syscall.NewCallback(callback), uintptr(userParam), 0)
}

// control the reporting of debug messages in a debug context
func DebugMessageControl(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	syscall.Syscall6(gpDebugMessageControl, 6, uintptr(source), uintptr(xtype), uintptr(severity), uintptr(count), uintptr(unsafe.Pointer(ids)), boolToUintptr(enabled))
}
func DebugMessageControlARB(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	syscall.Syscall6(gpDebugMessageControlARB, 6, uintptr(source), uintptr(xtype), uintptr(severity), uintptr(count), uintptr(unsafe.Pointer(ids)), boolToUintptr(enabled))
}
func DebugMessageControlKHR(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	syscall.Syscall6(gpDebugMessageControlKHR, 6, uintptr(source), uintptr(xtype), uintptr(severity), uintptr(count), uintptr(unsafe.Pointer(ids)), boolToUintptr(enabled))
}
func DebugMessageEnableAMD(category uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	syscall.Syscall6(gpDebugMessageEnableAMD, 5, uintptr(category), uintptr(severity), uintptr(count), uintptr(unsafe.Pointer(ids)), boolToUintptr(enabled), 0)
}

// inject an application-supplied message into the debug message queue
func DebugMessageInsert(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	syscall.Syscall6(gpDebugMessageInsert, 6, uintptr(source), uintptr(xtype), uintptr(id), uintptr(severity), uintptr(length), uintptr(unsafe.Pointer(buf)))
}
func DebugMessageInsertAMD(category uint32, severity uint32, id uint32, length int32, buf *uint8) {
	syscall.Syscall6(gpDebugMessageInsertAMD, 5, uintptr(category), uintptr(severity), uintptr(id), uintptr(length), uintptr(unsafe.Pointer(buf)), 0)
}
func DebugMessageInsertARB(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	syscall.Syscall6(gpDebugMessageInsertARB, 6, uintptr(source), uintptr(xtype), uintptr(id), uintptr(severity), uintptr(length), uintptr(unsafe.Pointer(buf)))
}
func DebugMessageInsertKHR(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	syscall.Syscall6(gpDebugMessageInsertKHR, 6, uintptr(source), uintptr(xtype), uintptr(id), uintptr(severity), uintptr(length), uintptr(unsafe.Pointer(buf)))
}
func DeformSGIX(mask uint32) {
	syscall.Syscall(gpDeformSGIX, 1, uintptr(mask), 0, 0)
}
func DeformationMap3dSGIX(target uint32, u1 float64, u2 float64, ustride int32, uorder int32, v1 float64, v2 float64, vstride int32, vorder int32, w1 float64, w2 float64, wstride int32, worder int32, points *float64) {
	syscall.Syscall15(gpDeformationMap3dSGIX, 14, uintptr(target), uintptr(math.Float64bits(u1)), uintptr(math.Float64bits(u2)), uintptr(ustride), uintptr(uorder), uintptr(math.Float64bits(v1)), uintptr(math.Float64bits(v2)), uintptr(vstride), uintptr(vorder), uintptr(math.Float64bits(w1)), uintptr(math.Float64bits(w2)), uintptr(wstride), uintptr(worder), uintptr(unsafe.Pointer(points)), 0)
}
func DeformationMap3fSGIX(target uint32, u1 float32, u2 float32, ustride int32, uorder int32, v1 float32, v2 float32, vstride int32, vorder int32, w1 float32, w2 float32, wstride int32, worder int32, points *float32) {
	syscall.Syscall15(gpDeformationMap3fSGIX, 14, uintptr(target), uintptr(math.Float32bits(u1)), uintptr(math.Float32bits(u2)), uintptr(ustride), uintptr(uorder), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), uintptr(vstride), uintptr(vorder), uintptr(math.Float32bits(w1)), uintptr(math.Float32bits(w2)), uintptr(wstride), uintptr(worder), uintptr(unsafe.Pointer(points)), 0)
}
func DeleteAsyncMarkersSGIX(marker uint32, xrange int32) {
	syscall.Syscall(gpDeleteAsyncMarkersSGIX, 2, uintptr(marker), uintptr(xrange), 0)
}

// delete named buffer objects
func DeleteBuffers(n int32, buffers *uint32) {
	syscall.Syscall(gpDeleteBuffers, 2, uintptr(n), uintptr(unsafe.Pointer(buffers)), 0)
}
func DeleteBuffersARB(n int32, buffers *uint32) {
	syscall.Syscall(gpDeleteBuffersARB, 2, uintptr(n), uintptr(unsafe.Pointer(buffers)), 0)
}
func DeleteCommandListsNV(n int32, lists *uint32) {
	syscall.Syscall(gpDeleteCommandListsNV, 2, uintptr(n), uintptr(unsafe.Pointer(lists)), 0)
}
func DeleteFencesAPPLE(n int32, fences *uint32) {
	syscall.Syscall(gpDeleteFencesAPPLE, 2, uintptr(n), uintptr(unsafe.Pointer(fences)), 0)
}
func DeleteFencesNV(n int32, fences *uint32) {
	syscall.Syscall(gpDeleteFencesNV, 2, uintptr(n), uintptr(unsafe.Pointer(fences)), 0)
}
func DeleteFragmentShaderATI(id uint32) {
	syscall.Syscall(gpDeleteFragmentShaderATI, 1, uintptr(id), 0, 0)
}

// delete framebuffer objects
func DeleteFramebuffers(n int32, framebuffers *uint32) {
	syscall.Syscall(gpDeleteFramebuffers, 2, uintptr(n), uintptr(unsafe.Pointer(framebuffers)), 0)
}
func DeleteFramebuffersEXT(n int32, framebuffers *uint32) {
	syscall.Syscall(gpDeleteFramebuffersEXT, 2, uintptr(n), uintptr(unsafe.Pointer(framebuffers)), 0)
}

// delete a contiguous group of display lists
func DeleteLists(list uint32, xrange int32) {
	syscall.Syscall(gpDeleteLists, 2, uintptr(list), uintptr(xrange), 0)
}
func DeleteMemoryObjectsEXT(n int32, memoryObjects *uint32) {
	syscall.Syscall(gpDeleteMemoryObjectsEXT, 2, uintptr(n), uintptr(unsafe.Pointer(memoryObjects)), 0)
}
func DeleteNamedStringARB(namelen int32, name *uint8) {
	syscall.Syscall(gpDeleteNamedStringARB, 2, uintptr(namelen), uintptr(unsafe.Pointer(name)), 0)
}
func DeleteNamesAMD(identifier uint32, num uint32, names *uint32) {
	syscall.Syscall(gpDeleteNamesAMD, 3, uintptr(identifier), uintptr(num), uintptr(unsafe.Pointer(names)))
}
func DeleteObjectARB(obj uintptr) {
	syscall.Syscall(gpDeleteObjectARB, 1, uintptr(obj), 0, 0)
}
func DeleteOcclusionQueriesNV(n int32, ids *uint32) {
	syscall.Syscall(gpDeleteOcclusionQueriesNV, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}
func DeletePathsNV(path uint32, xrange int32) {
	syscall.Syscall(gpDeletePathsNV, 2, uintptr(path), uintptr(xrange), 0)
}
func DeletePerfMonitorsAMD(n int32, monitors *uint32) {
	syscall.Syscall(gpDeletePerfMonitorsAMD, 2, uintptr(n), uintptr(unsafe.Pointer(monitors)), 0)
}
func DeletePerfQueryINTEL(queryHandle uint32) {
	syscall.Syscall(gpDeletePerfQueryINTEL, 1, uintptr(queryHandle), 0, 0)
}

// Deletes a program object
func DeleteProgram(program uint32) {
	syscall.Syscall(gpDeleteProgram, 1, uintptr(program), 0, 0)
}

// delete program pipeline objects
func DeleteProgramPipelines(n int32, pipelines *uint32) {
	syscall.Syscall(gpDeleteProgramPipelines, 2, uintptr(n), uintptr(unsafe.Pointer(pipelines)), 0)
}
func DeleteProgramPipelinesEXT(n int32, pipelines *uint32) {
	syscall.Syscall(gpDeleteProgramPipelinesEXT, 2, uintptr(n), uintptr(unsafe.Pointer(pipelines)), 0)
}
func DeleteProgramsARB(n int32, programs *uint32) {
	syscall.Syscall(gpDeleteProgramsARB, 2, uintptr(n), uintptr(unsafe.Pointer(programs)), 0)
}
func DeleteProgramsNV(n int32, programs *uint32) {
	syscall.Syscall(gpDeleteProgramsNV, 2, uintptr(n), uintptr(unsafe.Pointer(programs)), 0)
}

// delete named query objects
func DeleteQueries(n int32, ids *uint32) {
	syscall.Syscall(gpDeleteQueries, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}
func DeleteQueriesARB(n int32, ids *uint32) {
	syscall.Syscall(gpDeleteQueriesARB, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}
func DeleteQueryResourceTagNV(n int32, tagIds *int32) {
	syscall.Syscall(gpDeleteQueryResourceTagNV, 2, uintptr(n), uintptr(unsafe.Pointer(tagIds)), 0)
}

// delete renderbuffer objects
func DeleteRenderbuffers(n int32, renderbuffers *uint32) {
	syscall.Syscall(gpDeleteRenderbuffers, 2, uintptr(n), uintptr(unsafe.Pointer(renderbuffers)), 0)
}
func DeleteRenderbuffersEXT(n int32, renderbuffers *uint32) {
	syscall.Syscall(gpDeleteRenderbuffersEXT, 2, uintptr(n), uintptr(unsafe.Pointer(renderbuffers)), 0)
}

// delete named sampler objects
func DeleteSamplers(count int32, samplers *uint32) {
	syscall.Syscall(gpDeleteSamplers, 2, uintptr(count), uintptr(unsafe.Pointer(samplers)), 0)
}
func DeleteSemaphoresEXT(n int32, semaphores *uint32) {
	syscall.Syscall(gpDeleteSemaphoresEXT, 2, uintptr(n), uintptr(unsafe.Pointer(semaphores)), 0)
}

// Deletes a shader object
func DeleteShader(shader uint32) {
	syscall.Syscall(gpDeleteShader, 1, uintptr(shader), 0, 0)
}
func DeleteStatesNV(n int32, states *uint32) {
	syscall.Syscall(gpDeleteStatesNV, 2, uintptr(n), uintptr(unsafe.Pointer(states)), 0)
}

// delete a sync object
func DeleteSync(sync uintptr) {
	syscall.Syscall(gpDeleteSync, 1, uintptr(sync), 0, 0)
}

// delete named textures
func DeleteTextures(n int32, textures *uint32) {
	syscall.Syscall(gpDeleteTextures, 2, uintptr(n), uintptr(unsafe.Pointer(textures)), 0)
}
func DeleteTexturesEXT(n int32, textures *uint32) {
	syscall.Syscall(gpDeleteTexturesEXT, 2, uintptr(n), uintptr(unsafe.Pointer(textures)), 0)
}

// delete transform feedback objects
func DeleteTransformFeedbacks(n int32, ids *uint32) {
	syscall.Syscall(gpDeleteTransformFeedbacks, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}
func DeleteTransformFeedbacksNV(n int32, ids *uint32) {
	syscall.Syscall(gpDeleteTransformFeedbacksNV, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}

// delete vertex array objects
func DeleteVertexArrays(n int32, arrays *uint32) {
	syscall.Syscall(gpDeleteVertexArrays, 2, uintptr(n), uintptr(unsafe.Pointer(arrays)), 0)
}
func DeleteVertexArraysAPPLE(n int32, arrays *uint32) {
	syscall.Syscall(gpDeleteVertexArraysAPPLE, 2, uintptr(n), uintptr(unsafe.Pointer(arrays)), 0)
}
func DeleteVertexShaderEXT(id uint32) {
	syscall.Syscall(gpDeleteVertexShaderEXT, 1, uintptr(id), 0, 0)
}
func DepthBoundsEXT(zmin float64, zmax float64) {
	syscall.Syscall(gpDepthBoundsEXT, 2, uintptr(math.Float64bits(zmin)), uintptr(math.Float64bits(zmax)), 0)
}
func DepthBoundsdNV(zmin float64, zmax float64) {
	syscall.Syscall(gpDepthBoundsdNV, 2, uintptr(math.Float64bits(zmin)), uintptr(math.Float64bits(zmax)), 0)
}

// specify the value used for depth buffer comparisons
func DepthFunc(xfunc uint32) {
	syscall.Syscall(gpDepthFunc, 1, uintptr(xfunc), 0, 0)
}

// enable or disable writing into the depth buffer
func DepthMask(flag bool) {
	syscall.Syscall(gpDepthMask, 1, boolToUintptr(flag), 0, 0)
}

// specify mapping of depth values from normalized device coordinates to window coordinates
func DepthRange(n float64, f float64) {
	syscall.Syscall(gpDepthRange, 2, uintptr(math.Float64bits(n)), uintptr(math.Float64bits(f)), 0)
}
func DepthRangeArrayv(first uint32, count int32, v *float64) {
	syscall.Syscall(gpDepthRangeArrayv, 3, uintptr(first), uintptr(count), uintptr(unsafe.Pointer(v)))
}

// specify mapping of depth values from normalized device coordinates to window coordinates for a specified viewport
func DepthRangeIndexed(index uint32, n float64, f float64) {
	syscall.Syscall(gpDepthRangeIndexed, 3, uintptr(index), uintptr(math.Float64bits(n)), uintptr(math.Float64bits(f)))
}
func DepthRangedNV(zNear float64, zFar float64) {
	syscall.Syscall(gpDepthRangedNV, 2, uintptr(math.Float64bits(zNear)), uintptr(math.Float64bits(zFar)), 0)
}

// specify mapping of depth values from normalized device coordinates to window coordinates
func DepthRangef(n float32, f float32) {
	syscall.Syscall(gpDepthRangef, 2, uintptr(math.Float32bits(n)), uintptr(math.Float32bits(f)), 0)
}
func DepthRangefOES(n float32, f float32) {
	syscall.Syscall(gpDepthRangefOES, 2, uintptr(math.Float32bits(n)), uintptr(math.Float32bits(f)), 0)
}
func DepthRangexOES(n int32, f int32) {
	syscall.Syscall(gpDepthRangexOES, 2, uintptr(n), uintptr(f), 0)
}
func DetachObjectARB(containerObj uintptr, attachedObj uintptr) {
	syscall.Syscall(gpDetachObjectARB, 2, uintptr(containerObj), uintptr(attachedObj), 0)
}

// Detaches a shader object from a program object to which it is attached
func DetachShader(program uint32, shader uint32) {
	syscall.Syscall(gpDetachShader, 2, uintptr(program), uintptr(shader), 0)
}
func DetailTexFuncSGIS(target uint32, n int32, points *float32) {
	syscall.Syscall(gpDetailTexFuncSGIS, 3, uintptr(target), uintptr(n), uintptr(unsafe.Pointer(points)))
}
func Disable(cap uint32) {
	syscall.Syscall(gpDisable, 1, uintptr(cap), 0, 0)
}
func DisableClientState(array uint32) {
	syscall.Syscall(gpDisableClientState, 1, uintptr(array), 0, 0)
}
func DisableClientStateIndexedEXT(array uint32, index uint32) {
	syscall.Syscall(gpDisableClientStateIndexedEXT, 2, uintptr(array), uintptr(index), 0)
}
func DisableClientStateiEXT(array uint32, index uint32) {
	syscall.Syscall(gpDisableClientStateiEXT, 2, uintptr(array), uintptr(index), 0)
}
func DisableIndexedEXT(target uint32, index uint32) {
	syscall.Syscall(gpDisableIndexedEXT, 2, uintptr(target), uintptr(index), 0)
}
func DisableVariantClientStateEXT(id uint32) {
	syscall.Syscall(gpDisableVariantClientStateEXT, 1, uintptr(id), 0, 0)
}

// Enable or disable a generic vertex attribute     array
func DisableVertexArrayAttrib(vaobj uint32, index uint32) {
	syscall.Syscall(gpDisableVertexArrayAttrib, 2, uintptr(vaobj), uintptr(index), 0)
}
func DisableVertexArrayAttribEXT(vaobj uint32, index uint32) {
	syscall.Syscall(gpDisableVertexArrayAttribEXT, 2, uintptr(vaobj), uintptr(index), 0)
}
func DisableVertexArrayEXT(vaobj uint32, array uint32) {
	syscall.Syscall(gpDisableVertexArrayEXT, 2, uintptr(vaobj), uintptr(array), 0)
}
func DisableVertexAttribAPPLE(index uint32, pname uint32) {
	syscall.Syscall(gpDisableVertexAttribAPPLE, 2, uintptr(index), uintptr(pname), 0)
}

// Enable or disable a generic vertex attribute     array
func DisableVertexAttribArray(index uint32) {
	syscall.Syscall(gpDisableVertexAttribArray, 1, uintptr(index), 0, 0)
}
func DisableVertexAttribArrayARB(index uint32) {
	syscall.Syscall(gpDisableVertexAttribArrayARB, 1, uintptr(index), 0, 0)
}

// launch one or more compute work groups
func DispatchCompute(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32) {
	syscall.Syscall(gpDispatchCompute, 3, uintptr(num_groups_x), uintptr(num_groups_y), uintptr(num_groups_z))
}
func DispatchComputeGroupSizeARB(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32, group_size_x uint32, group_size_y uint32, group_size_z uint32) {
	syscall.Syscall6(gpDispatchComputeGroupSizeARB, 6, uintptr(num_groups_x), uintptr(num_groups_y), uintptr(num_groups_z), uintptr(group_size_x), uintptr(group_size_y), uintptr(group_size_z))
}

// launch one or more compute work groups using parameters stored in a buffer
func DispatchComputeIndirect(indirect int) {
	syscall.Syscall(gpDispatchComputeIndirect, 1, uintptr(indirect), 0, 0)
}

// render primitives from array data
func DrawArrays(mode uint32, first int32, count int32) {
	syscall.Syscall(gpDrawArrays, 3, uintptr(mode), uintptr(first), uintptr(count))
}
func DrawArraysEXT(mode uint32, first int32, count int32) {
	syscall.Syscall(gpDrawArraysEXT, 3, uintptr(mode), uintptr(first), uintptr(count))
}

// render primitives from array data, taking parameters from memory
func DrawArraysIndirect(mode uint32, indirect unsafe.Pointer) {
	syscall.Syscall(gpDrawArraysIndirect, 2, uintptr(mode), uintptr(indirect), 0)
}
func DrawArraysInstancedARB(mode uint32, first int32, count int32, primcount int32) {
	syscall.Syscall6(gpDrawArraysInstancedARB, 4, uintptr(mode), uintptr(first), uintptr(count), uintptr(primcount), 0, 0)
}

// draw multiple instances of a range of elements with offset applied to instanced attributes
func DrawArraysInstancedBaseInstance(mode uint32, first int32, count int32, instancecount int32, baseinstance uint32) {
	syscall.Syscall6(gpDrawArraysInstancedBaseInstance, 5, uintptr(mode), uintptr(first), uintptr(count), uintptr(instancecount), uintptr(baseinstance), 0)
}
func DrawArraysInstancedEXT(mode uint32, start int32, count int32, primcount int32) {
	syscall.Syscall6(gpDrawArraysInstancedEXT, 4, uintptr(mode), uintptr(start), uintptr(count), uintptr(primcount), 0, 0)
}

// specify which color buffers are to be drawn into
func DrawBuffer(buf uint32) {
	syscall.Syscall(gpDrawBuffer, 1, uintptr(buf), 0, 0)
}

// Specifies a list of color buffers to be drawn     into
func DrawBuffers(n int32, bufs *uint32) {
	syscall.Syscall(gpDrawBuffers, 2, uintptr(n), uintptr(unsafe.Pointer(bufs)), 0)
}
func DrawBuffersARB(n int32, bufs *uint32) {
	syscall.Syscall(gpDrawBuffersARB, 2, uintptr(n), uintptr(unsafe.Pointer(bufs)), 0)
}
func DrawBuffersATI(n int32, bufs *uint32) {
	syscall.Syscall(gpDrawBuffersATI, 2, uintptr(n), uintptr(unsafe.Pointer(bufs)), 0)
}
func DrawCommandsAddressNV(primitiveMode uint32, indirects *uint64, sizes *int32, count uint32) {
	syscall.Syscall6(gpDrawCommandsAddressNV, 4, uintptr(primitiveMode), uintptr(unsafe.Pointer(indirects)), uintptr(unsafe.Pointer(sizes)), uintptr(count), 0, 0)
}
func DrawCommandsNV(primitiveMode uint32, buffer uint32, indirects *int, sizes *int32, count uint32) {
	syscall.Syscall6(gpDrawCommandsNV, 5, uintptr(primitiveMode), uintptr(buffer), uintptr(unsafe.Pointer(indirects)), uintptr(unsafe.Pointer(sizes)), uintptr(count), 0)
}
func DrawCommandsStatesAddressNV(indirects *uint64, sizes *int32, states *uint32, fbos *uint32, count uint32) {
	syscall.Syscall6(gpDrawCommandsStatesAddressNV, 5, uintptr(unsafe.Pointer(indirects)), uintptr(unsafe.Pointer(sizes)), uintptr(unsafe.Pointer(states)), uintptr(unsafe.Pointer(fbos)), uintptr(count), 0)
}
func DrawCommandsStatesNV(buffer uint32, indirects *int, sizes *int32, states *uint32, fbos *uint32, count uint32) {
	syscall.Syscall6(gpDrawCommandsStatesNV, 6, uintptr(buffer), uintptr(unsafe.Pointer(indirects)), uintptr(unsafe.Pointer(sizes)), uintptr(unsafe.Pointer(states)), uintptr(unsafe.Pointer(fbos)), uintptr(count))
}
func DrawElementArrayAPPLE(mode uint32, first int32, count int32) {
	syscall.Syscall(gpDrawElementArrayAPPLE, 3, uintptr(mode), uintptr(first), uintptr(count))
}
func DrawElementArrayATI(mode uint32, count int32) {
	syscall.Syscall(gpDrawElementArrayATI, 2, uintptr(mode), uintptr(count), 0)
}

// render primitives from array data
func DrawElements(mode uint32, count int32, xtype uint32, indices uintptr) {
	syscall.Syscall6(gpDrawElements, 4, uintptr(mode), uintptr(count), uintptr(xtype), uintptr(indices), 0, 0)
}

// render primitives from array data with a per-element offset
func DrawElementsBaseVertex(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	syscall.Syscall6(gpDrawElementsBaseVertex, 5, uintptr(mode), uintptr(count), uintptr(xtype), uintptr(indices), uintptr(basevertex), 0)
}

// render indexed primitives from array data, taking parameters from memory
func DrawElementsIndirect(mode uint32, xtype uint32, indirect unsafe.Pointer) {
	syscall.Syscall(gpDrawElementsIndirect, 3, uintptr(mode), uintptr(xtype), uintptr(indirect))
}
func DrawElementsInstancedARB(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, primcount int32) {
	syscall.Syscall6(gpDrawElementsInstancedARB, 5, uintptr(mode), uintptr(count), uintptr(xtype), uintptr(indices), uintptr(primcount), 0)
}

// draw multiple instances of a set of elements with offset applied to instanced attributes
func DrawElementsInstancedBaseInstance(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, baseinstance uint32) {
	syscall.Syscall6(gpDrawElementsInstancedBaseInstance, 6, uintptr(mode), uintptr(count), uintptr(xtype), uintptr(indices), uintptr(instancecount), uintptr(baseinstance))
}

// render multiple instances of a set of primitives from array data with a per-element offset
func DrawElementsInstancedBaseVertex(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32) {
	syscall.Syscall6(gpDrawElementsInstancedBaseVertex, 6, uintptr(mode), uintptr(count), uintptr(xtype), uintptr(indices), uintptr(instancecount), uintptr(basevertex))
}

// render multiple instances of a set of primitives from array data with a per-element offset
func DrawElementsInstancedBaseVertexBaseInstance(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32, baseinstance uint32) {
	syscall.Syscall9(gpDrawElementsInstancedBaseVertexBaseInstance, 7, uintptr(mode), uintptr(count), uintptr(xtype), uintptr(indices), uintptr(instancecount), uintptr(basevertex), uintptr(baseinstance), 0, 0)
}
func DrawElementsInstancedEXT(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, primcount int32) {
	syscall.Syscall6(gpDrawElementsInstancedEXT, 5, uintptr(mode), uintptr(count), uintptr(xtype), uintptr(indices), uintptr(primcount), 0)
}
func DrawMeshArraysSUN(mode uint32, first int32, count int32, width int32) {
	syscall.Syscall6(gpDrawMeshArraysSUN, 4, uintptr(mode), uintptr(first), uintptr(count), uintptr(width), 0, 0)
}

// write a block of pixels to the frame buffer
func DrawPixels(width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall6(gpDrawPixels, 5, uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func DrawRangeElementArrayAPPLE(mode uint32, start uint32, end uint32, first int32, count int32) {
	syscall.Syscall6(gpDrawRangeElementArrayAPPLE, 5, uintptr(mode), uintptr(start), uintptr(end), uintptr(first), uintptr(count), 0)
}
func DrawRangeElementArrayATI(mode uint32, start uint32, end uint32, count int32) {
	syscall.Syscall6(gpDrawRangeElementArrayATI, 4, uintptr(mode), uintptr(start), uintptr(end), uintptr(count), 0, 0)
}

// render primitives from array data
func DrawRangeElements(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	syscall.Syscall6(gpDrawRangeElements, 6, uintptr(mode), uintptr(start), uintptr(end), uintptr(count), uintptr(xtype), uintptr(indices))
}

// render primitives from array data with a per-element offset
func DrawRangeElementsBaseVertex(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	syscall.Syscall9(gpDrawRangeElementsBaseVertex, 7, uintptr(mode), uintptr(start), uintptr(end), uintptr(count), uintptr(xtype), uintptr(indices), uintptr(basevertex), 0, 0)
}
func DrawRangeElementsEXT(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	syscall.Syscall6(gpDrawRangeElementsEXT, 6, uintptr(mode), uintptr(start), uintptr(end), uintptr(count), uintptr(xtype), uintptr(indices))
}
func DrawTextureNV(texture uint32, sampler uint32, x0 float32, y0 float32, x1 float32, y1 float32, z float32, s0 float32, t0 float32, s1 float32, t1 float32) {
	syscall.Syscall12(gpDrawTextureNV, 11, uintptr(texture), uintptr(sampler), uintptr(math.Float32bits(x0)), uintptr(math.Float32bits(y0)), uintptr(math.Float32bits(x1)), uintptr(math.Float32bits(y1)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(s0)), uintptr(math.Float32bits(t0)), uintptr(math.Float32bits(s1)), uintptr(math.Float32bits(t1)), 0)
}

// render primitives using a count derived from a transform feedback object
func DrawTransformFeedback(mode uint32, id uint32) {
	syscall.Syscall(gpDrawTransformFeedback, 2, uintptr(mode), uintptr(id), 0)
}

// render multiple instances of primitives using a count derived from a transform feedback object
func DrawTransformFeedbackInstanced(mode uint32, id uint32, instancecount int32) {
	syscall.Syscall(gpDrawTransformFeedbackInstanced, 3, uintptr(mode), uintptr(id), uintptr(instancecount))
}
func DrawTransformFeedbackNV(mode uint32, id uint32) {
	syscall.Syscall(gpDrawTransformFeedbackNV, 2, uintptr(mode), uintptr(id), 0)
}

// render primitives using a count derived from a specifed stream of a transform feedback object
func DrawTransformFeedbackStream(mode uint32, id uint32, stream uint32) {
	syscall.Syscall(gpDrawTransformFeedbackStream, 3, uintptr(mode), uintptr(id), uintptr(stream))
}

// render multiple instances of primitives using a count derived from a specifed stream of a transform feedback object
func DrawTransformFeedbackStreamInstanced(mode uint32, id uint32, stream uint32, instancecount int32) {
	syscall.Syscall6(gpDrawTransformFeedbackStreamInstanced, 4, uintptr(mode), uintptr(id), uintptr(stream), uintptr(instancecount), 0, 0)
}
func DrawVkImageNV(vkImage uint64, sampler uint32, x0 float32, y0 float32, x1 float32, y1 float32, z float32, s0 float32, t0 float32, s1 float32, t1 float32) {
	syscall.Syscall12(gpDrawVkImageNV, 11, uintptr(vkImage), uintptr(sampler), uintptr(math.Float32bits(x0)), uintptr(math.Float32bits(y0)), uintptr(math.Float32bits(x1)), uintptr(math.Float32bits(y1)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(s0)), uintptr(math.Float32bits(t0)), uintptr(math.Float32bits(s1)), uintptr(math.Float32bits(t1)), 0)
}

// Parameter image has type C.GLeglImageOES.
func EGLImageTargetTexStorageEXT(target uint32, image unsafe.Pointer, attrib_list *int32) {
	syscall.Syscall(gpEGLImageTargetTexStorageEXT, 3, uintptr(target), uintptr(image), uintptr(unsafe.Pointer(attrib_list)))
}

// Parameter image has type C.GLeglImageOES.
func EGLImageTargetTextureStorageEXT(texture uint32, image unsafe.Pointer, attrib_list *int32) {
	syscall.Syscall(gpEGLImageTargetTextureStorageEXT, 3, uintptr(texture), uintptr(image), uintptr(unsafe.Pointer(attrib_list)))
}

// flag edges as either boundary or nonboundary
func EdgeFlag(flag bool) {
	syscall.Syscall(gpEdgeFlag, 1, boolToUintptr(flag), 0, 0)
}
func EdgeFlagFormatNV(stride int32) {
	syscall.Syscall(gpEdgeFlagFormatNV, 1, uintptr(stride), 0, 0)
}

// define an array of edge flags
func EdgeFlagPointer(stride int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpEdgeFlagPointer, 2, uintptr(stride), uintptr(pointer), 0)
}
func EdgeFlagPointerEXT(stride int32, count int32, pointer *bool) {
	syscall.Syscall(gpEdgeFlagPointerEXT, 3, uintptr(stride), uintptr(count), uintptr(unsafe.Pointer(pointer)))
}
func EdgeFlagPointerListIBM(stride int32, pointer **bool, ptrstride int32) {
	syscall.Syscall(gpEdgeFlagPointerListIBM, 3, uintptr(stride), uintptr(unsafe.Pointer(pointer)), uintptr(ptrstride))
}
func EdgeFlagv(flag *bool) {
	syscall.Syscall(gpEdgeFlagv, 1, uintptr(unsafe.Pointer(flag)), 0, 0)
}
func ElementPointerAPPLE(xtype uint32, pointer unsafe.Pointer) {
	syscall.Syscall(gpElementPointerAPPLE, 2, uintptr(xtype), uintptr(pointer), 0)
}
func ElementPointerATI(xtype uint32, pointer unsafe.Pointer) {
	syscall.Syscall(gpElementPointerATI, 2, uintptr(xtype), uintptr(pointer), 0)
}

// enable or disable server-side GL capabilities
func Enable(cap uint32) {
	syscall.Syscall(gpEnable, 1, uintptr(cap), 0, 0)
}

// enable or disable client-side capability
func EnableClientState(array uint32) {
	syscall.Syscall(gpEnableClientState, 1, uintptr(array), 0, 0)
}
func EnableClientStateIndexedEXT(array uint32, index uint32) {
	syscall.Syscall(gpEnableClientStateIndexedEXT, 2, uintptr(array), uintptr(index), 0)
}
func EnableClientStateiEXT(array uint32, index uint32) {
	syscall.Syscall(gpEnableClientStateiEXT, 2, uintptr(array), uintptr(index), 0)
}
func EnableIndexedEXT(target uint32, index uint32) {
	syscall.Syscall(gpEnableIndexedEXT, 2, uintptr(target), uintptr(index), 0)
}
func EnableVariantClientStateEXT(id uint32) {
	syscall.Syscall(gpEnableVariantClientStateEXT, 1, uintptr(id), 0, 0)
}

// Enable or disable a generic vertex attribute     array
func EnableVertexArrayAttrib(vaobj uint32, index uint32) {
	syscall.Syscall(gpEnableVertexArrayAttrib, 2, uintptr(vaobj), uintptr(index), 0)
}
func EnableVertexArrayAttribEXT(vaobj uint32, index uint32) {
	syscall.Syscall(gpEnableVertexArrayAttribEXT, 2, uintptr(vaobj), uintptr(index), 0)
}
func EnableVertexArrayEXT(vaobj uint32, array uint32) {
	syscall.Syscall(gpEnableVertexArrayEXT, 2, uintptr(vaobj), uintptr(array), 0)
}
func EnableVertexAttribAPPLE(index uint32, pname uint32) {
	syscall.Syscall(gpEnableVertexAttribAPPLE, 2, uintptr(index), uintptr(pname), 0)
}

// Enable or disable a generic vertex attribute     array
func EnableVertexAttribArray(index uint32) {
	syscall.Syscall(gpEnableVertexAttribArray, 1, uintptr(index), 0, 0)
}
func EnableVertexAttribArrayARB(index uint32) {
	syscall.Syscall(gpEnableVertexAttribArrayARB, 1, uintptr(index), 0, 0)
}
func End() {
	syscall.Syscall(gpEnd, 0, 0, 0, 0)
}
func EndConditionalRenderNV() {
	syscall.Syscall(gpEndConditionalRenderNV, 0, 0, 0, 0)
}
func EndConditionalRenderNVX() {
	syscall.Syscall(gpEndConditionalRenderNVX, 0, 0, 0, 0)
}
func EndFragmentShaderATI() {
	syscall.Syscall(gpEndFragmentShaderATI, 0, 0, 0, 0)
}
func EndList() {
	syscall.Syscall(gpEndList, 0, 0, 0, 0)
}
func EndOcclusionQueryNV() {
	syscall.Syscall(gpEndOcclusionQueryNV, 0, 0, 0, 0)
}
func EndPerfMonitorAMD(monitor uint32) {
	syscall.Syscall(gpEndPerfMonitorAMD, 1, uintptr(monitor), 0, 0)
}
func EndPerfQueryINTEL(queryHandle uint32) {
	syscall.Syscall(gpEndPerfQueryINTEL, 1, uintptr(queryHandle), 0, 0)
}
func EndQuery(target uint32) {
	syscall.Syscall(gpEndQuery, 1, uintptr(target), 0, 0)
}
func EndQueryARB(target uint32) {
	syscall.Syscall(gpEndQueryARB, 1, uintptr(target), 0, 0)
}
func EndQueryIndexed(target uint32, index uint32) {
	syscall.Syscall(gpEndQueryIndexed, 2, uintptr(target), uintptr(index), 0)
}
func EndTransformFeedbackEXT() {
	syscall.Syscall(gpEndTransformFeedbackEXT, 0, 0, 0, 0)
}
func EndTransformFeedbackNV() {
	syscall.Syscall(gpEndTransformFeedbackNV, 0, 0, 0, 0)
}
func EndVertexShaderEXT() {
	syscall.Syscall(gpEndVertexShaderEXT, 0, 0, 0, 0)
}
func EndVideoCaptureNV(video_capture_slot uint32) {
	syscall.Syscall(gpEndVideoCaptureNV, 1, uintptr(video_capture_slot), 0, 0)
}
func EvalCoord1d(u float64) {
	syscall.Syscall(gpEvalCoord1d, 1, uintptr(math.Float64bits(u)), 0, 0)
}
func EvalCoord1dv(u *float64) {
	syscall.Syscall(gpEvalCoord1dv, 1, uintptr(unsafe.Pointer(u)), 0, 0)
}
func EvalCoord1f(u float32) {
	syscall.Syscall(gpEvalCoord1f, 1, uintptr(math.Float32bits(u)), 0, 0)
}
func EvalCoord1fv(u *float32) {
	syscall.Syscall(gpEvalCoord1fv, 1, uintptr(unsafe.Pointer(u)), 0, 0)
}
func EvalCoord1xOES(u int32) {
	syscall.Syscall(gpEvalCoord1xOES, 1, uintptr(u), 0, 0)
}
func EvalCoord1xvOES(coords *int32) {
	syscall.Syscall(gpEvalCoord1xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func EvalCoord2d(u float64, v float64) {
	syscall.Syscall(gpEvalCoord2d, 2, uintptr(math.Float64bits(u)), uintptr(math.Float64bits(v)), 0)
}
func EvalCoord2dv(u *float64) {
	syscall.Syscall(gpEvalCoord2dv, 1, uintptr(unsafe.Pointer(u)), 0, 0)
}
func EvalCoord2f(u float32, v float32) {
	syscall.Syscall(gpEvalCoord2f, 2, uintptr(math.Float32bits(u)), uintptr(math.Float32bits(v)), 0)
}
func EvalCoord2fv(u *float32) {
	syscall.Syscall(gpEvalCoord2fv, 1, uintptr(unsafe.Pointer(u)), 0, 0)
}
func EvalCoord2xOES(u int32, v int32) {
	syscall.Syscall(gpEvalCoord2xOES, 2, uintptr(u), uintptr(v), 0)
}
func EvalCoord2xvOES(coords *int32) {
	syscall.Syscall(gpEvalCoord2xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func EvalMapsNV(target uint32, mode uint32) {
	syscall.Syscall(gpEvalMapsNV, 2, uintptr(target), uintptr(mode), 0)
}
func EvalMesh1(mode uint32, i1 int32, i2 int32) {
	syscall.Syscall(gpEvalMesh1, 3, uintptr(mode), uintptr(i1), uintptr(i2))
}
func EvalMesh2(mode uint32, i1 int32, i2 int32, j1 int32, j2 int32) {
	syscall.Syscall6(gpEvalMesh2, 5, uintptr(mode), uintptr(i1), uintptr(i2), uintptr(j1), uintptr(j2), 0)
}
func EvalPoint1(i int32) {
	syscall.Syscall(gpEvalPoint1, 1, uintptr(i), 0, 0)
}
func EvalPoint2(i int32, j int32) {
	syscall.Syscall(gpEvalPoint2, 2, uintptr(i), uintptr(j), 0)
}
func EvaluateDepthValuesARB() {
	syscall.Syscall(gpEvaluateDepthValuesARB, 0, 0, 0, 0)
}
func ExecuteProgramNV(target uint32, id uint32, params *float32) {
	syscall.Syscall(gpExecuteProgramNV, 3, uintptr(target), uintptr(id), uintptr(unsafe.Pointer(params)))
}
func ExtractComponentEXT(res uint32, src uint32, num uint32) {
	syscall.Syscall(gpExtractComponentEXT, 3, uintptr(res), uintptr(src), uintptr(num))
}

// controls feedback mode
func FeedbackBuffer(size int32, xtype uint32, buffer *float32) {
	syscall.Syscall(gpFeedbackBuffer, 3, uintptr(size), uintptr(xtype), uintptr(unsafe.Pointer(buffer)))
}
func FeedbackBufferxOES(n int32, xtype uint32, buffer *int32) {
	syscall.Syscall(gpFeedbackBufferxOES, 3, uintptr(n), uintptr(xtype), uintptr(unsafe.Pointer(buffer)))
}

// create a new sync object and insert it into the GL command stream
func FenceSync(condition uint32, flags uint32) uintptr {
	ret, _, _ := syscall.Syscall(gpFenceSync, 2, uintptr(condition), uintptr(flags), 0)
	return (uintptr)(ret)
}
func FinalCombinerInputNV(variable uint32, input uint32, mapping uint32, componentUsage uint32) {
	syscall.Syscall6(gpFinalCombinerInputNV, 4, uintptr(variable), uintptr(input), uintptr(mapping), uintptr(componentUsage), 0, 0)
}

// block until all GL execution is complete
func Finish() {
	syscall.Syscall(gpFinish, 0, 0, 0, 0)
}
func FinishAsyncSGIX(markerp *uint32) int32 {
	ret, _, _ := syscall.Syscall(gpFinishAsyncSGIX, 1, uintptr(unsafe.Pointer(markerp)), 0, 0)
	return (int32)(ret)
}
func FinishFenceAPPLE(fence uint32) {
	syscall.Syscall(gpFinishFenceAPPLE, 1, uintptr(fence), 0, 0)
}
func FinishFenceNV(fence uint32) {
	syscall.Syscall(gpFinishFenceNV, 1, uintptr(fence), 0, 0)
}
func FinishObjectAPPLE(object uint32, name int32) {
	syscall.Syscall(gpFinishObjectAPPLE, 2, uintptr(object), uintptr(name), 0)
}
func FinishTextureSUNX() {
	syscall.Syscall(gpFinishTextureSUNX, 0, 0, 0, 0)
}

// force execution of GL commands in finite time
func Flush() {
	syscall.Syscall(gpFlush, 0, 0, 0, 0)
}

// indicate modifications to a range of a mapped buffer
func FlushMappedBufferRange(target uint32, offset int, length int) {
	syscall.Syscall(gpFlushMappedBufferRange, 3, uintptr(target), uintptr(offset), uintptr(length))
}
func FlushMappedBufferRangeAPPLE(target uint32, offset int, size int) {
	syscall.Syscall(gpFlushMappedBufferRangeAPPLE, 3, uintptr(target), uintptr(offset), uintptr(size))
}

// indicate modifications to a range of a mapped buffer
func FlushMappedNamedBufferRange(buffer uint32, offset int, length int) {
	syscall.Syscall(gpFlushMappedNamedBufferRange, 3, uintptr(buffer), uintptr(offset), uintptr(length))
}
func FlushMappedNamedBufferRangeEXT(buffer uint32, offset int, length int) {
	syscall.Syscall(gpFlushMappedNamedBufferRangeEXT, 3, uintptr(buffer), uintptr(offset), uintptr(length))
}
func FlushPixelDataRangeNV(target uint32) {
	syscall.Syscall(gpFlushPixelDataRangeNV, 1, uintptr(target), 0, 0)
}
func FlushRasterSGIX() {
	syscall.Syscall(gpFlushRasterSGIX, 0, 0, 0, 0)
}
func FlushStaticDataIBM(target uint32) {
	syscall.Syscall(gpFlushStaticDataIBM, 1, uintptr(target), 0, 0)
}
func FlushVertexArrayRangeAPPLE(length int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpFlushVertexArrayRangeAPPLE, 2, uintptr(length), uintptr(pointer), 0)
}
func FlushVertexArrayRangeNV() {
	syscall.Syscall(gpFlushVertexArrayRangeNV, 0, 0, 0, 0)
}
func FogCoordFormatNV(xtype uint32, stride int32) {
	syscall.Syscall(gpFogCoordFormatNV, 2, uintptr(xtype), uintptr(stride), 0)
}

// define an array of fog coordinates
func FogCoordPointer(xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpFogCoordPointer, 3, uintptr(xtype), uintptr(stride), uintptr(pointer))
}
func FogCoordPointerEXT(xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpFogCoordPointerEXT, 3, uintptr(xtype), uintptr(stride), uintptr(pointer))
}
func FogCoordPointerListIBM(xtype uint32, stride int32, pointer *unsafe.Pointer, ptrstride int32) {
	syscall.Syscall6(gpFogCoordPointerListIBM, 4, uintptr(xtype), uintptr(stride), uintptr(unsafe.Pointer(pointer)), uintptr(ptrstride), 0, 0)
}
func FogCoordd(coord float64) {
	syscall.Syscall(gpFogCoordd, 1, uintptr(math.Float64bits(coord)), 0, 0)
}
func FogCoorddEXT(coord float64) {
	syscall.Syscall(gpFogCoorddEXT, 1, uintptr(math.Float64bits(coord)), 0, 0)
}
func FogCoorddv(coord *float64) {
	syscall.Syscall(gpFogCoorddv, 1, uintptr(unsafe.Pointer(coord)), 0, 0)
}
func FogCoorddvEXT(coord *float64) {
	syscall.Syscall(gpFogCoorddvEXT, 1, uintptr(unsafe.Pointer(coord)), 0, 0)
}
func FogCoordf(coord float32) {
	syscall.Syscall(gpFogCoordf, 1, uintptr(math.Float32bits(coord)), 0, 0)
}
func FogCoordfEXT(coord float32) {
	syscall.Syscall(gpFogCoordfEXT, 1, uintptr(math.Float32bits(coord)), 0, 0)
}
func FogCoordfv(coord *float32) {
	syscall.Syscall(gpFogCoordfv, 1, uintptr(unsafe.Pointer(coord)), 0, 0)
}
func FogCoordfvEXT(coord *float32) {
	syscall.Syscall(gpFogCoordfvEXT, 1, uintptr(unsafe.Pointer(coord)), 0, 0)
}
func FogCoordhNV(fog uint16) {
	syscall.Syscall(gpFogCoordhNV, 1, uintptr(fog), 0, 0)
}
func FogCoordhvNV(fog *uint16) {
	syscall.Syscall(gpFogCoordhvNV, 1, uintptr(unsafe.Pointer(fog)), 0, 0)
}
func FogFuncSGIS(n int32, points *float32) {
	syscall.Syscall(gpFogFuncSGIS, 2, uintptr(n), uintptr(unsafe.Pointer(points)), 0)
}
func Fogf(pname uint32, param float32) {
	syscall.Syscall(gpFogf, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func Fogfv(pname uint32, params *float32) {
	syscall.Syscall(gpFogfv, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func Fogi(pname uint32, param int32) {
	syscall.Syscall(gpFogi, 2, uintptr(pname), uintptr(param), 0)
}
func Fogiv(pname uint32, params *int32) {
	syscall.Syscall(gpFogiv, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func FogxOES(pname uint32, param int32) {
	syscall.Syscall(gpFogxOES, 2, uintptr(pname), uintptr(param), 0)
}
func FogxvOES(pname uint32, param *int32) {
	syscall.Syscall(gpFogxvOES, 2, uintptr(pname), uintptr(unsafe.Pointer(param)), 0)
}
func FragmentColorMaterialSGIX(face uint32, mode uint32) {
	syscall.Syscall(gpFragmentColorMaterialSGIX, 2, uintptr(face), uintptr(mode), 0)
}
func FragmentCoverageColorNV(color uint32) {
	syscall.Syscall(gpFragmentCoverageColorNV, 1, uintptr(color), 0, 0)
}
func FragmentLightModelfSGIX(pname uint32, param float32) {
	syscall.Syscall(gpFragmentLightModelfSGIX, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func FragmentLightModelfvSGIX(pname uint32, params *float32) {
	syscall.Syscall(gpFragmentLightModelfvSGIX, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func FragmentLightModeliSGIX(pname uint32, param int32) {
	syscall.Syscall(gpFragmentLightModeliSGIX, 2, uintptr(pname), uintptr(param), 0)
}
func FragmentLightModelivSGIX(pname uint32, params *int32) {
	syscall.Syscall(gpFragmentLightModelivSGIX, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func FragmentLightfSGIX(light uint32, pname uint32, param float32) {
	syscall.Syscall(gpFragmentLightfSGIX, 3, uintptr(light), uintptr(pname), uintptr(math.Float32bits(param)))
}
func FragmentLightfvSGIX(light uint32, pname uint32, params *float32) {
	syscall.Syscall(gpFragmentLightfvSGIX, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func FragmentLightiSGIX(light uint32, pname uint32, param int32) {
	syscall.Syscall(gpFragmentLightiSGIX, 3, uintptr(light), uintptr(pname), uintptr(param))
}
func FragmentLightivSGIX(light uint32, pname uint32, params *int32) {
	syscall.Syscall(gpFragmentLightivSGIX, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func FragmentMaterialfSGIX(face uint32, pname uint32, param float32) {
	syscall.Syscall(gpFragmentMaterialfSGIX, 3, uintptr(face), uintptr(pname), uintptr(math.Float32bits(param)))
}
func FragmentMaterialfvSGIX(face uint32, pname uint32, params *float32) {
	syscall.Syscall(gpFragmentMaterialfvSGIX, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func FragmentMaterialiSGIX(face uint32, pname uint32, param int32) {
	syscall.Syscall(gpFragmentMaterialiSGIX, 3, uintptr(face), uintptr(pname), uintptr(param))
}
func FragmentMaterialivSGIX(face uint32, pname uint32, params *int32) {
	syscall.Syscall(gpFragmentMaterialivSGIX, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func FrameTerminatorGREMEDY() {
	syscall.Syscall(gpFrameTerminatorGREMEDY, 0, 0, 0, 0)
}
func FrameZoomSGIX(factor int32) {
	syscall.Syscall(gpFrameZoomSGIX, 1, uintptr(factor), 0, 0)
}
func FramebufferDrawBufferEXT(framebuffer uint32, mode uint32) {
	syscall.Syscall(gpFramebufferDrawBufferEXT, 2, uintptr(framebuffer), uintptr(mode), 0)
}
func FramebufferDrawBuffersEXT(framebuffer uint32, n int32, bufs *uint32) {
	syscall.Syscall(gpFramebufferDrawBuffersEXT, 3, uintptr(framebuffer), uintptr(n), uintptr(unsafe.Pointer(bufs)))
}
func FramebufferFetchBarrierEXT() {
	syscall.Syscall(gpFramebufferFetchBarrierEXT, 0, 0, 0, 0)
}

// set a named parameter of a framebuffer object
func FramebufferParameteri(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpFramebufferParameteri, 3, uintptr(target), uintptr(pname), uintptr(param))
}
func FramebufferReadBufferEXT(framebuffer uint32, mode uint32) {
	syscall.Syscall(gpFramebufferReadBufferEXT, 2, uintptr(framebuffer), uintptr(mode), 0)
}

// attach a renderbuffer as a logical buffer of a framebuffer object
func FramebufferRenderbuffer(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	syscall.Syscall6(gpFramebufferRenderbuffer, 4, uintptr(target), uintptr(attachment), uintptr(renderbuffertarget), uintptr(renderbuffer), 0, 0)
}
func FramebufferRenderbufferEXT(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	syscall.Syscall6(gpFramebufferRenderbufferEXT, 4, uintptr(target), uintptr(attachment), uintptr(renderbuffertarget), uintptr(renderbuffer), 0, 0)
}
func FramebufferSampleLocationsfvARB(target uint32, start uint32, count int32, v *float32) {
	syscall.Syscall6(gpFramebufferSampleLocationsfvARB, 4, uintptr(target), uintptr(start), uintptr(count), uintptr(unsafe.Pointer(v)), 0, 0)
}
func FramebufferSampleLocationsfvNV(target uint32, start uint32, count int32, v *float32) {
	syscall.Syscall6(gpFramebufferSampleLocationsfvNV, 4, uintptr(target), uintptr(start), uintptr(count), uintptr(unsafe.Pointer(v)), 0, 0)
}
func FramebufferSamplePositionsfvAMD(target uint32, numsamples uint32, pixelindex uint32, values *float32) {
	syscall.Syscall6(gpFramebufferSamplePositionsfvAMD, 4, uintptr(target), uintptr(numsamples), uintptr(pixelindex), uintptr(unsafe.Pointer(values)), 0, 0)
}
func FramebufferTexture1D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	syscall.Syscall6(gpFramebufferTexture1D, 5, uintptr(target), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), 0)
}
func FramebufferTexture1DEXT(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	syscall.Syscall6(gpFramebufferTexture1DEXT, 5, uintptr(target), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), 0)
}

// attach a level of a texture object as a logical buffer to the currently bound framebuffer object
func FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	syscall.Syscall6(gpFramebufferTexture2D, 5, uintptr(target), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), 0)
}
func FramebufferTexture2DEXT(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	syscall.Syscall6(gpFramebufferTexture2DEXT, 5, uintptr(target), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), 0)
}
func FramebufferTexture3D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32) {
	syscall.Syscall6(gpFramebufferTexture3D, 6, uintptr(target), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), uintptr(zoffset))
}
func FramebufferTexture3DEXT(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32) {
	syscall.Syscall6(gpFramebufferTexture3DEXT, 6, uintptr(target), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), uintptr(zoffset))
}
func FramebufferTextureARB(target uint32, attachment uint32, texture uint32, level int32) {
	syscall.Syscall6(gpFramebufferTextureARB, 4, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), 0, 0)
}
func FramebufferTextureEXT(target uint32, attachment uint32, texture uint32, level int32) {
	syscall.Syscall6(gpFramebufferTextureEXT, 4, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), 0, 0)
}
func FramebufferTextureFaceARB(target uint32, attachment uint32, texture uint32, level int32, face uint32) {
	syscall.Syscall6(gpFramebufferTextureFaceARB, 5, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(face), 0)
}
func FramebufferTextureFaceEXT(target uint32, attachment uint32, texture uint32, level int32, face uint32) {
	syscall.Syscall6(gpFramebufferTextureFaceEXT, 5, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(face), 0)
}

// attach a single layer of a texture object as a logical buffer of a framebuffer object
func FramebufferTextureLayer(target uint32, attachment uint32, texture uint32, level int32, layer int32) {
	syscall.Syscall6(gpFramebufferTextureLayer, 5, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(layer), 0)
}
func FramebufferTextureLayerARB(target uint32, attachment uint32, texture uint32, level int32, layer int32) {
	syscall.Syscall6(gpFramebufferTextureLayerARB, 5, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(layer), 0)
}
func FramebufferTextureLayerEXT(target uint32, attachment uint32, texture uint32, level int32, layer int32) {
	syscall.Syscall6(gpFramebufferTextureLayerEXT, 5, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(layer), 0)
}
func FramebufferTextureMultiviewOVR(target uint32, attachment uint32, texture uint32, level int32, baseViewIndex int32, numViews int32) {
	syscall.Syscall6(gpFramebufferTextureMultiviewOVR, 6, uintptr(target), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(baseViewIndex), uintptr(numViews))
}
func FreeObjectBufferATI(buffer uint32) {
	syscall.Syscall(gpFreeObjectBufferATI, 1, uintptr(buffer), 0, 0)
}

// define front- and back-facing polygons
func FrontFace(mode uint32) {
	syscall.Syscall(gpFrontFace, 1, uintptr(mode), 0, 0)
}

// multiply the current matrix by a perspective matrix
func Frustum(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	syscall.Syscall6(gpFrustum, 6, uintptr(math.Float64bits(left)), uintptr(math.Float64bits(right)), uintptr(math.Float64bits(bottom)), uintptr(math.Float64bits(top)), uintptr(math.Float64bits(zNear)), uintptr(math.Float64bits(zFar)))
}
func FrustumfOES(l float32, r float32, b float32, t float32, n float32, f float32) {
	syscall.Syscall6(gpFrustumfOES, 6, uintptr(math.Float32bits(l)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(n)), uintptr(math.Float32bits(f)))
}
func FrustumxOES(l int32, r int32, b int32, t int32, n int32, f int32) {
	syscall.Syscall6(gpFrustumxOES, 6, uintptr(l), uintptr(r), uintptr(b), uintptr(t), uintptr(n), uintptr(f))
}
func GenAsyncMarkersSGIX(xrange int32) uint32 {
	ret, _, _ := syscall.Syscall(gpGenAsyncMarkersSGIX, 1, uintptr(xrange), 0, 0)
	return (uint32)(ret)
}

// generate buffer object names
func GenBuffers(n int32, buffers *uint32) {
	syscall.Syscall(gpGenBuffers, 2, uintptr(n), uintptr(unsafe.Pointer(buffers)), 0)
}
func GenBuffersARB(n int32, buffers *uint32) {
	syscall.Syscall(gpGenBuffersARB, 2, uintptr(n), uintptr(unsafe.Pointer(buffers)), 0)
}
func GenFencesAPPLE(n int32, fences *uint32) {
	syscall.Syscall(gpGenFencesAPPLE, 2, uintptr(n), uintptr(unsafe.Pointer(fences)), 0)
}
func GenFencesNV(n int32, fences *uint32) {
	syscall.Syscall(gpGenFencesNV, 2, uintptr(n), uintptr(unsafe.Pointer(fences)), 0)
}
func GenFragmentShadersATI(xrange uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpGenFragmentShadersATI, 1, uintptr(xrange), 0, 0)
	return (uint32)(ret)
}

// generate framebuffer object names
func GenFramebuffers(n int32, framebuffers *uint32) {
	syscall.Syscall(gpGenFramebuffers, 2, uintptr(n), uintptr(unsafe.Pointer(framebuffers)), 0)
}
func GenFramebuffersEXT(n int32, framebuffers *uint32) {
	syscall.Syscall(gpGenFramebuffersEXT, 2, uintptr(n), uintptr(unsafe.Pointer(framebuffers)), 0)
}

// generate a contiguous set of empty display lists
func GenLists(xrange int32) uint32 {
	ret, _, _ := syscall.Syscall(gpGenLists, 1, uintptr(xrange), 0, 0)
	return (uint32)(ret)
}
func GenNamesAMD(identifier uint32, num uint32, names *uint32) {
	syscall.Syscall(gpGenNamesAMD, 3, uintptr(identifier), uintptr(num), uintptr(unsafe.Pointer(names)))
}
func GenOcclusionQueriesNV(n int32, ids *uint32) {
	syscall.Syscall(gpGenOcclusionQueriesNV, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}
func GenPathsNV(xrange int32) uint32 {
	ret, _, _ := syscall.Syscall(gpGenPathsNV, 1, uintptr(xrange), 0, 0)
	return (uint32)(ret)
}
func GenPerfMonitorsAMD(n int32, monitors *uint32) {
	syscall.Syscall(gpGenPerfMonitorsAMD, 2, uintptr(n), uintptr(unsafe.Pointer(monitors)), 0)
}

// reserve program pipeline object names
func GenProgramPipelines(n int32, pipelines *uint32) {
	syscall.Syscall(gpGenProgramPipelines, 2, uintptr(n), uintptr(unsafe.Pointer(pipelines)), 0)
}
func GenProgramPipelinesEXT(n int32, pipelines *uint32) {
	syscall.Syscall(gpGenProgramPipelinesEXT, 2, uintptr(n), uintptr(unsafe.Pointer(pipelines)), 0)
}
func GenProgramsARB(n int32, programs *uint32) {
	syscall.Syscall(gpGenProgramsARB, 2, uintptr(n), uintptr(unsafe.Pointer(programs)), 0)
}
func GenProgramsNV(n int32, programs *uint32) {
	syscall.Syscall(gpGenProgramsNV, 2, uintptr(n), uintptr(unsafe.Pointer(programs)), 0)
}

// generate query object names
func GenQueries(n int32, ids *uint32) {
	syscall.Syscall(gpGenQueries, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}
func GenQueriesARB(n int32, ids *uint32) {
	syscall.Syscall(gpGenQueriesARB, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}
func GenQueryResourceTagNV(n int32, tagIds *int32) {
	syscall.Syscall(gpGenQueryResourceTagNV, 2, uintptr(n), uintptr(unsafe.Pointer(tagIds)), 0)
}

// generate renderbuffer object names
func GenRenderbuffers(n int32, renderbuffers *uint32) {
	syscall.Syscall(gpGenRenderbuffers, 2, uintptr(n), uintptr(unsafe.Pointer(renderbuffers)), 0)
}
func GenRenderbuffersEXT(n int32, renderbuffers *uint32) {
	syscall.Syscall(gpGenRenderbuffersEXT, 2, uintptr(n), uintptr(unsafe.Pointer(renderbuffers)), 0)
}

// generate sampler object names
func GenSamplers(count int32, samplers *uint32) {
	syscall.Syscall(gpGenSamplers, 2, uintptr(count), uintptr(unsafe.Pointer(samplers)), 0)
}
func GenSemaphoresEXT(n int32, semaphores *uint32) {
	syscall.Syscall(gpGenSemaphoresEXT, 2, uintptr(n), uintptr(unsafe.Pointer(semaphores)), 0)
}
func GenSymbolsEXT(datatype uint32, storagetype uint32, xrange uint32, components uint32) uint32 {
	ret, _, _ := syscall.Syscall6(gpGenSymbolsEXT, 4, uintptr(datatype), uintptr(storagetype), uintptr(xrange), uintptr(components), 0, 0)
	return (uint32)(ret)
}

// generate texture names
func GenTextures(n int32, textures *uint32) {
	syscall.Syscall(gpGenTextures, 2, uintptr(n), uintptr(unsafe.Pointer(textures)), 0)
}
func GenTexturesEXT(n int32, textures *uint32) {
	syscall.Syscall(gpGenTexturesEXT, 2, uintptr(n), uintptr(unsafe.Pointer(textures)), 0)
}

// reserve transform feedback object names
func GenTransformFeedbacks(n int32, ids *uint32) {
	syscall.Syscall(gpGenTransformFeedbacks, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}
func GenTransformFeedbacksNV(n int32, ids *uint32) {
	syscall.Syscall(gpGenTransformFeedbacksNV, 2, uintptr(n), uintptr(unsafe.Pointer(ids)), 0)
}

// generate vertex array object names
func GenVertexArrays(n int32, arrays *uint32) {
	syscall.Syscall(gpGenVertexArrays, 2, uintptr(n), uintptr(unsafe.Pointer(arrays)), 0)
}
func GenVertexArraysAPPLE(n int32, arrays *uint32) {
	syscall.Syscall(gpGenVertexArraysAPPLE, 2, uintptr(n), uintptr(unsafe.Pointer(arrays)), 0)
}
func GenVertexShadersEXT(xrange uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpGenVertexShadersEXT, 1, uintptr(xrange), 0, 0)
	return (uint32)(ret)
}

// generate mipmaps for a specified texture object
func GenerateMipmap(target uint32) {
	syscall.Syscall(gpGenerateMipmap, 1, uintptr(target), 0, 0)
}
func GenerateMipmapEXT(target uint32) {
	syscall.Syscall(gpGenerateMipmapEXT, 1, uintptr(target), 0, 0)
}
func GenerateMultiTexMipmapEXT(texunit uint32, target uint32) {
	syscall.Syscall(gpGenerateMultiTexMipmapEXT, 2, uintptr(texunit), uintptr(target), 0)
}

// generate mipmaps for a specified texture object
func GenerateTextureMipmap(texture uint32) {
	syscall.Syscall(gpGenerateTextureMipmap, 1, uintptr(texture), 0, 0)
}
func GenerateTextureMipmapEXT(texture uint32, target uint32) {
	syscall.Syscall(gpGenerateTextureMipmapEXT, 2, uintptr(texture), uintptr(target), 0)
}

// retrieve information about the set of active atomic counter buffers for a program
func GetActiveAtomicCounterBufferiv(program uint32, bufferIndex uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetActiveAtomicCounterBufferiv, 4, uintptr(program), uintptr(bufferIndex), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}

// Returns information about an active attribute variable for the specified program object
func GetActiveAttrib(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	syscall.Syscall9(gpGetActiveAttrib, 7, uintptr(program), uintptr(index), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(xtype)), uintptr(unsafe.Pointer(name)), 0, 0)
}
func GetActiveAttribARB(programObj uintptr, index uint32, maxLength int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	syscall.Syscall9(gpGetActiveAttribARB, 7, uintptr(programObj), uintptr(index), uintptr(maxLength), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(xtype)), uintptr(unsafe.Pointer(name)), 0, 0)
}

// query the name of an active shader subroutine
func GetActiveSubroutineName(program uint32, shadertype uint32, index uint32, bufsize int32, length *int32, name *uint8) {
	syscall.Syscall6(gpGetActiveSubroutineName, 6, uintptr(program), uintptr(shadertype), uintptr(index), uintptr(bufsize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(name)))
}

// query the name of an active shader subroutine uniform
func GetActiveSubroutineUniformName(program uint32, shadertype uint32, index uint32, bufsize int32, length *int32, name *uint8) {
	syscall.Syscall6(gpGetActiveSubroutineUniformName, 6, uintptr(program), uintptr(shadertype), uintptr(index), uintptr(bufsize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(name)))
}
func GetActiveSubroutineUniformiv(program uint32, shadertype uint32, index uint32, pname uint32, values *int32) {
	syscall.Syscall6(gpGetActiveSubroutineUniformiv, 5, uintptr(program), uintptr(shadertype), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(values)), 0)
}

// Returns information about an active uniform variable for the specified program object
func GetActiveUniform(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	syscall.Syscall9(gpGetActiveUniform, 7, uintptr(program), uintptr(index), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(xtype)), uintptr(unsafe.Pointer(name)), 0, 0)
}
func GetActiveUniformARB(programObj uintptr, index uint32, maxLength int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	syscall.Syscall9(gpGetActiveUniformARB, 7, uintptr(programObj), uintptr(index), uintptr(maxLength), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(xtype)), uintptr(unsafe.Pointer(name)), 0, 0)
}

// retrieve the name of an active uniform block
func GetActiveUniformBlockName(program uint32, uniformBlockIndex uint32, bufSize int32, length *int32, uniformBlockName *uint8) {
	syscall.Syscall6(gpGetActiveUniformBlockName, 5, uintptr(program), uintptr(uniformBlockIndex), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(uniformBlockName)), 0)
}

// query information about an active uniform block
func GetActiveUniformBlockiv(program uint32, uniformBlockIndex uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetActiveUniformBlockiv, 4, uintptr(program), uintptr(uniformBlockIndex), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}

// query the name of an active uniform
func GetActiveUniformName(program uint32, uniformIndex uint32, bufSize int32, length *int32, uniformName *uint8) {
	syscall.Syscall6(gpGetActiveUniformName, 5, uintptr(program), uintptr(uniformIndex), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(uniformName)), 0)
}

// Returns information about several active uniform variables for the specified program object
func GetActiveUniformsiv(program uint32, uniformCount int32, uniformIndices *uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetActiveUniformsiv, 5, uintptr(program), uintptr(uniformCount), uintptr(unsafe.Pointer(uniformIndices)), uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetActiveVaryingNV(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	syscall.Syscall9(gpGetActiveVaryingNV, 7, uintptr(program), uintptr(index), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(xtype)), uintptr(unsafe.Pointer(name)), 0, 0)
}
func GetArrayObjectfvATI(array uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetArrayObjectfvATI, 3, uintptr(array), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetArrayObjectivATI(array uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetArrayObjectivATI, 3, uintptr(array), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetAttachedObjectsARB(containerObj uintptr, maxCount int32, count *int32, obj *uintptr) {
	syscall.Syscall6(gpGetAttachedObjectsARB, 4, uintptr(containerObj), uintptr(maxCount), uintptr(unsafe.Pointer(count)), uintptr(unsafe.Pointer(obj)), 0, 0)
}

// Returns the handles of the shader objects attached to a program object
func GetAttachedShaders(program uint32, maxCount int32, count *int32, shaders *uint32) {
	syscall.Syscall6(gpGetAttachedShaders, 4, uintptr(program), uintptr(maxCount), uintptr(unsafe.Pointer(count)), uintptr(unsafe.Pointer(shaders)), 0, 0)
}

// Returns the location of an attribute variable
func GetAttribLocation(program uint32, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetAttribLocation, 2, uintptr(program), uintptr(unsafe.Pointer(name)), 0)
	return (int32)(ret)
}
func GetAttribLocationARB(programObj uintptr, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetAttribLocationARB, 2, uintptr(programObj), uintptr(unsafe.Pointer(name)), 0)
	return (int32)(ret)
}
func GetBooleanIndexedvEXT(target uint32, index uint32, data *bool) {
	syscall.Syscall(gpGetBooleanIndexedvEXT, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetBooleanv(pname uint32, data *bool) {
	syscall.Syscall(gpGetBooleanv, 2, uintptr(pname), uintptr(unsafe.Pointer(data)), 0)
}

// return parameters of a buffer object
func GetBufferParameteriv(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetBufferParameteriv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetBufferParameterivARB(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetBufferParameterivARB, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetBufferParameterui64vNV(target uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpGetBufferParameterui64vNV, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// return the pointer to a mapped buffer object's data store
func GetBufferPointerv(target uint32, pname uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetBufferPointerv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetBufferPointervARB(target uint32, pname uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetBufferPointervARB, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// returns a subset of a buffer object's data store
func GetBufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpGetBufferSubData, 4, uintptr(target), uintptr(offset), uintptr(size), uintptr(data), 0, 0)
}
func GetBufferSubDataARB(target uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpGetBufferSubDataARB, 4, uintptr(target), uintptr(offset), uintptr(size), uintptr(data), 0, 0)
}

// return the coefficients of the specified clipping plane
func GetClipPlane(plane uint32, equation *float64) {
	syscall.Syscall(gpGetClipPlane, 2, uintptr(plane), uintptr(unsafe.Pointer(equation)), 0)
}
func GetClipPlanefOES(plane uint32, equation *float32) {
	syscall.Syscall(gpGetClipPlanefOES, 2, uintptr(plane), uintptr(unsafe.Pointer(equation)), 0)
}
func GetClipPlanexOES(plane uint32, equation *int32) {
	syscall.Syscall(gpGetClipPlanexOES, 2, uintptr(plane), uintptr(unsafe.Pointer(equation)), 0)
}
func GetColorTableEXT(target uint32, format uint32, xtype uint32, data unsafe.Pointer) {
	syscall.Syscall6(gpGetColorTableEXT, 4, uintptr(target), uintptr(format), uintptr(xtype), uintptr(data), 0, 0)
}
func GetColorTableParameterfvEXT(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetColorTableParameterfvEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetColorTableParameterfvSGI(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetColorTableParameterfvSGI, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetColorTableParameterivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetColorTableParameterivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetColorTableParameterivSGI(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetColorTableParameterivSGI, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetColorTableSGI(target uint32, format uint32, xtype uint32, table unsafe.Pointer) {
	syscall.Syscall6(gpGetColorTableSGI, 4, uintptr(target), uintptr(format), uintptr(xtype), uintptr(table), 0, 0)
}
func GetCombinerInputParameterfvNV(stage uint32, portion uint32, variable uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetCombinerInputParameterfvNV, 5, uintptr(stage), uintptr(portion), uintptr(variable), uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetCombinerInputParameterivNV(stage uint32, portion uint32, variable uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetCombinerInputParameterivNV, 5, uintptr(stage), uintptr(portion), uintptr(variable), uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetCombinerOutputParameterfvNV(stage uint32, portion uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetCombinerOutputParameterfvNV, 4, uintptr(stage), uintptr(portion), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetCombinerOutputParameterivNV(stage uint32, portion uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetCombinerOutputParameterivNV, 4, uintptr(stage), uintptr(portion), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetCombinerStageParameterfvNV(stage uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetCombinerStageParameterfvNV, 3, uintptr(stage), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetCommandHeaderNV(tokenID uint32, size uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpGetCommandHeaderNV, 2, uintptr(tokenID), uintptr(size), 0)
	return (uint32)(ret)
}
func GetCompressedMultiTexImageEXT(texunit uint32, target uint32, lod int32, img unsafe.Pointer) {
	syscall.Syscall6(gpGetCompressedMultiTexImageEXT, 4, uintptr(texunit), uintptr(target), uintptr(lod), uintptr(img), 0, 0)
}

// return a compressed texture image
func GetCompressedTexImage(target uint32, level int32, img unsafe.Pointer) {
	syscall.Syscall(gpGetCompressedTexImage, 3, uintptr(target), uintptr(level), uintptr(img))
}
func GetCompressedTexImageARB(target uint32, level int32, img unsafe.Pointer) {
	syscall.Syscall(gpGetCompressedTexImageARB, 3, uintptr(target), uintptr(level), uintptr(img))
}

// return a compressed texture image
func GetCompressedTextureImage(texture uint32, level int32, bufSize int32, pixels unsafe.Pointer) {
	syscall.Syscall6(gpGetCompressedTextureImage, 4, uintptr(texture), uintptr(level), uintptr(bufSize), uintptr(pixels), 0, 0)
}
func GetCompressedTextureImageEXT(texture uint32, target uint32, lod int32, img unsafe.Pointer) {
	syscall.Syscall6(gpGetCompressedTextureImageEXT, 4, uintptr(texture), uintptr(target), uintptr(lod), uintptr(img), 0, 0)
}

// retrieve a sub-region of a compressed texture image from a     compressed texture object
func GetCompressedTextureSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, bufSize int32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpGetCompressedTextureSubImage, 10, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(bufSize), uintptr(pixels), 0, 0)
}
func GetConvolutionFilterEXT(target uint32, format uint32, xtype uint32, image unsafe.Pointer) {
	syscall.Syscall6(gpGetConvolutionFilterEXT, 4, uintptr(target), uintptr(format), uintptr(xtype), uintptr(image), 0, 0)
}
func GetConvolutionParameterfvEXT(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetConvolutionParameterfvEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetConvolutionParameterivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetConvolutionParameterivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetConvolutionParameterxvOES(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetConvolutionParameterxvOES, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetCoverageModulationTableNV(bufsize int32, v *float32) {
	syscall.Syscall(gpGetCoverageModulationTableNV, 2, uintptr(bufsize), uintptr(unsafe.Pointer(v)), 0)
}

// retrieve messages from the debug message log
func GetDebugMessageLog(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret, _, _ := syscall.Syscall9(gpGetDebugMessageLog, 8, uintptr(count), uintptr(bufSize), uintptr(unsafe.Pointer(sources)), uintptr(unsafe.Pointer(types)), uintptr(unsafe.Pointer(ids)), uintptr(unsafe.Pointer(severities)), uintptr(unsafe.Pointer(lengths)), uintptr(unsafe.Pointer(messageLog)), 0)
	return (uint32)(ret)
}
func GetDebugMessageLogAMD(count uint32, bufsize int32, categories *uint32, severities *uint32, ids *uint32, lengths *int32, message *uint8) uint32 {
	ret, _, _ := syscall.Syscall9(gpGetDebugMessageLogAMD, 7, uintptr(count), uintptr(bufsize), uintptr(unsafe.Pointer(categories)), uintptr(unsafe.Pointer(severities)), uintptr(unsafe.Pointer(ids)), uintptr(unsafe.Pointer(lengths)), uintptr(unsafe.Pointer(message)), 0, 0)
	return (uint32)(ret)
}
func GetDebugMessageLogARB(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret, _, _ := syscall.Syscall9(gpGetDebugMessageLogARB, 8, uintptr(count), uintptr(bufSize), uintptr(unsafe.Pointer(sources)), uintptr(unsafe.Pointer(types)), uintptr(unsafe.Pointer(ids)), uintptr(unsafe.Pointer(severities)), uintptr(unsafe.Pointer(lengths)), uintptr(unsafe.Pointer(messageLog)), 0)
	return (uint32)(ret)
}
func GetDebugMessageLogKHR(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret, _, _ := syscall.Syscall9(gpGetDebugMessageLogKHR, 8, uintptr(count), uintptr(bufSize), uintptr(unsafe.Pointer(sources)), uintptr(unsafe.Pointer(types)), uintptr(unsafe.Pointer(ids)), uintptr(unsafe.Pointer(severities)), uintptr(unsafe.Pointer(lengths)), uintptr(unsafe.Pointer(messageLog)), 0)
	return (uint32)(ret)
}
func GetDetailTexFuncSGIS(target uint32, points *float32) {
	syscall.Syscall(gpGetDetailTexFuncSGIS, 2, uintptr(target), uintptr(unsafe.Pointer(points)), 0)
}
func GetDoubleIndexedvEXT(target uint32, index uint32, data *float64) {
	syscall.Syscall(gpGetDoubleIndexedvEXT, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetDoublei_v(target uint32, index uint32, data *float64) {
	syscall.Syscall(gpGetDoublei_v, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetDoublei_vEXT(pname uint32, index uint32, params *float64) {
	syscall.Syscall(gpGetDoublei_vEXT, 3, uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetDoublev(pname uint32, data *float64) {
	syscall.Syscall(gpGetDoublev, 2, uintptr(pname), uintptr(unsafe.Pointer(data)), 0)
}

// return error information
func GetError() uint32 {
	ret, _, _ := syscall.Syscall(gpGetError, 0, 0, 0, 0)
	return (uint32)(ret)
}
func GetFenceivNV(fence uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetFenceivNV, 3, uintptr(fence), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetFinalCombinerInputParameterfvNV(variable uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetFinalCombinerInputParameterfvNV, 3, uintptr(variable), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetFinalCombinerInputParameterivNV(variable uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetFinalCombinerInputParameterivNV, 3, uintptr(variable), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetFirstPerfQueryIdINTEL(queryId *uint32) {
	syscall.Syscall(gpGetFirstPerfQueryIdINTEL, 1, uintptr(unsafe.Pointer(queryId)), 0, 0)
}
func GetFixedvOES(pname uint32, params *int32) {
	syscall.Syscall(gpGetFixedvOES, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetFloatIndexedvEXT(target uint32, index uint32, data *float32) {
	syscall.Syscall(gpGetFloatIndexedvEXT, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetFloati_v(target uint32, index uint32, data *float32) {
	syscall.Syscall(gpGetFloati_v, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetFloati_vEXT(pname uint32, index uint32, params *float32) {
	syscall.Syscall(gpGetFloati_vEXT, 3, uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetFloatv(pname uint32, data *float32) {
	syscall.Syscall(gpGetFloatv, 2, uintptr(pname), uintptr(unsafe.Pointer(data)), 0)
}
func GetFogFuncSGIS(points *float32) {
	syscall.Syscall(gpGetFogFuncSGIS, 1, uintptr(unsafe.Pointer(points)), 0, 0)
}

// query the bindings of color indices to user-defined varying out variables
func GetFragDataIndex(program uint32, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetFragDataIndex, 2, uintptr(program), uintptr(unsafe.Pointer(name)), 0)
	return (int32)(ret)
}
func GetFragDataLocationEXT(program uint32, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetFragDataLocationEXT, 2, uintptr(program), uintptr(unsafe.Pointer(name)), 0)
	return (int32)(ret)
}
func GetFragmentLightfvSGIX(light uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetFragmentLightfvSGIX, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetFragmentLightivSGIX(light uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetFragmentLightivSGIX, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetFragmentMaterialfvSGIX(face uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetFragmentMaterialfvSGIX, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetFragmentMaterialivSGIX(face uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetFragmentMaterialivSGIX, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// retrieve information about attachments of a bound framebuffer object
func GetFramebufferAttachmentParameteriv(target uint32, attachment uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetFramebufferAttachmentParameteriv, 4, uintptr(target), uintptr(attachment), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetFramebufferAttachmentParameterivEXT(target uint32, attachment uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetFramebufferAttachmentParameterivEXT, 4, uintptr(target), uintptr(attachment), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetFramebufferParameterfvAMD(target uint32, pname uint32, numsamples uint32, pixelindex uint32, size int32, values *float32) {
	syscall.Syscall6(gpGetFramebufferParameterfvAMD, 6, uintptr(target), uintptr(pname), uintptr(numsamples), uintptr(pixelindex), uintptr(size), uintptr(unsafe.Pointer(values)))
}

// retrieve a named parameter from a framebuffer
func GetFramebufferParameteriv(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetFramebufferParameteriv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetFramebufferParameterivEXT(framebuffer uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetFramebufferParameterivEXT, 3, uintptr(framebuffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// check if the rendering context has not been lost due to software or hardware issues
func GetGraphicsResetStatus() uint32 {
	ret, _, _ := syscall.Syscall(gpGetGraphicsResetStatus, 0, 0, 0, 0)
	return (uint32)(ret)
}
func GetGraphicsResetStatusARB() uint32 {
	ret, _, _ := syscall.Syscall(gpGetGraphicsResetStatusARB, 0, 0, 0, 0)
	return (uint32)(ret)
}
func GetGraphicsResetStatusKHR() uint32 {
	ret, _, _ := syscall.Syscall(gpGetGraphicsResetStatusKHR, 0, 0, 0, 0)
	return (uint32)(ret)
}
func GetHandleARB(pname uint32) uintptr {
	ret, _, _ := syscall.Syscall(gpGetHandleARB, 1, uintptr(pname), 0, 0)
	return (uintptr)(ret)
}
func GetHistogramEXT(target uint32, reset bool, format uint32, xtype uint32, values unsafe.Pointer) {
	syscall.Syscall6(gpGetHistogramEXT, 5, uintptr(target), boolToUintptr(reset), uintptr(format), uintptr(xtype), uintptr(values), 0)
}
func GetHistogramParameterfvEXT(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetHistogramParameterfvEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetHistogramParameterivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetHistogramParameterivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetHistogramParameterxvOES(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetHistogramParameterxvOES, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetImageHandleARB(texture uint32, level int32, layered bool, layer int32, format uint32) uint64 {
	ret, _, _ := syscall.Syscall6(gpGetImageHandleARB, 5, uintptr(texture), uintptr(level), boolToUintptr(layered), uintptr(layer), uintptr(format), 0)
	return (uint64)(ret)
}
func GetImageHandleNV(texture uint32, level int32, layered bool, layer int32, format uint32) uint64 {
	ret, _, _ := syscall.Syscall6(gpGetImageHandleNV, 5, uintptr(texture), uintptr(level), boolToUintptr(layered), uintptr(layer), uintptr(format), 0)
	return (uint64)(ret)
}
func GetImageTransformParameterfvHP(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetImageTransformParameterfvHP, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetImageTransformParameterivHP(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetImageTransformParameterivHP, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetInfoLogARB(obj uintptr, maxLength int32, length *int32, infoLog *uint8) {
	syscall.Syscall6(gpGetInfoLogARB, 4, uintptr(obj), uintptr(maxLength), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(infoLog)), 0, 0)
}
func GetInstrumentsSGIX() int32 {
	ret, _, _ := syscall.Syscall(gpGetInstrumentsSGIX, 0, 0, 0, 0)
	return (int32)(ret)
}
func GetInteger64v(pname uint32, data *int64) {
	syscall.Syscall(gpGetInteger64v, 2, uintptr(pname), uintptr(unsafe.Pointer(data)), 0)
}
func GetIntegerIndexedvEXT(target uint32, index uint32, data *int32) {
	syscall.Syscall(gpGetIntegerIndexedvEXT, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetIntegeri_v(target uint32, index uint32, data *int32) {
	syscall.Syscall(gpGetIntegeri_v, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetIntegerui64i_vNV(value uint32, index uint32, result *uint64) {
	syscall.Syscall(gpGetIntegerui64i_vNV, 3, uintptr(value), uintptr(index), uintptr(unsafe.Pointer(result)))
}
func GetIntegerui64vNV(value uint32, result *uint64) {
	syscall.Syscall(gpGetIntegerui64vNV, 2, uintptr(value), uintptr(unsafe.Pointer(result)), 0)
}
func GetIntegerv(pname uint32, data *int32) {
	syscall.Syscall(gpGetIntegerv, 2, uintptr(pname), uintptr(unsafe.Pointer(data)), 0)
}
func GetInternalformatSampleivNV(target uint32, internalformat uint32, samples int32, pname uint32, bufSize int32, params *int32) {
	syscall.Syscall6(gpGetInternalformatSampleivNV, 6, uintptr(target), uintptr(internalformat), uintptr(samples), uintptr(pname), uintptr(bufSize), uintptr(unsafe.Pointer(params)))
}
func GetInternalformati64v(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int64) {
	syscall.Syscall6(gpGetInternalformati64v, 5, uintptr(target), uintptr(internalformat), uintptr(pname), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0)
}

// retrieve information about implementation-dependent support for internal formats
func GetInternalformativ(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int32) {
	syscall.Syscall6(gpGetInternalformativ, 5, uintptr(target), uintptr(internalformat), uintptr(pname), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0)
}
func GetInvariantBooleanvEXT(id uint32, value uint32, data *bool) {
	syscall.Syscall(gpGetInvariantBooleanvEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetInvariantFloatvEXT(id uint32, value uint32, data *float32) {
	syscall.Syscall(gpGetInvariantFloatvEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetInvariantIntegervEXT(id uint32, value uint32, data *int32) {
	syscall.Syscall(gpGetInvariantIntegervEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetLightfv(light uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetLightfv, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetLightiv(light uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetLightiv, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetLightxOES(light uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetLightxOES, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetLightxvOES(light uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetLightxvOES, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetListParameterfvSGIX(list uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetListParameterfvSGIX, 3, uintptr(list), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetListParameterivSGIX(list uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetListParameterivSGIX, 3, uintptr(list), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetLocalConstantBooleanvEXT(id uint32, value uint32, data *bool) {
	syscall.Syscall(gpGetLocalConstantBooleanvEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetLocalConstantFloatvEXT(id uint32, value uint32, data *float32) {
	syscall.Syscall(gpGetLocalConstantFloatvEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetLocalConstantIntegervEXT(id uint32, value uint32, data *int32) {
	syscall.Syscall(gpGetLocalConstantIntegervEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetMapAttribParameterfvNV(target uint32, index uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetMapAttribParameterfvNV, 4, uintptr(target), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMapAttribParameterivNV(target uint32, index uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetMapAttribParameterivNV, 4, uintptr(target), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMapControlPointsNV(target uint32, index uint32, xtype uint32, ustride int32, vstride int32, packed bool, points unsafe.Pointer) {
	syscall.Syscall9(gpGetMapControlPointsNV, 7, uintptr(target), uintptr(index), uintptr(xtype), uintptr(ustride), uintptr(vstride), boolToUintptr(packed), uintptr(points), 0, 0)
}
func GetMapParameterfvNV(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetMapParameterfvNV, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetMapParameterivNV(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetMapParameterivNV, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetMapdv(target uint32, query uint32, v *float64) {
	syscall.Syscall(gpGetMapdv, 3, uintptr(target), uintptr(query), uintptr(unsafe.Pointer(v)))
}
func GetMapfv(target uint32, query uint32, v *float32) {
	syscall.Syscall(gpGetMapfv, 3, uintptr(target), uintptr(query), uintptr(unsafe.Pointer(v)))
}
func GetMapiv(target uint32, query uint32, v *int32) {
	syscall.Syscall(gpGetMapiv, 3, uintptr(target), uintptr(query), uintptr(unsafe.Pointer(v)))
}
func GetMapxvOES(target uint32, query uint32, v *int32) {
	syscall.Syscall(gpGetMapxvOES, 3, uintptr(target), uintptr(query), uintptr(unsafe.Pointer(v)))
}
func GetMaterialfv(face uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetMaterialfv, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetMaterialiv(face uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetMaterialiv, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetMaterialxOES(face uint32, pname uint32, param int32) {
	syscall.Syscall(gpGetMaterialxOES, 3, uintptr(face), uintptr(pname), uintptr(param))
}
func GetMaterialxvOES(face uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetMaterialxvOES, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetMemoryObjectParameterivEXT(memoryObject uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetMemoryObjectParameterivEXT, 3, uintptr(memoryObject), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetMinmaxEXT(target uint32, reset bool, format uint32, xtype uint32, values unsafe.Pointer) {
	syscall.Syscall6(gpGetMinmaxEXT, 5, uintptr(target), boolToUintptr(reset), uintptr(format), uintptr(xtype), uintptr(values), 0)
}
func GetMinmaxParameterfvEXT(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetMinmaxParameterfvEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetMinmaxParameterivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetMinmaxParameterivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetMultiTexEnvfvEXT(texunit uint32, target uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetMultiTexEnvfvEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMultiTexEnvivEXT(texunit uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetMultiTexEnvivEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMultiTexGendvEXT(texunit uint32, coord uint32, pname uint32, params *float64) {
	syscall.Syscall6(gpGetMultiTexGendvEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMultiTexGenfvEXT(texunit uint32, coord uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetMultiTexGenfvEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMultiTexGenivEXT(texunit uint32, coord uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetMultiTexGenivEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMultiTexImageEXT(texunit uint32, target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall6(gpGetMultiTexImageEXT, 6, uintptr(texunit), uintptr(target), uintptr(level), uintptr(format), uintptr(xtype), uintptr(pixels))
}
func GetMultiTexLevelParameterfvEXT(texunit uint32, target uint32, level int32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetMultiTexLevelParameterfvEXT, 5, uintptr(texunit), uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetMultiTexLevelParameterivEXT(texunit uint32, target uint32, level int32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetMultiTexLevelParameterivEXT, 5, uintptr(texunit), uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetMultiTexParameterIivEXT(texunit uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetMultiTexParameterIivEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMultiTexParameterIuivEXT(texunit uint32, target uint32, pname uint32, params *uint32) {
	syscall.Syscall6(gpGetMultiTexParameterIuivEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMultiTexParameterfvEXT(texunit uint32, target uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetMultiTexParameterfvEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetMultiTexParameterivEXT(texunit uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetMultiTexParameterivEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}

// retrieve the location of a sample
func GetMultisamplefv(pname uint32, index uint32, val *float32) {
	syscall.Syscall(gpGetMultisamplefv, 3, uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(val)))
}
func GetMultisamplefvNV(pname uint32, index uint32, val *float32) {
	syscall.Syscall(gpGetMultisamplefvNV, 3, uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(val)))
}

// return parameters of a buffer object
func GetNamedBufferParameteri64v(buffer uint32, pname uint32, params *int64) {
	syscall.Syscall(gpGetNamedBufferParameteri64v, 3, uintptr(buffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// return parameters of a buffer object
func GetNamedBufferParameteriv(buffer uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetNamedBufferParameteriv, 3, uintptr(buffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetNamedBufferParameterivEXT(buffer uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetNamedBufferParameterivEXT, 3, uintptr(buffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetNamedBufferParameterui64vNV(buffer uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpGetNamedBufferParameterui64vNV, 3, uintptr(buffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// return the pointer to a mapped buffer object's data store
func GetNamedBufferPointerv(buffer uint32, pname uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetNamedBufferPointerv, 3, uintptr(buffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetNamedBufferPointervEXT(buffer uint32, pname uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetNamedBufferPointervEXT, 3, uintptr(buffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// returns a subset of a buffer object's data store
func GetNamedBufferSubData(buffer uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpGetNamedBufferSubData, 4, uintptr(buffer), uintptr(offset), uintptr(size), uintptr(data), 0, 0)
}
func GetNamedBufferSubDataEXT(buffer uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpGetNamedBufferSubDataEXT, 4, uintptr(buffer), uintptr(offset), uintptr(size), uintptr(data), 0, 0)
}

// retrieve information about attachments of a framebuffer object
func GetNamedFramebufferAttachmentParameteriv(framebuffer uint32, attachment uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetNamedFramebufferAttachmentParameteriv, 4, uintptr(framebuffer), uintptr(attachment), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetNamedFramebufferAttachmentParameterivEXT(framebuffer uint32, attachment uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetNamedFramebufferAttachmentParameterivEXT, 4, uintptr(framebuffer), uintptr(attachment), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetNamedFramebufferParameterfvAMD(framebuffer uint32, pname uint32, numsamples uint32, pixelindex uint32, size int32, values *float32) {
	syscall.Syscall6(gpGetNamedFramebufferParameterfvAMD, 6, uintptr(framebuffer), uintptr(pname), uintptr(numsamples), uintptr(pixelindex), uintptr(size), uintptr(unsafe.Pointer(values)))
}

// query a named parameter of a framebuffer object
func GetNamedFramebufferParameteriv(framebuffer uint32, pname uint32, param *int32) {
	syscall.Syscall(gpGetNamedFramebufferParameteriv, 3, uintptr(framebuffer), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func GetNamedFramebufferParameterivEXT(framebuffer uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetNamedFramebufferParameterivEXT, 3, uintptr(framebuffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetNamedProgramLocalParameterIivEXT(program uint32, target uint32, index uint32, params *int32) {
	syscall.Syscall6(gpGetNamedProgramLocalParameterIivEXT, 4, uintptr(program), uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetNamedProgramLocalParameterIuivEXT(program uint32, target uint32, index uint32, params *uint32) {
	syscall.Syscall6(gpGetNamedProgramLocalParameterIuivEXT, 4, uintptr(program), uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetNamedProgramLocalParameterdvEXT(program uint32, target uint32, index uint32, params *float64) {
	syscall.Syscall6(gpGetNamedProgramLocalParameterdvEXT, 4, uintptr(program), uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetNamedProgramLocalParameterfvEXT(program uint32, target uint32, index uint32, params *float32) {
	syscall.Syscall6(gpGetNamedProgramLocalParameterfvEXT, 4, uintptr(program), uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetNamedProgramStringEXT(program uint32, target uint32, pname uint32, xstring unsafe.Pointer) {
	syscall.Syscall6(gpGetNamedProgramStringEXT, 4, uintptr(program), uintptr(target), uintptr(pname), uintptr(xstring), 0, 0)
}
func GetNamedProgramivEXT(program uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetNamedProgramivEXT, 4, uintptr(program), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}

// query a named parameter of a renderbuffer object
func GetNamedRenderbufferParameteriv(renderbuffer uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetNamedRenderbufferParameteriv, 3, uintptr(renderbuffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetNamedRenderbufferParameterivEXT(renderbuffer uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetNamedRenderbufferParameterivEXT, 3, uintptr(renderbuffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetNamedStringARB(namelen int32, name *uint8, bufSize int32, stringlen *int32, xstring *uint8) {
	syscall.Syscall6(gpGetNamedStringARB, 5, uintptr(namelen), uintptr(unsafe.Pointer(name)), uintptr(bufSize), uintptr(unsafe.Pointer(stringlen)), uintptr(unsafe.Pointer(xstring)), 0)
}
func GetNamedStringivARB(namelen int32, name *uint8, pname uint32, params *int32) {
	syscall.Syscall6(gpGetNamedStringivARB, 4, uintptr(namelen), uintptr(unsafe.Pointer(name)), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetNextPerfQueryIdINTEL(queryId uint32, nextQueryId *uint32) {
	syscall.Syscall(gpGetNextPerfQueryIdINTEL, 2, uintptr(queryId), uintptr(unsafe.Pointer(nextQueryId)), 0)
}
func GetObjectBufferfvATI(buffer uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetObjectBufferfvATI, 3, uintptr(buffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetObjectBufferivATI(buffer uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetObjectBufferivATI, 3, uintptr(buffer), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// retrieve the label of a named object identified within a namespace
func GetObjectLabel(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8) {
	syscall.Syscall6(gpGetObjectLabel, 5, uintptr(identifier), uintptr(name), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(label)), 0)
}
func GetObjectLabelEXT(xtype uint32, object uint32, bufSize int32, length *int32, label *uint8) {
	syscall.Syscall6(gpGetObjectLabelEXT, 5, uintptr(xtype), uintptr(object), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(label)), 0)
}
func GetObjectLabelKHR(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8) {
	syscall.Syscall6(gpGetObjectLabelKHR, 5, uintptr(identifier), uintptr(name), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(label)), 0)
}
func GetObjectParameterfvARB(obj uintptr, pname uint32, params *float32) {
	syscall.Syscall(gpGetObjectParameterfvARB, 3, uintptr(obj), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetObjectParameterivAPPLE(objectType uint32, name uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetObjectParameterivAPPLE, 4, uintptr(objectType), uintptr(name), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetObjectParameterivARB(obj uintptr, pname uint32, params *int32) {
	syscall.Syscall(gpGetObjectParameterivARB, 3, uintptr(obj), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// retrieve the label of a sync object identified by a pointer
func GetObjectPtrLabel(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8) {
	syscall.Syscall6(gpGetObjectPtrLabel, 4, uintptr(ptr), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(label)), 0, 0)
}
func GetObjectPtrLabelKHR(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8) {
	syscall.Syscall6(gpGetObjectPtrLabelKHR, 4, uintptr(ptr), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(label)), 0, 0)
}
func GetOcclusionQueryivNV(id uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetOcclusionQueryivNV, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetOcclusionQueryuivNV(id uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpGetOcclusionQueryuivNV, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetPathCommandsNV(path uint32, commands *uint8) {
	syscall.Syscall(gpGetPathCommandsNV, 2, uintptr(path), uintptr(unsafe.Pointer(commands)), 0)
}
func GetPathCoordsNV(path uint32, coords *float32) {
	syscall.Syscall(gpGetPathCoordsNV, 2, uintptr(path), uintptr(unsafe.Pointer(coords)), 0)
}
func GetPathDashArrayNV(path uint32, dashArray *float32) {
	syscall.Syscall(gpGetPathDashArrayNV, 2, uintptr(path), uintptr(unsafe.Pointer(dashArray)), 0)
}
func GetPathLengthNV(path uint32, startSegment int32, numSegments int32) float32 {
	ret, _, _ := syscall.Syscall(gpGetPathLengthNV, 3, uintptr(path), uintptr(startSegment), uintptr(numSegments))
	return math.Float32frombits(uint32(ret))
}
func GetPathMetricRangeNV(metricQueryMask uint32, firstPathName uint32, numPaths int32, stride int32, metrics *float32) {
	syscall.Syscall6(gpGetPathMetricRangeNV, 5, uintptr(metricQueryMask), uintptr(firstPathName), uintptr(numPaths), uintptr(stride), uintptr(unsafe.Pointer(metrics)), 0)
}
func GetPathMetricsNV(metricQueryMask uint32, numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, stride int32, metrics *float32) {
	syscall.Syscall9(gpGetPathMetricsNV, 7, uintptr(metricQueryMask), uintptr(numPaths), uintptr(pathNameType), uintptr(paths), uintptr(pathBase), uintptr(stride), uintptr(unsafe.Pointer(metrics)), 0, 0)
}
func GetPathParameterfvNV(path uint32, pname uint32, value *float32) {
	syscall.Syscall(gpGetPathParameterfvNV, 3, uintptr(path), uintptr(pname), uintptr(unsafe.Pointer(value)))
}
func GetPathParameterivNV(path uint32, pname uint32, value *int32) {
	syscall.Syscall(gpGetPathParameterivNV, 3, uintptr(path), uintptr(pname), uintptr(unsafe.Pointer(value)))
}
func GetPathSpacingNV(pathListMode uint32, numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, advanceScale float32, kerningScale float32, transformType uint32, returnedSpacing *float32) {
	syscall.Syscall9(gpGetPathSpacingNV, 9, uintptr(pathListMode), uintptr(numPaths), uintptr(pathNameType), uintptr(paths), uintptr(pathBase), uintptr(math.Float32bits(advanceScale)), uintptr(math.Float32bits(kerningScale)), uintptr(transformType), uintptr(unsafe.Pointer(returnedSpacing)))
}
func GetPerfCounterInfoINTEL(queryId uint32, counterId uint32, counterNameLength uint32, counterName *uint8, counterDescLength uint32, counterDesc *uint8, counterOffset *uint32, counterDataSize *uint32, counterTypeEnum *uint32, counterDataTypeEnum *uint32, rawCounterMaxValue *uint64) {
	syscall.Syscall12(gpGetPerfCounterInfoINTEL, 11, uintptr(queryId), uintptr(counterId), uintptr(counterNameLength), uintptr(unsafe.Pointer(counterName)), uintptr(counterDescLength), uintptr(unsafe.Pointer(counterDesc)), uintptr(unsafe.Pointer(counterOffset)), uintptr(unsafe.Pointer(counterDataSize)), uintptr(unsafe.Pointer(counterTypeEnum)), uintptr(unsafe.Pointer(counterDataTypeEnum)), uintptr(unsafe.Pointer(rawCounterMaxValue)), 0)
}
func GetPerfMonitorCounterDataAMD(monitor uint32, pname uint32, dataSize int32, data *uint32, bytesWritten *int32) {
	syscall.Syscall6(gpGetPerfMonitorCounterDataAMD, 5, uintptr(monitor), uintptr(pname), uintptr(dataSize), uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(bytesWritten)), 0)
}
func GetPerfMonitorCounterInfoAMD(group uint32, counter uint32, pname uint32, data unsafe.Pointer) {
	syscall.Syscall6(gpGetPerfMonitorCounterInfoAMD, 4, uintptr(group), uintptr(counter), uintptr(pname), uintptr(data), 0, 0)
}
func GetPerfMonitorCounterStringAMD(group uint32, counter uint32, bufSize int32, length *int32, counterString *uint8) {
	syscall.Syscall6(gpGetPerfMonitorCounterStringAMD, 5, uintptr(group), uintptr(counter), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(counterString)), 0)
}
func GetPerfMonitorCountersAMD(group uint32, numCounters *int32, maxActiveCounters *int32, counterSize int32, counters *uint32) {
	syscall.Syscall6(gpGetPerfMonitorCountersAMD, 5, uintptr(group), uintptr(unsafe.Pointer(numCounters)), uintptr(unsafe.Pointer(maxActiveCounters)), uintptr(counterSize), uintptr(unsafe.Pointer(counters)), 0)
}
func GetPerfMonitorGroupStringAMD(group uint32, bufSize int32, length *int32, groupString *uint8) {
	syscall.Syscall6(gpGetPerfMonitorGroupStringAMD, 4, uintptr(group), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(groupString)), 0, 0)
}
func GetPerfMonitorGroupsAMD(numGroups *int32, groupsSize int32, groups *uint32) {
	syscall.Syscall(gpGetPerfMonitorGroupsAMD, 3, uintptr(unsafe.Pointer(numGroups)), uintptr(groupsSize), uintptr(unsafe.Pointer(groups)))
}
func GetPerfQueryDataINTEL(queryHandle uint32, flags uint32, dataSize int32, data unsafe.Pointer, bytesWritten *uint32) {
	syscall.Syscall6(gpGetPerfQueryDataINTEL, 5, uintptr(queryHandle), uintptr(flags), uintptr(dataSize), uintptr(data), uintptr(unsafe.Pointer(bytesWritten)), 0)
}
func GetPerfQueryIdByNameINTEL(queryName *uint8, queryId *uint32) {
	syscall.Syscall(gpGetPerfQueryIdByNameINTEL, 2, uintptr(unsafe.Pointer(queryName)), uintptr(unsafe.Pointer(queryId)), 0)
}
func GetPerfQueryInfoINTEL(queryId uint32, queryNameLength uint32, queryName *uint8, dataSize *uint32, noCounters *uint32, noInstances *uint32, capsMask *uint32) {
	syscall.Syscall9(gpGetPerfQueryInfoINTEL, 7, uintptr(queryId), uintptr(queryNameLength), uintptr(unsafe.Pointer(queryName)), uintptr(unsafe.Pointer(dataSize)), uintptr(unsafe.Pointer(noCounters)), uintptr(unsafe.Pointer(noInstances)), uintptr(unsafe.Pointer(capsMask)), 0, 0)
}
func GetPixelMapfv(xmap uint32, values *float32) {
	syscall.Syscall(gpGetPixelMapfv, 2, uintptr(xmap), uintptr(unsafe.Pointer(values)), 0)
}
func GetPixelMapuiv(xmap uint32, values *uint32) {
	syscall.Syscall(gpGetPixelMapuiv, 2, uintptr(xmap), uintptr(unsafe.Pointer(values)), 0)
}
func GetPixelMapusv(xmap uint32, values *uint16) {
	syscall.Syscall(gpGetPixelMapusv, 2, uintptr(xmap), uintptr(unsafe.Pointer(values)), 0)
}
func GetPixelMapxv(xmap uint32, size int32, values *int32) {
	syscall.Syscall(gpGetPixelMapxv, 3, uintptr(xmap), uintptr(size), uintptr(unsafe.Pointer(values)))
}
func GetPixelTexGenParameterfvSGIS(pname uint32, params *float32) {
	syscall.Syscall(gpGetPixelTexGenParameterfvSGIS, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetPixelTexGenParameterivSGIS(pname uint32, params *int32) {
	syscall.Syscall(gpGetPixelTexGenParameterivSGIS, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetPixelTransformParameterfvEXT(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetPixelTransformParameterfvEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetPixelTransformParameterivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetPixelTransformParameterivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetPointerIndexedvEXT(target uint32, index uint32, data *unsafe.Pointer) {
	syscall.Syscall(gpGetPointerIndexedvEXT, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetPointeri_vEXT(pname uint32, index uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetPointeri_vEXT, 3, uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(params)))
}

// return the address of the specified pointer
func GetPointerv(pname uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetPointerv, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetPointervEXT(pname uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetPointervEXT, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetPointervKHR(pname uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetPointervKHR, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}

// return the polygon stipple pattern
func GetPolygonStipple(mask *uint8) {
	syscall.Syscall(gpGetPolygonStipple, 1, uintptr(unsafe.Pointer(mask)), 0, 0)
}

// return a binary representation of a program object's compiled and linked executable source
func GetProgramBinary(program uint32, bufSize int32, length *int32, binaryFormat *uint32, binary unsafe.Pointer) {
	syscall.Syscall6(gpGetProgramBinary, 5, uintptr(program), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(binaryFormat)), uintptr(binary), 0)
}
func GetProgramEnvParameterIivNV(target uint32, index uint32, params *int32) {
	syscall.Syscall(gpGetProgramEnvParameterIivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetProgramEnvParameterIuivNV(target uint32, index uint32, params *uint32) {
	syscall.Syscall(gpGetProgramEnvParameterIuivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetProgramEnvParameterdvARB(target uint32, index uint32, params *float64) {
	syscall.Syscall(gpGetProgramEnvParameterdvARB, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetProgramEnvParameterfvARB(target uint32, index uint32, params *float32) {
	syscall.Syscall(gpGetProgramEnvParameterfvARB, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}

// Returns the information log for a program object
func GetProgramInfoLog(program uint32, bufSize int32, length *int32, infoLog *uint8) {
	syscall.Syscall6(gpGetProgramInfoLog, 4, uintptr(program), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(infoLog)), 0, 0)
}
func GetProgramInterfaceiv(program uint32, programInterface uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetProgramInterfaceiv, 4, uintptr(program), uintptr(programInterface), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetProgramLocalParameterIivNV(target uint32, index uint32, params *int32) {
	syscall.Syscall(gpGetProgramLocalParameterIivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetProgramLocalParameterIuivNV(target uint32, index uint32, params *uint32) {
	syscall.Syscall(gpGetProgramLocalParameterIuivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetProgramLocalParameterdvARB(target uint32, index uint32, params *float64) {
	syscall.Syscall(gpGetProgramLocalParameterdvARB, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetProgramLocalParameterfvARB(target uint32, index uint32, params *float32) {
	syscall.Syscall(gpGetProgramLocalParameterfvARB, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func GetProgramNamedParameterdvNV(id uint32, len int32, name *uint8, params *float64) {
	syscall.Syscall6(gpGetProgramNamedParameterdvNV, 4, uintptr(id), uintptr(len), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetProgramNamedParameterfvNV(id uint32, len int32, name *uint8, params *float32) {
	syscall.Syscall6(gpGetProgramNamedParameterfvNV, 4, uintptr(id), uintptr(len), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetProgramParameterdvNV(target uint32, index uint32, pname uint32, params *float64) {
	syscall.Syscall6(gpGetProgramParameterdvNV, 4, uintptr(target), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetProgramParameterfvNV(target uint32, index uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetProgramParameterfvNV, 4, uintptr(target), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}

// retrieve the info log string from a program pipeline object
func GetProgramPipelineInfoLog(pipeline uint32, bufSize int32, length *int32, infoLog *uint8) {
	syscall.Syscall6(gpGetProgramPipelineInfoLog, 4, uintptr(pipeline), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(infoLog)), 0, 0)
}
func GetProgramPipelineInfoLogEXT(pipeline uint32, bufSize int32, length *int32, infoLog *uint8) {
	syscall.Syscall6(gpGetProgramPipelineInfoLogEXT, 4, uintptr(pipeline), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(infoLog)), 0, 0)
}
func GetProgramPipelineiv(pipeline uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetProgramPipelineiv, 3, uintptr(pipeline), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetProgramPipelineivEXT(pipeline uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetProgramPipelineivEXT, 3, uintptr(pipeline), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// query the index of a named resource within a program
func GetProgramResourceIndex(program uint32, programInterface uint32, name *uint8) uint32 {
	ret, _, _ := syscall.Syscall(gpGetProgramResourceIndex, 3, uintptr(program), uintptr(programInterface), uintptr(unsafe.Pointer(name)))
	return (uint32)(ret)
}

// query the location of a named resource within a program
func GetProgramResourceLocation(program uint32, programInterface uint32, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetProgramResourceLocation, 3, uintptr(program), uintptr(programInterface), uintptr(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the fragment color index of a named variable within a program
func GetProgramResourceLocationIndex(program uint32, programInterface uint32, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetProgramResourceLocationIndex, 3, uintptr(program), uintptr(programInterface), uintptr(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the name of an indexed resource within a program
func GetProgramResourceName(program uint32, programInterface uint32, index uint32, bufSize int32, length *int32, name *uint8) {
	syscall.Syscall6(gpGetProgramResourceName, 6, uintptr(program), uintptr(programInterface), uintptr(index), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(name)))
}
func GetProgramResourcefvNV(program uint32, programInterface uint32, index uint32, propCount int32, props *uint32, bufSize int32, length *int32, params *float32) {
	syscall.Syscall9(gpGetProgramResourcefvNV, 8, uintptr(program), uintptr(programInterface), uintptr(index), uintptr(propCount), uintptr(unsafe.Pointer(props)), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(params)), 0)
}
func GetProgramResourceiv(program uint32, programInterface uint32, index uint32, propCount int32, props *uint32, bufSize int32, length *int32, params *int32) {
	syscall.Syscall9(gpGetProgramResourceiv, 8, uintptr(program), uintptr(programInterface), uintptr(index), uintptr(propCount), uintptr(unsafe.Pointer(props)), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(params)), 0)
}
func GetProgramStageiv(program uint32, shadertype uint32, pname uint32, values *int32) {
	syscall.Syscall6(gpGetProgramStageiv, 4, uintptr(program), uintptr(shadertype), uintptr(pname), uintptr(unsafe.Pointer(values)), 0, 0)
}
func GetProgramStringARB(target uint32, pname uint32, xstring unsafe.Pointer) {
	syscall.Syscall(gpGetProgramStringARB, 3, uintptr(target), uintptr(pname), uintptr(xstring))
}
func GetProgramStringNV(id uint32, pname uint32, program *uint8) {
	syscall.Syscall(gpGetProgramStringNV, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(program)))
}
func GetProgramSubroutineParameteruivNV(target uint32, index uint32, param *uint32) {
	syscall.Syscall(gpGetProgramSubroutineParameteruivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(param)))
}

// Returns a parameter from a program object
func GetProgramiv(program uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetProgramiv, 3, uintptr(program), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetProgramivARB(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetProgramivARB, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetProgramivNV(id uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetProgramivNV, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetQueryBufferObjecti64v(id uint32, buffer uint32, pname uint32, offset int) {
	syscall.Syscall6(gpGetQueryBufferObjecti64v, 4, uintptr(id), uintptr(buffer), uintptr(pname), uintptr(offset), 0, 0)
}
func GetQueryBufferObjectiv(id uint32, buffer uint32, pname uint32, offset int) {
	syscall.Syscall6(gpGetQueryBufferObjectiv, 4, uintptr(id), uintptr(buffer), uintptr(pname), uintptr(offset), 0, 0)
}
func GetQueryBufferObjectui64v(id uint32, buffer uint32, pname uint32, offset int) {
	syscall.Syscall6(gpGetQueryBufferObjectui64v, 4, uintptr(id), uintptr(buffer), uintptr(pname), uintptr(offset), 0, 0)
}
func GetQueryBufferObjectuiv(id uint32, buffer uint32, pname uint32, offset int) {
	syscall.Syscall6(gpGetQueryBufferObjectuiv, 4, uintptr(id), uintptr(buffer), uintptr(pname), uintptr(offset), 0, 0)
}

// return parameters of an indexed query object target
func GetQueryIndexediv(target uint32, index uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetQueryIndexediv, 4, uintptr(target), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetQueryObjecti64v(id uint32, pname uint32, params *int64) {
	syscall.Syscall(gpGetQueryObjecti64v, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetQueryObjecti64vEXT(id uint32, pname uint32, params *int64) {
	syscall.Syscall(gpGetQueryObjecti64vEXT, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetQueryObjectiv(id uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetQueryObjectiv, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetQueryObjectivARB(id uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetQueryObjectivARB, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetQueryObjectui64v(id uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpGetQueryObjectui64v, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetQueryObjectui64vEXT(id uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpGetQueryObjectui64vEXT, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// return parameters of a query object
func GetQueryObjectuiv(id uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpGetQueryObjectuiv, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetQueryObjectuivARB(id uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpGetQueryObjectuivARB, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// return parameters of a query object target
func GetQueryiv(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetQueryiv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetQueryivARB(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetQueryivARB, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// retrieve information about a bound renderbuffer object
func GetRenderbufferParameteriv(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetRenderbufferParameteriv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetRenderbufferParameterivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetRenderbufferParameterivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetSamplerParameterIiv(sampler uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetSamplerParameterIiv, 3, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetSamplerParameterIuiv(sampler uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpGetSamplerParameterIuiv, 3, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetSamplerParameterfv(sampler uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetSamplerParameterfv, 3, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetSamplerParameteriv(sampler uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetSamplerParameteriv, 3, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetSemaphoreParameterui64vEXT(semaphore uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpGetSemaphoreParameterui64vEXT, 3, uintptr(semaphore), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetSeparableFilterEXT(target uint32, format uint32, xtype uint32, row unsafe.Pointer, column unsafe.Pointer, span unsafe.Pointer) {
	syscall.Syscall6(gpGetSeparableFilterEXT, 6, uintptr(target), uintptr(format), uintptr(xtype), uintptr(row), uintptr(column), uintptr(span))
}

// Returns the information log for a shader object
func GetShaderInfoLog(shader uint32, bufSize int32, length *int32, infoLog *uint8) {
	syscall.Syscall6(gpGetShaderInfoLog, 4, uintptr(shader), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(infoLog)), 0, 0)
}

// retrieve the range and precision for numeric formats supported by the shader compiler
func GetShaderPrecisionFormat(shadertype uint32, precisiontype uint32, xrange *int32, precision *int32) {
	syscall.Syscall6(gpGetShaderPrecisionFormat, 4, uintptr(shadertype), uintptr(precisiontype), uintptr(unsafe.Pointer(xrange)), uintptr(unsafe.Pointer(precision)), 0, 0)
}

// Returns the source code string from a shader object
func GetShaderSource(shader uint32, bufSize int32, length *int32, source *uint8) {
	syscall.Syscall6(gpGetShaderSource, 4, uintptr(shader), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(source)), 0, 0)
}
func GetShaderSourceARB(obj uintptr, maxLength int32, length *int32, source *uint8) {
	syscall.Syscall6(gpGetShaderSourceARB, 4, uintptr(obj), uintptr(maxLength), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(source)), 0, 0)
}

// Returns a parameter from a shader object
func GetShaderiv(shader uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetShaderiv, 3, uintptr(shader), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetSharpenTexFuncSGIS(target uint32, points *float32) {
	syscall.Syscall(gpGetSharpenTexFuncSGIS, 2, uintptr(target), uintptr(unsafe.Pointer(points)), 0)
}
func GetStageIndexNV(shadertype uint32) uint16 {
	ret, _, _ := syscall.Syscall(gpGetStageIndexNV, 1, uintptr(shadertype), 0, 0)
	return (uint16)(ret)
}

// return a string describing the current GL connection
func GetString(name uint32) *uint8 {
	ret, _, _ := syscall.Syscall(gpGetString, 1, uintptr(name), 0, 0)
	return (*uint8)(unsafe.Pointer(ret))
}

// retrieve the index of a subroutine uniform of a given shader stage within a program
func GetSubroutineIndex(program uint32, shadertype uint32, name *uint8) uint32 {
	ret, _, _ := syscall.Syscall(gpGetSubroutineIndex, 3, uintptr(program), uintptr(shadertype), uintptr(unsafe.Pointer(name)))
	return (uint32)(ret)
}

// retrieve the location of a subroutine uniform of a given shader stage within a program
func GetSubroutineUniformLocation(program uint32, shadertype uint32, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetSubroutineUniformLocation, 3, uintptr(program), uintptr(shadertype), uintptr(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the properties of a sync object
func GetSynciv(sync uintptr, pname uint32, bufSize int32, length *int32, values *int32) {
	syscall.Syscall6(gpGetSynciv, 5, uintptr(sync), uintptr(pname), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(values)), 0)
}
func GetTexBumpParameterfvATI(pname uint32, param *float32) {
	syscall.Syscall(gpGetTexBumpParameterfvATI, 2, uintptr(pname), uintptr(unsafe.Pointer(param)), 0)
}
func GetTexBumpParameterivATI(pname uint32, param *int32) {
	syscall.Syscall(gpGetTexBumpParameterivATI, 2, uintptr(pname), uintptr(unsafe.Pointer(param)), 0)
}
func GetTexEnvfv(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetTexEnvfv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexEnviv(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTexEnviv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexEnvxvOES(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTexEnvxvOES, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexFilterFuncSGIS(target uint32, filter uint32, weights *float32) {
	syscall.Syscall(gpGetTexFilterFuncSGIS, 3, uintptr(target), uintptr(filter), uintptr(unsafe.Pointer(weights)))
}
func GetTexGendv(coord uint32, pname uint32, params *float64) {
	syscall.Syscall(gpGetTexGendv, 3, uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexGenfv(coord uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetTexGenfv, 3, uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexGeniv(coord uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTexGeniv, 3, uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexGenxvOES(coord uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTexGenxvOES, 3, uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// return a texture image
func GetTexImage(target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall6(gpGetTexImage, 5, uintptr(target), uintptr(level), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func GetTexLevelParameterfv(target uint32, level int32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetTexLevelParameterfv, 4, uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTexLevelParameteriv(target uint32, level int32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetTexLevelParameteriv, 4, uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTexLevelParameterxvOES(target uint32, level int32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetTexLevelParameterxvOES, 4, uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTexParameterIivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTexParameterIivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexParameterIuivEXT(target uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpGetTexParameterIuivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexParameterPointervAPPLE(target uint32, pname uint32, params *unsafe.Pointer) {
	syscall.Syscall(gpGetTexParameterPointervAPPLE, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexParameterfv(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetTexParameterfv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexParameteriv(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTexParameteriv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTexParameterxvOES(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTexParameterxvOES, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTextureHandleARB(texture uint32) uint64 {
	ret, _, _ := syscall.Syscall(gpGetTextureHandleARB, 1, uintptr(texture), 0, 0)
	return (uint64)(ret)
}
func GetTextureHandleNV(texture uint32) uint64 {
	ret, _, _ := syscall.Syscall(gpGetTextureHandleNV, 1, uintptr(texture), 0, 0)
	return (uint64)(ret)
}

// return a texture image
func GetTextureImage(texture uint32, level int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer) {
	syscall.Syscall6(gpGetTextureImage, 6, uintptr(texture), uintptr(level), uintptr(format), uintptr(xtype), uintptr(bufSize), uintptr(pixels))
}
func GetTextureImageEXT(texture uint32, target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall6(gpGetTextureImageEXT, 6, uintptr(texture), uintptr(target), uintptr(level), uintptr(format), uintptr(xtype), uintptr(pixels))
}
func GetTextureLevelParameterfv(texture uint32, level int32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetTextureLevelParameterfv, 4, uintptr(texture), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTextureLevelParameterfvEXT(texture uint32, target uint32, level int32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetTextureLevelParameterfvEXT, 5, uintptr(texture), uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetTextureLevelParameteriv(texture uint32, level int32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetTextureLevelParameteriv, 4, uintptr(texture), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTextureLevelParameterivEXT(texture uint32, target uint32, level int32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetTextureLevelParameterivEXT, 5, uintptr(texture), uintptr(target), uintptr(level), uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func GetTextureParameterIiv(texture uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTextureParameterIiv, 3, uintptr(texture), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTextureParameterIivEXT(texture uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetTextureParameterIivEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTextureParameterIuiv(texture uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpGetTextureParameterIuiv, 3, uintptr(texture), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTextureParameterIuivEXT(texture uint32, target uint32, pname uint32, params *uint32) {
	syscall.Syscall6(gpGetTextureParameterIuivEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTextureParameterfv(texture uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetTextureParameterfv, 3, uintptr(texture), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTextureParameterfvEXT(texture uint32, target uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetTextureParameterfvEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTextureParameteriv(texture uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetTextureParameteriv, 3, uintptr(texture), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetTextureParameterivEXT(texture uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetTextureParameterivEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTextureSamplerHandleARB(texture uint32, sampler uint32) uint64 {
	ret, _, _ := syscall.Syscall(gpGetTextureSamplerHandleARB, 2, uintptr(texture), uintptr(sampler), 0)
	return (uint64)(ret)
}
func GetTextureSamplerHandleNV(texture uint32, sampler uint32) uint64 {
	ret, _, _ := syscall.Syscall(gpGetTextureSamplerHandleNV, 2, uintptr(texture), uintptr(sampler), 0)
	return (uint64)(ret)
}

// retrieve a sub-region of a texture image from a texture     object
func GetTextureSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpGetTextureSubImage, 12, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(xtype), uintptr(bufSize), uintptr(pixels))
}
func GetTrackMatrixivNV(target uint32, address uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetTrackMatrixivNV, 4, uintptr(target), uintptr(address), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetTransformFeedbackVaryingEXT(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	syscall.Syscall9(gpGetTransformFeedbackVaryingEXT, 7, uintptr(program), uintptr(index), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(xtype)), uintptr(unsafe.Pointer(name)), 0, 0)
}
func GetTransformFeedbackVaryingNV(program uint32, index uint32, location *int32) {
	syscall.Syscall(gpGetTransformFeedbackVaryingNV, 3, uintptr(program), uintptr(index), uintptr(unsafe.Pointer(location)))
}
func GetTransformFeedbacki64_v(xfb uint32, pname uint32, index uint32, param *int64) {
	syscall.Syscall6(gpGetTransformFeedbacki64_v, 4, uintptr(xfb), uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(param)), 0, 0)
}
func GetTransformFeedbacki_v(xfb uint32, pname uint32, index uint32, param *int32) {
	syscall.Syscall6(gpGetTransformFeedbacki_v, 4, uintptr(xfb), uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(param)), 0, 0)
}

// query the state of a transform feedback object.
func GetTransformFeedbackiv(xfb uint32, pname uint32, param *int32) {
	syscall.Syscall(gpGetTransformFeedbackiv, 3, uintptr(xfb), uintptr(pname), uintptr(unsafe.Pointer(param)))
}

// retrieve the index of a named uniform block
func GetUniformBlockIndex(program uint32, uniformBlockName *uint8) uint32 {
	ret, _, _ := syscall.Syscall(gpGetUniformBlockIndex, 2, uintptr(program), uintptr(unsafe.Pointer(uniformBlockName)), 0)
	return (uint32)(ret)
}
func GetUniformBufferSizeEXT(program uint32, location int32) int32 {
	ret, _, _ := syscall.Syscall(gpGetUniformBufferSizeEXT, 2, uintptr(program), uintptr(location), 0)
	return (int32)(ret)
}

// retrieve the index of a named uniform block
func GetUniformIndices(program uint32, uniformCount int32, uniformNames **uint8, uniformIndices *uint32) {
	syscall.Syscall6(gpGetUniformIndices, 4, uintptr(program), uintptr(uniformCount), uintptr(unsafe.Pointer(uniformNames)), uintptr(unsafe.Pointer(uniformIndices)), 0, 0)
}

// Returns the location of a uniform variable
func GetUniformLocation(program uint32, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetUniformLocation, 2, uintptr(program), uintptr(unsafe.Pointer(name)), 0)
	return (int32)(ret)
}
func GetUniformLocationARB(programObj uintptr, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetUniformLocationARB, 2, uintptr(programObj), uintptr(unsafe.Pointer(name)), 0)
	return (int32)(ret)
}
func GetUniformOffsetEXT(program uint32, location int32) int {
	ret, _, _ := syscall.Syscall(gpGetUniformOffsetEXT, 2, uintptr(program), uintptr(location), 0)
	return (int)(ret)
}
func GetUniformSubroutineuiv(shadertype uint32, location int32, params *uint32) {
	syscall.Syscall(gpGetUniformSubroutineuiv, 3, uintptr(shadertype), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUniformdv(program uint32, location int32, params *float64) {
	syscall.Syscall(gpGetUniformdv, 3, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(params)))
}

// Returns the value of a uniform variable
func GetUniformfv(program uint32, location int32, params *float32) {
	syscall.Syscall(gpGetUniformfv, 3, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUniformfvARB(programObj uintptr, location int32, params *float32) {
	syscall.Syscall(gpGetUniformfvARB, 3, uintptr(programObj), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUniformi64vARB(program uint32, location int32, params *int64) {
	syscall.Syscall(gpGetUniformi64vARB, 3, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUniformi64vNV(program uint32, location int32, params *int64) {
	syscall.Syscall(gpGetUniformi64vNV, 3, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(params)))
}

// Returns the value of a uniform variable
func GetUniformiv(program uint32, location int32, params *int32) {
	syscall.Syscall(gpGetUniformiv, 3, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUniformivARB(programObj uintptr, location int32, params *int32) {
	syscall.Syscall(gpGetUniformivARB, 3, uintptr(programObj), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUniformui64vARB(program uint32, location int32, params *uint64) {
	syscall.Syscall(gpGetUniformui64vARB, 3, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUniformui64vNV(program uint32, location int32, params *uint64) {
	syscall.Syscall(gpGetUniformui64vNV, 3, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUniformuivEXT(program uint32, location int32, params *uint32) {
	syscall.Syscall(gpGetUniformuivEXT, 3, uintptr(program), uintptr(location), uintptr(unsafe.Pointer(params)))
}
func GetUnsignedBytei_vEXT(target uint32, index uint32, data *uint8) {
	syscall.Syscall(gpGetUnsignedBytei_vEXT, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(data)))
}
func GetUnsignedBytevEXT(pname uint32, data *uint8) {
	syscall.Syscall(gpGetUnsignedBytevEXT, 2, uintptr(pname), uintptr(unsafe.Pointer(data)), 0)
}
func GetVariantArrayObjectfvATI(id uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetVariantArrayObjectfvATI, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVariantArrayObjectivATI(id uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetVariantArrayObjectivATI, 3, uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVariantBooleanvEXT(id uint32, value uint32, data *bool) {
	syscall.Syscall(gpGetVariantBooleanvEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetVariantFloatvEXT(id uint32, value uint32, data *float32) {
	syscall.Syscall(gpGetVariantFloatvEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetVariantIntegervEXT(id uint32, value uint32, data *int32) {
	syscall.Syscall(gpGetVariantIntegervEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetVariantPointervEXT(id uint32, value uint32, data *unsafe.Pointer) {
	syscall.Syscall(gpGetVariantPointervEXT, 3, uintptr(id), uintptr(value), uintptr(unsafe.Pointer(data)))
}
func GetVaryingLocationNV(program uint32, name *uint8) int32 {
	ret, _, _ := syscall.Syscall(gpGetVaryingLocationNV, 2, uintptr(program), uintptr(unsafe.Pointer(name)), 0)
	return (int32)(ret)
}
func GetVertexArrayIndexed64iv(vaobj uint32, index uint32, pname uint32, param *int64) {
	syscall.Syscall6(gpGetVertexArrayIndexed64iv, 4, uintptr(vaobj), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(param)), 0, 0)
}
func GetVertexArrayIndexediv(vaobj uint32, index uint32, pname uint32, param *int32) {
	syscall.Syscall6(gpGetVertexArrayIndexediv, 4, uintptr(vaobj), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(param)), 0, 0)
}
func GetVertexArrayIntegeri_vEXT(vaobj uint32, index uint32, pname uint32, param *int32) {
	syscall.Syscall6(gpGetVertexArrayIntegeri_vEXT, 4, uintptr(vaobj), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(param)), 0, 0)
}
func GetVertexArrayIntegervEXT(vaobj uint32, pname uint32, param *int32) {
	syscall.Syscall(gpGetVertexArrayIntegervEXT, 3, uintptr(vaobj), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func GetVertexArrayPointeri_vEXT(vaobj uint32, index uint32, pname uint32, param *unsafe.Pointer) {
	syscall.Syscall6(gpGetVertexArrayPointeri_vEXT, 4, uintptr(vaobj), uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(param)), 0, 0)
}
func GetVertexArrayPointervEXT(vaobj uint32, pname uint32, param *unsafe.Pointer) {
	syscall.Syscall(gpGetVertexArrayPointervEXT, 3, uintptr(vaobj), uintptr(pname), uintptr(unsafe.Pointer(param)))
}

// retrieve parameters of a vertex array object
func GetVertexArrayiv(vaobj uint32, pname uint32, param *int32) {
	syscall.Syscall(gpGetVertexArrayiv, 3, uintptr(vaobj), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func GetVertexAttribArrayObjectfvATI(index uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetVertexAttribArrayObjectfvATI, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribArrayObjectivATI(index uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetVertexAttribArrayObjectivATI, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribIivEXT(index uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetVertexAttribIivEXT, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribIuivEXT(index uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpGetVertexAttribIuivEXT, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribLdv(index uint32, pname uint32, params *float64) {
	syscall.Syscall(gpGetVertexAttribLdv, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribLdvEXT(index uint32, pname uint32, params *float64) {
	syscall.Syscall(gpGetVertexAttribLdvEXT, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribLi64vNV(index uint32, pname uint32, params *int64) {
	syscall.Syscall(gpGetVertexAttribLi64vNV, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribLui64vARB(index uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpGetVertexAttribLui64vARB, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribLui64vNV(index uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpGetVertexAttribLui64vNV, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// return the address of the specified generic vertex attribute pointer
func GetVertexAttribPointerv(index uint32, pname uint32, pointer *unsafe.Pointer) {
	syscall.Syscall(gpGetVertexAttribPointerv, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(pointer)))
}
func GetVertexAttribPointervARB(index uint32, pname uint32, pointer *unsafe.Pointer) {
	syscall.Syscall(gpGetVertexAttribPointervARB, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(pointer)))
}
func GetVertexAttribPointervNV(index uint32, pname uint32, pointer *unsafe.Pointer) {
	syscall.Syscall(gpGetVertexAttribPointervNV, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(pointer)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribdv(index uint32, pname uint32, params *float64) {
	syscall.Syscall(gpGetVertexAttribdv, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribdvARB(index uint32, pname uint32, params *float64) {
	syscall.Syscall(gpGetVertexAttribdvARB, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribdvNV(index uint32, pname uint32, params *float64) {
	syscall.Syscall(gpGetVertexAttribdvNV, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribfv(index uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetVertexAttribfv, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribfvARB(index uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetVertexAttribfvARB, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribfvNV(index uint32, pname uint32, params *float32) {
	syscall.Syscall(gpGetVertexAttribfvNV, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribiv(index uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetVertexAttribiv, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribivARB(index uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetVertexAttribivARB, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVertexAttribivNV(index uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetVertexAttribivNV, 3, uintptr(index), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVideoCaptureStreamdvNV(video_capture_slot uint32, stream uint32, pname uint32, params *float64) {
	syscall.Syscall6(gpGetVideoCaptureStreamdvNV, 4, uintptr(video_capture_slot), uintptr(stream), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetVideoCaptureStreamfvNV(video_capture_slot uint32, stream uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpGetVideoCaptureStreamfvNV, 4, uintptr(video_capture_slot), uintptr(stream), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetVideoCaptureStreamivNV(video_capture_slot uint32, stream uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpGetVideoCaptureStreamivNV, 4, uintptr(video_capture_slot), uintptr(stream), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetVideoCaptureivNV(video_capture_slot uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetVideoCaptureivNV, 3, uintptr(video_capture_slot), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVideoi64vNV(video_slot uint32, pname uint32, params *int64) {
	syscall.Syscall(gpGetVideoi64vNV, 3, uintptr(video_slot), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVideoivNV(video_slot uint32, pname uint32, params *int32) {
	syscall.Syscall(gpGetVideoivNV, 3, uintptr(video_slot), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVideoui64vNV(video_slot uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpGetVideoui64vNV, 3, uintptr(video_slot), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func GetVideouivNV(video_slot uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpGetVideouivNV, 3, uintptr(video_slot), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// Return value has type C.GLVULKANPROCNV.
func GetVkProcAddrNV(name *uint8) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(gpGetVkProcAddrNV, 1, uintptr(unsafe.Pointer(name)), 0, 0)
	return (unsafe.Pointer)(ret)
}
func GetnCompressedTexImageARB(target uint32, lod int32, bufSize int32, img unsafe.Pointer) {
	syscall.Syscall6(gpGetnCompressedTexImageARB, 4, uintptr(target), uintptr(lod), uintptr(bufSize), uintptr(img), 0, 0)
}
func GetnTexImageARB(target uint32, level int32, format uint32, xtype uint32, bufSize int32, img unsafe.Pointer) {
	syscall.Syscall6(gpGetnTexImageARB, 6, uintptr(target), uintptr(level), uintptr(format), uintptr(xtype), uintptr(bufSize), uintptr(img))
}
func GetnUniformdvARB(program uint32, location int32, bufSize int32, params *float64) {
	syscall.Syscall6(gpGetnUniformdvARB, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformfv(program uint32, location int32, bufSize int32, params *float32) {
	syscall.Syscall6(gpGetnUniformfv, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformfvARB(program uint32, location int32, bufSize int32, params *float32) {
	syscall.Syscall6(gpGetnUniformfvARB, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformfvKHR(program uint32, location int32, bufSize int32, params *float32) {
	syscall.Syscall6(gpGetnUniformfvKHR, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformi64vARB(program uint32, location int32, bufSize int32, params *int64) {
	syscall.Syscall6(gpGetnUniformi64vARB, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformiv(program uint32, location int32, bufSize int32, params *int32) {
	syscall.Syscall6(gpGetnUniformiv, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformivARB(program uint32, location int32, bufSize int32, params *int32) {
	syscall.Syscall6(gpGetnUniformivARB, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformivKHR(program uint32, location int32, bufSize int32, params *int32) {
	syscall.Syscall6(gpGetnUniformivKHR, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformui64vARB(program uint32, location int32, bufSize int32, params *uint64) {
	syscall.Syscall6(gpGetnUniformui64vARB, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformuiv(program uint32, location int32, bufSize int32, params *uint32) {
	syscall.Syscall6(gpGetnUniformuiv, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformuivARB(program uint32, location int32, bufSize int32, params *uint32) {
	syscall.Syscall6(gpGetnUniformuivARB, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GetnUniformuivKHR(program uint32, location int32, bufSize int32, params *uint32) {
	syscall.Syscall6(gpGetnUniformuivKHR, 4, uintptr(program), uintptr(location), uintptr(bufSize), uintptr(unsafe.Pointer(params)), 0, 0)
}
func GlobalAlphaFactorbSUN(factor int8) {
	syscall.Syscall(gpGlobalAlphaFactorbSUN, 1, uintptr(factor), 0, 0)
}
func GlobalAlphaFactordSUN(factor float64) {
	syscall.Syscall(gpGlobalAlphaFactordSUN, 1, uintptr(math.Float64bits(factor)), 0, 0)
}
func GlobalAlphaFactorfSUN(factor float32) {
	syscall.Syscall(gpGlobalAlphaFactorfSUN, 1, uintptr(math.Float32bits(factor)), 0, 0)
}
func GlobalAlphaFactoriSUN(factor int32) {
	syscall.Syscall(gpGlobalAlphaFactoriSUN, 1, uintptr(factor), 0, 0)
}
func GlobalAlphaFactorsSUN(factor int16) {
	syscall.Syscall(gpGlobalAlphaFactorsSUN, 1, uintptr(factor), 0, 0)
}
func GlobalAlphaFactorubSUN(factor uint8) {
	syscall.Syscall(gpGlobalAlphaFactorubSUN, 1, uintptr(factor), 0, 0)
}
func GlobalAlphaFactoruiSUN(factor uint32) {
	syscall.Syscall(gpGlobalAlphaFactoruiSUN, 1, uintptr(factor), 0, 0)
}
func GlobalAlphaFactorusSUN(factor uint16) {
	syscall.Syscall(gpGlobalAlphaFactorusSUN, 1, uintptr(factor), 0, 0)
}

// specify implementation-specific hints
func Hint(target uint32, mode uint32) {
	syscall.Syscall(gpHint, 2, uintptr(target), uintptr(mode), 0)
}
func HintPGI(target uint32, mode int32) {
	syscall.Syscall(gpHintPGI, 2, uintptr(target), uintptr(mode), 0)
}
func HistogramEXT(target uint32, width int32, internalformat uint32, sink bool) {
	syscall.Syscall6(gpHistogramEXT, 4, uintptr(target), uintptr(width), uintptr(internalformat), boolToUintptr(sink), 0, 0)
}
func IglooInterfaceSGIX(pname uint32, params unsafe.Pointer) {
	syscall.Syscall(gpIglooInterfaceSGIX, 2, uintptr(pname), uintptr(params), 0)
}
func ImageTransformParameterfHP(target uint32, pname uint32, param float32) {
	syscall.Syscall(gpImageTransformParameterfHP, 3, uintptr(target), uintptr(pname), uintptr(math.Float32bits(param)))
}
func ImageTransformParameterfvHP(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpImageTransformParameterfvHP, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func ImageTransformParameteriHP(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpImageTransformParameteriHP, 3, uintptr(target), uintptr(pname), uintptr(param))
}
func ImageTransformParameterivHP(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpImageTransformParameterivHP, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func ImportMemoryFdEXT(memory uint32, size uint64, handleType uint32, fd int32) {
	syscall.Syscall6(gpImportMemoryFdEXT, 4, uintptr(memory), uintptr(size), uintptr(handleType), uintptr(fd), 0, 0)
}
func ImportMemoryWin32HandleEXT(memory uint32, size uint64, handleType uint32, handle unsafe.Pointer) {
	syscall.Syscall6(gpImportMemoryWin32HandleEXT, 4, uintptr(memory), uintptr(size), uintptr(handleType), uintptr(handle), 0, 0)
}
func ImportMemoryWin32NameEXT(memory uint32, size uint64, handleType uint32, name unsafe.Pointer) {
	syscall.Syscall6(gpImportMemoryWin32NameEXT, 4, uintptr(memory), uintptr(size), uintptr(handleType), uintptr(name), 0, 0)
}
func ImportSemaphoreFdEXT(semaphore uint32, handleType uint32, fd int32) {
	syscall.Syscall(gpImportSemaphoreFdEXT, 3, uintptr(semaphore), uintptr(handleType), uintptr(fd))
}
func ImportSemaphoreWin32HandleEXT(semaphore uint32, handleType uint32, handle unsafe.Pointer) {
	syscall.Syscall(gpImportSemaphoreWin32HandleEXT, 3, uintptr(semaphore), uintptr(handleType), uintptr(handle))
}
func ImportSemaphoreWin32NameEXT(semaphore uint32, handleType uint32, name unsafe.Pointer) {
	syscall.Syscall(gpImportSemaphoreWin32NameEXT, 3, uintptr(semaphore), uintptr(handleType), uintptr(name))
}
func ImportSyncEXT(external_sync_type uint32, external_sync int, flags uint32) uintptr {
	ret, _, _ := syscall.Syscall(gpImportSyncEXT, 3, uintptr(external_sync_type), uintptr(external_sync), uintptr(flags))
	return (uintptr)(ret)
}
func IndexFormatNV(xtype uint32, stride int32) {
	syscall.Syscall(gpIndexFormatNV, 2, uintptr(xtype), uintptr(stride), 0)
}
func IndexFuncEXT(xfunc uint32, ref float32) {
	syscall.Syscall(gpIndexFuncEXT, 2, uintptr(xfunc), uintptr(math.Float32bits(ref)), 0)
}

// control the writing of individual bits in the color index buffers
func IndexMask(mask uint32) {
	syscall.Syscall(gpIndexMask, 1, uintptr(mask), 0, 0)
}
func IndexMaterialEXT(face uint32, mode uint32) {
	syscall.Syscall(gpIndexMaterialEXT, 2, uintptr(face), uintptr(mode), 0)
}

// define an array of color indexes
func IndexPointer(xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpIndexPointer, 3, uintptr(xtype), uintptr(stride), uintptr(pointer))
}
func IndexPointerEXT(xtype uint32, stride int32, count int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpIndexPointerEXT, 4, uintptr(xtype), uintptr(stride), uintptr(count), uintptr(pointer), 0, 0)
}
func IndexPointerListIBM(xtype uint32, stride int32, pointer *unsafe.Pointer, ptrstride int32) {
	syscall.Syscall6(gpIndexPointerListIBM, 4, uintptr(xtype), uintptr(stride), uintptr(unsafe.Pointer(pointer)), uintptr(ptrstride), 0, 0)
}
func Indexd(c float64) {
	syscall.Syscall(gpIndexd, 1, uintptr(math.Float64bits(c)), 0, 0)
}
func Indexdv(c *float64) {
	syscall.Syscall(gpIndexdv, 1, uintptr(unsafe.Pointer(c)), 0, 0)
}
func Indexf(c float32) {
	syscall.Syscall(gpIndexf, 1, uintptr(math.Float32bits(c)), 0, 0)
}
func Indexfv(c *float32) {
	syscall.Syscall(gpIndexfv, 1, uintptr(unsafe.Pointer(c)), 0, 0)
}
func Indexi(c int32) {
	syscall.Syscall(gpIndexi, 1, uintptr(c), 0, 0)
}
func Indexiv(c *int32) {
	syscall.Syscall(gpIndexiv, 1, uintptr(unsafe.Pointer(c)), 0, 0)
}
func Indexs(c int16) {
	syscall.Syscall(gpIndexs, 1, uintptr(c), 0, 0)
}
func Indexsv(c *int16) {
	syscall.Syscall(gpIndexsv, 1, uintptr(unsafe.Pointer(c)), 0, 0)
}
func Indexub(c uint8) {
	syscall.Syscall(gpIndexub, 1, uintptr(c), 0, 0)
}
func Indexubv(c *uint8) {
	syscall.Syscall(gpIndexubv, 1, uintptr(unsafe.Pointer(c)), 0, 0)
}
func IndexxOES(component int32) {
	syscall.Syscall(gpIndexxOES, 1, uintptr(component), 0, 0)
}
func IndexxvOES(component *int32) {
	syscall.Syscall(gpIndexxvOES, 1, uintptr(unsafe.Pointer(component)), 0, 0)
}

// initialize the name stack
func InitNames() {
	syscall.Syscall(gpInitNames, 0, 0, 0, 0)
}
func InsertComponentEXT(res uint32, src uint32, num uint32) {
	syscall.Syscall(gpInsertComponentEXT, 3, uintptr(res), uintptr(src), uintptr(num))
}
func InsertEventMarkerEXT(length int32, marker *uint8) {
	syscall.Syscall(gpInsertEventMarkerEXT, 2, uintptr(length), uintptr(unsafe.Pointer(marker)), 0)
}
func InstrumentsBufferSGIX(size int32, buffer *int32) {
	syscall.Syscall(gpInstrumentsBufferSGIX, 2, uintptr(size), uintptr(unsafe.Pointer(buffer)), 0)
}

// simultaneously specify and enable several interleaved arrays
func InterleavedArrays(format uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpInterleavedArrays, 3, uintptr(format), uintptr(stride), uintptr(pointer))
}
func InterpolatePathsNV(resultPath uint32, pathA uint32, pathB uint32, weight float32) {
	syscall.Syscall6(gpInterpolatePathsNV, 4, uintptr(resultPath), uintptr(pathA), uintptr(pathB), uintptr(math.Float32bits(weight)), 0, 0)
}

// invalidate the content of a buffer object's data store
func InvalidateBufferData(buffer uint32) {
	syscall.Syscall(gpInvalidateBufferData, 1, uintptr(buffer), 0, 0)
}

// invalidate a region of a buffer object's data store
func InvalidateBufferSubData(buffer uint32, offset int, length int) {
	syscall.Syscall(gpInvalidateBufferSubData, 3, uintptr(buffer), uintptr(offset), uintptr(length))
}

// invalidate the content of some or all of a framebuffer's attachments
func InvalidateFramebuffer(target uint32, numAttachments int32, attachments *uint32) {
	syscall.Syscall(gpInvalidateFramebuffer, 3, uintptr(target), uintptr(numAttachments), uintptr(unsafe.Pointer(attachments)))
}

// invalidate the content of some or all of a framebuffer's attachments
func InvalidateNamedFramebufferData(framebuffer uint32, numAttachments int32, attachments *uint32) {
	syscall.Syscall(gpInvalidateNamedFramebufferData, 3, uintptr(framebuffer), uintptr(numAttachments), uintptr(unsafe.Pointer(attachments)))
}

// invalidate the content of a region of some or all of a framebuffer's attachments
func InvalidateNamedFramebufferSubData(framebuffer uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpInvalidateNamedFramebufferSubData, 7, uintptr(framebuffer), uintptr(numAttachments), uintptr(unsafe.Pointer(attachments)), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0, 0)
}

// invalidate the content of a region of some or all of a framebuffer's attachments
func InvalidateSubFramebuffer(target uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32) {
	syscall.Syscall9(gpInvalidateSubFramebuffer, 7, uintptr(target), uintptr(numAttachments), uintptr(unsafe.Pointer(attachments)), uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0, 0)
}

// invalidate the entirety a texture image
func InvalidateTexImage(texture uint32, level int32) {
	syscall.Syscall(gpInvalidateTexImage, 2, uintptr(texture), uintptr(level), 0)
}

// invalidate a region of a texture image
func InvalidateTexSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32) {
	syscall.Syscall9(gpInvalidateTexSubImage, 8, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), 0)
}
func IsAsyncMarkerSGIX(marker uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsAsyncMarkerSGIX, 1, uintptr(marker), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a buffer object
func IsBuffer(buffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsBuffer, 1, uintptr(buffer), 0, 0)
	return ret != 0
}
func IsBufferARB(buffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsBufferARB, 1, uintptr(buffer), 0, 0)
	return ret != 0
}
func IsBufferResidentNV(target uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsBufferResidentNV, 1, uintptr(target), 0, 0)
	return ret != 0
}
func IsCommandListNV(list uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsCommandListNV, 1, uintptr(list), 0, 0)
	return ret != 0
}
func IsEnabled(cap uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsEnabled, 1, uintptr(cap), 0, 0)
	return ret != 0
}
func IsEnabledIndexedEXT(target uint32, index uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsEnabledIndexedEXT, 2, uintptr(target), uintptr(index), 0)
	return ret != 0
}
func IsFenceAPPLE(fence uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsFenceAPPLE, 1, uintptr(fence), 0, 0)
	return ret != 0
}
func IsFenceNV(fence uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsFenceNV, 1, uintptr(fence), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a framebuffer object
func IsFramebuffer(framebuffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsFramebuffer, 1, uintptr(framebuffer), 0, 0)
	return ret != 0
}
func IsFramebufferEXT(framebuffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsFramebufferEXT, 1, uintptr(framebuffer), 0, 0)
	return ret != 0
}
func IsImageHandleResidentARB(handle uint64) bool {
	ret, _, _ := syscall.Syscall(gpIsImageHandleResidentARB, 1, uintptr(handle), 0, 0)
	return ret != 0
}
func IsImageHandleResidentNV(handle uint64) bool {
	ret, _, _ := syscall.Syscall(gpIsImageHandleResidentNV, 1, uintptr(handle), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a display list
func IsList(list uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsList, 1, uintptr(list), 0, 0)
	return ret != 0
}
func IsMemoryObjectEXT(memoryObject uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsMemoryObjectEXT, 1, uintptr(memoryObject), 0, 0)
	return ret != 0
}
func IsNameAMD(identifier uint32, name uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsNameAMD, 2, uintptr(identifier), uintptr(name), 0)
	return ret != 0
}
func IsNamedBufferResidentNV(buffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsNamedBufferResidentNV, 1, uintptr(buffer), 0, 0)
	return ret != 0
}
func IsNamedStringARB(namelen int32, name *uint8) bool {
	ret, _, _ := syscall.Syscall(gpIsNamedStringARB, 2, uintptr(namelen), uintptr(unsafe.Pointer(name)), 0)
	return ret != 0
}
func IsObjectBufferATI(buffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsObjectBufferATI, 1, uintptr(buffer), 0, 0)
	return ret != 0
}
func IsOcclusionQueryNV(id uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsOcclusionQueryNV, 1, uintptr(id), 0, 0)
	return ret != 0
}
func IsPathNV(path uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsPathNV, 1, uintptr(path), 0, 0)
	return ret != 0
}
func IsPointInFillPathNV(path uint32, mask uint32, x float32, y float32) bool {
	ret, _, _ := syscall.Syscall6(gpIsPointInFillPathNV, 4, uintptr(path), uintptr(mask), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), 0, 0)
	return ret != 0
}
func IsPointInStrokePathNV(path uint32, x float32, y float32) bool {
	ret, _, _ := syscall.Syscall(gpIsPointInStrokePathNV, 3, uintptr(path), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)))
	return ret != 0
}

// Determines if a name corresponds to a program object
func IsProgram(program uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsProgram, 1, uintptr(program), 0, 0)
	return ret != 0
}
func IsProgramARB(program uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsProgramARB, 1, uintptr(program), 0, 0)
	return ret != 0
}
func IsProgramNV(id uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsProgramNV, 1, uintptr(id), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a program pipeline object
func IsProgramPipeline(pipeline uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsProgramPipeline, 1, uintptr(pipeline), 0, 0)
	return ret != 0
}
func IsProgramPipelineEXT(pipeline uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsProgramPipelineEXT, 1, uintptr(pipeline), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a query object
func IsQuery(id uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsQuery, 1, uintptr(id), 0, 0)
	return ret != 0
}
func IsQueryARB(id uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsQueryARB, 1, uintptr(id), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a renderbuffer object
func IsRenderbuffer(renderbuffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsRenderbuffer, 1, uintptr(renderbuffer), 0, 0)
	return ret != 0
}
func IsRenderbufferEXT(renderbuffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsRenderbufferEXT, 1, uintptr(renderbuffer), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a sampler object
func IsSampler(sampler uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsSampler, 1, uintptr(sampler), 0, 0)
	return ret != 0
}
func IsSemaphoreEXT(semaphore uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsSemaphoreEXT, 1, uintptr(semaphore), 0, 0)
	return ret != 0
}

// Determines if a name corresponds to a shader object
func IsShader(shader uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsShader, 1, uintptr(shader), 0, 0)
	return ret != 0
}
func IsStateNV(state uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsStateNV, 1, uintptr(state), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a sync object
func IsSync(sync uintptr) bool {
	ret, _, _ := syscall.Syscall(gpIsSync, 1, uintptr(sync), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a texture
func IsTexture(texture uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsTexture, 1, uintptr(texture), 0, 0)
	return ret != 0
}
func IsTextureEXT(texture uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsTextureEXT, 1, uintptr(texture), 0, 0)
	return ret != 0
}
func IsTextureHandleResidentARB(handle uint64) bool {
	ret, _, _ := syscall.Syscall(gpIsTextureHandleResidentARB, 1, uintptr(handle), 0, 0)
	return ret != 0
}
func IsTextureHandleResidentNV(handle uint64) bool {
	ret, _, _ := syscall.Syscall(gpIsTextureHandleResidentNV, 1, uintptr(handle), 0, 0)
	return ret != 0
}

// determine if a name corresponds to a transform feedback object
func IsTransformFeedback(id uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsTransformFeedback, 1, uintptr(id), 0, 0)
	return ret != 0
}
func IsTransformFeedbackNV(id uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsTransformFeedbackNV, 1, uintptr(id), 0, 0)
	return ret != 0
}
func IsVariantEnabledEXT(id uint32, cap uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsVariantEnabledEXT, 2, uintptr(id), uintptr(cap), 0)
	return ret != 0
}

// determine if a name corresponds to a vertex array object
func IsVertexArray(array uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsVertexArray, 1, uintptr(array), 0, 0)
	return ret != 0
}
func IsVertexArrayAPPLE(array uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsVertexArrayAPPLE, 1, uintptr(array), 0, 0)
	return ret != 0
}
func IsVertexAttribEnabledAPPLE(index uint32, pname uint32) bool {
	ret, _, _ := syscall.Syscall(gpIsVertexAttribEnabledAPPLE, 2, uintptr(index), uintptr(pname), 0)
	return ret != 0
}
func LGPUCopyImageSubDataNVX(sourceGpu uint32, destinationGpuMask uint32, srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srxY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, width int32, height int32, depth int32) {
	panic("\"LGPUCopyImageSubDataNVX\" is not implemented")
}
func LGPUInterlockNVX() {
	syscall.Syscall(gpLGPUInterlockNVX, 0, 0, 0, 0)
}
func LGPUNamedBufferSubDataNVX(gpuMask uint32, buffer uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpLGPUNamedBufferSubDataNVX, 5, uintptr(gpuMask), uintptr(buffer), uintptr(offset), uintptr(size), uintptr(data), 0)
}
func LabelObjectEXT(xtype uint32, object uint32, length int32, label *uint8) {
	syscall.Syscall6(gpLabelObjectEXT, 4, uintptr(xtype), uintptr(object), uintptr(length), uintptr(unsafe.Pointer(label)), 0, 0)
}
func LightEnviSGIX(pname uint32, param int32) {
	syscall.Syscall(gpLightEnviSGIX, 2, uintptr(pname), uintptr(param), 0)
}
func LightModelf(pname uint32, param float32) {
	syscall.Syscall(gpLightModelf, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func LightModelfv(pname uint32, params *float32) {
	syscall.Syscall(gpLightModelfv, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func LightModeli(pname uint32, param int32) {
	syscall.Syscall(gpLightModeli, 2, uintptr(pname), uintptr(param), 0)
}
func LightModeliv(pname uint32, params *int32) {
	syscall.Syscall(gpLightModeliv, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func LightModelxOES(pname uint32, param int32) {
	syscall.Syscall(gpLightModelxOES, 2, uintptr(pname), uintptr(param), 0)
}
func LightModelxvOES(pname uint32, param *int32) {
	syscall.Syscall(gpLightModelxvOES, 2, uintptr(pname), uintptr(unsafe.Pointer(param)), 0)
}
func Lightf(light uint32, pname uint32, param float32) {
	syscall.Syscall(gpLightf, 3, uintptr(light), uintptr(pname), uintptr(math.Float32bits(param)))
}
func Lightfv(light uint32, pname uint32, params *float32) {
	syscall.Syscall(gpLightfv, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func Lighti(light uint32, pname uint32, param int32) {
	syscall.Syscall(gpLighti, 3, uintptr(light), uintptr(pname), uintptr(param))
}
func Lightiv(light uint32, pname uint32, params *int32) {
	syscall.Syscall(gpLightiv, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func LightxOES(light uint32, pname uint32, param int32) {
	syscall.Syscall(gpLightxOES, 3, uintptr(light), uintptr(pname), uintptr(param))
}
func LightxvOES(light uint32, pname uint32, params *int32) {
	syscall.Syscall(gpLightxvOES, 3, uintptr(light), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// specify the line stipple pattern
func LineStipple(factor int32, pattern uint16) {
	syscall.Syscall(gpLineStipple, 2, uintptr(factor), uintptr(pattern), 0)
}

// specify the width of rasterized lines
func LineWidth(width float32) {
	syscall.Syscall(gpLineWidth, 1, uintptr(math.Float32bits(width)), 0, 0)
}
func LineWidthxOES(width int32) {
	syscall.Syscall(gpLineWidthxOES, 1, uintptr(width), 0, 0)
}

// Links a program object
func LinkProgram(program uint32) {
	syscall.Syscall(gpLinkProgram, 1, uintptr(program), 0, 0)
}
func LinkProgramARB(programObj uintptr) {
	syscall.Syscall(gpLinkProgramARB, 1, uintptr(programObj), 0, 0)
}

// set the display-list base for
func ListBase(base uint32) {
	syscall.Syscall(gpListBase, 1, uintptr(base), 0, 0)
}
func ListDrawCommandsStatesClientNV(list uint32, segment uint32, indirects *unsafe.Pointer, sizes *int32, states *uint32, fbos *uint32, count uint32) {
	syscall.Syscall9(gpListDrawCommandsStatesClientNV, 7, uintptr(list), uintptr(segment), uintptr(unsafe.Pointer(indirects)), uintptr(unsafe.Pointer(sizes)), uintptr(unsafe.Pointer(states)), uintptr(unsafe.Pointer(fbos)), uintptr(count), 0, 0)
}
func ListParameterfSGIX(list uint32, pname uint32, param float32) {
	syscall.Syscall(gpListParameterfSGIX, 3, uintptr(list), uintptr(pname), uintptr(math.Float32bits(param)))
}
func ListParameterfvSGIX(list uint32, pname uint32, params *float32) {
	syscall.Syscall(gpListParameterfvSGIX, 3, uintptr(list), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func ListParameteriSGIX(list uint32, pname uint32, param int32) {
	syscall.Syscall(gpListParameteriSGIX, 3, uintptr(list), uintptr(pname), uintptr(param))
}
func ListParameterivSGIX(list uint32, pname uint32, params *int32) {
	syscall.Syscall(gpListParameterivSGIX, 3, uintptr(list), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// replace the current matrix with the identity matrix
func LoadIdentity() {
	syscall.Syscall(gpLoadIdentity, 0, 0, 0, 0)
}
func LoadIdentityDeformationMapSGIX(mask uint32) {
	syscall.Syscall(gpLoadIdentityDeformationMapSGIX, 1, uintptr(mask), 0, 0)
}
func LoadMatrixd(m *float64) {
	syscall.Syscall(gpLoadMatrixd, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func LoadMatrixf(m *float32) {
	syscall.Syscall(gpLoadMatrixf, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func LoadMatrixxOES(m *int32) {
	syscall.Syscall(gpLoadMatrixxOES, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}

// load a name onto the name stack
func LoadName(name uint32) {
	syscall.Syscall(gpLoadName, 1, uintptr(name), 0, 0)
}
func LoadProgramNV(target uint32, id uint32, len int32, program *uint8) {
	syscall.Syscall6(gpLoadProgramNV, 4, uintptr(target), uintptr(id), uintptr(len), uintptr(unsafe.Pointer(program)), 0, 0)
}
func LoadTransposeMatrixd(m *float64) {
	syscall.Syscall(gpLoadTransposeMatrixd, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func LoadTransposeMatrixdARB(m *float64) {
	syscall.Syscall(gpLoadTransposeMatrixdARB, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func LoadTransposeMatrixf(m *float32) {
	syscall.Syscall(gpLoadTransposeMatrixf, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func LoadTransposeMatrixfARB(m *float32) {
	syscall.Syscall(gpLoadTransposeMatrixfARB, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func LoadTransposeMatrixxOES(m *int32) {
	syscall.Syscall(gpLoadTransposeMatrixxOES, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func LockArraysEXT(first int32, count int32) {
	syscall.Syscall(gpLockArraysEXT, 2, uintptr(first), uintptr(count), 0)
}

// specify a logical pixel operation for rendering
func LogicOp(opcode uint32) {
	syscall.Syscall(gpLogicOp, 1, uintptr(opcode), 0, 0)
}
func MakeBufferNonResidentNV(target uint32) {
	syscall.Syscall(gpMakeBufferNonResidentNV, 1, uintptr(target), 0, 0)
}
func MakeBufferResidentNV(target uint32, access uint32) {
	syscall.Syscall(gpMakeBufferResidentNV, 2, uintptr(target), uintptr(access), 0)
}
func MakeImageHandleNonResidentARB(handle uint64) {
	syscall.Syscall(gpMakeImageHandleNonResidentARB, 1, uintptr(handle), 0, 0)
}
func MakeImageHandleNonResidentNV(handle uint64) {
	syscall.Syscall(gpMakeImageHandleNonResidentNV, 1, uintptr(handle), 0, 0)
}
func MakeImageHandleResidentARB(handle uint64, access uint32) {
	syscall.Syscall(gpMakeImageHandleResidentARB, 2, uintptr(handle), uintptr(access), 0)
}
func MakeImageHandleResidentNV(handle uint64, access uint32) {
	syscall.Syscall(gpMakeImageHandleResidentNV, 2, uintptr(handle), uintptr(access), 0)
}
func MakeNamedBufferNonResidentNV(buffer uint32) {
	syscall.Syscall(gpMakeNamedBufferNonResidentNV, 1, uintptr(buffer), 0, 0)
}
func MakeNamedBufferResidentNV(buffer uint32, access uint32) {
	syscall.Syscall(gpMakeNamedBufferResidentNV, 2, uintptr(buffer), uintptr(access), 0)
}
func MakeTextureHandleNonResidentARB(handle uint64) {
	syscall.Syscall(gpMakeTextureHandleNonResidentARB, 1, uintptr(handle), 0, 0)
}
func MakeTextureHandleNonResidentNV(handle uint64) {
	syscall.Syscall(gpMakeTextureHandleNonResidentNV, 1, uintptr(handle), 0, 0)
}
func MakeTextureHandleResidentARB(handle uint64) {
	syscall.Syscall(gpMakeTextureHandleResidentARB, 1, uintptr(handle), 0, 0)
}
func MakeTextureHandleResidentNV(handle uint64) {
	syscall.Syscall(gpMakeTextureHandleResidentNV, 1, uintptr(handle), 0, 0)
}
func Map1d(target uint32, u1 float64, u2 float64, stride int32, order int32, points *float64) {
	syscall.Syscall6(gpMap1d, 6, uintptr(target), uintptr(math.Float64bits(u1)), uintptr(math.Float64bits(u2)), uintptr(stride), uintptr(order), uintptr(unsafe.Pointer(points)))
}
func Map1f(target uint32, u1 float32, u2 float32, stride int32, order int32, points *float32) {
	syscall.Syscall6(gpMap1f, 6, uintptr(target), uintptr(math.Float32bits(u1)), uintptr(math.Float32bits(u2)), uintptr(stride), uintptr(order), uintptr(unsafe.Pointer(points)))
}
func Map1xOES(target uint32, u1 int32, u2 int32, stride int32, order int32, points int32) {
	syscall.Syscall6(gpMap1xOES, 6, uintptr(target), uintptr(u1), uintptr(u2), uintptr(stride), uintptr(order), uintptr(points))
}
func Map2d(target uint32, u1 float64, u2 float64, ustride int32, uorder int32, v1 float64, v2 float64, vstride int32, vorder int32, points *float64) {
	syscall.Syscall12(gpMap2d, 10, uintptr(target), uintptr(math.Float64bits(u1)), uintptr(math.Float64bits(u2)), uintptr(ustride), uintptr(uorder), uintptr(math.Float64bits(v1)), uintptr(math.Float64bits(v2)), uintptr(vstride), uintptr(vorder), uintptr(unsafe.Pointer(points)), 0, 0)
}
func Map2f(target uint32, u1 float32, u2 float32, ustride int32, uorder int32, v1 float32, v2 float32, vstride int32, vorder int32, points *float32) {
	syscall.Syscall12(gpMap2f, 10, uintptr(target), uintptr(math.Float32bits(u1)), uintptr(math.Float32bits(u2)), uintptr(ustride), uintptr(uorder), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), uintptr(vstride), uintptr(vorder), uintptr(unsafe.Pointer(points)), 0, 0)
}
func Map2xOES(target uint32, u1 int32, u2 int32, ustride int32, uorder int32, v1 int32, v2 int32, vstride int32, vorder int32, points int32) {
	syscall.Syscall12(gpMap2xOES, 10, uintptr(target), uintptr(u1), uintptr(u2), uintptr(ustride), uintptr(uorder), uintptr(v1), uintptr(v2), uintptr(vstride), uintptr(vorder), uintptr(points), 0, 0)
}

// map all of a buffer object's data store into the client's address space
func MapBuffer(target uint32, access uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(gpMapBuffer, 2, uintptr(target), uintptr(access), 0)
	return (unsafe.Pointer)(ret)
}
func MapBufferARB(target uint32, access uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(gpMapBufferARB, 2, uintptr(target), uintptr(access), 0)
	return (unsafe.Pointer)(ret)
}

// map all or part of a buffer object's data store into the client's address space
func MapBufferRange(target uint32, offset int, length int, access uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall6(gpMapBufferRange, 4, uintptr(target), uintptr(offset), uintptr(length), uintptr(access), 0, 0)
	return (unsafe.Pointer)(ret)
}
func MapControlPointsNV(target uint32, index uint32, xtype uint32, ustride int32, vstride int32, uorder int32, vorder int32, packed bool, points unsafe.Pointer) {
	syscall.Syscall9(gpMapControlPointsNV, 9, uintptr(target), uintptr(index), uintptr(xtype), uintptr(ustride), uintptr(vstride), uintptr(uorder), uintptr(vorder), boolToUintptr(packed), uintptr(points))
}
func MapGrid1d(un int32, u1 float64, u2 float64) {
	syscall.Syscall(gpMapGrid1d, 3, uintptr(un), uintptr(math.Float64bits(u1)), uintptr(math.Float64bits(u2)))
}
func MapGrid1f(un int32, u1 float32, u2 float32) {
	syscall.Syscall(gpMapGrid1f, 3, uintptr(un), uintptr(math.Float32bits(u1)), uintptr(math.Float32bits(u2)))
}
func MapGrid1xOES(n int32, u1 int32, u2 int32) {
	syscall.Syscall(gpMapGrid1xOES, 3, uintptr(n), uintptr(u1), uintptr(u2))
}
func MapGrid2d(un int32, u1 float64, u2 float64, vn int32, v1 float64, v2 float64) {
	syscall.Syscall6(gpMapGrid2d, 6, uintptr(un), uintptr(math.Float64bits(u1)), uintptr(math.Float64bits(u2)), uintptr(vn), uintptr(math.Float64bits(v1)), uintptr(math.Float64bits(v2)))
}
func MapGrid2f(un int32, u1 float32, u2 float32, vn int32, v1 float32, v2 float32) {
	syscall.Syscall6(gpMapGrid2f, 6, uintptr(un), uintptr(math.Float32bits(u1)), uintptr(math.Float32bits(u2)), uintptr(vn), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)))
}
func MapGrid2xOES(n int32, u1 int32, u2 int32, v1 int32, v2 int32) {
	syscall.Syscall6(gpMapGrid2xOES, 5, uintptr(n), uintptr(u1), uintptr(u2), uintptr(v1), uintptr(v2), 0)
}

// map all of a buffer object's data store into the client's address space
func MapNamedBuffer(buffer uint32, access uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(gpMapNamedBuffer, 2, uintptr(buffer), uintptr(access), 0)
	return (unsafe.Pointer)(ret)
}
func MapNamedBufferEXT(buffer uint32, access uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(gpMapNamedBufferEXT, 2, uintptr(buffer), uintptr(access), 0)
	return (unsafe.Pointer)(ret)
}

// map all or part of a buffer object's data store into the client's address space
func MapNamedBufferRange(buffer uint32, offset int, length int, access uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall6(gpMapNamedBufferRange, 4, uintptr(buffer), uintptr(offset), uintptr(length), uintptr(access), 0, 0)
	return (unsafe.Pointer)(ret)
}
func MapNamedBufferRangeEXT(buffer uint32, offset int, length int, access uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall6(gpMapNamedBufferRangeEXT, 4, uintptr(buffer), uintptr(offset), uintptr(length), uintptr(access), 0, 0)
	return (unsafe.Pointer)(ret)
}
func MapObjectBufferATI(buffer uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(gpMapObjectBufferATI, 1, uintptr(buffer), 0, 0)
	return (unsafe.Pointer)(ret)
}
func MapParameterfvNV(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpMapParameterfvNV, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func MapParameterivNV(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpMapParameterivNV, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func MapTexture2DINTEL(texture uint32, level int32, access uint32, stride *int32, layout *uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall6(gpMapTexture2DINTEL, 5, uintptr(texture), uintptr(level), uintptr(access), uintptr(unsafe.Pointer(stride)), uintptr(unsafe.Pointer(layout)), 0)
	return (unsafe.Pointer)(ret)
}
func MapVertexAttrib1dAPPLE(index uint32, size uint32, u1 float64, u2 float64, stride int32, order int32, points *float64) {
	syscall.Syscall9(gpMapVertexAttrib1dAPPLE, 7, uintptr(index), uintptr(size), uintptr(math.Float64bits(u1)), uintptr(math.Float64bits(u2)), uintptr(stride), uintptr(order), uintptr(unsafe.Pointer(points)), 0, 0)
}
func MapVertexAttrib1fAPPLE(index uint32, size uint32, u1 float32, u2 float32, stride int32, order int32, points *float32) {
	syscall.Syscall9(gpMapVertexAttrib1fAPPLE, 7, uintptr(index), uintptr(size), uintptr(math.Float32bits(u1)), uintptr(math.Float32bits(u2)), uintptr(stride), uintptr(order), uintptr(unsafe.Pointer(points)), 0, 0)
}
func MapVertexAttrib2dAPPLE(index uint32, size uint32, u1 float64, u2 float64, ustride int32, uorder int32, v1 float64, v2 float64, vstride int32, vorder int32, points *float64) {
	syscall.Syscall12(gpMapVertexAttrib2dAPPLE, 11, uintptr(index), uintptr(size), uintptr(math.Float64bits(u1)), uintptr(math.Float64bits(u2)), uintptr(ustride), uintptr(uorder), uintptr(math.Float64bits(v1)), uintptr(math.Float64bits(v2)), uintptr(vstride), uintptr(vorder), uintptr(unsafe.Pointer(points)), 0)
}
func MapVertexAttrib2fAPPLE(index uint32, size uint32, u1 float32, u2 float32, ustride int32, uorder int32, v1 float32, v2 float32, vstride int32, vorder int32, points *float32) {
	syscall.Syscall12(gpMapVertexAttrib2fAPPLE, 11, uintptr(index), uintptr(size), uintptr(math.Float32bits(u1)), uintptr(math.Float32bits(u2)), uintptr(ustride), uintptr(uorder), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), uintptr(vstride), uintptr(vorder), uintptr(unsafe.Pointer(points)), 0)
}
func Materialf(face uint32, pname uint32, param float32) {
	syscall.Syscall(gpMaterialf, 3, uintptr(face), uintptr(pname), uintptr(math.Float32bits(param)))
}
func Materialfv(face uint32, pname uint32, params *float32) {
	syscall.Syscall(gpMaterialfv, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func Materiali(face uint32, pname uint32, param int32) {
	syscall.Syscall(gpMateriali, 3, uintptr(face), uintptr(pname), uintptr(param))
}
func Materialiv(face uint32, pname uint32, params *int32) {
	syscall.Syscall(gpMaterialiv, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func MaterialxOES(face uint32, pname uint32, param int32) {
	syscall.Syscall(gpMaterialxOES, 3, uintptr(face), uintptr(pname), uintptr(param))
}
func MaterialxvOES(face uint32, pname uint32, param *int32) {
	syscall.Syscall(gpMaterialxvOES, 3, uintptr(face), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func MatrixFrustumEXT(mode uint32, left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	syscall.Syscall9(gpMatrixFrustumEXT, 7, uintptr(mode), uintptr(math.Float64bits(left)), uintptr(math.Float64bits(right)), uintptr(math.Float64bits(bottom)), uintptr(math.Float64bits(top)), uintptr(math.Float64bits(zNear)), uintptr(math.Float64bits(zFar)), 0, 0)
}
func MatrixIndexPointerARB(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpMatrixIndexPointerARB, 4, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0, 0)
}
func MatrixIndexubvARB(size int32, indices *uint8) {
	syscall.Syscall(gpMatrixIndexubvARB, 2, uintptr(size), uintptr(unsafe.Pointer(indices)), 0)
}
func MatrixIndexuivARB(size int32, indices *uint32) {
	syscall.Syscall(gpMatrixIndexuivARB, 2, uintptr(size), uintptr(unsafe.Pointer(indices)), 0)
}
func MatrixIndexusvARB(size int32, indices *uint16) {
	syscall.Syscall(gpMatrixIndexusvARB, 2, uintptr(size), uintptr(unsafe.Pointer(indices)), 0)
}
func MatrixLoad3x2fNV(matrixMode uint32, m *float32) {
	syscall.Syscall(gpMatrixLoad3x2fNV, 2, uintptr(matrixMode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixLoad3x3fNV(matrixMode uint32, m *float32) {
	syscall.Syscall(gpMatrixLoad3x3fNV, 2, uintptr(matrixMode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixLoadIdentityEXT(mode uint32) {
	syscall.Syscall(gpMatrixLoadIdentityEXT, 1, uintptr(mode), 0, 0)
}
func MatrixLoadTranspose3x3fNV(matrixMode uint32, m *float32) {
	syscall.Syscall(gpMatrixLoadTranspose3x3fNV, 2, uintptr(matrixMode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixLoadTransposedEXT(mode uint32, m *float64) {
	syscall.Syscall(gpMatrixLoadTransposedEXT, 2, uintptr(mode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixLoadTransposefEXT(mode uint32, m *float32) {
	syscall.Syscall(gpMatrixLoadTransposefEXT, 2, uintptr(mode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixLoaddEXT(mode uint32, m *float64) {
	syscall.Syscall(gpMatrixLoaddEXT, 2, uintptr(mode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixLoadfEXT(mode uint32, m *float32) {
	syscall.Syscall(gpMatrixLoadfEXT, 2, uintptr(mode), uintptr(unsafe.Pointer(m)), 0)
}

// specify which matrix is the current matrix
func MatrixMode(mode uint32) {
	syscall.Syscall(gpMatrixMode, 1, uintptr(mode), 0, 0)
}
func MatrixMult3x2fNV(matrixMode uint32, m *float32) {
	syscall.Syscall(gpMatrixMult3x2fNV, 2, uintptr(matrixMode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixMult3x3fNV(matrixMode uint32, m *float32) {
	syscall.Syscall(gpMatrixMult3x3fNV, 2, uintptr(matrixMode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixMultTranspose3x3fNV(matrixMode uint32, m *float32) {
	syscall.Syscall(gpMatrixMultTranspose3x3fNV, 2, uintptr(matrixMode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixMultTransposedEXT(mode uint32, m *float64) {
	syscall.Syscall(gpMatrixMultTransposedEXT, 2, uintptr(mode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixMultTransposefEXT(mode uint32, m *float32) {
	syscall.Syscall(gpMatrixMultTransposefEXT, 2, uintptr(mode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixMultdEXT(mode uint32, m *float64) {
	syscall.Syscall(gpMatrixMultdEXT, 2, uintptr(mode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixMultfEXT(mode uint32, m *float32) {
	syscall.Syscall(gpMatrixMultfEXT, 2, uintptr(mode), uintptr(unsafe.Pointer(m)), 0)
}
func MatrixOrthoEXT(mode uint32, left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	syscall.Syscall9(gpMatrixOrthoEXT, 7, uintptr(mode), uintptr(math.Float64bits(left)), uintptr(math.Float64bits(right)), uintptr(math.Float64bits(bottom)), uintptr(math.Float64bits(top)), uintptr(math.Float64bits(zNear)), uintptr(math.Float64bits(zFar)), 0, 0)
}
func MatrixPopEXT(mode uint32) {
	syscall.Syscall(gpMatrixPopEXT, 1, uintptr(mode), 0, 0)
}
func MatrixPushEXT(mode uint32) {
	syscall.Syscall(gpMatrixPushEXT, 1, uintptr(mode), 0, 0)
}
func MatrixRotatedEXT(mode uint32, angle float64, x float64, y float64, z float64) {
	syscall.Syscall6(gpMatrixRotatedEXT, 5, uintptr(mode), uintptr(math.Float64bits(angle)), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0)
}
func MatrixRotatefEXT(mode uint32, angle float32, x float32, y float32, z float32) {
	syscall.Syscall6(gpMatrixRotatefEXT, 5, uintptr(mode), uintptr(math.Float32bits(angle)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0)
}
func MatrixScaledEXT(mode uint32, x float64, y float64, z float64) {
	syscall.Syscall6(gpMatrixScaledEXT, 4, uintptr(mode), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func MatrixScalefEXT(mode uint32, x float32, y float32, z float32) {
	syscall.Syscall6(gpMatrixScalefEXT, 4, uintptr(mode), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func MatrixTranslatedEXT(mode uint32, x float64, y float64, z float64) {
	syscall.Syscall6(gpMatrixTranslatedEXT, 4, uintptr(mode), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func MatrixTranslatefEXT(mode uint32, x float32, y float32, z float32) {
	syscall.Syscall6(gpMatrixTranslatefEXT, 4, uintptr(mode), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func MaxShaderCompilerThreadsARB(count uint32) {
	syscall.Syscall(gpMaxShaderCompilerThreadsARB, 1, uintptr(count), 0, 0)
}
func MaxShaderCompilerThreadsKHR(count uint32) {
	syscall.Syscall(gpMaxShaderCompilerThreadsKHR, 1, uintptr(count), 0, 0)
}

// defines a barrier ordering memory transactions
func MemoryBarrier(barriers uint32) {
	syscall.Syscall(gpMemoryBarrier, 1, uintptr(barriers), 0, 0)
}
func MemoryBarrierByRegion(barriers uint32) {
	syscall.Syscall(gpMemoryBarrierByRegion, 1, uintptr(barriers), 0, 0)
}
func MemoryBarrierEXT(barriers uint32) {
	syscall.Syscall(gpMemoryBarrierEXT, 1, uintptr(barriers), 0, 0)
}
func MemoryObjectParameterivEXT(memoryObject uint32, pname uint32, params *int32) {
	syscall.Syscall(gpMemoryObjectParameterivEXT, 3, uintptr(memoryObject), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func MinSampleShadingARB(value float32) {
	syscall.Syscall(gpMinSampleShadingARB, 1, uintptr(math.Float32bits(value)), 0, 0)
}
func MinmaxEXT(target uint32, internalformat uint32, sink bool) {
	syscall.Syscall(gpMinmaxEXT, 3, uintptr(target), uintptr(internalformat), boolToUintptr(sink))
}
func MultMatrixd(m *float64) {
	syscall.Syscall(gpMultMatrixd, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func MultMatrixf(m *float32) {
	syscall.Syscall(gpMultMatrixf, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func MultMatrixxOES(m *int32) {
	syscall.Syscall(gpMultMatrixxOES, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func MultTransposeMatrixd(m *float64) {
	syscall.Syscall(gpMultTransposeMatrixd, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func MultTransposeMatrixdARB(m *float64) {
	syscall.Syscall(gpMultTransposeMatrixdARB, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func MultTransposeMatrixf(m *float32) {
	syscall.Syscall(gpMultTransposeMatrixf, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func MultTransposeMatrixfARB(m *float32) {
	syscall.Syscall(gpMultTransposeMatrixfARB, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
func MultTransposeMatrixxOES(m *int32) {
	syscall.Syscall(gpMultTransposeMatrixxOES, 1, uintptr(unsafe.Pointer(m)), 0, 0)
}

// render multiple sets of primitives from array data
func MultiDrawArrays(mode uint32, first *int32, count *int32, drawcount int32) {
	syscall.Syscall6(gpMultiDrawArrays, 4, uintptr(mode), uintptr(unsafe.Pointer(first)), uintptr(unsafe.Pointer(count)), uintptr(drawcount), 0, 0)
}
func MultiDrawArraysEXT(mode uint32, first *int32, count *int32, primcount int32) {
	syscall.Syscall6(gpMultiDrawArraysEXT, 4, uintptr(mode), uintptr(unsafe.Pointer(first)), uintptr(unsafe.Pointer(count)), uintptr(primcount), 0, 0)
}

// render multiple sets of primitives from array data, taking parameters from memory
func MultiDrawArraysIndirect(mode uint32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	syscall.Syscall6(gpMultiDrawArraysIndirect, 4, uintptr(mode), uintptr(indirect), uintptr(drawcount), uintptr(stride), 0, 0)
}
func MultiDrawArraysIndirectAMD(mode uint32, indirect unsafe.Pointer, primcount int32, stride int32) {
	syscall.Syscall6(gpMultiDrawArraysIndirectAMD, 4, uintptr(mode), uintptr(indirect), uintptr(primcount), uintptr(stride), 0, 0)
}
func MultiDrawArraysIndirectBindlessCountNV(mode uint32, indirect unsafe.Pointer, drawCount int32, maxDrawCount int32, stride int32, vertexBufferCount int32) {
	syscall.Syscall6(gpMultiDrawArraysIndirectBindlessCountNV, 6, uintptr(mode), uintptr(indirect), uintptr(drawCount), uintptr(maxDrawCount), uintptr(stride), uintptr(vertexBufferCount))
}
func MultiDrawArraysIndirectBindlessNV(mode uint32, indirect unsafe.Pointer, drawCount int32, stride int32, vertexBufferCount int32) {
	syscall.Syscall6(gpMultiDrawArraysIndirectBindlessNV, 5, uintptr(mode), uintptr(indirect), uintptr(drawCount), uintptr(stride), uintptr(vertexBufferCount), 0)
}
func MultiDrawArraysIndirectCountARB(mode uint32, indirect unsafe.Pointer, drawcount int, maxdrawcount int32, stride int32) {
	syscall.Syscall6(gpMultiDrawArraysIndirectCountARB, 5, uintptr(mode), uintptr(indirect), uintptr(drawcount), uintptr(maxdrawcount), uintptr(stride), 0)
}
func MultiDrawElementArrayAPPLE(mode uint32, first *int32, count *int32, primcount int32) {
	syscall.Syscall6(gpMultiDrawElementArrayAPPLE, 4, uintptr(mode), uintptr(unsafe.Pointer(first)), uintptr(unsafe.Pointer(count)), uintptr(primcount), 0, 0)
}

// render multiple sets of primitives by specifying indices of array data elements
func MultiDrawElements(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32) {
	syscall.Syscall6(gpMultiDrawElements, 5, uintptr(mode), uintptr(unsafe.Pointer(count)), uintptr(xtype), uintptr(unsafe.Pointer(indices)), uintptr(drawcount), 0)
}

// render multiple sets of primitives by specifying indices of array data elements and an index to apply to each index
func MultiDrawElementsBaseVertex(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32, basevertex *int32) {
	syscall.Syscall6(gpMultiDrawElementsBaseVertex, 6, uintptr(mode), uintptr(unsafe.Pointer(count)), uintptr(xtype), uintptr(unsafe.Pointer(indices)), uintptr(drawcount), uintptr(unsafe.Pointer(basevertex)))
}
func MultiDrawElementsEXT(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, primcount int32) {
	syscall.Syscall6(gpMultiDrawElementsEXT, 5, uintptr(mode), uintptr(unsafe.Pointer(count)), uintptr(xtype), uintptr(unsafe.Pointer(indices)), uintptr(primcount), 0)
}

// render indexed primitives from array data, taking parameters from memory
func MultiDrawElementsIndirect(mode uint32, xtype uint32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	syscall.Syscall6(gpMultiDrawElementsIndirect, 5, uintptr(mode), uintptr(xtype), uintptr(indirect), uintptr(drawcount), uintptr(stride), 0)
}
func MultiDrawElementsIndirectAMD(mode uint32, xtype uint32, indirect unsafe.Pointer, primcount int32, stride int32) {
	syscall.Syscall6(gpMultiDrawElementsIndirectAMD, 5, uintptr(mode), uintptr(xtype), uintptr(indirect), uintptr(primcount), uintptr(stride), 0)
}
func MultiDrawElementsIndirectBindlessCountNV(mode uint32, xtype uint32, indirect unsafe.Pointer, drawCount int32, maxDrawCount int32, stride int32, vertexBufferCount int32) {
	syscall.Syscall9(gpMultiDrawElementsIndirectBindlessCountNV, 7, uintptr(mode), uintptr(xtype), uintptr(indirect), uintptr(drawCount), uintptr(maxDrawCount), uintptr(stride), uintptr(vertexBufferCount), 0, 0)
}
func MultiDrawElementsIndirectBindlessNV(mode uint32, xtype uint32, indirect unsafe.Pointer, drawCount int32, stride int32, vertexBufferCount int32) {
	syscall.Syscall6(gpMultiDrawElementsIndirectBindlessNV, 6, uintptr(mode), uintptr(xtype), uintptr(indirect), uintptr(drawCount), uintptr(stride), uintptr(vertexBufferCount))
}
func MultiDrawElementsIndirectCountARB(mode uint32, xtype uint32, indirect unsafe.Pointer, drawcount int, maxdrawcount int32, stride int32) {
	syscall.Syscall6(gpMultiDrawElementsIndirectCountARB, 6, uintptr(mode), uintptr(xtype), uintptr(indirect), uintptr(drawcount), uintptr(maxdrawcount), uintptr(stride))
}
func MultiDrawRangeElementArrayAPPLE(mode uint32, start uint32, end uint32, first *int32, count *int32, primcount int32) {
	syscall.Syscall6(gpMultiDrawRangeElementArrayAPPLE, 6, uintptr(mode), uintptr(start), uintptr(end), uintptr(unsafe.Pointer(first)), uintptr(unsafe.Pointer(count)), uintptr(primcount))
}
func MultiModeDrawArraysIBM(mode *uint32, first *int32, count *int32, primcount int32, modestride int32) {
	syscall.Syscall6(gpMultiModeDrawArraysIBM, 5, uintptr(unsafe.Pointer(mode)), uintptr(unsafe.Pointer(first)), uintptr(unsafe.Pointer(count)), uintptr(primcount), uintptr(modestride), 0)
}
func MultiModeDrawElementsIBM(mode *uint32, count *int32, xtype uint32, indices *unsafe.Pointer, primcount int32, modestride int32) {
	syscall.Syscall6(gpMultiModeDrawElementsIBM, 6, uintptr(unsafe.Pointer(mode)), uintptr(unsafe.Pointer(count)), uintptr(xtype), uintptr(unsafe.Pointer(indices)), uintptr(primcount), uintptr(modestride))
}
func MultiTexBufferEXT(texunit uint32, target uint32, internalformat uint32, buffer uint32) {
	syscall.Syscall6(gpMultiTexBufferEXT, 4, uintptr(texunit), uintptr(target), uintptr(internalformat), uintptr(buffer), 0, 0)
}
func MultiTexCoord1bOES(texture uint32, s int8) {
	syscall.Syscall(gpMultiTexCoord1bOES, 2, uintptr(texture), uintptr(s), 0)
}
func MultiTexCoord1bvOES(texture uint32, coords *int8) {
	syscall.Syscall(gpMultiTexCoord1bvOES, 2, uintptr(texture), uintptr(unsafe.Pointer(coords)), 0)
}
func MultiTexCoord1d(target uint32, s float64) {
	syscall.Syscall(gpMultiTexCoord1d, 2, uintptr(target), uintptr(math.Float64bits(s)), 0)
}
func MultiTexCoord1dARB(target uint32, s float64) {
	syscall.Syscall(gpMultiTexCoord1dARB, 2, uintptr(target), uintptr(math.Float64bits(s)), 0)
}
func MultiTexCoord1dv(target uint32, v *float64) {
	syscall.Syscall(gpMultiTexCoord1dv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1dvARB(target uint32, v *float64) {
	syscall.Syscall(gpMultiTexCoord1dvARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1f(target uint32, s float32) {
	syscall.Syscall(gpMultiTexCoord1f, 2, uintptr(target), uintptr(math.Float32bits(s)), 0)
}
func MultiTexCoord1fARB(target uint32, s float32) {
	syscall.Syscall(gpMultiTexCoord1fARB, 2, uintptr(target), uintptr(math.Float32bits(s)), 0)
}
func MultiTexCoord1fv(target uint32, v *float32) {
	syscall.Syscall(gpMultiTexCoord1fv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1fvARB(target uint32, v *float32) {
	syscall.Syscall(gpMultiTexCoord1fvARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1hNV(target uint32, s uint16) {
	syscall.Syscall(gpMultiTexCoord1hNV, 2, uintptr(target), uintptr(s), 0)
}
func MultiTexCoord1hvNV(target uint32, v *uint16) {
	syscall.Syscall(gpMultiTexCoord1hvNV, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1i(target uint32, s int32) {
	syscall.Syscall(gpMultiTexCoord1i, 2, uintptr(target), uintptr(s), 0)
}
func MultiTexCoord1iARB(target uint32, s int32) {
	syscall.Syscall(gpMultiTexCoord1iARB, 2, uintptr(target), uintptr(s), 0)
}
func MultiTexCoord1iv(target uint32, v *int32) {
	syscall.Syscall(gpMultiTexCoord1iv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1ivARB(target uint32, v *int32) {
	syscall.Syscall(gpMultiTexCoord1ivARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1s(target uint32, s int16) {
	syscall.Syscall(gpMultiTexCoord1s, 2, uintptr(target), uintptr(s), 0)
}
func MultiTexCoord1sARB(target uint32, s int16) {
	syscall.Syscall(gpMultiTexCoord1sARB, 2, uintptr(target), uintptr(s), 0)
}
func MultiTexCoord1sv(target uint32, v *int16) {
	syscall.Syscall(gpMultiTexCoord1sv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1svARB(target uint32, v *int16) {
	syscall.Syscall(gpMultiTexCoord1svARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord1xOES(texture uint32, s int32) {
	syscall.Syscall(gpMultiTexCoord1xOES, 2, uintptr(texture), uintptr(s), 0)
}
func MultiTexCoord1xvOES(texture uint32, coords *int32) {
	syscall.Syscall(gpMultiTexCoord1xvOES, 2, uintptr(texture), uintptr(unsafe.Pointer(coords)), 0)
}
func MultiTexCoord2bOES(texture uint32, s int8, t int8) {
	syscall.Syscall(gpMultiTexCoord2bOES, 3, uintptr(texture), uintptr(s), uintptr(t))
}
func MultiTexCoord2bvOES(texture uint32, coords *int8) {
	syscall.Syscall(gpMultiTexCoord2bvOES, 2, uintptr(texture), uintptr(unsafe.Pointer(coords)), 0)
}
func MultiTexCoord2d(target uint32, s float64, t float64) {
	syscall.Syscall(gpMultiTexCoord2d, 3, uintptr(target), uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)))
}
func MultiTexCoord2dARB(target uint32, s float64, t float64) {
	syscall.Syscall(gpMultiTexCoord2dARB, 3, uintptr(target), uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)))
}
func MultiTexCoord2dv(target uint32, v *float64) {
	syscall.Syscall(gpMultiTexCoord2dv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2dvARB(target uint32, v *float64) {
	syscall.Syscall(gpMultiTexCoord2dvARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2f(target uint32, s float32, t float32) {
	syscall.Syscall(gpMultiTexCoord2f, 3, uintptr(target), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)))
}
func MultiTexCoord2fARB(target uint32, s float32, t float32) {
	syscall.Syscall(gpMultiTexCoord2fARB, 3, uintptr(target), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)))
}
func MultiTexCoord2fv(target uint32, v *float32) {
	syscall.Syscall(gpMultiTexCoord2fv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2fvARB(target uint32, v *float32) {
	syscall.Syscall(gpMultiTexCoord2fvARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2hNV(target uint32, s uint16, t uint16) {
	syscall.Syscall(gpMultiTexCoord2hNV, 3, uintptr(target), uintptr(s), uintptr(t))
}
func MultiTexCoord2hvNV(target uint32, v *uint16) {
	syscall.Syscall(gpMultiTexCoord2hvNV, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2i(target uint32, s int32, t int32) {
	syscall.Syscall(gpMultiTexCoord2i, 3, uintptr(target), uintptr(s), uintptr(t))
}
func MultiTexCoord2iARB(target uint32, s int32, t int32) {
	syscall.Syscall(gpMultiTexCoord2iARB, 3, uintptr(target), uintptr(s), uintptr(t))
}
func MultiTexCoord2iv(target uint32, v *int32) {
	syscall.Syscall(gpMultiTexCoord2iv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2ivARB(target uint32, v *int32) {
	syscall.Syscall(gpMultiTexCoord2ivARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2s(target uint32, s int16, t int16) {
	syscall.Syscall(gpMultiTexCoord2s, 3, uintptr(target), uintptr(s), uintptr(t))
}
func MultiTexCoord2sARB(target uint32, s int16, t int16) {
	syscall.Syscall(gpMultiTexCoord2sARB, 3, uintptr(target), uintptr(s), uintptr(t))
}
func MultiTexCoord2sv(target uint32, v *int16) {
	syscall.Syscall(gpMultiTexCoord2sv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2svARB(target uint32, v *int16) {
	syscall.Syscall(gpMultiTexCoord2svARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord2xOES(texture uint32, s int32, t int32) {
	syscall.Syscall(gpMultiTexCoord2xOES, 3, uintptr(texture), uintptr(s), uintptr(t))
}
func MultiTexCoord2xvOES(texture uint32, coords *int32) {
	syscall.Syscall(gpMultiTexCoord2xvOES, 2, uintptr(texture), uintptr(unsafe.Pointer(coords)), 0)
}
func MultiTexCoord3bOES(texture uint32, s int8, t int8, r int8) {
	syscall.Syscall6(gpMultiTexCoord3bOES, 4, uintptr(texture), uintptr(s), uintptr(t), uintptr(r), 0, 0)
}
func MultiTexCoord3bvOES(texture uint32, coords *int8) {
	syscall.Syscall(gpMultiTexCoord3bvOES, 2, uintptr(texture), uintptr(unsafe.Pointer(coords)), 0)
}
func MultiTexCoord3d(target uint32, s float64, t float64, r float64) {
	syscall.Syscall6(gpMultiTexCoord3d, 4, uintptr(target), uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)), uintptr(math.Float64bits(r)), 0, 0)
}
func MultiTexCoord3dARB(target uint32, s float64, t float64, r float64) {
	syscall.Syscall6(gpMultiTexCoord3dARB, 4, uintptr(target), uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)), uintptr(math.Float64bits(r)), 0, 0)
}
func MultiTexCoord3dv(target uint32, v *float64) {
	syscall.Syscall(gpMultiTexCoord3dv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3dvARB(target uint32, v *float64) {
	syscall.Syscall(gpMultiTexCoord3dvARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3f(target uint32, s float32, t float32, r float32) {
	syscall.Syscall6(gpMultiTexCoord3f, 4, uintptr(target), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)), 0, 0)
}
func MultiTexCoord3fARB(target uint32, s float32, t float32, r float32) {
	syscall.Syscall6(gpMultiTexCoord3fARB, 4, uintptr(target), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)), 0, 0)
}
func MultiTexCoord3fv(target uint32, v *float32) {
	syscall.Syscall(gpMultiTexCoord3fv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3fvARB(target uint32, v *float32) {
	syscall.Syscall(gpMultiTexCoord3fvARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3hNV(target uint32, s uint16, t uint16, r uint16) {
	syscall.Syscall6(gpMultiTexCoord3hNV, 4, uintptr(target), uintptr(s), uintptr(t), uintptr(r), 0, 0)
}
func MultiTexCoord3hvNV(target uint32, v *uint16) {
	syscall.Syscall(gpMultiTexCoord3hvNV, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3i(target uint32, s int32, t int32, r int32) {
	syscall.Syscall6(gpMultiTexCoord3i, 4, uintptr(target), uintptr(s), uintptr(t), uintptr(r), 0, 0)
}
func MultiTexCoord3iARB(target uint32, s int32, t int32, r int32) {
	syscall.Syscall6(gpMultiTexCoord3iARB, 4, uintptr(target), uintptr(s), uintptr(t), uintptr(r), 0, 0)
}
func MultiTexCoord3iv(target uint32, v *int32) {
	syscall.Syscall(gpMultiTexCoord3iv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3ivARB(target uint32, v *int32) {
	syscall.Syscall(gpMultiTexCoord3ivARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3s(target uint32, s int16, t int16, r int16) {
	syscall.Syscall6(gpMultiTexCoord3s, 4, uintptr(target), uintptr(s), uintptr(t), uintptr(r), 0, 0)
}
func MultiTexCoord3sARB(target uint32, s int16, t int16, r int16) {
	syscall.Syscall6(gpMultiTexCoord3sARB, 4, uintptr(target), uintptr(s), uintptr(t), uintptr(r), 0, 0)
}
func MultiTexCoord3sv(target uint32, v *int16) {
	syscall.Syscall(gpMultiTexCoord3sv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3svARB(target uint32, v *int16) {
	syscall.Syscall(gpMultiTexCoord3svARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord3xOES(texture uint32, s int32, t int32, r int32) {
	syscall.Syscall6(gpMultiTexCoord3xOES, 4, uintptr(texture), uintptr(s), uintptr(t), uintptr(r), 0, 0)
}
func MultiTexCoord3xvOES(texture uint32, coords *int32) {
	syscall.Syscall(gpMultiTexCoord3xvOES, 2, uintptr(texture), uintptr(unsafe.Pointer(coords)), 0)
}
func MultiTexCoord4bOES(texture uint32, s int8, t int8, r int8, q int8) {
	syscall.Syscall6(gpMultiTexCoord4bOES, 5, uintptr(texture), uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0)
}
func MultiTexCoord4bvOES(texture uint32, coords *int8) {
	syscall.Syscall(gpMultiTexCoord4bvOES, 2, uintptr(texture), uintptr(unsafe.Pointer(coords)), 0)
}
func MultiTexCoord4d(target uint32, s float64, t float64, r float64, q float64) {
	syscall.Syscall6(gpMultiTexCoord4d, 5, uintptr(target), uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)), uintptr(math.Float64bits(r)), uintptr(math.Float64bits(q)), 0)
}
func MultiTexCoord4dARB(target uint32, s float64, t float64, r float64, q float64) {
	syscall.Syscall6(gpMultiTexCoord4dARB, 5, uintptr(target), uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)), uintptr(math.Float64bits(r)), uintptr(math.Float64bits(q)), 0)
}
func MultiTexCoord4dv(target uint32, v *float64) {
	syscall.Syscall(gpMultiTexCoord4dv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4dvARB(target uint32, v *float64) {
	syscall.Syscall(gpMultiTexCoord4dvARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4f(target uint32, s float32, t float32, r float32, q float32) {
	syscall.Syscall6(gpMultiTexCoord4f, 5, uintptr(target), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(q)), 0)
}
func MultiTexCoord4fARB(target uint32, s float32, t float32, r float32, q float32) {
	syscall.Syscall6(gpMultiTexCoord4fARB, 5, uintptr(target), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(q)), 0)
}
func MultiTexCoord4fv(target uint32, v *float32) {
	syscall.Syscall(gpMultiTexCoord4fv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4fvARB(target uint32, v *float32) {
	syscall.Syscall(gpMultiTexCoord4fvARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4hNV(target uint32, s uint16, t uint16, r uint16, q uint16) {
	syscall.Syscall6(gpMultiTexCoord4hNV, 5, uintptr(target), uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0)
}
func MultiTexCoord4hvNV(target uint32, v *uint16) {
	syscall.Syscall(gpMultiTexCoord4hvNV, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4i(target uint32, s int32, t int32, r int32, q int32) {
	syscall.Syscall6(gpMultiTexCoord4i, 5, uintptr(target), uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0)
}
func MultiTexCoord4iARB(target uint32, s int32, t int32, r int32, q int32) {
	syscall.Syscall6(gpMultiTexCoord4iARB, 5, uintptr(target), uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0)
}
func MultiTexCoord4iv(target uint32, v *int32) {
	syscall.Syscall(gpMultiTexCoord4iv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4ivARB(target uint32, v *int32) {
	syscall.Syscall(gpMultiTexCoord4ivARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4s(target uint32, s int16, t int16, r int16, q int16) {
	syscall.Syscall6(gpMultiTexCoord4s, 5, uintptr(target), uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0)
}
func MultiTexCoord4sARB(target uint32, s int16, t int16, r int16, q int16) {
	syscall.Syscall6(gpMultiTexCoord4sARB, 5, uintptr(target), uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0)
}
func MultiTexCoord4sv(target uint32, v *int16) {
	syscall.Syscall(gpMultiTexCoord4sv, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4svARB(target uint32, v *int16) {
	syscall.Syscall(gpMultiTexCoord4svARB, 2, uintptr(target), uintptr(unsafe.Pointer(v)), 0)
}
func MultiTexCoord4xOES(texture uint32, s int32, t int32, r int32, q int32) {
	syscall.Syscall6(gpMultiTexCoord4xOES, 5, uintptr(texture), uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0)
}
func MultiTexCoord4xvOES(texture uint32, coords *int32) {
	syscall.Syscall(gpMultiTexCoord4xvOES, 2, uintptr(texture), uintptr(unsafe.Pointer(coords)), 0)
}
func MultiTexCoordPointerEXT(texunit uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpMultiTexCoordPointerEXT, 5, uintptr(texunit), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0)
}
func MultiTexEnvfEXT(texunit uint32, target uint32, pname uint32, param float32) {
	syscall.Syscall6(gpMultiTexEnvfEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(math.Float32bits(param)), 0, 0)
}
func MultiTexEnvfvEXT(texunit uint32, target uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpMultiTexEnvfvEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexEnviEXT(texunit uint32, target uint32, pname uint32, param int32) {
	syscall.Syscall6(gpMultiTexEnviEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(param), 0, 0)
}
func MultiTexEnvivEXT(texunit uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpMultiTexEnvivEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexGendEXT(texunit uint32, coord uint32, pname uint32, param float64) {
	syscall.Syscall6(gpMultiTexGendEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(math.Float64bits(param)), 0, 0)
}
func MultiTexGendvEXT(texunit uint32, coord uint32, pname uint32, params *float64) {
	syscall.Syscall6(gpMultiTexGendvEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexGenfEXT(texunit uint32, coord uint32, pname uint32, param float32) {
	syscall.Syscall6(gpMultiTexGenfEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(math.Float32bits(param)), 0, 0)
}
func MultiTexGenfvEXT(texunit uint32, coord uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpMultiTexGenfvEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexGeniEXT(texunit uint32, coord uint32, pname uint32, param int32) {
	syscall.Syscall6(gpMultiTexGeniEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(param), 0, 0)
}
func MultiTexGenivEXT(texunit uint32, coord uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpMultiTexGenivEXT, 4, uintptr(texunit), uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexImage1DEXT(texunit uint32, target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpMultiTexImage1DEXT, 9, uintptr(texunit), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels))
}
func MultiTexImage2DEXT(texunit uint32, target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpMultiTexImage2DEXT, 10, uintptr(texunit), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}
func MultiTexImage3DEXT(texunit uint32, target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpMultiTexImage3DEXT, 11, uintptr(texunit), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func MultiTexParameterIivEXT(texunit uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpMultiTexParameterIivEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexParameterIuivEXT(texunit uint32, target uint32, pname uint32, params *uint32) {
	syscall.Syscall6(gpMultiTexParameterIuivEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexParameterfEXT(texunit uint32, target uint32, pname uint32, param float32) {
	syscall.Syscall6(gpMultiTexParameterfEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(math.Float32bits(param)), 0, 0)
}
func MultiTexParameterfvEXT(texunit uint32, target uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpMultiTexParameterfvEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexParameteriEXT(texunit uint32, target uint32, pname uint32, param int32) {
	syscall.Syscall6(gpMultiTexParameteriEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(param), 0, 0)
}
func MultiTexParameterivEXT(texunit uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpMultiTexParameterivEXT, 4, uintptr(texunit), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MultiTexRenderbufferEXT(texunit uint32, target uint32, renderbuffer uint32) {
	syscall.Syscall(gpMultiTexRenderbufferEXT, 3, uintptr(texunit), uintptr(target), uintptr(renderbuffer))
}
func MultiTexSubImage1DEXT(texunit uint32, target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpMultiTexSubImage1DEXT, 8, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func MultiTexSubImage2DEXT(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpMultiTexSubImage2DEXT, 10, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}
func MultiTexSubImage3DEXT(texunit uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpMultiTexSubImage3DEXT, 12, uintptr(texunit), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(xtype), uintptr(pixels))
}
func MulticastBarrierNV() {
	syscall.Syscall(gpMulticastBarrierNV, 0, 0, 0, 0)
}
func MulticastBlitFramebufferNV(srcGpu uint32, dstGpu uint32, srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	syscall.Syscall12(gpMulticastBlitFramebufferNV, 12, uintptr(srcGpu), uintptr(dstGpu), uintptr(srcX0), uintptr(srcY0), uintptr(srcX1), uintptr(srcY1), uintptr(dstX0), uintptr(dstY0), uintptr(dstX1), uintptr(dstY1), uintptr(mask), uintptr(filter))
}
func MulticastBufferSubDataNV(gpuMask uint32, buffer uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpMulticastBufferSubDataNV, 5, uintptr(gpuMask), uintptr(buffer), uintptr(offset), uintptr(size), uintptr(data), 0)
}
func MulticastCopyBufferSubDataNV(readGpu uint32, writeGpuMask uint32, readBuffer uint32, writeBuffer uint32, readOffset int, writeOffset int, size int) {
	syscall.Syscall9(gpMulticastCopyBufferSubDataNV, 7, uintptr(readGpu), uintptr(writeGpuMask), uintptr(readBuffer), uintptr(writeBuffer), uintptr(readOffset), uintptr(writeOffset), uintptr(size), 0, 0)
}
func MulticastCopyImageSubDataNV(srcGpu uint32, dstGpuMask uint32, srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, srcWidth int32, srcHeight int32, srcDepth int32) {
	panic("\"MulticastCopyImageSubDataNV\" is not implemented")
}
func MulticastFramebufferSampleLocationsfvNV(gpu uint32, framebuffer uint32, start uint32, count int32, v *float32) {
	syscall.Syscall6(gpMulticastFramebufferSampleLocationsfvNV, 5, uintptr(gpu), uintptr(framebuffer), uintptr(start), uintptr(count), uintptr(unsafe.Pointer(v)), 0)
}
func MulticastGetQueryObjecti64vNV(gpu uint32, id uint32, pname uint32, params *int64) {
	syscall.Syscall6(gpMulticastGetQueryObjecti64vNV, 4, uintptr(gpu), uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MulticastGetQueryObjectivNV(gpu uint32, id uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpMulticastGetQueryObjectivNV, 4, uintptr(gpu), uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MulticastGetQueryObjectui64vNV(gpu uint32, id uint32, pname uint32, params *uint64) {
	syscall.Syscall6(gpMulticastGetQueryObjectui64vNV, 4, uintptr(gpu), uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MulticastGetQueryObjectuivNV(gpu uint32, id uint32, pname uint32, params *uint32) {
	syscall.Syscall6(gpMulticastGetQueryObjectuivNV, 4, uintptr(gpu), uintptr(id), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func MulticastWaitSyncNV(signalGpu uint32, waitGpuMask uint32) {
	syscall.Syscall(gpMulticastWaitSyncNV, 2, uintptr(signalGpu), uintptr(waitGpuMask), 0)
}

// creates and initializes a buffer object's data     store
func NamedBufferData(buffer uint32, size int, data unsafe.Pointer, usage uint32) {
	syscall.Syscall6(gpNamedBufferData, 4, uintptr(buffer), uintptr(size), uintptr(data), uintptr(usage), 0, 0)
}
func NamedBufferDataEXT(buffer uint32, size int, data unsafe.Pointer, usage uint32) {
	syscall.Syscall6(gpNamedBufferDataEXT, 4, uintptr(buffer), uintptr(size), uintptr(data), uintptr(usage), 0, 0)
}
func NamedBufferPageCommitmentARB(buffer uint32, offset int, size int, commit bool) {
	syscall.Syscall6(gpNamedBufferPageCommitmentARB, 4, uintptr(buffer), uintptr(offset), uintptr(size), boolToUintptr(commit), 0, 0)
}
func NamedBufferPageCommitmentEXT(buffer uint32, offset int, size int, commit bool) {
	syscall.Syscall6(gpNamedBufferPageCommitmentEXT, 4, uintptr(buffer), uintptr(offset), uintptr(size), boolToUintptr(commit), 0, 0)
}

// creates and initializes a buffer object's immutable data     store
func NamedBufferStorage(buffer uint32, size int, data unsafe.Pointer, flags uint32) {
	syscall.Syscall6(gpNamedBufferStorage, 4, uintptr(buffer), uintptr(size), uintptr(data), uintptr(flags), 0, 0)
}
func NamedBufferStorageEXT(buffer uint32, size int, data unsafe.Pointer, flags uint32) {
	syscall.Syscall6(gpNamedBufferStorageEXT, 4, uintptr(buffer), uintptr(size), uintptr(data), uintptr(flags), 0, 0)
}

// Parameter clientBuffer has type C.GLeglClientBufferEXT.
func NamedBufferStorageExternalEXT(buffer uint32, offset int, size int, clientBuffer unsafe.Pointer, flags uint32) {
	syscall.Syscall6(gpNamedBufferStorageExternalEXT, 5, uintptr(buffer), uintptr(offset), uintptr(size), uintptr(clientBuffer), uintptr(flags), 0)
}
func NamedBufferStorageMemEXT(buffer uint32, size int, memory uint32, offset uint64) {
	syscall.Syscall6(gpNamedBufferStorageMemEXT, 4, uintptr(buffer), uintptr(size), uintptr(memory), uintptr(offset), 0, 0)
}

// updates a subset of a buffer object's data store
func NamedBufferSubData(buffer uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpNamedBufferSubData, 4, uintptr(buffer), uintptr(offset), uintptr(size), uintptr(data), 0, 0)
}
func NamedBufferSubDataEXT(buffer uint32, offset int, size int, data unsafe.Pointer) {
	syscall.Syscall6(gpNamedBufferSubDataEXT, 4, uintptr(buffer), uintptr(offset), uintptr(size), uintptr(data), 0, 0)
}
func NamedCopyBufferSubDataEXT(readBuffer uint32, writeBuffer uint32, readOffset int, writeOffset int, size int) {
	syscall.Syscall6(gpNamedCopyBufferSubDataEXT, 5, uintptr(readBuffer), uintptr(writeBuffer), uintptr(readOffset), uintptr(writeOffset), uintptr(size), 0)
}

// specify which color buffers are to be drawn into
func NamedFramebufferDrawBuffer(framebuffer uint32, buf uint32) {
	syscall.Syscall(gpNamedFramebufferDrawBuffer, 2, uintptr(framebuffer), uintptr(buf), 0)
}

// Specifies a list of color buffers to be drawn     into
func NamedFramebufferDrawBuffers(framebuffer uint32, n int32, bufs *uint32) {
	syscall.Syscall(gpNamedFramebufferDrawBuffers, 3, uintptr(framebuffer), uintptr(n), uintptr(unsafe.Pointer(bufs)))
}

// set a named parameter of a framebuffer object
func NamedFramebufferParameteri(framebuffer uint32, pname uint32, param int32) {
	syscall.Syscall(gpNamedFramebufferParameteri, 3, uintptr(framebuffer), uintptr(pname), uintptr(param))
}
func NamedFramebufferParameteriEXT(framebuffer uint32, pname uint32, param int32) {
	syscall.Syscall(gpNamedFramebufferParameteriEXT, 3, uintptr(framebuffer), uintptr(pname), uintptr(param))
}

// select a color buffer source for pixels
func NamedFramebufferReadBuffer(framebuffer uint32, src uint32) {
	syscall.Syscall(gpNamedFramebufferReadBuffer, 2, uintptr(framebuffer), uintptr(src), 0)
}

// attach a renderbuffer as a logical buffer of a framebuffer object
func NamedFramebufferRenderbuffer(framebuffer uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	syscall.Syscall6(gpNamedFramebufferRenderbuffer, 4, uintptr(framebuffer), uintptr(attachment), uintptr(renderbuffertarget), uintptr(renderbuffer), 0, 0)
}
func NamedFramebufferRenderbufferEXT(framebuffer uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	syscall.Syscall6(gpNamedFramebufferRenderbufferEXT, 4, uintptr(framebuffer), uintptr(attachment), uintptr(renderbuffertarget), uintptr(renderbuffer), 0, 0)
}
func NamedFramebufferSampleLocationsfvARB(framebuffer uint32, start uint32, count int32, v *float32) {
	syscall.Syscall6(gpNamedFramebufferSampleLocationsfvARB, 4, uintptr(framebuffer), uintptr(start), uintptr(count), uintptr(unsafe.Pointer(v)), 0, 0)
}
func NamedFramebufferSampleLocationsfvNV(framebuffer uint32, start uint32, count int32, v *float32) {
	syscall.Syscall6(gpNamedFramebufferSampleLocationsfvNV, 4, uintptr(framebuffer), uintptr(start), uintptr(count), uintptr(unsafe.Pointer(v)), 0, 0)
}
func NamedFramebufferSamplePositionsfvAMD(framebuffer uint32, numsamples uint32, pixelindex uint32, values *float32) {
	syscall.Syscall6(gpNamedFramebufferSamplePositionsfvAMD, 4, uintptr(framebuffer), uintptr(numsamples), uintptr(pixelindex), uintptr(unsafe.Pointer(values)), 0, 0)
}
func NamedFramebufferTexture(framebuffer uint32, attachment uint32, texture uint32, level int32) {
	syscall.Syscall6(gpNamedFramebufferTexture, 4, uintptr(framebuffer), uintptr(attachment), uintptr(texture), uintptr(level), 0, 0)
}
func NamedFramebufferTexture1DEXT(framebuffer uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	syscall.Syscall6(gpNamedFramebufferTexture1DEXT, 5, uintptr(framebuffer), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), 0)
}
func NamedFramebufferTexture2DEXT(framebuffer uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	syscall.Syscall6(gpNamedFramebufferTexture2DEXT, 5, uintptr(framebuffer), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), 0)
}
func NamedFramebufferTexture3DEXT(framebuffer uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32) {
	syscall.Syscall6(gpNamedFramebufferTexture3DEXT, 6, uintptr(framebuffer), uintptr(attachment), uintptr(textarget), uintptr(texture), uintptr(level), uintptr(zoffset))
}
func NamedFramebufferTextureEXT(framebuffer uint32, attachment uint32, texture uint32, level int32) {
	syscall.Syscall6(gpNamedFramebufferTextureEXT, 4, uintptr(framebuffer), uintptr(attachment), uintptr(texture), uintptr(level), 0, 0)
}
func NamedFramebufferTextureFaceEXT(framebuffer uint32, attachment uint32, texture uint32, level int32, face uint32) {
	syscall.Syscall6(gpNamedFramebufferTextureFaceEXT, 5, uintptr(framebuffer), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(face), 0)
}

// attach a single layer of a texture object as a logical buffer of a framebuffer object
func NamedFramebufferTextureLayer(framebuffer uint32, attachment uint32, texture uint32, level int32, layer int32) {
	syscall.Syscall6(gpNamedFramebufferTextureLayer, 5, uintptr(framebuffer), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(layer), 0)
}
func NamedFramebufferTextureLayerEXT(framebuffer uint32, attachment uint32, texture uint32, level int32, layer int32) {
	syscall.Syscall6(gpNamedFramebufferTextureLayerEXT, 5, uintptr(framebuffer), uintptr(attachment), uintptr(texture), uintptr(level), uintptr(layer), 0)
}
func NamedProgramLocalParameter4dEXT(program uint32, target uint32, index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall9(gpNamedProgramLocalParameter4dEXT, 7, uintptr(program), uintptr(target), uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0, 0)
}
func NamedProgramLocalParameter4dvEXT(program uint32, target uint32, index uint32, params *float64) {
	syscall.Syscall6(gpNamedProgramLocalParameter4dvEXT, 4, uintptr(program), uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)), 0, 0)
}
func NamedProgramLocalParameter4fEXT(program uint32, target uint32, index uint32, x float32, y float32, z float32, w float32) {
	syscall.Syscall9(gpNamedProgramLocalParameter4fEXT, 7, uintptr(program), uintptr(target), uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0, 0)
}
func NamedProgramLocalParameter4fvEXT(program uint32, target uint32, index uint32, params *float32) {
	syscall.Syscall6(gpNamedProgramLocalParameter4fvEXT, 4, uintptr(program), uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)), 0, 0)
}
func NamedProgramLocalParameterI4iEXT(program uint32, target uint32, index uint32, x int32, y int32, z int32, w int32) {
	syscall.Syscall9(gpNamedProgramLocalParameterI4iEXT, 7, uintptr(program), uintptr(target), uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func NamedProgramLocalParameterI4ivEXT(program uint32, target uint32, index uint32, params *int32) {
	syscall.Syscall6(gpNamedProgramLocalParameterI4ivEXT, 4, uintptr(program), uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)), 0, 0)
}
func NamedProgramLocalParameterI4uiEXT(program uint32, target uint32, index uint32, x uint32, y uint32, z uint32, w uint32) {
	syscall.Syscall9(gpNamedProgramLocalParameterI4uiEXT, 7, uintptr(program), uintptr(target), uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func NamedProgramLocalParameterI4uivEXT(program uint32, target uint32, index uint32, params *uint32) {
	syscall.Syscall6(gpNamedProgramLocalParameterI4uivEXT, 4, uintptr(program), uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)), 0, 0)
}
func NamedProgramLocalParameters4fvEXT(program uint32, target uint32, index uint32, count int32, params *float32) {
	syscall.Syscall6(gpNamedProgramLocalParameters4fvEXT, 5, uintptr(program), uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0)
}
func NamedProgramLocalParametersI4ivEXT(program uint32, target uint32, index uint32, count int32, params *int32) {
	syscall.Syscall6(gpNamedProgramLocalParametersI4ivEXT, 5, uintptr(program), uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0)
}
func NamedProgramLocalParametersI4uivEXT(program uint32, target uint32, index uint32, count int32, params *uint32) {
	syscall.Syscall6(gpNamedProgramLocalParametersI4uivEXT, 5, uintptr(program), uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0)
}
func NamedProgramStringEXT(program uint32, target uint32, format uint32, len int32, xstring unsafe.Pointer) {
	syscall.Syscall6(gpNamedProgramStringEXT, 5, uintptr(program), uintptr(target), uintptr(format), uintptr(len), uintptr(xstring), 0)
}

// establish data storage, format and dimensions of a     renderbuffer object's image
func NamedRenderbufferStorage(renderbuffer uint32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpNamedRenderbufferStorage, 4, uintptr(renderbuffer), uintptr(internalformat), uintptr(width), uintptr(height), 0, 0)
}
func NamedRenderbufferStorageEXT(renderbuffer uint32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpNamedRenderbufferStorageEXT, 4, uintptr(renderbuffer), uintptr(internalformat), uintptr(width), uintptr(height), 0, 0)
}

// establish data storage, format, dimensions and sample count of     a renderbuffer object's image
func NamedRenderbufferStorageMultisample(renderbuffer uint32, samples int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpNamedRenderbufferStorageMultisample, 5, uintptr(renderbuffer), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), 0)
}
func NamedRenderbufferStorageMultisampleCoverageEXT(renderbuffer uint32, coverageSamples int32, colorSamples int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpNamedRenderbufferStorageMultisampleCoverageEXT, 6, uintptr(renderbuffer), uintptr(coverageSamples), uintptr(colorSamples), uintptr(internalformat), uintptr(width), uintptr(height))
}
func NamedRenderbufferStorageMultisampleEXT(renderbuffer uint32, samples int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpNamedRenderbufferStorageMultisampleEXT, 5, uintptr(renderbuffer), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), 0)
}
func NamedStringARB(xtype uint32, namelen int32, name *uint8, stringlen int32, xstring *uint8) {
	syscall.Syscall6(gpNamedStringARB, 5, uintptr(xtype), uintptr(namelen), uintptr(unsafe.Pointer(name)), uintptr(stringlen), uintptr(unsafe.Pointer(xstring)), 0)
}

// create or replace a display list
func NewList(list uint32, mode uint32) {
	syscall.Syscall(gpNewList, 2, uintptr(list), uintptr(mode), 0)
}
func NewObjectBufferATI(size int32, pointer unsafe.Pointer, usage uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpNewObjectBufferATI, 3, uintptr(size), uintptr(pointer), uintptr(usage))
	return (uint32)(ret)
}
func Normal3b(nx int8, ny int8, nz int8) {
	syscall.Syscall(gpNormal3b, 3, uintptr(nx), uintptr(ny), uintptr(nz))
}
func Normal3bv(v *int8) {
	syscall.Syscall(gpNormal3bv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Normal3d(nx float64, ny float64, nz float64) {
	syscall.Syscall(gpNormal3d, 3, uintptr(math.Float64bits(nx)), uintptr(math.Float64bits(ny)), uintptr(math.Float64bits(nz)))
}
func Normal3dv(v *float64) {
	syscall.Syscall(gpNormal3dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Normal3f(nx float32, ny float32, nz float32) {
	syscall.Syscall(gpNormal3f, 3, uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)))
}
func Normal3fVertex3fSUN(nx float32, ny float32, nz float32, x float32, y float32, z float32) {
	syscall.Syscall6(gpNormal3fVertex3fSUN, 6, uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func Normal3fVertex3fvSUN(n *float32, v *float32) {
	syscall.Syscall(gpNormal3fVertex3fvSUN, 2, uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)), 0)
}
func Normal3fv(v *float32) {
	syscall.Syscall(gpNormal3fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Normal3hNV(nx uint16, ny uint16, nz uint16) {
	syscall.Syscall(gpNormal3hNV, 3, uintptr(nx), uintptr(ny), uintptr(nz))
}
func Normal3hvNV(v *uint16) {
	syscall.Syscall(gpNormal3hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Normal3i(nx int32, ny int32, nz int32) {
	syscall.Syscall(gpNormal3i, 3, uintptr(nx), uintptr(ny), uintptr(nz))
}
func Normal3iv(v *int32) {
	syscall.Syscall(gpNormal3iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Normal3s(nx int16, ny int16, nz int16) {
	syscall.Syscall(gpNormal3s, 3, uintptr(nx), uintptr(ny), uintptr(nz))
}
func Normal3sv(v *int16) {
	syscall.Syscall(gpNormal3sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Normal3xOES(nx int32, ny int32, nz int32) {
	syscall.Syscall(gpNormal3xOES, 3, uintptr(nx), uintptr(ny), uintptr(nz))
}
func Normal3xvOES(coords *int32) {
	syscall.Syscall(gpNormal3xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func NormalFormatNV(xtype uint32, stride int32) {
	syscall.Syscall(gpNormalFormatNV, 2, uintptr(xtype), uintptr(stride), 0)
}

// define an array of normals
func NormalPointer(xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpNormalPointer, 3, uintptr(xtype), uintptr(stride), uintptr(pointer))
}
func NormalPointerEXT(xtype uint32, stride int32, count int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpNormalPointerEXT, 4, uintptr(xtype), uintptr(stride), uintptr(count), uintptr(pointer), 0, 0)
}
func NormalPointerListIBM(xtype uint32, stride int32, pointer *unsafe.Pointer, ptrstride int32) {
	syscall.Syscall6(gpNormalPointerListIBM, 4, uintptr(xtype), uintptr(stride), uintptr(unsafe.Pointer(pointer)), uintptr(ptrstride), 0, 0)
}
func NormalPointervINTEL(xtype uint32, pointer *unsafe.Pointer) {
	syscall.Syscall(gpNormalPointervINTEL, 2, uintptr(xtype), uintptr(unsafe.Pointer(pointer)), 0)
}
func NormalStream3bATI(stream uint32, nx int8, ny int8, nz int8) {
	syscall.Syscall6(gpNormalStream3bATI, 4, uintptr(stream), uintptr(nx), uintptr(ny), uintptr(nz), 0, 0)
}
func NormalStream3bvATI(stream uint32, coords *int8) {
	syscall.Syscall(gpNormalStream3bvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func NormalStream3dATI(stream uint32, nx float64, ny float64, nz float64) {
	syscall.Syscall6(gpNormalStream3dATI, 4, uintptr(stream), uintptr(math.Float64bits(nx)), uintptr(math.Float64bits(ny)), uintptr(math.Float64bits(nz)), 0, 0)
}
func NormalStream3dvATI(stream uint32, coords *float64) {
	syscall.Syscall(gpNormalStream3dvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func NormalStream3fATI(stream uint32, nx float32, ny float32, nz float32) {
	syscall.Syscall6(gpNormalStream3fATI, 4, uintptr(stream), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), 0, 0)
}
func NormalStream3fvATI(stream uint32, coords *float32) {
	syscall.Syscall(gpNormalStream3fvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func NormalStream3iATI(stream uint32, nx int32, ny int32, nz int32) {
	syscall.Syscall6(gpNormalStream3iATI, 4, uintptr(stream), uintptr(nx), uintptr(ny), uintptr(nz), 0, 0)
}
func NormalStream3ivATI(stream uint32, coords *int32) {
	syscall.Syscall(gpNormalStream3ivATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func NormalStream3sATI(stream uint32, nx int16, ny int16, nz int16) {
	syscall.Syscall6(gpNormalStream3sATI, 4, uintptr(stream), uintptr(nx), uintptr(ny), uintptr(nz), 0, 0)
}
func NormalStream3svATI(stream uint32, coords *int16) {
	syscall.Syscall(gpNormalStream3svATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}

// label a named object identified within a namespace
func ObjectLabel(identifier uint32, name uint32, length int32, label *uint8) {
	syscall.Syscall6(gpObjectLabel, 4, uintptr(identifier), uintptr(name), uintptr(length), uintptr(unsafe.Pointer(label)), 0, 0)
}
func ObjectLabelKHR(identifier uint32, name uint32, length int32, label *uint8) {
	syscall.Syscall6(gpObjectLabelKHR, 4, uintptr(identifier), uintptr(name), uintptr(length), uintptr(unsafe.Pointer(label)), 0, 0)
}

// label a a sync object identified by a pointer
func ObjectPtrLabel(ptr unsafe.Pointer, length int32, label *uint8) {
	syscall.Syscall(gpObjectPtrLabel, 3, uintptr(ptr), uintptr(length), uintptr(unsafe.Pointer(label)))
}
func ObjectPtrLabelKHR(ptr unsafe.Pointer, length int32, label *uint8) {
	syscall.Syscall(gpObjectPtrLabelKHR, 3, uintptr(ptr), uintptr(length), uintptr(unsafe.Pointer(label)))
}
func ObjectPurgeableAPPLE(objectType uint32, name uint32, option uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpObjectPurgeableAPPLE, 3, uintptr(objectType), uintptr(name), uintptr(option))
	return (uint32)(ret)
}
func ObjectUnpurgeableAPPLE(objectType uint32, name uint32, option uint32) uint32 {
	ret, _, _ := syscall.Syscall(gpObjectUnpurgeableAPPLE, 3, uintptr(objectType), uintptr(name), uintptr(option))
	return (uint32)(ret)
}

// multiply the current matrix with an orthographic matrix
func Ortho(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	syscall.Syscall6(gpOrtho, 6, uintptr(math.Float64bits(left)), uintptr(math.Float64bits(right)), uintptr(math.Float64bits(bottom)), uintptr(math.Float64bits(top)), uintptr(math.Float64bits(zNear)), uintptr(math.Float64bits(zFar)))
}
func OrthofOES(l float32, r float32, b float32, t float32, n float32, f float32) {
	syscall.Syscall6(gpOrthofOES, 6, uintptr(math.Float32bits(l)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(n)), uintptr(math.Float32bits(f)))
}
func OrthoxOES(l int32, r int32, b int32, t int32, n int32, f int32) {
	syscall.Syscall6(gpOrthoxOES, 6, uintptr(l), uintptr(r), uintptr(b), uintptr(t), uintptr(n), uintptr(f))
}
func PNTrianglesfATI(pname uint32, param float32) {
	syscall.Syscall(gpPNTrianglesfATI, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func PNTrianglesiATI(pname uint32, param int32) {
	syscall.Syscall(gpPNTrianglesiATI, 2, uintptr(pname), uintptr(param), 0)
}
func PassTexCoordATI(dst uint32, coord uint32, swizzle uint32) {
	syscall.Syscall(gpPassTexCoordATI, 3, uintptr(dst), uintptr(coord), uintptr(swizzle))
}

// place a marker in the feedback buffer
func PassThrough(token float32) {
	syscall.Syscall(gpPassThrough, 1, uintptr(math.Float32bits(token)), 0, 0)
}
func PassThroughxOES(token int32) {
	syscall.Syscall(gpPassThroughxOES, 1, uintptr(token), 0, 0)
}
func PatchParameterfv(pname uint32, values *float32) {
	syscall.Syscall(gpPatchParameterfv, 2, uintptr(pname), uintptr(unsafe.Pointer(values)), 0)
}

// specifies the parameters for patch primitives
func PatchParameteri(pname uint32, value int32) {
	syscall.Syscall(gpPatchParameteri, 2, uintptr(pname), uintptr(value), 0)
}
func PathCommandsNV(path uint32, numCommands int32, commands *uint8, numCoords int32, coordType uint32, coords unsafe.Pointer) {
	syscall.Syscall6(gpPathCommandsNV, 6, uintptr(path), uintptr(numCommands), uintptr(unsafe.Pointer(commands)), uintptr(numCoords), uintptr(coordType), uintptr(coords))
}
func PathCoordsNV(path uint32, numCoords int32, coordType uint32, coords unsafe.Pointer) {
	syscall.Syscall6(gpPathCoordsNV, 4, uintptr(path), uintptr(numCoords), uintptr(coordType), uintptr(coords), 0, 0)
}
func PathCoverDepthFuncNV(xfunc uint32) {
	syscall.Syscall(gpPathCoverDepthFuncNV, 1, uintptr(xfunc), 0, 0)
}
func PathDashArrayNV(path uint32, dashCount int32, dashArray *float32) {
	syscall.Syscall(gpPathDashArrayNV, 3, uintptr(path), uintptr(dashCount), uintptr(unsafe.Pointer(dashArray)))
}
func PathGlyphIndexArrayNV(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, firstGlyphIndex uint32, numGlyphs int32, pathParameterTemplate uint32, emScale float32) uint32 {
	ret, _, _ := syscall.Syscall9(gpPathGlyphIndexArrayNV, 8, uintptr(firstPathName), uintptr(fontTarget), uintptr(fontName), uintptr(fontStyle), uintptr(firstGlyphIndex), uintptr(numGlyphs), uintptr(pathParameterTemplate), uintptr(math.Float32bits(emScale)), 0)
	return (uint32)(ret)
}
func PathGlyphIndexRangeNV(fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, pathParameterTemplate uint32, emScale float32, baseAndCount *uint32) uint32 {
	ret, _, _ := syscall.Syscall6(gpPathGlyphIndexRangeNV, 6, uintptr(fontTarget), uintptr(fontName), uintptr(fontStyle), uintptr(pathParameterTemplate), uintptr(math.Float32bits(emScale)), uintptr(unsafe.Pointer(baseAndCount)))
	return (uint32)(ret)
}
func PathGlyphRangeNV(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, firstGlyph uint32, numGlyphs int32, handleMissingGlyphs uint32, pathParameterTemplate uint32, emScale float32) {
	syscall.Syscall9(gpPathGlyphRangeNV, 9, uintptr(firstPathName), uintptr(fontTarget), uintptr(fontName), uintptr(fontStyle), uintptr(firstGlyph), uintptr(numGlyphs), uintptr(handleMissingGlyphs), uintptr(pathParameterTemplate), uintptr(math.Float32bits(emScale)))
}
func PathGlyphsNV(firstPathName uint32, fontTarget uint32, fontName unsafe.Pointer, fontStyle uint32, numGlyphs int32, xtype uint32, charcodes unsafe.Pointer, handleMissingGlyphs uint32, pathParameterTemplate uint32, emScale float32) {
	syscall.Syscall12(gpPathGlyphsNV, 10, uintptr(firstPathName), uintptr(fontTarget), uintptr(fontName), uintptr(fontStyle), uintptr(numGlyphs), uintptr(xtype), uintptr(charcodes), uintptr(handleMissingGlyphs), uintptr(pathParameterTemplate), uintptr(math.Float32bits(emScale)), 0, 0)
}
func PathMemoryGlyphIndexArrayNV(firstPathName uint32, fontTarget uint32, fontSize int, fontData unsafe.Pointer, faceIndex int32, firstGlyphIndex uint32, numGlyphs int32, pathParameterTemplate uint32, emScale float32) uint32 {
	ret, _, _ := syscall.Syscall9(gpPathMemoryGlyphIndexArrayNV, 9, uintptr(firstPathName), uintptr(fontTarget), uintptr(fontSize), uintptr(fontData), uintptr(faceIndex), uintptr(firstGlyphIndex), uintptr(numGlyphs), uintptr(pathParameterTemplate), uintptr(math.Float32bits(emScale)))
	return (uint32)(ret)
}
func PathParameterfNV(path uint32, pname uint32, value float32) {
	syscall.Syscall(gpPathParameterfNV, 3, uintptr(path), uintptr(pname), uintptr(math.Float32bits(value)))
}
func PathParameterfvNV(path uint32, pname uint32, value *float32) {
	syscall.Syscall(gpPathParameterfvNV, 3, uintptr(path), uintptr(pname), uintptr(unsafe.Pointer(value)))
}
func PathParameteriNV(path uint32, pname uint32, value int32) {
	syscall.Syscall(gpPathParameteriNV, 3, uintptr(path), uintptr(pname), uintptr(value))
}
func PathParameterivNV(path uint32, pname uint32, value *int32) {
	syscall.Syscall(gpPathParameterivNV, 3, uintptr(path), uintptr(pname), uintptr(unsafe.Pointer(value)))
}
func PathStencilDepthOffsetNV(factor float32, units float32) {
	syscall.Syscall(gpPathStencilDepthOffsetNV, 2, uintptr(math.Float32bits(factor)), uintptr(math.Float32bits(units)), 0)
}
func PathStencilFuncNV(xfunc uint32, ref int32, mask uint32) {
	syscall.Syscall(gpPathStencilFuncNV, 3, uintptr(xfunc), uintptr(ref), uintptr(mask))
}
func PathStringNV(path uint32, format uint32, length int32, pathString unsafe.Pointer) {
	syscall.Syscall6(gpPathStringNV, 4, uintptr(path), uintptr(format), uintptr(length), uintptr(pathString), 0, 0)
}
func PathSubCommandsNV(path uint32, commandStart int32, commandsToDelete int32, numCommands int32, commands *uint8, numCoords int32, coordType uint32, coords unsafe.Pointer) {
	syscall.Syscall9(gpPathSubCommandsNV, 8, uintptr(path), uintptr(commandStart), uintptr(commandsToDelete), uintptr(numCommands), uintptr(unsafe.Pointer(commands)), uintptr(numCoords), uintptr(coordType), uintptr(coords), 0)
}
func PathSubCoordsNV(path uint32, coordStart int32, numCoords int32, coordType uint32, coords unsafe.Pointer) {
	syscall.Syscall6(gpPathSubCoordsNV, 5, uintptr(path), uintptr(coordStart), uintptr(numCoords), uintptr(coordType), uintptr(coords), 0)
}

// pause transform feedback operations
func PauseTransformFeedback() {
	syscall.Syscall(gpPauseTransformFeedback, 0, 0, 0, 0)
}
func PauseTransformFeedbackNV() {
	syscall.Syscall(gpPauseTransformFeedbackNV, 0, 0, 0, 0)
}
func PixelDataRangeNV(target uint32, length int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpPixelDataRangeNV, 3, uintptr(target), uintptr(length), uintptr(pointer))
}
func PixelMapfv(xmap uint32, mapsize int32, values *float32) {
	syscall.Syscall(gpPixelMapfv, 3, uintptr(xmap), uintptr(mapsize), uintptr(unsafe.Pointer(values)))
}
func PixelMapuiv(xmap uint32, mapsize int32, values *uint32) {
	syscall.Syscall(gpPixelMapuiv, 3, uintptr(xmap), uintptr(mapsize), uintptr(unsafe.Pointer(values)))
}
func PixelMapusv(xmap uint32, mapsize int32, values *uint16) {
	syscall.Syscall(gpPixelMapusv, 3, uintptr(xmap), uintptr(mapsize), uintptr(unsafe.Pointer(values)))
}
func PixelMapx(xmap uint32, size int32, values *int32) {
	syscall.Syscall(gpPixelMapx, 3, uintptr(xmap), uintptr(size), uintptr(unsafe.Pointer(values)))
}
func PixelStoref(pname uint32, param float32) {
	syscall.Syscall(gpPixelStoref, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}

// set pixel storage modes
func PixelStorei(pname uint32, param int32) {
	syscall.Syscall(gpPixelStorei, 2, uintptr(pname), uintptr(param), 0)
}
func PixelStorex(pname uint32, param int32) {
	syscall.Syscall(gpPixelStorex, 2, uintptr(pname), uintptr(param), 0)
}
func PixelTexGenParameterfSGIS(pname uint32, param float32) {
	syscall.Syscall(gpPixelTexGenParameterfSGIS, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func PixelTexGenParameterfvSGIS(pname uint32, params *float32) {
	syscall.Syscall(gpPixelTexGenParameterfvSGIS, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func PixelTexGenParameteriSGIS(pname uint32, param int32) {
	syscall.Syscall(gpPixelTexGenParameteriSGIS, 2, uintptr(pname), uintptr(param), 0)
}
func PixelTexGenParameterivSGIS(pname uint32, params *int32) {
	syscall.Syscall(gpPixelTexGenParameterivSGIS, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func PixelTexGenSGIX(mode uint32) {
	syscall.Syscall(gpPixelTexGenSGIX, 1, uintptr(mode), 0, 0)
}
func PixelTransferf(pname uint32, param float32) {
	syscall.Syscall(gpPixelTransferf, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func PixelTransferi(pname uint32, param int32) {
	syscall.Syscall(gpPixelTransferi, 2, uintptr(pname), uintptr(param), 0)
}
func PixelTransferxOES(pname uint32, param int32) {
	syscall.Syscall(gpPixelTransferxOES, 2, uintptr(pname), uintptr(param), 0)
}
func PixelTransformParameterfEXT(target uint32, pname uint32, param float32) {
	syscall.Syscall(gpPixelTransformParameterfEXT, 3, uintptr(target), uintptr(pname), uintptr(math.Float32bits(param)))
}
func PixelTransformParameterfvEXT(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpPixelTransformParameterfvEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func PixelTransformParameteriEXT(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpPixelTransformParameteriEXT, 3, uintptr(target), uintptr(pname), uintptr(param))
}
func PixelTransformParameterivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpPixelTransformParameterivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// specify the pixel zoom factors
func PixelZoom(xfactor float32, yfactor float32) {
	syscall.Syscall(gpPixelZoom, 2, uintptr(math.Float32bits(xfactor)), uintptr(math.Float32bits(yfactor)), 0)
}
func PixelZoomxOES(xfactor int32, yfactor int32) {
	syscall.Syscall(gpPixelZoomxOES, 2, uintptr(xfactor), uintptr(yfactor), 0)
}
func PointAlongPathNV(path uint32, startSegment int32, numSegments int32, distance float32, x *float32, y *float32, tangentX *float32, tangentY *float32) bool {
	ret, _, _ := syscall.Syscall9(gpPointAlongPathNV, 8, uintptr(path), uintptr(startSegment), uintptr(numSegments), uintptr(math.Float32bits(distance)), uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y)), uintptr(unsafe.Pointer(tangentX)), uintptr(unsafe.Pointer(tangentY)), 0)
	return ret != 0
}
func PointParameterf(pname uint32, param float32) {
	syscall.Syscall(gpPointParameterf, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func PointParameterfARB(pname uint32, param float32) {
	syscall.Syscall(gpPointParameterfARB, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func PointParameterfEXT(pname uint32, param float32) {
	syscall.Syscall(gpPointParameterfEXT, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func PointParameterfSGIS(pname uint32, param float32) {
	syscall.Syscall(gpPointParameterfSGIS, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func PointParameterfv(pname uint32, params *float32) {
	syscall.Syscall(gpPointParameterfv, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func PointParameterfvARB(pname uint32, params *float32) {
	syscall.Syscall(gpPointParameterfvARB, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func PointParameterfvEXT(pname uint32, params *float32) {
	syscall.Syscall(gpPointParameterfvEXT, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func PointParameterfvSGIS(pname uint32, params *float32) {
	syscall.Syscall(gpPointParameterfvSGIS, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func PointParameteri(pname uint32, param int32) {
	syscall.Syscall(gpPointParameteri, 2, uintptr(pname), uintptr(param), 0)
}
func PointParameteriNV(pname uint32, param int32) {
	syscall.Syscall(gpPointParameteriNV, 2, uintptr(pname), uintptr(param), 0)
}
func PointParameteriv(pname uint32, params *int32) {
	syscall.Syscall(gpPointParameteriv, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func PointParameterivNV(pname uint32, params *int32) {
	syscall.Syscall(gpPointParameterivNV, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func PointParameterxOES(pname uint32, param int32) {
	syscall.Syscall(gpPointParameterxOES, 2, uintptr(pname), uintptr(param), 0)
}
func PointParameterxvOES(pname uint32, params *int32) {
	syscall.Syscall(gpPointParameterxvOES, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}

// specify the diameter of rasterized points
func PointSize(size float32) {
	syscall.Syscall(gpPointSize, 1, uintptr(math.Float32bits(size)), 0, 0)
}
func PointSizexOES(size int32) {
	syscall.Syscall(gpPointSizexOES, 1, uintptr(size), 0, 0)
}
func PollAsyncSGIX(markerp *uint32) int32 {
	ret, _, _ := syscall.Syscall(gpPollAsyncSGIX, 1, uintptr(unsafe.Pointer(markerp)), 0, 0)
	return (int32)(ret)
}
func PollInstrumentsSGIX(marker_p *int32) int32 {
	ret, _, _ := syscall.Syscall(gpPollInstrumentsSGIX, 1, uintptr(unsafe.Pointer(marker_p)), 0, 0)
	return (int32)(ret)
}

// select a polygon rasterization mode
func PolygonMode(face uint32, mode uint32) {
	syscall.Syscall(gpPolygonMode, 2, uintptr(face), uintptr(mode), 0)
}

// set the scale and units used to calculate depth values
func PolygonOffset(factor float32, units float32) {
	syscall.Syscall(gpPolygonOffset, 2, uintptr(math.Float32bits(factor)), uintptr(math.Float32bits(units)), 0)
}
func PolygonOffsetClamp(factor float32, units float32, clamp float32) {
	syscall.Syscall(gpPolygonOffsetClamp, 3, uintptr(math.Float32bits(factor)), uintptr(math.Float32bits(units)), uintptr(math.Float32bits(clamp)))
}
func PolygonOffsetClampEXT(factor float32, units float32, clamp float32) {
	syscall.Syscall(gpPolygonOffsetClampEXT, 3, uintptr(math.Float32bits(factor)), uintptr(math.Float32bits(units)), uintptr(math.Float32bits(clamp)))
}
func PolygonOffsetEXT(factor float32, bias float32) {
	syscall.Syscall(gpPolygonOffsetEXT, 2, uintptr(math.Float32bits(factor)), uintptr(math.Float32bits(bias)), 0)
}
func PolygonOffsetxOES(factor int32, units int32) {
	syscall.Syscall(gpPolygonOffsetxOES, 2, uintptr(factor), uintptr(units), 0)
}

// set the polygon stippling pattern
func PolygonStipple(mask *uint8) {
	syscall.Syscall(gpPolygonStipple, 1, uintptr(unsafe.Pointer(mask)), 0, 0)
}
func PopAttrib() {
	syscall.Syscall(gpPopAttrib, 0, 0, 0, 0)
}
func PopClientAttrib() {
	syscall.Syscall(gpPopClientAttrib, 0, 0, 0, 0)
}

// pop the active debug group
func PopDebugGroup() {
	syscall.Syscall(gpPopDebugGroup, 0, 0, 0, 0)
}
func PopDebugGroupKHR() {
	syscall.Syscall(gpPopDebugGroupKHR, 0, 0, 0, 0)
}
func PopGroupMarkerEXT() {
	syscall.Syscall(gpPopGroupMarkerEXT, 0, 0, 0, 0)
}
func PopMatrix() {
	syscall.Syscall(gpPopMatrix, 0, 0, 0, 0)
}
func PopName() {
	syscall.Syscall(gpPopName, 0, 0, 0, 0)
}
func PresentFrameDualFillNV(video_slot uint32, minPresentTime uint64, beginPresentTimeId uint32, presentDurationId uint32, xtype uint32, target0 uint32, fill0 uint32, target1 uint32, fill1 uint32, target2 uint32, fill2 uint32, target3 uint32, fill3 uint32) {
	syscall.Syscall15(gpPresentFrameDualFillNV, 13, uintptr(video_slot), uintptr(minPresentTime), uintptr(beginPresentTimeId), uintptr(presentDurationId), uintptr(xtype), uintptr(target0), uintptr(fill0), uintptr(target1), uintptr(fill1), uintptr(target2), uintptr(fill2), uintptr(target3), uintptr(fill3), 0, 0)
}
func PresentFrameKeyedNV(video_slot uint32, minPresentTime uint64, beginPresentTimeId uint32, presentDurationId uint32, xtype uint32, target0 uint32, fill0 uint32, key0 uint32, target1 uint32, fill1 uint32, key1 uint32) {
	syscall.Syscall12(gpPresentFrameKeyedNV, 11, uintptr(video_slot), uintptr(minPresentTime), uintptr(beginPresentTimeId), uintptr(presentDurationId), uintptr(xtype), uintptr(target0), uintptr(fill0), uintptr(key0), uintptr(target1), uintptr(fill1), uintptr(key1), 0)
}
func PrimitiveBoundingBoxARB(minX float32, minY float32, minZ float32, minW float32, maxX float32, maxY float32, maxZ float32, maxW float32) {
	syscall.Syscall9(gpPrimitiveBoundingBoxARB, 8, uintptr(math.Float32bits(minX)), uintptr(math.Float32bits(minY)), uintptr(math.Float32bits(minZ)), uintptr(math.Float32bits(minW)), uintptr(math.Float32bits(maxX)), uintptr(math.Float32bits(maxY)), uintptr(math.Float32bits(maxZ)), uintptr(math.Float32bits(maxW)), 0)
}
func PrimitiveRestartIndexNV(index uint32) {
	syscall.Syscall(gpPrimitiveRestartIndexNV, 1, uintptr(index), 0, 0)
}
func PrimitiveRestartNV() {
	syscall.Syscall(gpPrimitiveRestartNV, 0, 0, 0, 0)
}

// set texture residence priority
func PrioritizeTextures(n int32, textures *uint32, priorities *float32) {
	syscall.Syscall(gpPrioritizeTextures, 3, uintptr(n), uintptr(unsafe.Pointer(textures)), uintptr(unsafe.Pointer(priorities)))
}
func PrioritizeTexturesEXT(n int32, textures *uint32, priorities *float32) {
	syscall.Syscall(gpPrioritizeTexturesEXT, 3, uintptr(n), uintptr(unsafe.Pointer(textures)), uintptr(unsafe.Pointer(priorities)))
}
func PrioritizeTexturesxOES(n int32, textures *uint32, priorities *int32) {
	syscall.Syscall(gpPrioritizeTexturesxOES, 3, uintptr(n), uintptr(unsafe.Pointer(textures)), uintptr(unsafe.Pointer(priorities)))
}

// load a program object with a program binary
func ProgramBinary(program uint32, binaryFormat uint32, binary unsafe.Pointer, length int32) {
	syscall.Syscall6(gpProgramBinary, 4, uintptr(program), uintptr(binaryFormat), uintptr(binary), uintptr(length), 0, 0)
}
func ProgramBufferParametersIivNV(target uint32, bindingIndex uint32, wordIndex uint32, count int32, params *int32) {
	syscall.Syscall6(gpProgramBufferParametersIivNV, 5, uintptr(target), uintptr(bindingIndex), uintptr(wordIndex), uintptr(count), uintptr(unsafe.Pointer(params)), 0)
}
func ProgramBufferParametersIuivNV(target uint32, bindingIndex uint32, wordIndex uint32, count int32, params *uint32) {
	syscall.Syscall6(gpProgramBufferParametersIuivNV, 5, uintptr(target), uintptr(bindingIndex), uintptr(wordIndex), uintptr(count), uintptr(unsafe.Pointer(params)), 0)
}
func ProgramBufferParametersfvNV(target uint32, bindingIndex uint32, wordIndex uint32, count int32, params *float32) {
	syscall.Syscall6(gpProgramBufferParametersfvNV, 5, uintptr(target), uintptr(bindingIndex), uintptr(wordIndex), uintptr(count), uintptr(unsafe.Pointer(params)), 0)
}
func ProgramEnvParameter4dARB(target uint32, index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpProgramEnvParameter4dARB, 6, uintptr(target), uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)))
}
func ProgramEnvParameter4dvARB(target uint32, index uint32, params *float64) {
	syscall.Syscall(gpProgramEnvParameter4dvARB, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func ProgramEnvParameter4fARB(target uint32, index uint32, x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpProgramEnvParameter4fARB, 6, uintptr(target), uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)))
}
func ProgramEnvParameter4fvARB(target uint32, index uint32, params *float32) {
	syscall.Syscall(gpProgramEnvParameter4fvARB, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func ProgramEnvParameterI4iNV(target uint32, index uint32, x int32, y int32, z int32, w int32) {
	syscall.Syscall6(gpProgramEnvParameterI4iNV, 6, uintptr(target), uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
}
func ProgramEnvParameterI4ivNV(target uint32, index uint32, params *int32) {
	syscall.Syscall(gpProgramEnvParameterI4ivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func ProgramEnvParameterI4uiNV(target uint32, index uint32, x uint32, y uint32, z uint32, w uint32) {
	syscall.Syscall6(gpProgramEnvParameterI4uiNV, 6, uintptr(target), uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
}
func ProgramEnvParameterI4uivNV(target uint32, index uint32, params *uint32) {
	syscall.Syscall(gpProgramEnvParameterI4uivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func ProgramEnvParameters4fvEXT(target uint32, index uint32, count int32, params *float32) {
	syscall.Syscall6(gpProgramEnvParameters4fvEXT, 4, uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0, 0)
}
func ProgramEnvParametersI4ivNV(target uint32, index uint32, count int32, params *int32) {
	syscall.Syscall6(gpProgramEnvParametersI4ivNV, 4, uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0, 0)
}
func ProgramEnvParametersI4uivNV(target uint32, index uint32, count int32, params *uint32) {
	syscall.Syscall6(gpProgramEnvParametersI4uivNV, 4, uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0, 0)
}
func ProgramLocalParameter4dARB(target uint32, index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpProgramLocalParameter4dARB, 6, uintptr(target), uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)))
}
func ProgramLocalParameter4dvARB(target uint32, index uint32, params *float64) {
	syscall.Syscall(gpProgramLocalParameter4dvARB, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func ProgramLocalParameter4fARB(target uint32, index uint32, x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpProgramLocalParameter4fARB, 6, uintptr(target), uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)))
}
func ProgramLocalParameter4fvARB(target uint32, index uint32, params *float32) {
	syscall.Syscall(gpProgramLocalParameter4fvARB, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func ProgramLocalParameterI4iNV(target uint32, index uint32, x int32, y int32, z int32, w int32) {
	syscall.Syscall6(gpProgramLocalParameterI4iNV, 6, uintptr(target), uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
}
func ProgramLocalParameterI4ivNV(target uint32, index uint32, params *int32) {
	syscall.Syscall(gpProgramLocalParameterI4ivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func ProgramLocalParameterI4uiNV(target uint32, index uint32, x uint32, y uint32, z uint32, w uint32) {
	syscall.Syscall6(gpProgramLocalParameterI4uiNV, 6, uintptr(target), uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
}
func ProgramLocalParameterI4uivNV(target uint32, index uint32, params *uint32) {
	syscall.Syscall(gpProgramLocalParameterI4uivNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(params)))
}
func ProgramLocalParameters4fvEXT(target uint32, index uint32, count int32, params *float32) {
	syscall.Syscall6(gpProgramLocalParameters4fvEXT, 4, uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0, 0)
}
func ProgramLocalParametersI4ivNV(target uint32, index uint32, count int32, params *int32) {
	syscall.Syscall6(gpProgramLocalParametersI4ivNV, 4, uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0, 0)
}
func ProgramLocalParametersI4uivNV(target uint32, index uint32, count int32, params *uint32) {
	syscall.Syscall6(gpProgramLocalParametersI4uivNV, 4, uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(params)), 0, 0)
}
func ProgramNamedParameter4dNV(id uint32, len int32, name *uint8, x float64, y float64, z float64, w float64) {
	syscall.Syscall9(gpProgramNamedParameter4dNV, 7, uintptr(id), uintptr(len), uintptr(unsafe.Pointer(name)), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0, 0)
}
func ProgramNamedParameter4dvNV(id uint32, len int32, name *uint8, v *float64) {
	syscall.Syscall6(gpProgramNamedParameter4dvNV, 4, uintptr(id), uintptr(len), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(v)), 0, 0)
}
func ProgramNamedParameter4fNV(id uint32, len int32, name *uint8, x float32, y float32, z float32, w float32) {
	syscall.Syscall9(gpProgramNamedParameter4fNV, 7, uintptr(id), uintptr(len), uintptr(unsafe.Pointer(name)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0, 0)
}
func ProgramNamedParameter4fvNV(id uint32, len int32, name *uint8, v *float32) {
	syscall.Syscall6(gpProgramNamedParameter4fvNV, 4, uintptr(id), uintptr(len), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(v)), 0, 0)
}
func ProgramParameter4dNV(target uint32, index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpProgramParameter4dNV, 6, uintptr(target), uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)))
}
func ProgramParameter4dvNV(target uint32, index uint32, v *float64) {
	syscall.Syscall(gpProgramParameter4dvNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(v)))
}
func ProgramParameter4fNV(target uint32, index uint32, x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpProgramParameter4fNV, 6, uintptr(target), uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)))
}
func ProgramParameter4fvNV(target uint32, index uint32, v *float32) {
	syscall.Syscall(gpProgramParameter4fvNV, 3, uintptr(target), uintptr(index), uintptr(unsafe.Pointer(v)))
}

// specify a parameter for a program object
func ProgramParameteri(program uint32, pname uint32, value int32) {
	syscall.Syscall(gpProgramParameteri, 3, uintptr(program), uintptr(pname), uintptr(value))
}
func ProgramParameteriARB(program uint32, pname uint32, value int32) {
	syscall.Syscall(gpProgramParameteriARB, 3, uintptr(program), uintptr(pname), uintptr(value))
}
func ProgramParameteriEXT(program uint32, pname uint32, value int32) {
	syscall.Syscall(gpProgramParameteriEXT, 3, uintptr(program), uintptr(pname), uintptr(value))
}
func ProgramParameters4dvNV(target uint32, index uint32, count int32, v *float64) {
	syscall.Syscall6(gpProgramParameters4dvNV, 4, uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)), 0, 0)
}
func ProgramParameters4fvNV(target uint32, index uint32, count int32, v *float32) {
	syscall.Syscall6(gpProgramParameters4fvNV, 4, uintptr(target), uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)), 0, 0)
}
func ProgramPathFragmentInputGenNV(program uint32, location int32, genMode uint32, components int32, coeffs *float32) {
	syscall.Syscall6(gpProgramPathFragmentInputGenNV, 5, uintptr(program), uintptr(location), uintptr(genMode), uintptr(components), uintptr(unsafe.Pointer(coeffs)), 0)
}
func ProgramStringARB(target uint32, format uint32, len int32, xstring unsafe.Pointer) {
	syscall.Syscall6(gpProgramStringARB, 4, uintptr(target), uintptr(format), uintptr(len), uintptr(xstring), 0, 0)
}
func ProgramSubroutineParametersuivNV(target uint32, count int32, params *uint32) {
	syscall.Syscall(gpProgramSubroutineParametersuivNV, 3, uintptr(target), uintptr(count), uintptr(unsafe.Pointer(params)))
}
func ProgramUniform1d(program uint32, location int32, v0 float64) {
	syscall.Syscall(gpProgramUniform1d, 3, uintptr(program), uintptr(location), uintptr(math.Float64bits(v0)))
}
func ProgramUniform1dEXT(program uint32, location int32, x float64) {
	syscall.Syscall(gpProgramUniform1dEXT, 3, uintptr(program), uintptr(location), uintptr(math.Float64bits(x)))
}
func ProgramUniform1dv(program uint32, location int32, count int32, value *float64) {
	syscall.Syscall6(gpProgramUniform1dv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform1dvEXT(program uint32, location int32, count int32, value *float64) {
	syscall.Syscall6(gpProgramUniform1dvEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1f(program uint32, location int32, v0 float32) {
	syscall.Syscall(gpProgramUniform1f, 3, uintptr(program), uintptr(location), uintptr(math.Float32bits(v0)))
}
func ProgramUniform1fEXT(program uint32, location int32, v0 float32) {
	syscall.Syscall(gpProgramUniform1fEXT, 3, uintptr(program), uintptr(location), uintptr(math.Float32bits(v0)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1fv(program uint32, location int32, count int32, value *float32) {
	syscall.Syscall6(gpProgramUniform1fv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform1fvEXT(program uint32, location int32, count int32, value *float32) {
	syscall.Syscall6(gpProgramUniform1fvEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1i(program uint32, location int32, v0 int32) {
	syscall.Syscall(gpProgramUniform1i, 3, uintptr(program), uintptr(location), uintptr(v0))
}
func ProgramUniform1i64ARB(program uint32, location int32, x int64) {
	syscall.Syscall(gpProgramUniform1i64ARB, 3, uintptr(program), uintptr(location), uintptr(x))
}
func ProgramUniform1i64NV(program uint32, location int32, x int64) {
	syscall.Syscall(gpProgramUniform1i64NV, 3, uintptr(program), uintptr(location), uintptr(x))
}
func ProgramUniform1i64vARB(program uint32, location int32, count int32, value *int64) {
	syscall.Syscall6(gpProgramUniform1i64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform1i64vNV(program uint32, location int32, count int32, value *int64) {
	syscall.Syscall6(gpProgramUniform1i64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform1iEXT(program uint32, location int32, v0 int32) {
	syscall.Syscall(gpProgramUniform1iEXT, 3, uintptr(program), uintptr(location), uintptr(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1iv(program uint32, location int32, count int32, value *int32) {
	syscall.Syscall6(gpProgramUniform1iv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform1ivEXT(program uint32, location int32, count int32, value *int32) {
	syscall.Syscall6(gpProgramUniform1ivEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1ui(program uint32, location int32, v0 uint32) {
	syscall.Syscall(gpProgramUniform1ui, 3, uintptr(program), uintptr(location), uintptr(v0))
}
func ProgramUniform1ui64ARB(program uint32, location int32, x uint64) {
	syscall.Syscall(gpProgramUniform1ui64ARB, 3, uintptr(program), uintptr(location), uintptr(x))
}
func ProgramUniform1ui64NV(program uint32, location int32, x uint64) {
	syscall.Syscall(gpProgramUniform1ui64NV, 3, uintptr(program), uintptr(location), uintptr(x))
}
func ProgramUniform1ui64vARB(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniform1ui64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform1ui64vNV(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniform1ui64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform1uiEXT(program uint32, location int32, v0 uint32) {
	syscall.Syscall(gpProgramUniform1uiEXT, 3, uintptr(program), uintptr(location), uintptr(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1uiv(program uint32, location int32, count int32, value *uint32) {
	syscall.Syscall6(gpProgramUniform1uiv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform1uivEXT(program uint32, location int32, count int32, value *uint32) {
	syscall.Syscall6(gpProgramUniform1uivEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2d(program uint32, location int32, v0 float64, v1 float64) {
	syscall.Syscall6(gpProgramUniform2d, 4, uintptr(program), uintptr(location), uintptr(math.Float64bits(v0)), uintptr(math.Float64bits(v1)), 0, 0)
}
func ProgramUniform2dEXT(program uint32, location int32, x float64, y float64) {
	syscall.Syscall6(gpProgramUniform2dEXT, 4, uintptr(program), uintptr(location), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), 0, 0)
}
func ProgramUniform2dv(program uint32, location int32, count int32, value *float64) {
	syscall.Syscall6(gpProgramUniform2dv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2dvEXT(program uint32, location int32, count int32, value *float64) {
	syscall.Syscall6(gpProgramUniform2dvEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2f(program uint32, location int32, v0 float32, v1 float32) {
	syscall.Syscall6(gpProgramUniform2f, 4, uintptr(program), uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), 0, 0)
}
func ProgramUniform2fEXT(program uint32, location int32, v0 float32, v1 float32) {
	syscall.Syscall6(gpProgramUniform2fEXT, 4, uintptr(program), uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2fv(program uint32, location int32, count int32, value *float32) {
	syscall.Syscall6(gpProgramUniform2fv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2fvEXT(program uint32, location int32, count int32, value *float32) {
	syscall.Syscall6(gpProgramUniform2fvEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2i(program uint32, location int32, v0 int32, v1 int32) {
	syscall.Syscall6(gpProgramUniform2i, 4, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), 0, 0)
}
func ProgramUniform2i64ARB(program uint32, location int32, x int64, y int64) {
	syscall.Syscall6(gpProgramUniform2i64ARB, 4, uintptr(program), uintptr(location), uintptr(x), uintptr(y), 0, 0)
}
func ProgramUniform2i64NV(program uint32, location int32, x int64, y int64) {
	syscall.Syscall6(gpProgramUniform2i64NV, 4, uintptr(program), uintptr(location), uintptr(x), uintptr(y), 0, 0)
}
func ProgramUniform2i64vARB(program uint32, location int32, count int32, value *int64) {
	syscall.Syscall6(gpProgramUniform2i64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2i64vNV(program uint32, location int32, count int32, value *int64) {
	syscall.Syscall6(gpProgramUniform2i64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2iEXT(program uint32, location int32, v0 int32, v1 int32) {
	syscall.Syscall6(gpProgramUniform2iEXT, 4, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2iv(program uint32, location int32, count int32, value *int32) {
	syscall.Syscall6(gpProgramUniform2iv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2ivEXT(program uint32, location int32, count int32, value *int32) {
	syscall.Syscall6(gpProgramUniform2ivEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2ui(program uint32, location int32, v0 uint32, v1 uint32) {
	syscall.Syscall6(gpProgramUniform2ui, 4, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), 0, 0)
}
func ProgramUniform2ui64ARB(program uint32, location int32, x uint64, y uint64) {
	syscall.Syscall6(gpProgramUniform2ui64ARB, 4, uintptr(program), uintptr(location), uintptr(x), uintptr(y), 0, 0)
}
func ProgramUniform2ui64NV(program uint32, location int32, x uint64, y uint64) {
	syscall.Syscall6(gpProgramUniform2ui64NV, 4, uintptr(program), uintptr(location), uintptr(x), uintptr(y), 0, 0)
}
func ProgramUniform2ui64vARB(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniform2ui64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2ui64vNV(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniform2ui64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2uiEXT(program uint32, location int32, v0 uint32, v1 uint32) {
	syscall.Syscall6(gpProgramUniform2uiEXT, 4, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2uiv(program uint32, location int32, count int32, value *uint32) {
	syscall.Syscall6(gpProgramUniform2uiv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform2uivEXT(program uint32, location int32, count int32, value *uint32) {
	syscall.Syscall6(gpProgramUniform2uivEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3d(program uint32, location int32, v0 float64, v1 float64, v2 float64) {
	syscall.Syscall6(gpProgramUniform3d, 5, uintptr(program), uintptr(location), uintptr(math.Float64bits(v0)), uintptr(math.Float64bits(v1)), uintptr(math.Float64bits(v2)), 0)
}
func ProgramUniform3dEXT(program uint32, location int32, x float64, y float64, z float64) {
	syscall.Syscall6(gpProgramUniform3dEXT, 5, uintptr(program), uintptr(location), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0)
}
func ProgramUniform3dv(program uint32, location int32, count int32, value *float64) {
	syscall.Syscall6(gpProgramUniform3dv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3dvEXT(program uint32, location int32, count int32, value *float64) {
	syscall.Syscall6(gpProgramUniform3dvEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3f(program uint32, location int32, v0 float32, v1 float32, v2 float32) {
	syscall.Syscall6(gpProgramUniform3f, 5, uintptr(program), uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), 0)
}
func ProgramUniform3fEXT(program uint32, location int32, v0 float32, v1 float32, v2 float32) {
	syscall.Syscall6(gpProgramUniform3fEXT, 5, uintptr(program), uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3fv(program uint32, location int32, count int32, value *float32) {
	syscall.Syscall6(gpProgramUniform3fv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3fvEXT(program uint32, location int32, count int32, value *float32) {
	syscall.Syscall6(gpProgramUniform3fvEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3i(program uint32, location int32, v0 int32, v1 int32, v2 int32) {
	syscall.Syscall6(gpProgramUniform3i, 5, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), 0)
}
func ProgramUniform3i64ARB(program uint32, location int32, x int64, y int64, z int64) {
	syscall.Syscall6(gpProgramUniform3i64ARB, 5, uintptr(program), uintptr(location), uintptr(x), uintptr(y), uintptr(z), 0)
}
func ProgramUniform3i64NV(program uint32, location int32, x int64, y int64, z int64) {
	syscall.Syscall6(gpProgramUniform3i64NV, 5, uintptr(program), uintptr(location), uintptr(x), uintptr(y), uintptr(z), 0)
}
func ProgramUniform3i64vARB(program uint32, location int32, count int32, value *int64) {
	syscall.Syscall6(gpProgramUniform3i64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3i64vNV(program uint32, location int32, count int32, value *int64) {
	syscall.Syscall6(gpProgramUniform3i64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3iEXT(program uint32, location int32, v0 int32, v1 int32, v2 int32) {
	syscall.Syscall6(gpProgramUniform3iEXT, 5, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3iv(program uint32, location int32, count int32, value *int32) {
	syscall.Syscall6(gpProgramUniform3iv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3ivEXT(program uint32, location int32, count int32, value *int32) {
	syscall.Syscall6(gpProgramUniform3ivEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32) {
	syscall.Syscall6(gpProgramUniform3ui, 5, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), 0)
}
func ProgramUniform3ui64ARB(program uint32, location int32, x uint64, y uint64, z uint64) {
	syscall.Syscall6(gpProgramUniform3ui64ARB, 5, uintptr(program), uintptr(location), uintptr(x), uintptr(y), uintptr(z), 0)
}
func ProgramUniform3ui64NV(program uint32, location int32, x uint64, y uint64, z uint64) {
	syscall.Syscall6(gpProgramUniform3ui64NV, 5, uintptr(program), uintptr(location), uintptr(x), uintptr(y), uintptr(z), 0)
}
func ProgramUniform3ui64vARB(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniform3ui64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3ui64vNV(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniform3ui64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3uiEXT(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32) {
	syscall.Syscall6(gpProgramUniform3uiEXT, 5, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3uiv(program uint32, location int32, count int32, value *uint32) {
	syscall.Syscall6(gpProgramUniform3uiv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform3uivEXT(program uint32, location int32, count int32, value *uint32) {
	syscall.Syscall6(gpProgramUniform3uivEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4d(program uint32, location int32, v0 float64, v1 float64, v2 float64, v3 float64) {
	syscall.Syscall6(gpProgramUniform4d, 6, uintptr(program), uintptr(location), uintptr(math.Float64bits(v0)), uintptr(math.Float64bits(v1)), uintptr(math.Float64bits(v2)), uintptr(math.Float64bits(v3)))
}
func ProgramUniform4dEXT(program uint32, location int32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpProgramUniform4dEXT, 6, uintptr(program), uintptr(location), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)))
}
func ProgramUniform4dv(program uint32, location int32, count int32, value *float64) {
	syscall.Syscall6(gpProgramUniform4dv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4dvEXT(program uint32, location int32, count int32, value *float64) {
	syscall.Syscall6(gpProgramUniform4dvEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4f(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	syscall.Syscall6(gpProgramUniform4f, 6, uintptr(program), uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), uintptr(math.Float32bits(v3)))
}
func ProgramUniform4fEXT(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	syscall.Syscall6(gpProgramUniform4fEXT, 6, uintptr(program), uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), uintptr(math.Float32bits(v3)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4fv(program uint32, location int32, count int32, value *float32) {
	syscall.Syscall6(gpProgramUniform4fv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4fvEXT(program uint32, location int32, count int32, value *float32) {
	syscall.Syscall6(gpProgramUniform4fvEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4i(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	syscall.Syscall6(gpProgramUniform4i, 6, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3))
}
func ProgramUniform4i64ARB(program uint32, location int32, x int64, y int64, z int64, w int64) {
	syscall.Syscall6(gpProgramUniform4i64ARB, 6, uintptr(program), uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
}
func ProgramUniform4i64NV(program uint32, location int32, x int64, y int64, z int64, w int64) {
	syscall.Syscall6(gpProgramUniform4i64NV, 6, uintptr(program), uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
}
func ProgramUniform4i64vARB(program uint32, location int32, count int32, value *int64) {
	syscall.Syscall6(gpProgramUniform4i64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4i64vNV(program uint32, location int32, count int32, value *int64) {
	syscall.Syscall6(gpProgramUniform4i64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4iEXT(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	syscall.Syscall6(gpProgramUniform4iEXT, 6, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4iv(program uint32, location int32, count int32, value *int32) {
	syscall.Syscall6(gpProgramUniform4iv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4ivEXT(program uint32, location int32, count int32, value *int32) {
	syscall.Syscall6(gpProgramUniform4ivEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	syscall.Syscall6(gpProgramUniform4ui, 6, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3))
}
func ProgramUniform4ui64ARB(program uint32, location int32, x uint64, y uint64, z uint64, w uint64) {
	syscall.Syscall6(gpProgramUniform4ui64ARB, 6, uintptr(program), uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
}
func ProgramUniform4ui64NV(program uint32, location int32, x uint64, y uint64, z uint64, w uint64) {
	syscall.Syscall6(gpProgramUniform4ui64NV, 6, uintptr(program), uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w))
}
func ProgramUniform4ui64vARB(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniform4ui64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4ui64vNV(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniform4ui64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4uiEXT(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	syscall.Syscall6(gpProgramUniform4uiEXT, 6, uintptr(program), uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4uiv(program uint32, location int32, count int32, value *uint32) {
	syscall.Syscall6(gpProgramUniform4uiv, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniform4uivEXT(program uint32, location int32, count int32, value *uint32) {
	syscall.Syscall6(gpProgramUniform4uivEXT, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramUniformHandleui64ARB(program uint32, location int32, value uint64) {
	syscall.Syscall(gpProgramUniformHandleui64ARB, 3, uintptr(program), uintptr(location), uintptr(value))
}
func ProgramUniformHandleui64NV(program uint32, location int32, value uint64) {
	syscall.Syscall(gpProgramUniformHandleui64NV, 3, uintptr(program), uintptr(location), uintptr(value))
}
func ProgramUniformHandleui64vARB(program uint32, location int32, count int32, values *uint64) {
	syscall.Syscall6(gpProgramUniformHandleui64vARB, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(values)), 0, 0)
}
func ProgramUniformHandleui64vNV(program uint32, location int32, count int32, values *uint64) {
	syscall.Syscall6(gpProgramUniformHandleui64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(values)), 0, 0)
}
func ProgramUniformMatrix2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix2dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix2dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix2dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix2fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix2fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix2x3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix2x3dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix2x3dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix2x3dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix2x3fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix2x3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix2x3fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix2x4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix2x4dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix2x4dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix2x4dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix2x4fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix2x4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix2x4fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix3dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix3dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix3fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix3fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3x2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix3x2dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3x2dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix3x2dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix3x2fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3x2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix3x2fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3x4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix3x4dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3x4dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix3x4dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix3x4fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix3x4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix3x4fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix4dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix4dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix4fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix4fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4x2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix4x2dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4x2dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix4x2dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix4x2fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4x2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix4x2fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4x3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix4x3dv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4x3dvEXT(program uint32, location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpProgramUniformMatrix4x3dvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix4x3fv, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformMatrix4x3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpProgramUniformMatrix4x3fvEXT, 5, uintptr(program), uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0)
}
func ProgramUniformui64NV(program uint32, location int32, value uint64) {
	syscall.Syscall(gpProgramUniformui64NV, 3, uintptr(program), uintptr(location), uintptr(value))
}
func ProgramUniformui64vNV(program uint32, location int32, count int32, value *uint64) {
	syscall.Syscall6(gpProgramUniformui64vNV, 4, uintptr(program), uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)), 0, 0)
}
func ProgramVertexLimitNV(target uint32, limit int32) {
	syscall.Syscall(gpProgramVertexLimitNV, 2, uintptr(target), uintptr(limit), 0)
}

// specifiy the vertex to be used as the source of data for flat shaded varyings
func ProvokingVertex(mode uint32) {
	syscall.Syscall(gpProvokingVertex, 1, uintptr(mode), 0, 0)
}
func ProvokingVertexEXT(mode uint32) {
	syscall.Syscall(gpProvokingVertexEXT, 1, uintptr(mode), 0, 0)
}

// push and pop the server attribute stack
func PushAttrib(mask uint32) {
	syscall.Syscall(gpPushAttrib, 1, uintptr(mask), 0, 0)
}

// push and pop the client attribute stack
func PushClientAttrib(mask uint32) {
	syscall.Syscall(gpPushClientAttrib, 1, uintptr(mask), 0, 0)
}
func PushClientAttribDefaultEXT(mask uint32) {
	syscall.Syscall(gpPushClientAttribDefaultEXT, 1, uintptr(mask), 0, 0)
}

// push a named debug group into the command stream
func PushDebugGroup(source uint32, id uint32, length int32, message *uint8) {
	syscall.Syscall6(gpPushDebugGroup, 4, uintptr(source), uintptr(id), uintptr(length), uintptr(unsafe.Pointer(message)), 0, 0)
}
func PushDebugGroupKHR(source uint32, id uint32, length int32, message *uint8) {
	syscall.Syscall6(gpPushDebugGroupKHR, 4, uintptr(source), uintptr(id), uintptr(length), uintptr(unsafe.Pointer(message)), 0, 0)
}
func PushGroupMarkerEXT(length int32, marker *uint8) {
	syscall.Syscall(gpPushGroupMarkerEXT, 2, uintptr(length), uintptr(unsafe.Pointer(marker)), 0)
}

// push and pop the current matrix stack
func PushMatrix() {
	syscall.Syscall(gpPushMatrix, 0, 0, 0, 0)
}

// push and pop the name stack
func PushName(name uint32) {
	syscall.Syscall(gpPushName, 1, uintptr(name), 0, 0)
}

// record the GL time into a query object after all previous commands have reached the GL server but have not yet necessarily executed.
func QueryCounter(id uint32, target uint32) {
	syscall.Syscall(gpQueryCounter, 2, uintptr(id), uintptr(target), 0)
}

// return the values of the current matrix
func QueryMatrixxOES(mantissa *int32, exponent *int32) uint32 {
	ret, _, _ := syscall.Syscall(gpQueryMatrixxOES, 2, uintptr(unsafe.Pointer(mantissa)), uintptr(unsafe.Pointer(exponent)), 0)
	return (uint32)(ret)
}
func QueryObjectParameteruiAMD(target uint32, id uint32, pname uint32, param uint32) {
	syscall.Syscall6(gpQueryObjectParameteruiAMD, 4, uintptr(target), uintptr(id), uintptr(pname), uintptr(param), 0, 0)
}
func QueryResourceNV(queryType uint32, tagId int32, bufSize uint32, buffer *int32) int32 {
	ret, _, _ := syscall.Syscall6(gpQueryResourceNV, 4, uintptr(queryType), uintptr(tagId), uintptr(bufSize), uintptr(unsafe.Pointer(buffer)), 0, 0)
	return (int32)(ret)
}
func QueryResourceTagNV(tagId int32, tagString *uint8) {
	syscall.Syscall(gpQueryResourceTagNV, 2, uintptr(tagId), uintptr(unsafe.Pointer(tagString)), 0)
}
func RasterPos2d(x float64, y float64) {
	syscall.Syscall(gpRasterPos2d, 2, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), 0)
}
func RasterPos2dv(v *float64) {
	syscall.Syscall(gpRasterPos2dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos2f(x float32, y float32) {
	syscall.Syscall(gpRasterPos2f, 2, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), 0)
}
func RasterPos2fv(v *float32) {
	syscall.Syscall(gpRasterPos2fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos2i(x int32, y int32) {
	syscall.Syscall(gpRasterPos2i, 2, uintptr(x), uintptr(y), 0)
}
func RasterPos2iv(v *int32) {
	syscall.Syscall(gpRasterPos2iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos2s(x int16, y int16) {
	syscall.Syscall(gpRasterPos2s, 2, uintptr(x), uintptr(y), 0)
}
func RasterPos2sv(v *int16) {
	syscall.Syscall(gpRasterPos2sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos2xOES(x int32, y int32) {
	syscall.Syscall(gpRasterPos2xOES, 2, uintptr(x), uintptr(y), 0)
}
func RasterPos2xvOES(coords *int32) {
	syscall.Syscall(gpRasterPos2xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func RasterPos3d(x float64, y float64, z float64) {
	syscall.Syscall(gpRasterPos3d, 3, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)))
}
func RasterPos3dv(v *float64) {
	syscall.Syscall(gpRasterPos3dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos3f(x float32, y float32, z float32) {
	syscall.Syscall(gpRasterPos3f, 3, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func RasterPos3fv(v *float32) {
	syscall.Syscall(gpRasterPos3fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos3i(x int32, y int32, z int32) {
	syscall.Syscall(gpRasterPos3i, 3, uintptr(x), uintptr(y), uintptr(z))
}
func RasterPos3iv(v *int32) {
	syscall.Syscall(gpRasterPos3iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos3s(x int16, y int16, z int16) {
	syscall.Syscall(gpRasterPos3s, 3, uintptr(x), uintptr(y), uintptr(z))
}
func RasterPos3sv(v *int16) {
	syscall.Syscall(gpRasterPos3sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos3xOES(x int32, y int32, z int32) {
	syscall.Syscall(gpRasterPos3xOES, 3, uintptr(x), uintptr(y), uintptr(z))
}
func RasterPos3xvOES(coords *int32) {
	syscall.Syscall(gpRasterPos3xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func RasterPos4d(x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpRasterPos4d, 4, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0, 0)
}
func RasterPos4dv(v *float64) {
	syscall.Syscall(gpRasterPos4dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos4f(x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpRasterPos4f, 4, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0, 0)
}
func RasterPos4fv(v *float32) {
	syscall.Syscall(gpRasterPos4fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos4i(x int32, y int32, z int32, w int32) {
	syscall.Syscall6(gpRasterPos4i, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func RasterPos4iv(v *int32) {
	syscall.Syscall(gpRasterPos4iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos4s(x int16, y int16, z int16, w int16) {
	syscall.Syscall6(gpRasterPos4s, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func RasterPos4sv(v *int16) {
	syscall.Syscall(gpRasterPos4sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func RasterPos4xOES(x int32, y int32, z int32, w int32) {
	syscall.Syscall6(gpRasterPos4xOES, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func RasterPos4xvOES(coords *int32) {
	syscall.Syscall(gpRasterPos4xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func RasterSamplesEXT(samples uint32, fixedsamplelocations bool) {
	syscall.Syscall(gpRasterSamplesEXT, 2, uintptr(samples), boolToUintptr(fixedsamplelocations), 0)
}

// select a color buffer source for pixels
func ReadBuffer(src uint32) {
	syscall.Syscall(gpReadBuffer, 1, uintptr(src), 0, 0)
}
func ReadInstrumentsSGIX(marker int32) {
	syscall.Syscall(gpReadInstrumentsSGIX, 1, uintptr(marker), 0, 0)
}

// read a block of pixels from the frame buffer
func ReadPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpReadPixels, 7, uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}

// read a block of pixels from the frame buffer
func ReadnPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpReadnPixels, 8, uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(bufSize), uintptr(data), 0)
}
func ReadnPixelsARB(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpReadnPixelsARB, 8, uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(bufSize), uintptr(data), 0)
}
func ReadnPixelsKHR(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	syscall.Syscall9(gpReadnPixelsKHR, 8, uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(bufSize), uintptr(data), 0)
}
func Rectd(x1 float64, y1 float64, x2 float64, y2 float64) {
	syscall.Syscall6(gpRectd, 4, uintptr(math.Float64bits(x1)), uintptr(math.Float64bits(y1)), uintptr(math.Float64bits(x2)), uintptr(math.Float64bits(y2)), 0, 0)
}
func Rectdv(v1 *float64, v2 *float64) {
	syscall.Syscall(gpRectdv, 2, uintptr(unsafe.Pointer(v1)), uintptr(unsafe.Pointer(v2)), 0)
}
func Rectf(x1 float32, y1 float32, x2 float32, y2 float32) {
	syscall.Syscall6(gpRectf, 4, uintptr(math.Float32bits(x1)), uintptr(math.Float32bits(y1)), uintptr(math.Float32bits(x2)), uintptr(math.Float32bits(y2)), 0, 0)
}
func Rectfv(v1 *float32, v2 *float32) {
	syscall.Syscall(gpRectfv, 2, uintptr(unsafe.Pointer(v1)), uintptr(unsafe.Pointer(v2)), 0)
}
func Recti(x1 int32, y1 int32, x2 int32, y2 int32) {
	syscall.Syscall6(gpRecti, 4, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), 0, 0)
}
func Rectiv(v1 *int32, v2 *int32) {
	syscall.Syscall(gpRectiv, 2, uintptr(unsafe.Pointer(v1)), uintptr(unsafe.Pointer(v2)), 0)
}
func Rects(x1 int16, y1 int16, x2 int16, y2 int16) {
	syscall.Syscall6(gpRects, 4, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), 0, 0)
}
func Rectsv(v1 *int16, v2 *int16) {
	syscall.Syscall(gpRectsv, 2, uintptr(unsafe.Pointer(v1)), uintptr(unsafe.Pointer(v2)), 0)
}
func RectxOES(x1 int32, y1 int32, x2 int32, y2 int32) {
	syscall.Syscall6(gpRectxOES, 4, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), 0, 0)
}
func RectxvOES(v1 *int32, v2 *int32) {
	syscall.Syscall(gpRectxvOES, 2, uintptr(unsafe.Pointer(v1)), uintptr(unsafe.Pointer(v2)), 0)
}
func ReferencePlaneSGIX(equation *float64) {
	syscall.Syscall(gpReferencePlaneSGIX, 1, uintptr(unsafe.Pointer(equation)), 0, 0)
}
func ReleaseKeyedMutexWin32EXT(memory uint32, key uint64) bool {
	ret, _, _ := syscall.Syscall(gpReleaseKeyedMutexWin32EXT, 2, uintptr(memory), uintptr(key), 0)
	return ret != 0
}

// release resources consumed by the implementation's shader compiler
func ReleaseShaderCompiler() {
	syscall.Syscall(gpReleaseShaderCompiler, 0, 0, 0, 0)
}
func RenderGpuMaskNV(mask uint32) {
	syscall.Syscall(gpRenderGpuMaskNV, 1, uintptr(mask), 0, 0)
}

// set rasterization mode
func RenderMode(mode uint32) int32 {
	ret, _, _ := syscall.Syscall(gpRenderMode, 1, uintptr(mode), 0, 0)
	return (int32)(ret)
}

// establish data storage, format and dimensions of a     renderbuffer object's image
func RenderbufferStorage(target uint32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpRenderbufferStorage, 4, uintptr(target), uintptr(internalformat), uintptr(width), uintptr(height), 0, 0)
}
func RenderbufferStorageEXT(target uint32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpRenderbufferStorageEXT, 4, uintptr(target), uintptr(internalformat), uintptr(width), uintptr(height), 0, 0)
}

// establish data storage, format, dimensions and sample count of     a renderbuffer object's image
func RenderbufferStorageMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpRenderbufferStorageMultisample, 5, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), 0)
}
func RenderbufferStorageMultisampleCoverageNV(target uint32, coverageSamples int32, colorSamples int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpRenderbufferStorageMultisampleCoverageNV, 6, uintptr(target), uintptr(coverageSamples), uintptr(colorSamples), uintptr(internalformat), uintptr(width), uintptr(height))
}
func RenderbufferStorageMultisampleEXT(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpRenderbufferStorageMultisampleEXT, 5, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), 0)
}
func ReplacementCodePointerSUN(xtype uint32, stride int32, pointer *unsafe.Pointer) {
	syscall.Syscall(gpReplacementCodePointerSUN, 3, uintptr(xtype), uintptr(stride), uintptr(unsafe.Pointer(pointer)))
}
func ReplacementCodeubSUN(code uint8) {
	syscall.Syscall(gpReplacementCodeubSUN, 1, uintptr(code), 0, 0)
}
func ReplacementCodeubvSUN(code *uint8) {
	syscall.Syscall(gpReplacementCodeubvSUN, 1, uintptr(unsafe.Pointer(code)), 0, 0)
}
func ReplacementCodeuiColor3fVertex3fSUN(rc uint32, r float32, g float32, b float32, x float32, y float32, z float32) {
	syscall.Syscall9(gpReplacementCodeuiColor3fVertex3fSUN, 7, uintptr(rc), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(g)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func ReplacementCodeuiColor3fVertex3fvSUN(rc *uint32, c *float32, v *float32) {
	syscall.Syscall(gpReplacementCodeuiColor3fVertex3fvSUN, 3, uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(v)))
}
func ReplacementCodeuiColor4fNormal3fVertex3fSUN(rc uint32, r float32, g float32, b float32, a float32, nx float32, ny float32, nz float32, x float32, y float32, z float32) {
	syscall.Syscall12(gpReplacementCodeuiColor4fNormal3fVertex3fSUN, 11, uintptr(rc), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(g)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(a)), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0)
}
func ReplacementCodeuiColor4fNormal3fVertex3fvSUN(rc *uint32, c *float32, n *float32, v *float32) {
	syscall.Syscall6(gpReplacementCodeuiColor4fNormal3fVertex3fvSUN, 4, uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)), 0, 0)
}
func ReplacementCodeuiColor4ubVertex3fSUN(rc uint32, r uint8, g uint8, b uint8, a uint8, x float32, y float32, z float32) {
	syscall.Syscall9(gpReplacementCodeuiColor4ubVertex3fSUN, 8, uintptr(rc), uintptr(r), uintptr(g), uintptr(b), uintptr(a), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0)
}
func ReplacementCodeuiColor4ubVertex3fvSUN(rc *uint32, c *uint8, v *float32) {
	syscall.Syscall(gpReplacementCodeuiColor4ubVertex3fvSUN, 3, uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(v)))
}
func ReplacementCodeuiNormal3fVertex3fSUN(rc uint32, nx float32, ny float32, nz float32, x float32, y float32, z float32) {
	syscall.Syscall9(gpReplacementCodeuiNormal3fVertex3fSUN, 7, uintptr(rc), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func ReplacementCodeuiNormal3fVertex3fvSUN(rc *uint32, n *float32, v *float32) {
	syscall.Syscall(gpReplacementCodeuiNormal3fVertex3fvSUN, 3, uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)))
}
func ReplacementCodeuiSUN(code uint32) {
	syscall.Syscall(gpReplacementCodeuiSUN, 1, uintptr(code), 0, 0)
}
func ReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN(rc uint32, s float32, t float32, r float32, g float32, b float32, a float32, nx float32, ny float32, nz float32, x float32, y float32, z float32) {
	syscall.Syscall15(gpReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN, 13, uintptr(rc), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(g)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(a)), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func ReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN(rc *uint32, tc *float32, c *float32, n *float32, v *float32) {
	syscall.Syscall6(gpReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN, 5, uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)), 0)
}
func ReplacementCodeuiTexCoord2fNormal3fVertex3fSUN(rc uint32, s float32, t float32, nx float32, ny float32, nz float32, x float32, y float32, z float32) {
	syscall.Syscall9(gpReplacementCodeuiTexCoord2fNormal3fVertex3fSUN, 9, uintptr(rc), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func ReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN(rc *uint32, tc *float32, n *float32, v *float32) {
	syscall.Syscall6(gpReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN, 4, uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)), 0, 0)
}
func ReplacementCodeuiTexCoord2fVertex3fSUN(rc uint32, s float32, t float32, x float32, y float32, z float32) {
	syscall.Syscall6(gpReplacementCodeuiTexCoord2fVertex3fSUN, 6, uintptr(rc), uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func ReplacementCodeuiTexCoord2fVertex3fvSUN(rc *uint32, tc *float32, v *float32) {
	syscall.Syscall(gpReplacementCodeuiTexCoord2fVertex3fvSUN, 3, uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(v)))
}
func ReplacementCodeuiVertex3fSUN(rc uint32, x float32, y float32, z float32) {
	syscall.Syscall6(gpReplacementCodeuiVertex3fSUN, 4, uintptr(rc), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func ReplacementCodeuiVertex3fvSUN(rc *uint32, v *float32) {
	syscall.Syscall(gpReplacementCodeuiVertex3fvSUN, 2, uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(v)), 0)
}
func ReplacementCodeuivSUN(code *uint32) {
	syscall.Syscall(gpReplacementCodeuivSUN, 1, uintptr(unsafe.Pointer(code)), 0, 0)
}
func ReplacementCodeusSUN(code uint16) {
	syscall.Syscall(gpReplacementCodeusSUN, 1, uintptr(code), 0, 0)
}
func ReplacementCodeusvSUN(code *uint16) {
	syscall.Syscall(gpReplacementCodeusvSUN, 1, uintptr(unsafe.Pointer(code)), 0, 0)
}
func RequestResidentProgramsNV(n int32, programs *uint32) {
	syscall.Syscall(gpRequestResidentProgramsNV, 2, uintptr(n), uintptr(unsafe.Pointer(programs)), 0)
}
func ResetHistogramEXT(target uint32) {
	syscall.Syscall(gpResetHistogramEXT, 1, uintptr(target), 0, 0)
}
func ResetMinmaxEXT(target uint32) {
	syscall.Syscall(gpResetMinmaxEXT, 1, uintptr(target), 0, 0)
}
func ResizeBuffersMESA() {
	syscall.Syscall(gpResizeBuffersMESA, 0, 0, 0, 0)
}
func ResolveDepthValuesNV() {
	syscall.Syscall(gpResolveDepthValuesNV, 0, 0, 0, 0)
}

// resume transform feedback operations
func ResumeTransformFeedback() {
	syscall.Syscall(gpResumeTransformFeedback, 0, 0, 0, 0)
}
func ResumeTransformFeedbackNV() {
	syscall.Syscall(gpResumeTransformFeedbackNV, 0, 0, 0, 0)
}
func Rotated(angle float64, x float64, y float64, z float64) {
	syscall.Syscall6(gpRotated, 4, uintptr(math.Float64bits(angle)), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func Rotatef(angle float32, x float32, y float32, z float32) {
	syscall.Syscall6(gpRotatef, 4, uintptr(math.Float32bits(angle)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func RotatexOES(angle int32, x int32, y int32, z int32) {
	syscall.Syscall6(gpRotatexOES, 4, uintptr(angle), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}

// specify multisample coverage parameters
func SampleCoverage(value float32, invert bool) {
	syscall.Syscall(gpSampleCoverage, 2, uintptr(math.Float32bits(value)), boolToUintptr(invert), 0)
}
func SampleCoverageARB(value float32, invert bool) {
	syscall.Syscall(gpSampleCoverageARB, 2, uintptr(math.Float32bits(value)), boolToUintptr(invert), 0)
}
func SampleCoveragexOES(value int32, invert bool) {
	syscall.Syscall(gpSampleCoveragexOES, 2, uintptr(value), boolToUintptr(invert), 0)
}
func SampleMapATI(dst uint32, interp uint32, swizzle uint32) {
	syscall.Syscall(gpSampleMapATI, 3, uintptr(dst), uintptr(interp), uintptr(swizzle))
}
func SampleMaskEXT(value float32, invert bool) {
	syscall.Syscall(gpSampleMaskEXT, 2, uintptr(math.Float32bits(value)), boolToUintptr(invert), 0)
}
func SampleMaskIndexedNV(index uint32, mask uint32) {
	syscall.Syscall(gpSampleMaskIndexedNV, 2, uintptr(index), uintptr(mask), 0)
}
func SampleMaskSGIS(value float32, invert bool) {
	syscall.Syscall(gpSampleMaskSGIS, 2, uintptr(math.Float32bits(value)), boolToUintptr(invert), 0)
}

// set the value of a sub-word of the sample mask
func SampleMaski(maskNumber uint32, mask uint32) {
	syscall.Syscall(gpSampleMaski, 2, uintptr(maskNumber), uintptr(mask), 0)
}
func SamplePatternEXT(pattern uint32) {
	syscall.Syscall(gpSamplePatternEXT, 1, uintptr(pattern), 0, 0)
}
func SamplePatternSGIS(pattern uint32) {
	syscall.Syscall(gpSamplePatternSGIS, 1, uintptr(pattern), 0, 0)
}
func SamplerParameterIiv(sampler uint32, pname uint32, param *int32) {
	syscall.Syscall(gpSamplerParameterIiv, 3, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func SamplerParameterIuiv(sampler uint32, pname uint32, param *uint32) {
	syscall.Syscall(gpSamplerParameterIuiv, 3, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func SamplerParameterf(sampler uint32, pname uint32, param float32) {
	syscall.Syscall(gpSamplerParameterf, 3, uintptr(sampler), uintptr(pname), uintptr(math.Float32bits(param)))
}
func SamplerParameterfv(sampler uint32, pname uint32, param *float32) {
	syscall.Syscall(gpSamplerParameterfv, 3, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func SamplerParameteri(sampler uint32, pname uint32, param int32) {
	syscall.Syscall(gpSamplerParameteri, 3, uintptr(sampler), uintptr(pname), uintptr(param))
}
func SamplerParameteriv(sampler uint32, pname uint32, param *int32) {
	syscall.Syscall(gpSamplerParameteriv, 3, uintptr(sampler), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func Scaled(x float64, y float64, z float64) {
	syscall.Syscall(gpScaled, 3, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)))
}
func Scalef(x float32, y float32, z float32) {
	syscall.Syscall(gpScalef, 3, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func ScalexOES(x int32, y int32, z int32) {
	syscall.Syscall(gpScalexOES, 3, uintptr(x), uintptr(y), uintptr(z))
}

// define the scissor box
func Scissor(x int32, y int32, width int32, height int32) {
	syscall.Syscall6(gpScissor, 4, uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0, 0)
}
func ScissorArrayv(first uint32, count int32, v *int32) {
	syscall.Syscall(gpScissorArrayv, 3, uintptr(first), uintptr(count), uintptr(unsafe.Pointer(v)))
}

// define the scissor box for a specific viewport
func ScissorIndexed(index uint32, left int32, bottom int32, width int32, height int32) {
	syscall.Syscall6(gpScissorIndexed, 5, uintptr(index), uintptr(left), uintptr(bottom), uintptr(width), uintptr(height), 0)
}
func ScissorIndexedv(index uint32, v *int32) {
	syscall.Syscall(gpScissorIndexedv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func SecondaryColor3b(red int8, green int8, blue int8) {
	syscall.Syscall(gpSecondaryColor3b, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3bEXT(red int8, green int8, blue int8) {
	syscall.Syscall(gpSecondaryColor3bEXT, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3bv(v *int8) {
	syscall.Syscall(gpSecondaryColor3bv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3bvEXT(v *int8) {
	syscall.Syscall(gpSecondaryColor3bvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3d(red float64, green float64, blue float64) {
	syscall.Syscall(gpSecondaryColor3d, 3, uintptr(math.Float64bits(red)), uintptr(math.Float64bits(green)), uintptr(math.Float64bits(blue)))
}
func SecondaryColor3dEXT(red float64, green float64, blue float64) {
	syscall.Syscall(gpSecondaryColor3dEXT, 3, uintptr(math.Float64bits(red)), uintptr(math.Float64bits(green)), uintptr(math.Float64bits(blue)))
}
func SecondaryColor3dv(v *float64) {
	syscall.Syscall(gpSecondaryColor3dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3dvEXT(v *float64) {
	syscall.Syscall(gpSecondaryColor3dvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3f(red float32, green float32, blue float32) {
	syscall.Syscall(gpSecondaryColor3f, 3, uintptr(math.Float32bits(red)), uintptr(math.Float32bits(green)), uintptr(math.Float32bits(blue)))
}
func SecondaryColor3fEXT(red float32, green float32, blue float32) {
	syscall.Syscall(gpSecondaryColor3fEXT, 3, uintptr(math.Float32bits(red)), uintptr(math.Float32bits(green)), uintptr(math.Float32bits(blue)))
}
func SecondaryColor3fv(v *float32) {
	syscall.Syscall(gpSecondaryColor3fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3fvEXT(v *float32) {
	syscall.Syscall(gpSecondaryColor3fvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3hNV(red uint16, green uint16, blue uint16) {
	syscall.Syscall(gpSecondaryColor3hNV, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3hvNV(v *uint16) {
	syscall.Syscall(gpSecondaryColor3hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3i(red int32, green int32, blue int32) {
	syscall.Syscall(gpSecondaryColor3i, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3iEXT(red int32, green int32, blue int32) {
	syscall.Syscall(gpSecondaryColor3iEXT, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3iv(v *int32) {
	syscall.Syscall(gpSecondaryColor3iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3ivEXT(v *int32) {
	syscall.Syscall(gpSecondaryColor3ivEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3s(red int16, green int16, blue int16) {
	syscall.Syscall(gpSecondaryColor3s, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3sEXT(red int16, green int16, blue int16) {
	syscall.Syscall(gpSecondaryColor3sEXT, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3sv(v *int16) {
	syscall.Syscall(gpSecondaryColor3sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3svEXT(v *int16) {
	syscall.Syscall(gpSecondaryColor3svEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3ub(red uint8, green uint8, blue uint8) {
	syscall.Syscall(gpSecondaryColor3ub, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3ubEXT(red uint8, green uint8, blue uint8) {
	syscall.Syscall(gpSecondaryColor3ubEXT, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3ubv(v *uint8) {
	syscall.Syscall(gpSecondaryColor3ubv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3ubvEXT(v *uint8) {
	syscall.Syscall(gpSecondaryColor3ubvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3ui(red uint32, green uint32, blue uint32) {
	syscall.Syscall(gpSecondaryColor3ui, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3uiEXT(red uint32, green uint32, blue uint32) {
	syscall.Syscall(gpSecondaryColor3uiEXT, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3uiv(v *uint32) {
	syscall.Syscall(gpSecondaryColor3uiv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3uivEXT(v *uint32) {
	syscall.Syscall(gpSecondaryColor3uivEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3us(red uint16, green uint16, blue uint16) {
	syscall.Syscall(gpSecondaryColor3us, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3usEXT(red uint16, green uint16, blue uint16) {
	syscall.Syscall(gpSecondaryColor3usEXT, 3, uintptr(red), uintptr(green), uintptr(blue))
}
func SecondaryColor3usv(v *uint16) {
	syscall.Syscall(gpSecondaryColor3usv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColor3usvEXT(v *uint16) {
	syscall.Syscall(gpSecondaryColor3usvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func SecondaryColorFormatNV(size int32, xtype uint32, stride int32) {
	syscall.Syscall(gpSecondaryColorFormatNV, 3, uintptr(size), uintptr(xtype), uintptr(stride))
}

// define an array of secondary colors
func SecondaryColorPointer(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpSecondaryColorPointer, 4, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0, 0)
}
func SecondaryColorPointerEXT(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpSecondaryColorPointerEXT, 4, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0, 0)
}
func SecondaryColorPointerListIBM(size int32, xtype uint32, stride int32, pointer *unsafe.Pointer, ptrstride int32) {
	syscall.Syscall6(gpSecondaryColorPointerListIBM, 5, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(unsafe.Pointer(pointer)), uintptr(ptrstride), 0)
}

// establish a buffer for selection mode values
func SelectBuffer(size int32, buffer *uint32) {
	syscall.Syscall(gpSelectBuffer, 2, uintptr(size), uintptr(unsafe.Pointer(buffer)), 0)
}
func SelectPerfMonitorCountersAMD(monitor uint32, enable bool, group uint32, numCounters int32, counterList *uint32) {
	syscall.Syscall6(gpSelectPerfMonitorCountersAMD, 5, uintptr(monitor), boolToUintptr(enable), uintptr(group), uintptr(numCounters), uintptr(unsafe.Pointer(counterList)), 0)
}
func SemaphoreParameterui64vEXT(semaphore uint32, pname uint32, params *uint64) {
	syscall.Syscall(gpSemaphoreParameterui64vEXT, 3, uintptr(semaphore), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func SeparableFilter2DEXT(target uint32, internalformat uint32, width int32, height int32, format uint32, xtype uint32, row unsafe.Pointer, column unsafe.Pointer) {
	syscall.Syscall9(gpSeparableFilter2DEXT, 8, uintptr(target), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(row), uintptr(column), 0)
}
func SetFenceAPPLE(fence uint32) {
	syscall.Syscall(gpSetFenceAPPLE, 1, uintptr(fence), 0, 0)
}
func SetFenceNV(fence uint32, condition uint32) {
	syscall.Syscall(gpSetFenceNV, 2, uintptr(fence), uintptr(condition), 0)
}
func SetFragmentShaderConstantATI(dst uint32, value *float32) {
	syscall.Syscall(gpSetFragmentShaderConstantATI, 2, uintptr(dst), uintptr(unsafe.Pointer(value)), 0)
}
func SetInvariantEXT(id uint32, xtype uint32, addr unsafe.Pointer) {
	syscall.Syscall(gpSetInvariantEXT, 3, uintptr(id), uintptr(xtype), uintptr(addr))
}
func SetLocalConstantEXT(id uint32, xtype uint32, addr unsafe.Pointer) {
	syscall.Syscall(gpSetLocalConstantEXT, 3, uintptr(id), uintptr(xtype), uintptr(addr))
}
func SetMultisamplefvAMD(pname uint32, index uint32, val *float32) {
	syscall.Syscall(gpSetMultisamplefvAMD, 3, uintptr(pname), uintptr(index), uintptr(unsafe.Pointer(val)))
}

// select flat or smooth shading
func ShadeModel(mode uint32) {
	syscall.Syscall(gpShadeModel, 1, uintptr(mode), 0, 0)
}

// load pre-compiled shader binaries
func ShaderBinary(count int32, shaders *uint32, binaryformat uint32, binary unsafe.Pointer, length int32) {
	syscall.Syscall6(gpShaderBinary, 5, uintptr(count), uintptr(unsafe.Pointer(shaders)), uintptr(binaryformat), uintptr(binary), uintptr(length), 0)
}
func ShaderOp1EXT(op uint32, res uint32, arg1 uint32) {
	syscall.Syscall(gpShaderOp1EXT, 3, uintptr(op), uintptr(res), uintptr(arg1))
}
func ShaderOp2EXT(op uint32, res uint32, arg1 uint32, arg2 uint32) {
	syscall.Syscall6(gpShaderOp2EXT, 4, uintptr(op), uintptr(res), uintptr(arg1), uintptr(arg2), 0, 0)
}
func ShaderOp3EXT(op uint32, res uint32, arg1 uint32, arg2 uint32, arg3 uint32) {
	syscall.Syscall6(gpShaderOp3EXT, 5, uintptr(op), uintptr(res), uintptr(arg1), uintptr(arg2), uintptr(arg3), 0)
}

// Replaces the source code in a shader object
func ShaderSource(shader uint32, count int32, xstring **uint8, length *int32) {
	syscall.Syscall6(gpShaderSource, 4, uintptr(shader), uintptr(count), uintptr(unsafe.Pointer(xstring)), uintptr(unsafe.Pointer(length)), 0, 0)
}
func ShaderSourceARB(shaderObj uintptr, count int32, xstring **uint8, length *int32) {
	syscall.Syscall6(gpShaderSourceARB, 4, uintptr(shaderObj), uintptr(count), uintptr(unsafe.Pointer(xstring)), uintptr(unsafe.Pointer(length)), 0, 0)
}

// change an active shader storage block binding
func ShaderStorageBlockBinding(program uint32, storageBlockIndex uint32, storageBlockBinding uint32) {
	syscall.Syscall(gpShaderStorageBlockBinding, 3, uintptr(program), uintptr(storageBlockIndex), uintptr(storageBlockBinding))
}
func SharpenTexFuncSGIS(target uint32, n int32, points *float32) {
	syscall.Syscall(gpSharpenTexFuncSGIS, 3, uintptr(target), uintptr(n), uintptr(unsafe.Pointer(points)))
}
func SignalSemaphoreEXT(semaphore uint32, numBufferBarriers uint32, buffers *uint32, numTextureBarriers uint32, textures *uint32, dstLayouts *uint32) {
	syscall.Syscall6(gpSignalSemaphoreEXT, 6, uintptr(semaphore), uintptr(numBufferBarriers), uintptr(unsafe.Pointer(buffers)), uintptr(numTextureBarriers), uintptr(unsafe.Pointer(textures)), uintptr(unsafe.Pointer(dstLayouts)))
}
func SignalVkFenceNV(vkFence uint64) {
	syscall.Syscall(gpSignalVkFenceNV, 1, uintptr(vkFence), 0, 0)
}
func SignalVkSemaphoreNV(vkSemaphore uint64) {
	syscall.Syscall(gpSignalVkSemaphoreNV, 1, uintptr(vkSemaphore), 0, 0)
}
func SpecializeShaderARB(shader uint32, pEntryPoint *uint8, numSpecializationConstants uint32, pConstantIndex *uint32, pConstantValue *uint32) {
	syscall.Syscall6(gpSpecializeShaderARB, 5, uintptr(shader), uintptr(unsafe.Pointer(pEntryPoint)), uintptr(numSpecializationConstants), uintptr(unsafe.Pointer(pConstantIndex)), uintptr(unsafe.Pointer(pConstantValue)), 0)
}
func SpriteParameterfSGIX(pname uint32, param float32) {
	syscall.Syscall(gpSpriteParameterfSGIX, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func SpriteParameterfvSGIX(pname uint32, params *float32) {
	syscall.Syscall(gpSpriteParameterfvSGIX, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func SpriteParameteriSGIX(pname uint32, param int32) {
	syscall.Syscall(gpSpriteParameteriSGIX, 2, uintptr(pname), uintptr(param), 0)
}
func SpriteParameterivSGIX(pname uint32, params *int32) {
	syscall.Syscall(gpSpriteParameterivSGIX, 2, uintptr(pname), uintptr(unsafe.Pointer(params)), 0)
}
func StartInstrumentsSGIX() {
	syscall.Syscall(gpStartInstrumentsSGIX, 0, 0, 0, 0)
}
func StateCaptureNV(state uint32, mode uint32) {
	syscall.Syscall(gpStateCaptureNV, 2, uintptr(state), uintptr(mode), 0)
}
func StencilClearTagEXT(stencilTagBits int32, stencilClearTag uint32) {
	syscall.Syscall(gpStencilClearTagEXT, 2, uintptr(stencilTagBits), uintptr(stencilClearTag), 0)
}
func StencilFillPathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, fillMode uint32, mask uint32, transformType uint32, transformValues *float32) {
	syscall.Syscall9(gpStencilFillPathInstancedNV, 8, uintptr(numPaths), uintptr(pathNameType), uintptr(paths), uintptr(pathBase), uintptr(fillMode), uintptr(mask), uintptr(transformType), uintptr(unsafe.Pointer(transformValues)), 0)
}
func StencilFillPathNV(path uint32, fillMode uint32, mask uint32) {
	syscall.Syscall(gpStencilFillPathNV, 3, uintptr(path), uintptr(fillMode), uintptr(mask))
}

// set front and back function and reference value for stencil testing
func StencilFunc(xfunc uint32, ref int32, mask uint32) {
	syscall.Syscall(gpStencilFunc, 3, uintptr(xfunc), uintptr(ref), uintptr(mask))
}

// set front and/or back function and reference value for stencil testing
func StencilFuncSeparate(face uint32, xfunc uint32, ref int32, mask uint32) {
	syscall.Syscall6(gpStencilFuncSeparate, 4, uintptr(face), uintptr(xfunc), uintptr(ref), uintptr(mask), 0, 0)
}
func StencilFuncSeparateATI(frontfunc uint32, backfunc uint32, ref int32, mask uint32) {
	syscall.Syscall6(gpStencilFuncSeparateATI, 4, uintptr(frontfunc), uintptr(backfunc), uintptr(ref), uintptr(mask), 0, 0)
}

// control the front and back writing of individual bits in the stencil planes
func StencilMask(mask uint32) {
	syscall.Syscall(gpStencilMask, 1, uintptr(mask), 0, 0)
}

// control the front and/or back writing of individual bits in the stencil planes
func StencilMaskSeparate(face uint32, mask uint32) {
	syscall.Syscall(gpStencilMaskSeparate, 2, uintptr(face), uintptr(mask), 0)
}

// set front and back stencil test actions
func StencilOp(fail uint32, zfail uint32, zpass uint32) {
	syscall.Syscall(gpStencilOp, 3, uintptr(fail), uintptr(zfail), uintptr(zpass))
}

// set front and/or back stencil test actions
func StencilOpSeparate(face uint32, sfail uint32, dpfail uint32, dppass uint32) {
	syscall.Syscall6(gpStencilOpSeparate, 4, uintptr(face), uintptr(sfail), uintptr(dpfail), uintptr(dppass), 0, 0)
}
func StencilOpSeparateATI(face uint32, sfail uint32, dpfail uint32, dppass uint32) {
	syscall.Syscall6(gpStencilOpSeparateATI, 4, uintptr(face), uintptr(sfail), uintptr(dpfail), uintptr(dppass), 0, 0)
}
func StencilOpValueAMD(face uint32, value uint32) {
	syscall.Syscall(gpStencilOpValueAMD, 2, uintptr(face), uintptr(value), 0)
}
func StencilStrokePathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, reference int32, mask uint32, transformType uint32, transformValues *float32) {
	syscall.Syscall9(gpStencilStrokePathInstancedNV, 8, uintptr(numPaths), uintptr(pathNameType), uintptr(paths), uintptr(pathBase), uintptr(reference), uintptr(mask), uintptr(transformType), uintptr(unsafe.Pointer(transformValues)), 0)
}
func StencilStrokePathNV(path uint32, reference int32, mask uint32) {
	syscall.Syscall(gpStencilStrokePathNV, 3, uintptr(path), uintptr(reference), uintptr(mask))
}
func StencilThenCoverFillPathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, fillMode uint32, mask uint32, coverMode uint32, transformType uint32, transformValues *float32) {
	syscall.Syscall9(gpStencilThenCoverFillPathInstancedNV, 9, uintptr(numPaths), uintptr(pathNameType), uintptr(paths), uintptr(pathBase), uintptr(fillMode), uintptr(mask), uintptr(coverMode), uintptr(transformType), uintptr(unsafe.Pointer(transformValues)))
}
func StencilThenCoverFillPathNV(path uint32, fillMode uint32, mask uint32, coverMode uint32) {
	syscall.Syscall6(gpStencilThenCoverFillPathNV, 4, uintptr(path), uintptr(fillMode), uintptr(mask), uintptr(coverMode), 0, 0)
}
func StencilThenCoverStrokePathInstancedNV(numPaths int32, pathNameType uint32, paths unsafe.Pointer, pathBase uint32, reference int32, mask uint32, coverMode uint32, transformType uint32, transformValues *float32) {
	syscall.Syscall9(gpStencilThenCoverStrokePathInstancedNV, 9, uintptr(numPaths), uintptr(pathNameType), uintptr(paths), uintptr(pathBase), uintptr(reference), uintptr(mask), uintptr(coverMode), uintptr(transformType), uintptr(unsafe.Pointer(transformValues)))
}
func StencilThenCoverStrokePathNV(path uint32, reference int32, mask uint32, coverMode uint32) {
	syscall.Syscall6(gpStencilThenCoverStrokePathNV, 4, uintptr(path), uintptr(reference), uintptr(mask), uintptr(coverMode), 0, 0)
}
func StopInstrumentsSGIX(marker int32) {
	syscall.Syscall(gpStopInstrumentsSGIX, 1, uintptr(marker), 0, 0)
}
func StringMarkerGREMEDY(len int32, xstring unsafe.Pointer) {
	syscall.Syscall(gpStringMarkerGREMEDY, 2, uintptr(len), uintptr(xstring), 0)
}
func SubpixelPrecisionBiasNV(xbits uint32, ybits uint32) {
	syscall.Syscall(gpSubpixelPrecisionBiasNV, 2, uintptr(xbits), uintptr(ybits), 0)
}
func SwizzleEXT(res uint32, in uint32, outX uint32, outY uint32, outZ uint32, outW uint32) {
	syscall.Syscall6(gpSwizzleEXT, 6, uintptr(res), uintptr(in), uintptr(outX), uintptr(outY), uintptr(outZ), uintptr(outW))
}
func SyncTextureINTEL(texture uint32) {
	syscall.Syscall(gpSyncTextureINTEL, 1, uintptr(texture), 0, 0)
}
func TagSampleBufferSGIX() {
	syscall.Syscall(gpTagSampleBufferSGIX, 0, 0, 0, 0)
}
func Tangent3bEXT(tx int8, ty int8, tz int8) {
	syscall.Syscall(gpTangent3bEXT, 3, uintptr(tx), uintptr(ty), uintptr(tz))
}
func Tangent3bvEXT(v *int8) {
	syscall.Syscall(gpTangent3bvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Tangent3dEXT(tx float64, ty float64, tz float64) {
	syscall.Syscall(gpTangent3dEXT, 3, uintptr(math.Float64bits(tx)), uintptr(math.Float64bits(ty)), uintptr(math.Float64bits(tz)))
}
func Tangent3dvEXT(v *float64) {
	syscall.Syscall(gpTangent3dvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Tangent3fEXT(tx float32, ty float32, tz float32) {
	syscall.Syscall(gpTangent3fEXT, 3, uintptr(math.Float32bits(tx)), uintptr(math.Float32bits(ty)), uintptr(math.Float32bits(tz)))
}
func Tangent3fvEXT(v *float32) {
	syscall.Syscall(gpTangent3fvEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Tangent3iEXT(tx int32, ty int32, tz int32) {
	syscall.Syscall(gpTangent3iEXT, 3, uintptr(tx), uintptr(ty), uintptr(tz))
}
func Tangent3ivEXT(v *int32) {
	syscall.Syscall(gpTangent3ivEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Tangent3sEXT(tx int16, ty int16, tz int16) {
	syscall.Syscall(gpTangent3sEXT, 3, uintptr(tx), uintptr(ty), uintptr(tz))
}
func Tangent3svEXT(v *int16) {
	syscall.Syscall(gpTangent3svEXT, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TangentPointerEXT(xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpTangentPointerEXT, 3, uintptr(xtype), uintptr(stride), uintptr(pointer))
}
func TbufferMask3DFX(mask uint32) {
	syscall.Syscall(gpTbufferMask3DFX, 1, uintptr(mask), 0, 0)
}
func TessellationFactorAMD(factor float32) {
	syscall.Syscall(gpTessellationFactorAMD, 1, uintptr(math.Float32bits(factor)), 0, 0)
}
func TessellationModeAMD(mode uint32) {
	syscall.Syscall(gpTessellationModeAMD, 1, uintptr(mode), 0, 0)
}
func TestFenceAPPLE(fence uint32) bool {
	ret, _, _ := syscall.Syscall(gpTestFenceAPPLE, 1, uintptr(fence), 0, 0)
	return ret != 0
}
func TestFenceNV(fence uint32) bool {
	ret, _, _ := syscall.Syscall(gpTestFenceNV, 1, uintptr(fence), 0, 0)
	return ret != 0
}
func TestObjectAPPLE(object uint32, name uint32) bool {
	ret, _, _ := syscall.Syscall(gpTestObjectAPPLE, 2, uintptr(object), uintptr(name), 0)
	return ret != 0
}
func TexBufferARB(target uint32, internalformat uint32, buffer uint32) {
	syscall.Syscall(gpTexBufferARB, 3, uintptr(target), uintptr(internalformat), uintptr(buffer))
}
func TexBufferEXT(target uint32, internalformat uint32, buffer uint32) {
	syscall.Syscall(gpTexBufferEXT, 3, uintptr(target), uintptr(internalformat), uintptr(buffer))
}

// attach a range of a buffer object's data store to a buffer texture object
func TexBufferRange(target uint32, internalformat uint32, buffer uint32, offset int, size int) {
	syscall.Syscall6(gpTexBufferRange, 5, uintptr(target), uintptr(internalformat), uintptr(buffer), uintptr(offset), uintptr(size), 0)
}
func TexBumpParameterfvATI(pname uint32, param *float32) {
	syscall.Syscall(gpTexBumpParameterfvATI, 2, uintptr(pname), uintptr(unsafe.Pointer(param)), 0)
}
func TexBumpParameterivATI(pname uint32, param *int32) {
	syscall.Syscall(gpTexBumpParameterivATI, 2, uintptr(pname), uintptr(unsafe.Pointer(param)), 0)
}
func TexCoord1bOES(s int8) {
	syscall.Syscall(gpTexCoord1bOES, 1, uintptr(s), 0, 0)
}
func TexCoord1bvOES(coords *int8) {
	syscall.Syscall(gpTexCoord1bvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func TexCoord1d(s float64) {
	syscall.Syscall(gpTexCoord1d, 1, uintptr(math.Float64bits(s)), 0, 0)
}
func TexCoord1dv(v *float64) {
	syscall.Syscall(gpTexCoord1dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord1f(s float32) {
	syscall.Syscall(gpTexCoord1f, 1, uintptr(math.Float32bits(s)), 0, 0)
}
func TexCoord1fv(v *float32) {
	syscall.Syscall(gpTexCoord1fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord1hNV(s uint16) {
	syscall.Syscall(gpTexCoord1hNV, 1, uintptr(s), 0, 0)
}
func TexCoord1hvNV(v *uint16) {
	syscall.Syscall(gpTexCoord1hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord1i(s int32) {
	syscall.Syscall(gpTexCoord1i, 1, uintptr(s), 0, 0)
}
func TexCoord1iv(v *int32) {
	syscall.Syscall(gpTexCoord1iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord1s(s int16) {
	syscall.Syscall(gpTexCoord1s, 1, uintptr(s), 0, 0)
}
func TexCoord1sv(v *int16) {
	syscall.Syscall(gpTexCoord1sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord1xOES(s int32) {
	syscall.Syscall(gpTexCoord1xOES, 1, uintptr(s), 0, 0)
}
func TexCoord1xvOES(coords *int32) {
	syscall.Syscall(gpTexCoord1xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func TexCoord2bOES(s int8, t int8) {
	syscall.Syscall(gpTexCoord2bOES, 2, uintptr(s), uintptr(t), 0)
}
func TexCoord2bvOES(coords *int8) {
	syscall.Syscall(gpTexCoord2bvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func TexCoord2d(s float64, t float64) {
	syscall.Syscall(gpTexCoord2d, 2, uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)), 0)
}
func TexCoord2dv(v *float64) {
	syscall.Syscall(gpTexCoord2dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord2f(s float32, t float32) {
	syscall.Syscall(gpTexCoord2f, 2, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), 0)
}
func TexCoord2fColor3fVertex3fSUN(s float32, t float32, r float32, g float32, b float32, x float32, y float32, z float32) {
	syscall.Syscall9(gpTexCoord2fColor3fVertex3fSUN, 8, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(g)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0)
}
func TexCoord2fColor3fVertex3fvSUN(tc *float32, c *float32, v *float32) {
	syscall.Syscall(gpTexCoord2fColor3fVertex3fvSUN, 3, uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(v)))
}
func TexCoord2fColor4fNormal3fVertex3fSUN(s float32, t float32, r float32, g float32, b float32, a float32, nx float32, ny float32, nz float32, x float32, y float32, z float32) {
	syscall.Syscall12(gpTexCoord2fColor4fNormal3fVertex3fSUN, 12, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(g)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(a)), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func TexCoord2fColor4fNormal3fVertex3fvSUN(tc *float32, c *float32, n *float32, v *float32) {
	syscall.Syscall6(gpTexCoord2fColor4fNormal3fVertex3fvSUN, 4, uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord2fColor4ubVertex3fSUN(s float32, t float32, r uint8, g uint8, b uint8, a uint8, x float32, y float32, z float32) {
	syscall.Syscall9(gpTexCoord2fColor4ubVertex3fSUN, 9, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(r), uintptr(g), uintptr(b), uintptr(a), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func TexCoord2fColor4ubVertex3fvSUN(tc *float32, c *uint8, v *float32) {
	syscall.Syscall(gpTexCoord2fColor4ubVertex3fvSUN, 3, uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(v)))
}
func TexCoord2fNormal3fVertex3fSUN(s float32, t float32, nx float32, ny float32, nz float32, x float32, y float32, z float32) {
	syscall.Syscall9(gpTexCoord2fNormal3fVertex3fSUN, 8, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0)
}
func TexCoord2fNormal3fVertex3fvSUN(tc *float32, n *float32, v *float32) {
	syscall.Syscall(gpTexCoord2fNormal3fVertex3fvSUN, 3, uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)))
}
func TexCoord2fVertex3fSUN(s float32, t float32, x float32, y float32, z float32) {
	syscall.Syscall6(gpTexCoord2fVertex3fSUN, 5, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0)
}
func TexCoord2fVertex3fvSUN(tc *float32, v *float32) {
	syscall.Syscall(gpTexCoord2fVertex3fvSUN, 2, uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(v)), 0)
}
func TexCoord2fv(v *float32) {
	syscall.Syscall(gpTexCoord2fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord2hNV(s uint16, t uint16) {
	syscall.Syscall(gpTexCoord2hNV, 2, uintptr(s), uintptr(t), 0)
}
func TexCoord2hvNV(v *uint16) {
	syscall.Syscall(gpTexCoord2hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord2i(s int32, t int32) {
	syscall.Syscall(gpTexCoord2i, 2, uintptr(s), uintptr(t), 0)
}
func TexCoord2iv(v *int32) {
	syscall.Syscall(gpTexCoord2iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord2s(s int16, t int16) {
	syscall.Syscall(gpTexCoord2s, 2, uintptr(s), uintptr(t), 0)
}
func TexCoord2sv(v *int16) {
	syscall.Syscall(gpTexCoord2sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord2xOES(s int32, t int32) {
	syscall.Syscall(gpTexCoord2xOES, 2, uintptr(s), uintptr(t), 0)
}
func TexCoord2xvOES(coords *int32) {
	syscall.Syscall(gpTexCoord2xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func TexCoord3bOES(s int8, t int8, r int8) {
	syscall.Syscall(gpTexCoord3bOES, 3, uintptr(s), uintptr(t), uintptr(r))
}
func TexCoord3bvOES(coords *int8) {
	syscall.Syscall(gpTexCoord3bvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func TexCoord3d(s float64, t float64, r float64) {
	syscall.Syscall(gpTexCoord3d, 3, uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)), uintptr(math.Float64bits(r)))
}
func TexCoord3dv(v *float64) {
	syscall.Syscall(gpTexCoord3dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord3f(s float32, t float32, r float32) {
	syscall.Syscall(gpTexCoord3f, 3, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)))
}
func TexCoord3fv(v *float32) {
	syscall.Syscall(gpTexCoord3fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord3hNV(s uint16, t uint16, r uint16) {
	syscall.Syscall(gpTexCoord3hNV, 3, uintptr(s), uintptr(t), uintptr(r))
}
func TexCoord3hvNV(v *uint16) {
	syscall.Syscall(gpTexCoord3hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord3i(s int32, t int32, r int32) {
	syscall.Syscall(gpTexCoord3i, 3, uintptr(s), uintptr(t), uintptr(r))
}
func TexCoord3iv(v *int32) {
	syscall.Syscall(gpTexCoord3iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord3s(s int16, t int16, r int16) {
	syscall.Syscall(gpTexCoord3s, 3, uintptr(s), uintptr(t), uintptr(r))
}
func TexCoord3sv(v *int16) {
	syscall.Syscall(gpTexCoord3sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord3xOES(s int32, t int32, r int32) {
	syscall.Syscall(gpTexCoord3xOES, 3, uintptr(s), uintptr(t), uintptr(r))
}
func TexCoord3xvOES(coords *int32) {
	syscall.Syscall(gpTexCoord3xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func TexCoord4bOES(s int8, t int8, r int8, q int8) {
	syscall.Syscall6(gpTexCoord4bOES, 4, uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0, 0)
}
func TexCoord4bvOES(coords *int8) {
	syscall.Syscall(gpTexCoord4bvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func TexCoord4d(s float64, t float64, r float64, q float64) {
	syscall.Syscall6(gpTexCoord4d, 4, uintptr(math.Float64bits(s)), uintptr(math.Float64bits(t)), uintptr(math.Float64bits(r)), uintptr(math.Float64bits(q)), 0, 0)
}
func TexCoord4dv(v *float64) {
	syscall.Syscall(gpTexCoord4dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord4f(s float32, t float32, r float32, q float32) {
	syscall.Syscall6(gpTexCoord4f, 4, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(q)), 0, 0)
}
func TexCoord4fColor4fNormal3fVertex4fSUN(s float32, t float32, p float32, q float32, r float32, g float32, b float32, a float32, nx float32, ny float32, nz float32, x float32, y float32, z float32, w float32) {
	syscall.Syscall15(gpTexCoord4fColor4fNormal3fVertex4fSUN, 15, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(p)), uintptr(math.Float32bits(q)), uintptr(math.Float32bits(r)), uintptr(math.Float32bits(g)), uintptr(math.Float32bits(b)), uintptr(math.Float32bits(a)), uintptr(math.Float32bits(nx)), uintptr(math.Float32bits(ny)), uintptr(math.Float32bits(nz)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)))
}
func TexCoord4fColor4fNormal3fVertex4fvSUN(tc *float32, c *float32, n *float32, v *float32) {
	syscall.Syscall6(gpTexCoord4fColor4fNormal3fVertex4fvSUN, 4, uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(c)), uintptr(unsafe.Pointer(n)), uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord4fVertex4fSUN(s float32, t float32, p float32, q float32, x float32, y float32, z float32, w float32) {
	syscall.Syscall9(gpTexCoord4fVertex4fSUN, 8, uintptr(math.Float32bits(s)), uintptr(math.Float32bits(t)), uintptr(math.Float32bits(p)), uintptr(math.Float32bits(q)), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0)
}
func TexCoord4fVertex4fvSUN(tc *float32, v *float32) {
	syscall.Syscall(gpTexCoord4fVertex4fvSUN, 2, uintptr(unsafe.Pointer(tc)), uintptr(unsafe.Pointer(v)), 0)
}
func TexCoord4fv(v *float32) {
	syscall.Syscall(gpTexCoord4fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord4hNV(s uint16, t uint16, r uint16, q uint16) {
	syscall.Syscall6(gpTexCoord4hNV, 4, uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0, 0)
}
func TexCoord4hvNV(v *uint16) {
	syscall.Syscall(gpTexCoord4hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord4i(s int32, t int32, r int32, q int32) {
	syscall.Syscall6(gpTexCoord4i, 4, uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0, 0)
}
func TexCoord4iv(v *int32) {
	syscall.Syscall(gpTexCoord4iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord4s(s int16, t int16, r int16, q int16) {
	syscall.Syscall6(gpTexCoord4s, 4, uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0, 0)
}
func TexCoord4sv(v *int16) {
	syscall.Syscall(gpTexCoord4sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func TexCoord4xOES(s int32, t int32, r int32, q int32) {
	syscall.Syscall6(gpTexCoord4xOES, 4, uintptr(s), uintptr(t), uintptr(r), uintptr(q), 0, 0)
}
func TexCoord4xvOES(coords *int32) {
	syscall.Syscall(gpTexCoord4xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func TexCoordFormatNV(size int32, xtype uint32, stride int32) {
	syscall.Syscall(gpTexCoordFormatNV, 3, uintptr(size), uintptr(xtype), uintptr(stride))
}

// define an array of texture coordinates
func TexCoordPointer(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpTexCoordPointer, 4, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0, 0)
}
func TexCoordPointerEXT(size int32, xtype uint32, stride int32, count int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpTexCoordPointerEXT, 5, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(count), uintptr(pointer), 0)
}
func TexCoordPointerListIBM(size int32, xtype uint32, stride int32, pointer *unsafe.Pointer, ptrstride int32) {
	syscall.Syscall6(gpTexCoordPointerListIBM, 5, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(unsafe.Pointer(pointer)), uintptr(ptrstride), 0)
}
func TexCoordPointervINTEL(size int32, xtype uint32, pointer *unsafe.Pointer) {
	syscall.Syscall(gpTexCoordPointervINTEL, 3, uintptr(size), uintptr(xtype), uintptr(unsafe.Pointer(pointer)))
}
func TexEnvf(target uint32, pname uint32, param float32) {
	syscall.Syscall(gpTexEnvf, 3, uintptr(target), uintptr(pname), uintptr(math.Float32bits(param)))
}
func TexEnvfv(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpTexEnvfv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexEnvi(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpTexEnvi, 3, uintptr(target), uintptr(pname), uintptr(param))
}
func TexEnviv(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpTexEnviv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexEnvxOES(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpTexEnvxOES, 3, uintptr(target), uintptr(pname), uintptr(param))
}
func TexEnvxvOES(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpTexEnvxvOES, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexFilterFuncSGIS(target uint32, filter uint32, n int32, weights *float32) {
	syscall.Syscall6(gpTexFilterFuncSGIS, 4, uintptr(target), uintptr(filter), uintptr(n), uintptr(unsafe.Pointer(weights)), 0, 0)
}
func TexGend(coord uint32, pname uint32, param float64) {
	syscall.Syscall(gpTexGend, 3, uintptr(coord), uintptr(pname), uintptr(math.Float64bits(param)))
}
func TexGendv(coord uint32, pname uint32, params *float64) {
	syscall.Syscall(gpTexGendv, 3, uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexGenf(coord uint32, pname uint32, param float32) {
	syscall.Syscall(gpTexGenf, 3, uintptr(coord), uintptr(pname), uintptr(math.Float32bits(param)))
}
func TexGenfv(coord uint32, pname uint32, params *float32) {
	syscall.Syscall(gpTexGenfv, 3, uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexGeni(coord uint32, pname uint32, param int32) {
	syscall.Syscall(gpTexGeni, 3, uintptr(coord), uintptr(pname), uintptr(param))
}
func TexGeniv(coord uint32, pname uint32, params *int32) {
	syscall.Syscall(gpTexGeniv, 3, uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexGenxOES(coord uint32, pname uint32, param int32) {
	syscall.Syscall(gpTexGenxOES, 3, uintptr(coord), uintptr(pname), uintptr(param))
}
func TexGenxvOES(coord uint32, pname uint32, params *int32) {
	syscall.Syscall(gpTexGenxvOES, 3, uintptr(coord), uintptr(pname), uintptr(unsafe.Pointer(params)))
}

// specify a one-dimensional texture image
func TexImage1D(target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTexImage1D, 8, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}

// specify a two-dimensional texture image
func TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTexImage2D, 9, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels))
}

// establish the data storage, format, dimensions, and number of samples of a multisample texture's image
func TexImage2DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	syscall.Syscall6(gpTexImage2DMultisample, 6, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), boolToUintptr(fixedsamplelocations))
}
func TexImage2DMultisampleCoverageNV(target uint32, coverageSamples int32, colorSamples int32, internalFormat int32, width int32, height int32, fixedSampleLocations bool) {
	syscall.Syscall9(gpTexImage2DMultisampleCoverageNV, 7, uintptr(target), uintptr(coverageSamples), uintptr(colorSamples), uintptr(internalFormat), uintptr(width), uintptr(height), boolToUintptr(fixedSampleLocations), 0, 0)
}

// specify a three-dimensional texture image
func TexImage3D(target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTexImage3D, 10, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}
func TexImage3DEXT(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTexImage3DEXT, 10, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}

// establish the data storage, format, dimensions, and number of samples of a multisample texture's image
func TexImage3DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	syscall.Syscall9(gpTexImage3DMultisample, 7, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedsamplelocations), 0, 0)
}
func TexImage3DMultisampleCoverageNV(target uint32, coverageSamples int32, colorSamples int32, internalFormat int32, width int32, height int32, depth int32, fixedSampleLocations bool) {
	syscall.Syscall9(gpTexImage3DMultisampleCoverageNV, 8, uintptr(target), uintptr(coverageSamples), uintptr(colorSamples), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedSampleLocations), 0)
}
func TexImage4DSGIS(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, size4d int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTexImage4DSGIS, 11, uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(size4d), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func TexPageCommitmentARB(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, commit bool) {
	syscall.Syscall9(gpTexPageCommitmentARB, 9, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(commit))
}
func TexParameterIivEXT(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpTexParameterIivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexParameterIuivEXT(target uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpTexParameterIuivEXT, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexParameterf(target uint32, pname uint32, param float32) {
	syscall.Syscall(gpTexParameterf, 3, uintptr(target), uintptr(pname), uintptr(math.Float32bits(param)))
}
func TexParameterfv(target uint32, pname uint32, params *float32) {
	syscall.Syscall(gpTexParameterfv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexParameteri(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpTexParameteri, 3, uintptr(target), uintptr(pname), uintptr(param))
}
func TexParameteriv(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpTexParameteriv, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexParameterxOES(target uint32, pname uint32, param int32) {
	syscall.Syscall(gpTexParameterxOES, 3, uintptr(target), uintptr(pname), uintptr(param))
}
func TexParameterxvOES(target uint32, pname uint32, params *int32) {
	syscall.Syscall(gpTexParameterxvOES, 3, uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TexRenderbufferNV(target uint32, renderbuffer uint32) {
	syscall.Syscall(gpTexRenderbufferNV, 2, uintptr(target), uintptr(renderbuffer), 0)
}

// simultaneously specify storage for all levels of a one-dimensional texture
func TexStorage1D(target uint32, levels int32, internalformat uint32, width int32) {
	syscall.Syscall6(gpTexStorage1D, 4, uintptr(target), uintptr(levels), uintptr(internalformat), uintptr(width), 0, 0)
}

// simultaneously specify storage for all levels of a two-dimensional or one-dimensional array texture
func TexStorage2D(target uint32, levels int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpTexStorage2D, 5, uintptr(target), uintptr(levels), uintptr(internalformat), uintptr(width), uintptr(height), 0)
}

// specify storage for a two-dimensional multisample texture
func TexStorage2DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	syscall.Syscall6(gpTexStorage2DMultisample, 6, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), boolToUintptr(fixedsamplelocations))
}

// simultaneously specify storage for all levels of a three-dimensional, two-dimensional array or cube-map array texture
func TexStorage3D(target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	syscall.Syscall6(gpTexStorage3D, 6, uintptr(target), uintptr(levels), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth))
}

// specify storage for a two-dimensional multisample array texture
func TexStorage3DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	syscall.Syscall9(gpTexStorage3DMultisample, 7, uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedsamplelocations), 0, 0)
}
func TexStorageMem1DEXT(target uint32, levels int32, internalFormat uint32, width int32, memory uint32, offset uint64) {
	syscall.Syscall6(gpTexStorageMem1DEXT, 6, uintptr(target), uintptr(levels), uintptr(internalFormat), uintptr(width), uintptr(memory), uintptr(offset))
}
func TexStorageMem2DEXT(target uint32, levels int32, internalFormat uint32, width int32, height int32, memory uint32, offset uint64) {
	syscall.Syscall9(gpTexStorageMem2DEXT, 7, uintptr(target), uintptr(levels), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(memory), uintptr(offset), 0, 0)
}
func TexStorageMem2DMultisampleEXT(target uint32, samples int32, internalFormat uint32, width int32, height int32, fixedSampleLocations bool, memory uint32, offset uint64) {
	syscall.Syscall9(gpTexStorageMem2DMultisampleEXT, 8, uintptr(target), uintptr(samples), uintptr(internalFormat), uintptr(width), uintptr(height), boolToUintptr(fixedSampleLocations), uintptr(memory), uintptr(offset), 0)
}
func TexStorageMem3DEXT(target uint32, levels int32, internalFormat uint32, width int32, height int32, depth int32, memory uint32, offset uint64) {
	syscall.Syscall9(gpTexStorageMem3DEXT, 8, uintptr(target), uintptr(levels), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), uintptr(memory), uintptr(offset), 0)
}
func TexStorageMem3DMultisampleEXT(target uint32, samples int32, internalFormat uint32, width int32, height int32, depth int32, fixedSampleLocations bool, memory uint32, offset uint64) {
	syscall.Syscall9(gpTexStorageMem3DMultisampleEXT, 9, uintptr(target), uintptr(samples), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedSampleLocations), uintptr(memory), uintptr(offset))
}
func TexStorageSparseAMD(target uint32, internalFormat uint32, width int32, height int32, depth int32, layers int32, flags uint32) {
	syscall.Syscall9(gpTexStorageSparseAMD, 7, uintptr(target), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), uintptr(layers), uintptr(flags), 0, 0)
}

// specify a one-dimensional texture subimage
func TexSubImage1D(target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTexSubImage1D, 7, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}
func TexSubImage1DEXT(target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTexSubImage1DEXT, 7, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}

// specify a two-dimensional texture subimage
func TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTexSubImage2D, 9, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(pixels))
}
func TexSubImage2DEXT(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTexSubImage2DEXT, 9, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(pixels))
}

// specify a three-dimensional texture subimage
func TexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTexSubImage3D, 11, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func TexSubImage3DEXT(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTexSubImage3DEXT, 11, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func TexSubImage4DSGIS(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, woffset int32, width int32, height int32, depth int32, size4d int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall15(gpTexSubImage4DSGIS, 13, uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(woffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(size4d), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}

// controls the ordering of reads and writes to rendered fragments across drawing commands
func TextureBarrier() {
	syscall.Syscall(gpTextureBarrier, 0, 0, 0, 0)
}
func TextureBarrierNV() {
	syscall.Syscall(gpTextureBarrierNV, 0, 0, 0, 0)
}

// attach a buffer object's data store to a buffer texture object
func TextureBuffer(texture uint32, internalformat uint32, buffer uint32) {
	syscall.Syscall(gpTextureBuffer, 3, uintptr(texture), uintptr(internalformat), uintptr(buffer))
}
func TextureBufferEXT(texture uint32, target uint32, internalformat uint32, buffer uint32) {
	syscall.Syscall6(gpTextureBufferEXT, 4, uintptr(texture), uintptr(target), uintptr(internalformat), uintptr(buffer), 0, 0)
}

// attach a range of a buffer object's data store to a buffer texture object
func TextureBufferRange(texture uint32, internalformat uint32, buffer uint32, offset int, size int) {
	syscall.Syscall6(gpTextureBufferRange, 5, uintptr(texture), uintptr(internalformat), uintptr(buffer), uintptr(offset), uintptr(size), 0)
}
func TextureBufferRangeEXT(texture uint32, target uint32, internalformat uint32, buffer uint32, offset int, size int) {
	syscall.Syscall6(gpTextureBufferRangeEXT, 6, uintptr(texture), uintptr(target), uintptr(internalformat), uintptr(buffer), uintptr(offset), uintptr(size))
}
func TextureColorMaskSGIS(red bool, green bool, blue bool, alpha bool) {
	syscall.Syscall6(gpTextureColorMaskSGIS, 4, boolToUintptr(red), boolToUintptr(green), boolToUintptr(blue), boolToUintptr(alpha), 0, 0)
}
func TextureImage1DEXT(texture uint32, target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTextureImage1DEXT, 9, uintptr(texture), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels))
}
func TextureImage2DEXT(texture uint32, target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTextureImage2DEXT, 10, uintptr(texture), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}
func TextureImage2DMultisampleCoverageNV(texture uint32, target uint32, coverageSamples int32, colorSamples int32, internalFormat int32, width int32, height int32, fixedSampleLocations bool) {
	syscall.Syscall9(gpTextureImage2DMultisampleCoverageNV, 8, uintptr(texture), uintptr(target), uintptr(coverageSamples), uintptr(colorSamples), uintptr(internalFormat), uintptr(width), uintptr(height), boolToUintptr(fixedSampleLocations), 0)
}
func TextureImage2DMultisampleNV(texture uint32, target uint32, samples int32, internalFormat int32, width int32, height int32, fixedSampleLocations bool) {
	syscall.Syscall9(gpTextureImage2DMultisampleNV, 7, uintptr(texture), uintptr(target), uintptr(samples), uintptr(internalFormat), uintptr(width), uintptr(height), boolToUintptr(fixedSampleLocations), 0, 0)
}
func TextureImage3DEXT(texture uint32, target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTextureImage3DEXT, 11, uintptr(texture), uintptr(target), uintptr(level), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), uintptr(border), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func TextureImage3DMultisampleCoverageNV(texture uint32, target uint32, coverageSamples int32, colorSamples int32, internalFormat int32, width int32, height int32, depth int32, fixedSampleLocations bool) {
	syscall.Syscall9(gpTextureImage3DMultisampleCoverageNV, 9, uintptr(texture), uintptr(target), uintptr(coverageSamples), uintptr(colorSamples), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedSampleLocations))
}
func TextureImage3DMultisampleNV(texture uint32, target uint32, samples int32, internalFormat int32, width int32, height int32, depth int32, fixedSampleLocations bool) {
	syscall.Syscall9(gpTextureImage3DMultisampleNV, 8, uintptr(texture), uintptr(target), uintptr(samples), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedSampleLocations), 0)
}
func TextureLightEXT(pname uint32) {
	syscall.Syscall(gpTextureLightEXT, 1, uintptr(pname), 0, 0)
}
func TextureMaterialEXT(face uint32, mode uint32) {
	syscall.Syscall(gpTextureMaterialEXT, 2, uintptr(face), uintptr(mode), 0)
}
func TextureNormalEXT(mode uint32) {
	syscall.Syscall(gpTextureNormalEXT, 1, uintptr(mode), 0, 0)
}
func TexturePageCommitmentEXT(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, commit bool) {
	syscall.Syscall9(gpTexturePageCommitmentEXT, 9, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(commit))
}
func TextureParameterIiv(texture uint32, pname uint32, params *int32) {
	syscall.Syscall(gpTextureParameterIiv, 3, uintptr(texture), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TextureParameterIivEXT(texture uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpTextureParameterIivEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func TextureParameterIuiv(texture uint32, pname uint32, params *uint32) {
	syscall.Syscall(gpTextureParameterIuiv, 3, uintptr(texture), uintptr(pname), uintptr(unsafe.Pointer(params)))
}
func TextureParameterIuivEXT(texture uint32, target uint32, pname uint32, params *uint32) {
	syscall.Syscall6(gpTextureParameterIuivEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func TextureParameterf(texture uint32, pname uint32, param float32) {
	syscall.Syscall(gpTextureParameterf, 3, uintptr(texture), uintptr(pname), uintptr(math.Float32bits(param)))
}
func TextureParameterfEXT(texture uint32, target uint32, pname uint32, param float32) {
	syscall.Syscall6(gpTextureParameterfEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(math.Float32bits(param)), 0, 0)
}
func TextureParameterfv(texture uint32, pname uint32, param *float32) {
	syscall.Syscall(gpTextureParameterfv, 3, uintptr(texture), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func TextureParameterfvEXT(texture uint32, target uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpTextureParameterfvEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func TextureParameteri(texture uint32, pname uint32, param int32) {
	syscall.Syscall(gpTextureParameteri, 3, uintptr(texture), uintptr(pname), uintptr(param))
}
func TextureParameteriEXT(texture uint32, target uint32, pname uint32, param int32) {
	syscall.Syscall6(gpTextureParameteriEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(param), 0, 0)
}
func TextureParameteriv(texture uint32, pname uint32, param *int32) {
	syscall.Syscall(gpTextureParameteriv, 3, uintptr(texture), uintptr(pname), uintptr(unsafe.Pointer(param)))
}
func TextureParameterivEXT(texture uint32, target uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpTextureParameterivEXT, 4, uintptr(texture), uintptr(target), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func TextureRangeAPPLE(target uint32, length int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpTextureRangeAPPLE, 3, uintptr(target), uintptr(length), uintptr(pointer))
}
func TextureRenderbufferEXT(texture uint32, target uint32, renderbuffer uint32) {
	syscall.Syscall(gpTextureRenderbufferEXT, 3, uintptr(texture), uintptr(target), uintptr(renderbuffer))
}

// simultaneously specify storage for all levels of a one-dimensional texture
func TextureStorage1D(texture uint32, levels int32, internalformat uint32, width int32) {
	syscall.Syscall6(gpTextureStorage1D, 4, uintptr(texture), uintptr(levels), uintptr(internalformat), uintptr(width), 0, 0)
}
func TextureStorage1DEXT(texture uint32, target uint32, levels int32, internalformat uint32, width int32) {
	syscall.Syscall6(gpTextureStorage1DEXT, 5, uintptr(texture), uintptr(target), uintptr(levels), uintptr(internalformat), uintptr(width), 0)
}

// simultaneously specify storage for all levels of a two-dimensional or one-dimensional array texture
func TextureStorage2D(texture uint32, levels int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpTextureStorage2D, 5, uintptr(texture), uintptr(levels), uintptr(internalformat), uintptr(width), uintptr(height), 0)
}
func TextureStorage2DEXT(texture uint32, target uint32, levels int32, internalformat uint32, width int32, height int32) {
	syscall.Syscall6(gpTextureStorage2DEXT, 6, uintptr(texture), uintptr(target), uintptr(levels), uintptr(internalformat), uintptr(width), uintptr(height))
}

// specify storage for a two-dimensional multisample texture
func TextureStorage2DMultisample(texture uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	syscall.Syscall6(gpTextureStorage2DMultisample, 6, uintptr(texture), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), boolToUintptr(fixedsamplelocations))
}
func TextureStorage2DMultisampleEXT(texture uint32, target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	syscall.Syscall9(gpTextureStorage2DMultisampleEXT, 7, uintptr(texture), uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), boolToUintptr(fixedsamplelocations), 0, 0)
}

// simultaneously specify storage for all levels of a three-dimensional, two-dimensional array or cube-map array texture
func TextureStorage3D(texture uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	syscall.Syscall6(gpTextureStorage3D, 6, uintptr(texture), uintptr(levels), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth))
}
func TextureStorage3DEXT(texture uint32, target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	syscall.Syscall9(gpTextureStorage3DEXT, 7, uintptr(texture), uintptr(target), uintptr(levels), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), 0, 0)
}

// specify storage for a two-dimensional multisample array texture
func TextureStorage3DMultisample(texture uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	syscall.Syscall9(gpTextureStorage3DMultisample, 7, uintptr(texture), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedsamplelocations), 0, 0)
}
func TextureStorage3DMultisampleEXT(texture uint32, target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	syscall.Syscall9(gpTextureStorage3DMultisampleEXT, 8, uintptr(texture), uintptr(target), uintptr(samples), uintptr(internalformat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedsamplelocations), 0)
}
func TextureStorageMem1DEXT(texture uint32, levels int32, internalFormat uint32, width int32, memory uint32, offset uint64) {
	syscall.Syscall6(gpTextureStorageMem1DEXT, 6, uintptr(texture), uintptr(levels), uintptr(internalFormat), uintptr(width), uintptr(memory), uintptr(offset))
}
func TextureStorageMem2DEXT(texture uint32, levels int32, internalFormat uint32, width int32, height int32, memory uint32, offset uint64) {
	syscall.Syscall9(gpTextureStorageMem2DEXT, 7, uintptr(texture), uintptr(levels), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(memory), uintptr(offset), 0, 0)
}
func TextureStorageMem2DMultisampleEXT(texture uint32, samples int32, internalFormat uint32, width int32, height int32, fixedSampleLocations bool, memory uint32, offset uint64) {
	syscall.Syscall9(gpTextureStorageMem2DMultisampleEXT, 8, uintptr(texture), uintptr(samples), uintptr(internalFormat), uintptr(width), uintptr(height), boolToUintptr(fixedSampleLocations), uintptr(memory), uintptr(offset), 0)
}
func TextureStorageMem3DEXT(texture uint32, levels int32, internalFormat uint32, width int32, height int32, depth int32, memory uint32, offset uint64) {
	syscall.Syscall9(gpTextureStorageMem3DEXT, 8, uintptr(texture), uintptr(levels), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), uintptr(memory), uintptr(offset), 0)
}
func TextureStorageMem3DMultisampleEXT(texture uint32, samples int32, internalFormat uint32, width int32, height int32, depth int32, fixedSampleLocations bool, memory uint32, offset uint64) {
	syscall.Syscall9(gpTextureStorageMem3DMultisampleEXT, 9, uintptr(texture), uintptr(samples), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), boolToUintptr(fixedSampleLocations), uintptr(memory), uintptr(offset))
}
func TextureStorageSparseAMD(texture uint32, target uint32, internalFormat uint32, width int32, height int32, depth int32, layers int32, flags uint32) {
	syscall.Syscall9(gpTextureStorageSparseAMD, 8, uintptr(texture), uintptr(target), uintptr(internalFormat), uintptr(width), uintptr(height), uintptr(depth), uintptr(layers), uintptr(flags), 0)
}

// specify a one-dimensional texture subimage
func TextureSubImage1D(texture uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTextureSubImage1D, 7, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}
func TextureSubImage1DEXT(texture uint32, target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTextureSubImage1DEXT, 8, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(width), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}

// specify a two-dimensional texture subimage
func TextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall9(gpTextureSubImage2D, 9, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(pixels))
}
func TextureSubImage2DEXT(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTextureSubImage2DEXT, 10, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(width), uintptr(height), uintptr(format), uintptr(xtype), uintptr(pixels), 0, 0)
}

// specify a three-dimensional texture subimage
func TextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTextureSubImage3D, 11, uintptr(texture), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(xtype), uintptr(pixels), 0)
}
func TextureSubImage3DEXT(texture uint32, target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	syscall.Syscall12(gpTextureSubImage3DEXT, 12, uintptr(texture), uintptr(target), uintptr(level), uintptr(xoffset), uintptr(yoffset), uintptr(zoffset), uintptr(width), uintptr(height), uintptr(depth), uintptr(format), uintptr(xtype), uintptr(pixels))
}

// initialize a texture as a data alias of another texture's data store
func TextureView(texture uint32, target uint32, origtexture uint32, internalformat uint32, minlevel uint32, numlevels uint32, minlayer uint32, numlayers uint32) {
	syscall.Syscall9(gpTextureView, 8, uintptr(texture), uintptr(target), uintptr(origtexture), uintptr(internalformat), uintptr(minlevel), uintptr(numlevels), uintptr(minlayer), uintptr(numlayers), 0)
}
func TrackMatrixNV(target uint32, address uint32, matrix uint32, transform uint32) {
	syscall.Syscall6(gpTrackMatrixNV, 4, uintptr(target), uintptr(address), uintptr(matrix), uintptr(transform), 0, 0)
}
func TransformFeedbackAttribsNV(count int32, attribs *int32, bufferMode uint32) {
	syscall.Syscall(gpTransformFeedbackAttribsNV, 3, uintptr(count), uintptr(unsafe.Pointer(attribs)), uintptr(bufferMode))
}

// bind a buffer object to a transform feedback buffer object
func TransformFeedbackBufferBase(xfb uint32, index uint32, buffer uint32) {
	syscall.Syscall(gpTransformFeedbackBufferBase, 3, uintptr(xfb), uintptr(index), uintptr(buffer))
}

// bind a range within a buffer object to a transform feedback buffer object
func TransformFeedbackBufferRange(xfb uint32, index uint32, buffer uint32, offset int, size int) {
	syscall.Syscall6(gpTransformFeedbackBufferRange, 5, uintptr(xfb), uintptr(index), uintptr(buffer), uintptr(offset), uintptr(size), 0)
}
func TransformFeedbackStreamAttribsNV(count int32, attribs *int32, nbuffers int32, bufstreams *int32, bufferMode uint32) {
	syscall.Syscall6(gpTransformFeedbackStreamAttribsNV, 5, uintptr(count), uintptr(unsafe.Pointer(attribs)), uintptr(nbuffers), uintptr(unsafe.Pointer(bufstreams)), uintptr(bufferMode), 0)
}
func TransformFeedbackVaryingsEXT(program uint32, count int32, varyings **uint8, bufferMode uint32) {
	syscall.Syscall6(gpTransformFeedbackVaryingsEXT, 4, uintptr(program), uintptr(count), uintptr(unsafe.Pointer(varyings)), uintptr(bufferMode), 0, 0)
}
func TransformFeedbackVaryingsNV(program uint32, count int32, locations *int32, bufferMode uint32) {
	syscall.Syscall6(gpTransformFeedbackVaryingsNV, 4, uintptr(program), uintptr(count), uintptr(unsafe.Pointer(locations)), uintptr(bufferMode), 0, 0)
}
func TransformPathNV(resultPath uint32, srcPath uint32, transformType uint32, transformValues *float32) {
	syscall.Syscall6(gpTransformPathNV, 4, uintptr(resultPath), uintptr(srcPath), uintptr(transformType), uintptr(unsafe.Pointer(transformValues)), 0, 0)
}
func Translated(x float64, y float64, z float64) {
	syscall.Syscall(gpTranslated, 3, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)))
}
func Translatef(x float32, y float32, z float32) {
	syscall.Syscall(gpTranslatef, 3, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func TranslatexOES(x int32, y int32, z int32) {
	syscall.Syscall(gpTranslatexOES, 3, uintptr(x), uintptr(y), uintptr(z))
}
func Uniform1d(location int32, x float64) {
	syscall.Syscall(gpUniform1d, 2, uintptr(location), uintptr(math.Float64bits(x)), 0)
}
func Uniform1dv(location int32, count int32, value *float64) {
	syscall.Syscall(gpUniform1dv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1f(location int32, v0 float32) {
	syscall.Syscall(gpUniform1f, 2, uintptr(location), uintptr(math.Float32bits(v0)), 0)
}
func Uniform1fARB(location int32, v0 float32) {
	syscall.Syscall(gpUniform1fARB, 2, uintptr(location), uintptr(math.Float32bits(v0)), 0)
}

// Specify the value of a uniform variable for the current program object
func Uniform1fv(location int32, count int32, value *float32) {
	syscall.Syscall(gpUniform1fv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform1fvARB(location int32, count int32, value *float32) {
	syscall.Syscall(gpUniform1fvARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1i(location int32, v0 int32) {
	syscall.Syscall(gpUniform1i, 2, uintptr(location), uintptr(v0), 0)
}
func Uniform1i64ARB(location int32, x int64) {
	syscall.Syscall(gpUniform1i64ARB, 2, uintptr(location), uintptr(x), 0)
}
func Uniform1i64NV(location int32, x int64) {
	syscall.Syscall(gpUniform1i64NV, 2, uintptr(location), uintptr(x), 0)
}
func Uniform1i64vARB(location int32, count int32, value *int64) {
	syscall.Syscall(gpUniform1i64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform1i64vNV(location int32, count int32, value *int64) {
	syscall.Syscall(gpUniform1i64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform1iARB(location int32, v0 int32) {
	syscall.Syscall(gpUniform1iARB, 2, uintptr(location), uintptr(v0), 0)
}

// Specify the value of a uniform variable for the current program object
func Uniform1iv(location int32, count int32, value *int32) {
	syscall.Syscall(gpUniform1iv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform1ivARB(location int32, count int32, value *int32) {
	syscall.Syscall(gpUniform1ivARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform1ui64ARB(location int32, x uint64) {
	syscall.Syscall(gpUniform1ui64ARB, 2, uintptr(location), uintptr(x), 0)
}
func Uniform1ui64NV(location int32, x uint64) {
	syscall.Syscall(gpUniform1ui64NV, 2, uintptr(location), uintptr(x), 0)
}
func Uniform1ui64vARB(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniform1ui64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform1ui64vNV(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniform1ui64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform1uiEXT(location int32, v0 uint32) {
	syscall.Syscall(gpUniform1uiEXT, 2, uintptr(location), uintptr(v0), 0)
}
func Uniform1uivEXT(location int32, count int32, value *uint32) {
	syscall.Syscall(gpUniform1uivEXT, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform2d(location int32, x float64, y float64) {
	syscall.Syscall(gpUniform2d, 3, uintptr(location), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
}
func Uniform2dv(location int32, count int32, value *float64) {
	syscall.Syscall(gpUniform2dv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2f(location int32, v0 float32, v1 float32) {
	syscall.Syscall(gpUniform2f, 3, uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)))
}
func Uniform2fARB(location int32, v0 float32, v1 float32) {
	syscall.Syscall(gpUniform2fARB, 3, uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2fv(location int32, count int32, value *float32) {
	syscall.Syscall(gpUniform2fv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform2fvARB(location int32, count int32, value *float32) {
	syscall.Syscall(gpUniform2fvARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2i(location int32, v0 int32, v1 int32) {
	syscall.Syscall(gpUniform2i, 3, uintptr(location), uintptr(v0), uintptr(v1))
}
func Uniform2i64ARB(location int32, x int64, y int64) {
	syscall.Syscall(gpUniform2i64ARB, 3, uintptr(location), uintptr(x), uintptr(y))
}
func Uniform2i64NV(location int32, x int64, y int64) {
	syscall.Syscall(gpUniform2i64NV, 3, uintptr(location), uintptr(x), uintptr(y))
}
func Uniform2i64vARB(location int32, count int32, value *int64) {
	syscall.Syscall(gpUniform2i64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform2i64vNV(location int32, count int32, value *int64) {
	syscall.Syscall(gpUniform2i64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform2iARB(location int32, v0 int32, v1 int32) {
	syscall.Syscall(gpUniform2iARB, 3, uintptr(location), uintptr(v0), uintptr(v1))
}

// Specify the value of a uniform variable for the current program object
func Uniform2iv(location int32, count int32, value *int32) {
	syscall.Syscall(gpUniform2iv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform2ivARB(location int32, count int32, value *int32) {
	syscall.Syscall(gpUniform2ivARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform2ui64ARB(location int32, x uint64, y uint64) {
	syscall.Syscall(gpUniform2ui64ARB, 3, uintptr(location), uintptr(x), uintptr(y))
}
func Uniform2ui64NV(location int32, x uint64, y uint64) {
	syscall.Syscall(gpUniform2ui64NV, 3, uintptr(location), uintptr(x), uintptr(y))
}
func Uniform2ui64vARB(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniform2ui64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform2ui64vNV(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniform2ui64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform2uiEXT(location int32, v0 uint32, v1 uint32) {
	syscall.Syscall(gpUniform2uiEXT, 3, uintptr(location), uintptr(v0), uintptr(v1))
}
func Uniform2uivEXT(location int32, count int32, value *uint32) {
	syscall.Syscall(gpUniform2uivEXT, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform3d(location int32, x float64, y float64, z float64) {
	syscall.Syscall6(gpUniform3d, 4, uintptr(location), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func Uniform3dv(location int32, count int32, value *float64) {
	syscall.Syscall(gpUniform3dv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3f(location int32, v0 float32, v1 float32, v2 float32) {
	syscall.Syscall6(gpUniform3f, 4, uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), 0, 0)
}
func Uniform3fARB(location int32, v0 float32, v1 float32, v2 float32) {
	syscall.Syscall6(gpUniform3fARB, 4, uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func Uniform3fv(location int32, count int32, value *float32) {
	syscall.Syscall(gpUniform3fv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform3fvARB(location int32, count int32, value *float32) {
	syscall.Syscall(gpUniform3fvARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3i(location int32, v0 int32, v1 int32, v2 int32) {
	syscall.Syscall6(gpUniform3i, 4, uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), 0, 0)
}
func Uniform3i64ARB(location int32, x int64, y int64, z int64) {
	syscall.Syscall6(gpUniform3i64ARB, 4, uintptr(location), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func Uniform3i64NV(location int32, x int64, y int64, z int64) {
	syscall.Syscall6(gpUniform3i64NV, 4, uintptr(location), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func Uniform3i64vARB(location int32, count int32, value *int64) {
	syscall.Syscall(gpUniform3i64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform3i64vNV(location int32, count int32, value *int64) {
	syscall.Syscall(gpUniform3i64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform3iARB(location int32, v0 int32, v1 int32, v2 int32) {
	syscall.Syscall6(gpUniform3iARB, 4, uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func Uniform3iv(location int32, count int32, value *int32) {
	syscall.Syscall(gpUniform3iv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform3ivARB(location int32, count int32, value *int32) {
	syscall.Syscall(gpUniform3ivARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform3ui64ARB(location int32, x uint64, y uint64, z uint64) {
	syscall.Syscall6(gpUniform3ui64ARB, 4, uintptr(location), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func Uniform3ui64NV(location int32, x uint64, y uint64, z uint64) {
	syscall.Syscall6(gpUniform3ui64NV, 4, uintptr(location), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func Uniform3ui64vARB(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniform3ui64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform3ui64vNV(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniform3ui64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform3uiEXT(location int32, v0 uint32, v1 uint32, v2 uint32) {
	syscall.Syscall6(gpUniform3uiEXT, 4, uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), 0, 0)
}
func Uniform3uivEXT(location int32, count int32, value *uint32) {
	syscall.Syscall(gpUniform3uivEXT, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform4d(location int32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpUniform4d, 5, uintptr(location), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0)
}
func Uniform4dv(location int32, count int32, value *float64) {
	syscall.Syscall(gpUniform4dv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	syscall.Syscall6(gpUniform4f, 5, uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), uintptr(math.Float32bits(v3)), 0)
}
func Uniform4fARB(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	syscall.Syscall6(gpUniform4fARB, 5, uintptr(location), uintptr(math.Float32bits(v0)), uintptr(math.Float32bits(v1)), uintptr(math.Float32bits(v2)), uintptr(math.Float32bits(v3)), 0)
}

// Specify the value of a uniform variable for the current program object
func Uniform4fv(location int32, count int32, value *float32) {
	syscall.Syscall(gpUniform4fv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform4fvARB(location int32, count int32, value *float32) {
	syscall.Syscall(gpUniform4fvARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4i(location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	syscall.Syscall6(gpUniform4i, 5, uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3), 0)
}
func Uniform4i64ARB(location int32, x int64, y int64, z int64, w int64) {
	syscall.Syscall6(gpUniform4i64ARB, 5, uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func Uniform4i64NV(location int32, x int64, y int64, z int64, w int64) {
	syscall.Syscall6(gpUniform4i64NV, 5, uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func Uniform4i64vARB(location int32, count int32, value *int64) {
	syscall.Syscall(gpUniform4i64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform4i64vNV(location int32, count int32, value *int64) {
	syscall.Syscall(gpUniform4i64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform4iARB(location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	syscall.Syscall6(gpUniform4iARB, 5, uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3), 0)
}

// Specify the value of a uniform variable for the current program object
func Uniform4iv(location int32, count int32, value *int32) {
	syscall.Syscall(gpUniform4iv, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform4ivARB(location int32, count int32, value *int32) {
	syscall.Syscall(gpUniform4ivARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform4ui64ARB(location int32, x uint64, y uint64, z uint64, w uint64) {
	syscall.Syscall6(gpUniform4ui64ARB, 5, uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func Uniform4ui64NV(location int32, x uint64, y uint64, z uint64, w uint64) {
	syscall.Syscall6(gpUniform4ui64NV, 5, uintptr(location), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func Uniform4ui64vARB(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniform4ui64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform4ui64vNV(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniform4ui64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func Uniform4uiEXT(location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	syscall.Syscall6(gpUniform4uiEXT, 5, uintptr(location), uintptr(v0), uintptr(v1), uintptr(v2), uintptr(v3), 0)
}
func Uniform4uivEXT(location int32, count int32, value *uint32) {
	syscall.Syscall(gpUniform4uivEXT, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}

// assign a binding point to an active uniform block
func UniformBlockBinding(program uint32, uniformBlockIndex uint32, uniformBlockBinding uint32) {
	syscall.Syscall(gpUniformBlockBinding, 3, uintptr(program), uintptr(uniformBlockIndex), uintptr(uniformBlockBinding))
}
func UniformBufferEXT(program uint32, location int32, buffer uint32) {
	syscall.Syscall(gpUniformBufferEXT, 3, uintptr(program), uintptr(location), uintptr(buffer))
}
func UniformHandleui64ARB(location int32, value uint64) {
	syscall.Syscall(gpUniformHandleui64ARB, 2, uintptr(location), uintptr(value), 0)
}
func UniformHandleui64NV(location int32, value uint64) {
	syscall.Syscall(gpUniformHandleui64NV, 2, uintptr(location), uintptr(value), 0)
}
func UniformHandleui64vARB(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniformHandleui64vARB, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func UniformHandleui64vNV(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniformHandleui64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func UniformMatrix2dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix2dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix2fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix2fvARB(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix2fvARB, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix2x3dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix2x3dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2x3fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix2x3fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix2x4dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix2x4dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2x4fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix2x4fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix3dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix3dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix3fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix3fvARB(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix3fvARB, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix3x2dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix3x2dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3x2fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix3x2fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix3x4dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix3x4dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3x4fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix3x4fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix4dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix4dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix4fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix4fvARB(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix4fvARB, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix4x2dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix4x2dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4x2fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix4x2fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformMatrix4x3dv(location int32, count int32, transpose bool, value *float64) {
	syscall.Syscall6(gpUniformMatrix4x3dv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4x3fv(location int32, count int32, transpose bool, value *float32) {
	syscall.Syscall6(gpUniformMatrix4x3fv, 4, uintptr(location), uintptr(count), boolToUintptr(transpose), uintptr(unsafe.Pointer(value)), 0, 0)
}
func UniformSubroutinesuiv(shadertype uint32, count int32, indices *uint32) {
	syscall.Syscall(gpUniformSubroutinesuiv, 3, uintptr(shadertype), uintptr(count), uintptr(unsafe.Pointer(indices)))
}
func Uniformui64NV(location int32, value uint64) {
	syscall.Syscall(gpUniformui64NV, 2, uintptr(location), uintptr(value), 0)
}
func Uniformui64vNV(location int32, count int32, value *uint64) {
	syscall.Syscall(gpUniformui64vNV, 3, uintptr(location), uintptr(count), uintptr(unsafe.Pointer(value)))
}
func UnlockArraysEXT() {
	syscall.Syscall(gpUnlockArraysEXT, 0, 0, 0, 0)
}

// release the mapping of a buffer object's data store into the client's address space
func UnmapBuffer(target uint32) bool {
	ret, _, _ := syscall.Syscall(gpUnmapBuffer, 1, uintptr(target), 0, 0)
	return ret != 0
}
func UnmapBufferARB(target uint32) bool {
	ret, _, _ := syscall.Syscall(gpUnmapBufferARB, 1, uintptr(target), 0, 0)
	return ret != 0
}

// release the mapping of a buffer object's data store into the client's address space
func UnmapNamedBuffer(buffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpUnmapNamedBuffer, 1, uintptr(buffer), 0, 0)
	return ret != 0
}
func UnmapNamedBufferEXT(buffer uint32) bool {
	ret, _, _ := syscall.Syscall(gpUnmapNamedBufferEXT, 1, uintptr(buffer), 0, 0)
	return ret != 0
}
func UnmapObjectBufferATI(buffer uint32) {
	syscall.Syscall(gpUnmapObjectBufferATI, 1, uintptr(buffer), 0, 0)
}
func UnmapTexture2DINTEL(texture uint32, level int32) {
	syscall.Syscall(gpUnmapTexture2DINTEL, 2, uintptr(texture), uintptr(level), 0)
}
func UpdateObjectBufferATI(buffer uint32, offset uint32, size int32, pointer unsafe.Pointer, preserve uint32) {
	syscall.Syscall6(gpUpdateObjectBufferATI, 5, uintptr(buffer), uintptr(offset), uintptr(size), uintptr(pointer), uintptr(preserve), 0)
}

// Installs a program object as part of current rendering state
func UseProgram(program uint32) {
	syscall.Syscall(gpUseProgram, 1, uintptr(program), 0, 0)
}
func UseProgramObjectARB(programObj uintptr) {
	syscall.Syscall(gpUseProgramObjectARB, 1, uintptr(programObj), 0, 0)
}

// bind stages of a program object to a program pipeline
func UseProgramStages(pipeline uint32, stages uint32, program uint32) {
	syscall.Syscall(gpUseProgramStages, 3, uintptr(pipeline), uintptr(stages), uintptr(program))
}
func UseProgramStagesEXT(pipeline uint32, stages uint32, program uint32) {
	syscall.Syscall(gpUseProgramStagesEXT, 3, uintptr(pipeline), uintptr(stages), uintptr(program))
}
func UseShaderProgramEXT(xtype uint32, program uint32) {
	syscall.Syscall(gpUseShaderProgramEXT, 2, uintptr(xtype), uintptr(program), 0)
}
func VDPAUFiniNV() {
	syscall.Syscall(gpVDPAUFiniNV, 0, 0, 0, 0)
}
func VDPAUGetSurfaceivNV(surface uintptr, pname uint32, bufSize int32, length *int32, values *int32) {
	syscall.Syscall6(gpVDPAUGetSurfaceivNV, 5, uintptr(surface), uintptr(pname), uintptr(bufSize), uintptr(unsafe.Pointer(length)), uintptr(unsafe.Pointer(values)), 0)
}
func VDPAUInitNV(vdpDevice unsafe.Pointer, getProcAddress unsafe.Pointer) {
	syscall.Syscall(gpVDPAUInitNV, 2, uintptr(vdpDevice), uintptr(getProcAddress), 0)
}
func VDPAUIsSurfaceNV(surface uintptr) bool {
	ret, _, _ := syscall.Syscall(gpVDPAUIsSurfaceNV, 1, uintptr(surface), 0, 0)
	return ret != 0
}
func VDPAUMapSurfacesNV(numSurfaces int32, surfaces *uintptr) {
	syscall.Syscall(gpVDPAUMapSurfacesNV, 2, uintptr(numSurfaces), uintptr(unsafe.Pointer(surfaces)), 0)
}
func VDPAURegisterOutputSurfaceNV(vdpSurface unsafe.Pointer, target uint32, numTextureNames int32, textureNames *uint32) uintptr {
	ret, _, _ := syscall.Syscall6(gpVDPAURegisterOutputSurfaceNV, 4, uintptr(vdpSurface), uintptr(target), uintptr(numTextureNames), uintptr(unsafe.Pointer(textureNames)), 0, 0)
	return (uintptr)(ret)
}
func VDPAURegisterVideoSurfaceNV(vdpSurface unsafe.Pointer, target uint32, numTextureNames int32, textureNames *uint32) uintptr {
	ret, _, _ := syscall.Syscall6(gpVDPAURegisterVideoSurfaceNV, 4, uintptr(vdpSurface), uintptr(target), uintptr(numTextureNames), uintptr(unsafe.Pointer(textureNames)), 0, 0)
	return (uintptr)(ret)
}
func VDPAUSurfaceAccessNV(surface uintptr, access uint32) {
	syscall.Syscall(gpVDPAUSurfaceAccessNV, 2, uintptr(surface), uintptr(access), 0)
}
func VDPAUUnmapSurfacesNV(numSurface int32, surfaces *uintptr) {
	syscall.Syscall(gpVDPAUUnmapSurfacesNV, 2, uintptr(numSurface), uintptr(unsafe.Pointer(surfaces)), 0)
}
func VDPAUUnregisterSurfaceNV(surface uintptr) {
	syscall.Syscall(gpVDPAUUnregisterSurfaceNV, 1, uintptr(surface), 0, 0)
}

// Validates a program object
func ValidateProgram(program uint32) {
	syscall.Syscall(gpValidateProgram, 1, uintptr(program), 0, 0)
}
func ValidateProgramARB(programObj uintptr) {
	syscall.Syscall(gpValidateProgramARB, 1, uintptr(programObj), 0, 0)
}

// validate a program pipeline object against current GL state
func ValidateProgramPipeline(pipeline uint32) {
	syscall.Syscall(gpValidateProgramPipeline, 1, uintptr(pipeline), 0, 0)
}
func ValidateProgramPipelineEXT(pipeline uint32) {
	syscall.Syscall(gpValidateProgramPipelineEXT, 1, uintptr(pipeline), 0, 0)
}
func VariantArrayObjectATI(id uint32, xtype uint32, stride int32, buffer uint32, offset uint32) {
	syscall.Syscall6(gpVariantArrayObjectATI, 5, uintptr(id), uintptr(xtype), uintptr(stride), uintptr(buffer), uintptr(offset), 0)
}
func VariantPointerEXT(id uint32, xtype uint32, stride uint32, addr unsafe.Pointer) {
	syscall.Syscall6(gpVariantPointerEXT, 4, uintptr(id), uintptr(xtype), uintptr(stride), uintptr(addr), 0, 0)
}
func VariantbvEXT(id uint32, addr *int8) {
	syscall.Syscall(gpVariantbvEXT, 2, uintptr(id), uintptr(unsafe.Pointer(addr)), 0)
}
func VariantdvEXT(id uint32, addr *float64) {
	syscall.Syscall(gpVariantdvEXT, 2, uintptr(id), uintptr(unsafe.Pointer(addr)), 0)
}
func VariantfvEXT(id uint32, addr *float32) {
	syscall.Syscall(gpVariantfvEXT, 2, uintptr(id), uintptr(unsafe.Pointer(addr)), 0)
}
func VariantivEXT(id uint32, addr *int32) {
	syscall.Syscall(gpVariantivEXT, 2, uintptr(id), uintptr(unsafe.Pointer(addr)), 0)
}
func VariantsvEXT(id uint32, addr *int16) {
	syscall.Syscall(gpVariantsvEXT, 2, uintptr(id), uintptr(unsafe.Pointer(addr)), 0)
}
func VariantubvEXT(id uint32, addr *uint8) {
	syscall.Syscall(gpVariantubvEXT, 2, uintptr(id), uintptr(unsafe.Pointer(addr)), 0)
}
func VariantuivEXT(id uint32, addr *uint32) {
	syscall.Syscall(gpVariantuivEXT, 2, uintptr(id), uintptr(unsafe.Pointer(addr)), 0)
}
func VariantusvEXT(id uint32, addr *uint16) {
	syscall.Syscall(gpVariantusvEXT, 2, uintptr(id), uintptr(unsafe.Pointer(addr)), 0)
}
func Vertex2bOES(x int8, y int8) {
	syscall.Syscall(gpVertex2bOES, 2, uintptr(x), uintptr(y), 0)
}
func Vertex2bvOES(coords *int8) {
	syscall.Syscall(gpVertex2bvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func Vertex2d(x float64, y float64) {
	syscall.Syscall(gpVertex2d, 2, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), 0)
}
func Vertex2dv(v *float64) {
	syscall.Syscall(gpVertex2dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex2f(x float32, y float32) {
	syscall.Syscall(gpVertex2f, 2, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), 0)
}
func Vertex2fv(v *float32) {
	syscall.Syscall(gpVertex2fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex2hNV(x uint16, y uint16) {
	syscall.Syscall(gpVertex2hNV, 2, uintptr(x), uintptr(y), 0)
}
func Vertex2hvNV(v *uint16) {
	syscall.Syscall(gpVertex2hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex2i(x int32, y int32) {
	syscall.Syscall(gpVertex2i, 2, uintptr(x), uintptr(y), 0)
}
func Vertex2iv(v *int32) {
	syscall.Syscall(gpVertex2iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex2s(x int16, y int16) {
	syscall.Syscall(gpVertex2s, 2, uintptr(x), uintptr(y), 0)
}
func Vertex2sv(v *int16) {
	syscall.Syscall(gpVertex2sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex2xOES(x int32) {
	syscall.Syscall(gpVertex2xOES, 1, uintptr(x), 0, 0)
}
func Vertex2xvOES(coords *int32) {
	syscall.Syscall(gpVertex2xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func Vertex3bOES(x int8, y int8, z int8) {
	syscall.Syscall(gpVertex3bOES, 3, uintptr(x), uintptr(y), uintptr(z))
}
func Vertex3bvOES(coords *int8) {
	syscall.Syscall(gpVertex3bvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func Vertex3d(x float64, y float64, z float64) {
	syscall.Syscall(gpVertex3d, 3, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)))
}
func Vertex3dv(v *float64) {
	syscall.Syscall(gpVertex3dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex3f(x float32, y float32, z float32) {
	syscall.Syscall(gpVertex3f, 3, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func Vertex3fv(v *float32) {
	syscall.Syscall(gpVertex3fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex3hNV(x uint16, y uint16, z uint16) {
	syscall.Syscall(gpVertex3hNV, 3, uintptr(x), uintptr(y), uintptr(z))
}
func Vertex3hvNV(v *uint16) {
	syscall.Syscall(gpVertex3hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex3i(x int32, y int32, z int32) {
	syscall.Syscall(gpVertex3i, 3, uintptr(x), uintptr(y), uintptr(z))
}
func Vertex3iv(v *int32) {
	syscall.Syscall(gpVertex3iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex3s(x int16, y int16, z int16) {
	syscall.Syscall(gpVertex3s, 3, uintptr(x), uintptr(y), uintptr(z))
}
func Vertex3sv(v *int16) {
	syscall.Syscall(gpVertex3sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex3xOES(x int32, y int32) {
	syscall.Syscall(gpVertex3xOES, 2, uintptr(x), uintptr(y), 0)
}
func Vertex3xvOES(coords *int32) {
	syscall.Syscall(gpVertex3xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func Vertex4bOES(x int8, y int8, z int8, w int8) {
	syscall.Syscall6(gpVertex4bOES, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func Vertex4bvOES(coords *int8) {
	syscall.Syscall(gpVertex4bvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func Vertex4d(x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpVertex4d, 4, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0, 0)
}
func Vertex4dv(v *float64) {
	syscall.Syscall(gpVertex4dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex4f(x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpVertex4f, 4, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0, 0)
}
func Vertex4fv(v *float32) {
	syscall.Syscall(gpVertex4fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex4hNV(x uint16, y uint16, z uint16, w uint16) {
	syscall.Syscall6(gpVertex4hNV, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func Vertex4hvNV(v *uint16) {
	syscall.Syscall(gpVertex4hvNV, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex4i(x int32, y int32, z int32, w int32) {
	syscall.Syscall6(gpVertex4i, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func Vertex4iv(v *int32) {
	syscall.Syscall(gpVertex4iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex4s(x int16, y int16, z int16, w int16) {
	syscall.Syscall6(gpVertex4s, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func Vertex4sv(v *int16) {
	syscall.Syscall(gpVertex4sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func Vertex4xOES(x int32, y int32, z int32) {
	syscall.Syscall(gpVertex4xOES, 3, uintptr(x), uintptr(y), uintptr(z))
}
func Vertex4xvOES(coords *int32) {
	syscall.Syscall(gpVertex4xvOES, 1, uintptr(unsafe.Pointer(coords)), 0, 0)
}
func VertexArrayAttribBinding(vaobj uint32, attribindex uint32, bindingindex uint32) {
	syscall.Syscall(gpVertexArrayAttribBinding, 3, uintptr(vaobj), uintptr(attribindex), uintptr(bindingindex))
}

// specify the organization of vertex arrays
func VertexArrayAttribFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32) {
	syscall.Syscall6(gpVertexArrayAttribFormat, 6, uintptr(vaobj), uintptr(attribindex), uintptr(size), uintptr(xtype), boolToUintptr(normalized), uintptr(relativeoffset))
}
func VertexArrayAttribIFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	syscall.Syscall6(gpVertexArrayAttribIFormat, 5, uintptr(vaobj), uintptr(attribindex), uintptr(size), uintptr(xtype), uintptr(relativeoffset), 0)
}
func VertexArrayAttribLFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	syscall.Syscall6(gpVertexArrayAttribLFormat, 5, uintptr(vaobj), uintptr(attribindex), uintptr(size), uintptr(xtype), uintptr(relativeoffset), 0)
}
func VertexArrayBindVertexBufferEXT(vaobj uint32, bindingindex uint32, buffer uint32, offset int, stride int32) {
	syscall.Syscall6(gpVertexArrayBindVertexBufferEXT, 5, uintptr(vaobj), uintptr(bindingindex), uintptr(buffer), uintptr(offset), uintptr(stride), 0)
}

// modify the rate at which generic vertex attributes     advance
func VertexArrayBindingDivisor(vaobj uint32, bindingindex uint32, divisor uint32) {
	syscall.Syscall(gpVertexArrayBindingDivisor, 3, uintptr(vaobj), uintptr(bindingindex), uintptr(divisor))
}
func VertexArrayColorOffsetEXT(vaobj uint32, buffer uint32, size int32, xtype uint32, stride int32, offset int) {
	syscall.Syscall6(gpVertexArrayColorOffsetEXT, 6, uintptr(vaobj), uintptr(buffer), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(offset))
}
func VertexArrayEdgeFlagOffsetEXT(vaobj uint32, buffer uint32, stride int32, offset int) {
	syscall.Syscall6(gpVertexArrayEdgeFlagOffsetEXT, 4, uintptr(vaobj), uintptr(buffer), uintptr(stride), uintptr(offset), 0, 0)
}

// configures element array buffer binding of a vertex array object
func VertexArrayElementBuffer(vaobj uint32, buffer uint32) {
	syscall.Syscall(gpVertexArrayElementBuffer, 2, uintptr(vaobj), uintptr(buffer), 0)
}
func VertexArrayFogCoordOffsetEXT(vaobj uint32, buffer uint32, xtype uint32, stride int32, offset int) {
	syscall.Syscall6(gpVertexArrayFogCoordOffsetEXT, 5, uintptr(vaobj), uintptr(buffer), uintptr(xtype), uintptr(stride), uintptr(offset), 0)
}
func VertexArrayIndexOffsetEXT(vaobj uint32, buffer uint32, xtype uint32, stride int32, offset int) {
	syscall.Syscall6(gpVertexArrayIndexOffsetEXT, 5, uintptr(vaobj), uintptr(buffer), uintptr(xtype), uintptr(stride), uintptr(offset), 0)
}
func VertexArrayMultiTexCoordOffsetEXT(vaobj uint32, buffer uint32, texunit uint32, size int32, xtype uint32, stride int32, offset int) {
	syscall.Syscall9(gpVertexArrayMultiTexCoordOffsetEXT, 7, uintptr(vaobj), uintptr(buffer), uintptr(texunit), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(offset), 0, 0)
}
func VertexArrayNormalOffsetEXT(vaobj uint32, buffer uint32, xtype uint32, stride int32, offset int) {
	syscall.Syscall6(gpVertexArrayNormalOffsetEXT, 5, uintptr(vaobj), uintptr(buffer), uintptr(xtype), uintptr(stride), uintptr(offset), 0)
}
func VertexArrayParameteriAPPLE(pname uint32, param int32) {
	syscall.Syscall(gpVertexArrayParameteriAPPLE, 2, uintptr(pname), uintptr(param), 0)
}
func VertexArrayRangeAPPLE(length int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpVertexArrayRangeAPPLE, 2, uintptr(length), uintptr(pointer), 0)
}
func VertexArrayRangeNV(length int32, pointer unsafe.Pointer) {
	syscall.Syscall(gpVertexArrayRangeNV, 2, uintptr(length), uintptr(pointer), 0)
}
func VertexArraySecondaryColorOffsetEXT(vaobj uint32, buffer uint32, size int32, xtype uint32, stride int32, offset int) {
	syscall.Syscall6(gpVertexArraySecondaryColorOffsetEXT, 6, uintptr(vaobj), uintptr(buffer), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(offset))
}
func VertexArrayTexCoordOffsetEXT(vaobj uint32, buffer uint32, size int32, xtype uint32, stride int32, offset int) {
	syscall.Syscall6(gpVertexArrayTexCoordOffsetEXT, 6, uintptr(vaobj), uintptr(buffer), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(offset))
}
func VertexArrayVertexAttribBindingEXT(vaobj uint32, attribindex uint32, bindingindex uint32) {
	syscall.Syscall(gpVertexArrayVertexAttribBindingEXT, 3, uintptr(vaobj), uintptr(attribindex), uintptr(bindingindex))
}
func VertexArrayVertexAttribDivisorEXT(vaobj uint32, index uint32, divisor uint32) {
	syscall.Syscall(gpVertexArrayVertexAttribDivisorEXT, 3, uintptr(vaobj), uintptr(index), uintptr(divisor))
}
func VertexArrayVertexAttribFormatEXT(vaobj uint32, attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32) {
	syscall.Syscall6(gpVertexArrayVertexAttribFormatEXT, 6, uintptr(vaobj), uintptr(attribindex), uintptr(size), uintptr(xtype), boolToUintptr(normalized), uintptr(relativeoffset))
}
func VertexArrayVertexAttribIFormatEXT(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	syscall.Syscall6(gpVertexArrayVertexAttribIFormatEXT, 5, uintptr(vaobj), uintptr(attribindex), uintptr(size), uintptr(xtype), uintptr(relativeoffset), 0)
}
func VertexArrayVertexAttribIOffsetEXT(vaobj uint32, buffer uint32, index uint32, size int32, xtype uint32, stride int32, offset int) {
	syscall.Syscall9(gpVertexArrayVertexAttribIOffsetEXT, 7, uintptr(vaobj), uintptr(buffer), uintptr(index), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(offset), 0, 0)
}
func VertexArrayVertexAttribLFormatEXT(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	syscall.Syscall6(gpVertexArrayVertexAttribLFormatEXT, 5, uintptr(vaobj), uintptr(attribindex), uintptr(size), uintptr(xtype), uintptr(relativeoffset), 0)
}
func VertexArrayVertexAttribLOffsetEXT(vaobj uint32, buffer uint32, index uint32, size int32, xtype uint32, stride int32, offset int) {
	syscall.Syscall9(gpVertexArrayVertexAttribLOffsetEXT, 7, uintptr(vaobj), uintptr(buffer), uintptr(index), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(offset), 0, 0)
}
func VertexArrayVertexAttribOffsetEXT(vaobj uint32, buffer uint32, index uint32, size int32, xtype uint32, normalized bool, stride int32, offset int) {
	syscall.Syscall9(gpVertexArrayVertexAttribOffsetEXT, 8, uintptr(vaobj), uintptr(buffer), uintptr(index), uintptr(size), uintptr(xtype), boolToUintptr(normalized), uintptr(stride), uintptr(offset), 0)
}
func VertexArrayVertexBindingDivisorEXT(vaobj uint32, bindingindex uint32, divisor uint32) {
	syscall.Syscall(gpVertexArrayVertexBindingDivisorEXT, 3, uintptr(vaobj), uintptr(bindingindex), uintptr(divisor))
}

// bind a buffer to a vertex buffer bind point
func VertexArrayVertexBuffer(vaobj uint32, bindingindex uint32, buffer uint32, offset int, stride int32) {
	syscall.Syscall6(gpVertexArrayVertexBuffer, 5, uintptr(vaobj), uintptr(bindingindex), uintptr(buffer), uintptr(offset), uintptr(stride), 0)
}

// attach multiple buffer objects to a vertex array object
func VertexArrayVertexBuffers(vaobj uint32, first uint32, count int32, buffers *uint32, offsets *int, strides *int32) {
	syscall.Syscall6(gpVertexArrayVertexBuffers, 6, uintptr(vaobj), uintptr(first), uintptr(count), uintptr(unsafe.Pointer(buffers)), uintptr(unsafe.Pointer(offsets)), uintptr(unsafe.Pointer(strides)))
}
func VertexArrayVertexOffsetEXT(vaobj uint32, buffer uint32, size int32, xtype uint32, stride int32, offset int) {
	syscall.Syscall6(gpVertexArrayVertexOffsetEXT, 6, uintptr(vaobj), uintptr(buffer), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(offset))
}
func VertexAttrib1d(index uint32, x float64) {
	syscall.Syscall(gpVertexAttrib1d, 2, uintptr(index), uintptr(math.Float64bits(x)), 0)
}
func VertexAttrib1dARB(index uint32, x float64) {
	syscall.Syscall(gpVertexAttrib1dARB, 2, uintptr(index), uintptr(math.Float64bits(x)), 0)
}
func VertexAttrib1dNV(index uint32, x float64) {
	syscall.Syscall(gpVertexAttrib1dNV, 2, uintptr(index), uintptr(math.Float64bits(x)), 0)
}
func VertexAttrib1dv(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib1dv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1dvARB(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib1dvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1dvNV(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib1dvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1f(index uint32, x float32) {
	syscall.Syscall(gpVertexAttrib1f, 2, uintptr(index), uintptr(math.Float32bits(x)), 0)
}
func VertexAttrib1fARB(index uint32, x float32) {
	syscall.Syscall(gpVertexAttrib1fARB, 2, uintptr(index), uintptr(math.Float32bits(x)), 0)
}
func VertexAttrib1fNV(index uint32, x float32) {
	syscall.Syscall(gpVertexAttrib1fNV, 2, uintptr(index), uintptr(math.Float32bits(x)), 0)
}
func VertexAttrib1fv(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib1fv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1fvARB(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib1fvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1fvNV(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib1fvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1hNV(index uint32, x uint16) {
	syscall.Syscall(gpVertexAttrib1hNV, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttrib1hvNV(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttrib1hvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1s(index uint32, x int16) {
	syscall.Syscall(gpVertexAttrib1s, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttrib1sARB(index uint32, x int16) {
	syscall.Syscall(gpVertexAttrib1sARB, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttrib1sNV(index uint32, x int16) {
	syscall.Syscall(gpVertexAttrib1sNV, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttrib1sv(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib1sv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1svARB(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib1svARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib1svNV(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib1svNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2d(index uint32, x float64, y float64) {
	syscall.Syscall(gpVertexAttrib2d, 3, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
}
func VertexAttrib2dARB(index uint32, x float64, y float64) {
	syscall.Syscall(gpVertexAttrib2dARB, 3, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
}
func VertexAttrib2dNV(index uint32, x float64, y float64) {
	syscall.Syscall(gpVertexAttrib2dNV, 3, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
}
func VertexAttrib2dv(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib2dv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2dvARB(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib2dvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2dvNV(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib2dvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2f(index uint32, x float32, y float32) {
	syscall.Syscall(gpVertexAttrib2f, 3, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)))
}
func VertexAttrib2fARB(index uint32, x float32, y float32) {
	syscall.Syscall(gpVertexAttrib2fARB, 3, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)))
}
func VertexAttrib2fNV(index uint32, x float32, y float32) {
	syscall.Syscall(gpVertexAttrib2fNV, 3, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)))
}
func VertexAttrib2fv(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib2fv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2fvARB(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib2fvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2fvNV(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib2fvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2hNV(index uint32, x uint16, y uint16) {
	syscall.Syscall(gpVertexAttrib2hNV, 3, uintptr(index), uintptr(x), uintptr(y))
}
func VertexAttrib2hvNV(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttrib2hvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2s(index uint32, x int16, y int16) {
	syscall.Syscall(gpVertexAttrib2s, 3, uintptr(index), uintptr(x), uintptr(y))
}
func VertexAttrib2sARB(index uint32, x int16, y int16) {
	syscall.Syscall(gpVertexAttrib2sARB, 3, uintptr(index), uintptr(x), uintptr(y))
}
func VertexAttrib2sNV(index uint32, x int16, y int16) {
	syscall.Syscall(gpVertexAttrib2sNV, 3, uintptr(index), uintptr(x), uintptr(y))
}
func VertexAttrib2sv(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib2sv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2svARB(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib2svARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib2svNV(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib2svNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3d(index uint32, x float64, y float64, z float64) {
	syscall.Syscall6(gpVertexAttrib3d, 4, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func VertexAttrib3dARB(index uint32, x float64, y float64, z float64) {
	syscall.Syscall6(gpVertexAttrib3dARB, 4, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func VertexAttrib3dNV(index uint32, x float64, y float64, z float64) {
	syscall.Syscall6(gpVertexAttrib3dNV, 4, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func VertexAttrib3dv(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib3dv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3dvARB(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib3dvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3dvNV(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib3dvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3f(index uint32, x float32, y float32, z float32) {
	syscall.Syscall6(gpVertexAttrib3f, 4, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func VertexAttrib3fARB(index uint32, x float32, y float32, z float32) {
	syscall.Syscall6(gpVertexAttrib3fARB, 4, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func VertexAttrib3fNV(index uint32, x float32, y float32, z float32) {
	syscall.Syscall6(gpVertexAttrib3fNV, 4, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func VertexAttrib3fv(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib3fv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3fvARB(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib3fvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3fvNV(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib3fvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3hNV(index uint32, x uint16, y uint16, z uint16) {
	syscall.Syscall6(gpVertexAttrib3hNV, 4, uintptr(index), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexAttrib3hvNV(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttrib3hvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3s(index uint32, x int16, y int16, z int16) {
	syscall.Syscall6(gpVertexAttrib3s, 4, uintptr(index), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexAttrib3sARB(index uint32, x int16, y int16, z int16) {
	syscall.Syscall6(gpVertexAttrib3sARB, 4, uintptr(index), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexAttrib3sNV(index uint32, x int16, y int16, z int16) {
	syscall.Syscall6(gpVertexAttrib3sNV, 4, uintptr(index), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexAttrib3sv(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib3sv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3svARB(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib3svARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib3svNV(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib3svNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4Nbv(index uint32, v *int8) {
	syscall.Syscall(gpVertexAttrib4Nbv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4NbvARB(index uint32, v *int8) {
	syscall.Syscall(gpVertexAttrib4NbvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4Niv(index uint32, v *int32) {
	syscall.Syscall(gpVertexAttrib4Niv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4NivARB(index uint32, v *int32) {
	syscall.Syscall(gpVertexAttrib4NivARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4Nsv(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib4Nsv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4NsvARB(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib4NsvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4Nub(index uint32, x uint8, y uint8, z uint8, w uint8) {
	syscall.Syscall6(gpVertexAttrib4Nub, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttrib4NubARB(index uint32, x uint8, y uint8, z uint8, w uint8) {
	syscall.Syscall6(gpVertexAttrib4NubARB, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttrib4Nubv(index uint32, v *uint8) {
	syscall.Syscall(gpVertexAttrib4Nubv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4NubvARB(index uint32, v *uint8) {
	syscall.Syscall(gpVertexAttrib4NubvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4Nuiv(index uint32, v *uint32) {
	syscall.Syscall(gpVertexAttrib4Nuiv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4NuivARB(index uint32, v *uint32) {
	syscall.Syscall(gpVertexAttrib4NuivARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4Nusv(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttrib4Nusv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4NusvARB(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttrib4NusvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4bv(index uint32, v *int8) {
	syscall.Syscall(gpVertexAttrib4bv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4bvARB(index uint32, v *int8) {
	syscall.Syscall(gpVertexAttrib4bvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4d(index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpVertexAttrib4d, 5, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0)
}
func VertexAttrib4dARB(index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpVertexAttrib4dARB, 5, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0)
}
func VertexAttrib4dNV(index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpVertexAttrib4dNV, 5, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0)
}
func VertexAttrib4dv(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib4dv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4dvARB(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib4dvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4dvNV(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttrib4dvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4f(index uint32, x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpVertexAttrib4f, 5, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0)
}
func VertexAttrib4fARB(index uint32, x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpVertexAttrib4fARB, 5, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0)
}
func VertexAttrib4fNV(index uint32, x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpVertexAttrib4fNV, 5, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0)
}
func VertexAttrib4fv(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib4fv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4fvARB(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib4fvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4fvNV(index uint32, v *float32) {
	syscall.Syscall(gpVertexAttrib4fvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4hNV(index uint32, x uint16, y uint16, z uint16, w uint16) {
	syscall.Syscall6(gpVertexAttrib4hNV, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttrib4hvNV(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttrib4hvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4iv(index uint32, v *int32) {
	syscall.Syscall(gpVertexAttrib4iv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4ivARB(index uint32, v *int32) {
	syscall.Syscall(gpVertexAttrib4ivARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4s(index uint32, x int16, y int16, z int16, w int16) {
	syscall.Syscall6(gpVertexAttrib4s, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttrib4sARB(index uint32, x int16, y int16, z int16, w int16) {
	syscall.Syscall6(gpVertexAttrib4sARB, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttrib4sNV(index uint32, x int16, y int16, z int16, w int16) {
	syscall.Syscall6(gpVertexAttrib4sNV, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttrib4sv(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib4sv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4svARB(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib4svARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4svNV(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttrib4svNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4ubNV(index uint32, x uint8, y uint8, z uint8, w uint8) {
	syscall.Syscall6(gpVertexAttrib4ubNV, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttrib4ubv(index uint32, v *uint8) {
	syscall.Syscall(gpVertexAttrib4ubv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4ubvARB(index uint32, v *uint8) {
	syscall.Syscall(gpVertexAttrib4ubvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4ubvNV(index uint32, v *uint8) {
	syscall.Syscall(gpVertexAttrib4ubvNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4uiv(index uint32, v *uint32) {
	syscall.Syscall(gpVertexAttrib4uiv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4uivARB(index uint32, v *uint32) {
	syscall.Syscall(gpVertexAttrib4uivARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4usv(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttrib4usv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttrib4usvARB(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttrib4usvARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribArrayObjectATI(index uint32, size int32, xtype uint32, normalized bool, stride int32, buffer uint32, offset uint32) {
	syscall.Syscall9(gpVertexAttribArrayObjectATI, 7, uintptr(index), uintptr(size), uintptr(xtype), boolToUintptr(normalized), uintptr(stride), uintptr(buffer), uintptr(offset), 0, 0)
}

// associate a vertex attribute and a vertex buffer binding for a vertex array object
func VertexAttribBinding(attribindex uint32, bindingindex uint32) {
	syscall.Syscall(gpVertexAttribBinding, 2, uintptr(attribindex), uintptr(bindingindex), 0)
}
func VertexAttribDivisorARB(index uint32, divisor uint32) {
	syscall.Syscall(gpVertexAttribDivisorARB, 2, uintptr(index), uintptr(divisor), 0)
}

// specify the organization of vertex arrays
func VertexAttribFormat(attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32) {
	syscall.Syscall6(gpVertexAttribFormat, 5, uintptr(attribindex), uintptr(size), uintptr(xtype), boolToUintptr(normalized), uintptr(relativeoffset), 0)
}
func VertexAttribFormatNV(index uint32, size int32, xtype uint32, normalized bool, stride int32) {
	syscall.Syscall6(gpVertexAttribFormatNV, 5, uintptr(index), uintptr(size), uintptr(xtype), boolToUintptr(normalized), uintptr(stride), 0)
}
func VertexAttribI1iEXT(index uint32, x int32) {
	syscall.Syscall(gpVertexAttribI1iEXT, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttribI1ivEXT(index uint32, v *int32) {
	syscall.Syscall(gpVertexAttribI1ivEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI1uiEXT(index uint32, x uint32) {
	syscall.Syscall(gpVertexAttribI1uiEXT, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttribI1uivEXT(index uint32, v *uint32) {
	syscall.Syscall(gpVertexAttribI1uivEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI2iEXT(index uint32, x int32, y int32) {
	syscall.Syscall(gpVertexAttribI2iEXT, 3, uintptr(index), uintptr(x), uintptr(y))
}
func VertexAttribI2ivEXT(index uint32, v *int32) {
	syscall.Syscall(gpVertexAttribI2ivEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI2uiEXT(index uint32, x uint32, y uint32) {
	syscall.Syscall(gpVertexAttribI2uiEXT, 3, uintptr(index), uintptr(x), uintptr(y))
}
func VertexAttribI2uivEXT(index uint32, v *uint32) {
	syscall.Syscall(gpVertexAttribI2uivEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI3iEXT(index uint32, x int32, y int32, z int32) {
	syscall.Syscall6(gpVertexAttribI3iEXT, 4, uintptr(index), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexAttribI3ivEXT(index uint32, v *int32) {
	syscall.Syscall(gpVertexAttribI3ivEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI3uiEXT(index uint32, x uint32, y uint32, z uint32) {
	syscall.Syscall6(gpVertexAttribI3uiEXT, 4, uintptr(index), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexAttribI3uivEXT(index uint32, v *uint32) {
	syscall.Syscall(gpVertexAttribI3uivEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI4bvEXT(index uint32, v *int8) {
	syscall.Syscall(gpVertexAttribI4bvEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI4iEXT(index uint32, x int32, y int32, z int32, w int32) {
	syscall.Syscall6(gpVertexAttribI4iEXT, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttribI4ivEXT(index uint32, v *int32) {
	syscall.Syscall(gpVertexAttribI4ivEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI4svEXT(index uint32, v *int16) {
	syscall.Syscall(gpVertexAttribI4svEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI4ubvEXT(index uint32, v *uint8) {
	syscall.Syscall(gpVertexAttribI4ubvEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI4uiEXT(index uint32, x uint32, y uint32, z uint32, w uint32) {
	syscall.Syscall6(gpVertexAttribI4uiEXT, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttribI4uivEXT(index uint32, v *uint32) {
	syscall.Syscall(gpVertexAttribI4uivEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribI4usvEXT(index uint32, v *uint16) {
	syscall.Syscall(gpVertexAttribI4usvEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribIFormat(attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	syscall.Syscall6(gpVertexAttribIFormat, 4, uintptr(attribindex), uintptr(size), uintptr(xtype), uintptr(relativeoffset), 0, 0)
}
func VertexAttribIFormatNV(index uint32, size int32, xtype uint32, stride int32) {
	syscall.Syscall6(gpVertexAttribIFormatNV, 4, uintptr(index), uintptr(size), uintptr(xtype), uintptr(stride), 0, 0)
}
func VertexAttribIPointerEXT(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpVertexAttribIPointerEXT, 5, uintptr(index), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0)
}
func VertexAttribL1d(index uint32, x float64) {
	syscall.Syscall(gpVertexAttribL1d, 2, uintptr(index), uintptr(math.Float64bits(x)), 0)
}
func VertexAttribL1dEXT(index uint32, x float64) {
	syscall.Syscall(gpVertexAttribL1dEXT, 2, uintptr(index), uintptr(math.Float64bits(x)), 0)
}
func VertexAttribL1dv(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttribL1dv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL1dvEXT(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttribL1dvEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL1i64NV(index uint32, x int64) {
	syscall.Syscall(gpVertexAttribL1i64NV, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttribL1i64vNV(index uint32, v *int64) {
	syscall.Syscall(gpVertexAttribL1i64vNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL1ui64ARB(index uint32, x uint64) {
	syscall.Syscall(gpVertexAttribL1ui64ARB, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttribL1ui64NV(index uint32, x uint64) {
	syscall.Syscall(gpVertexAttribL1ui64NV, 2, uintptr(index), uintptr(x), 0)
}
func VertexAttribL1ui64vARB(index uint32, v *uint64) {
	syscall.Syscall(gpVertexAttribL1ui64vARB, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL1ui64vNV(index uint32, v *uint64) {
	syscall.Syscall(gpVertexAttribL1ui64vNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL2d(index uint32, x float64, y float64) {
	syscall.Syscall(gpVertexAttribL2d, 3, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
}
func VertexAttribL2dEXT(index uint32, x float64, y float64) {
	syscall.Syscall(gpVertexAttribL2dEXT, 3, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
}
func VertexAttribL2dv(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttribL2dv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL2dvEXT(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttribL2dvEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL2i64NV(index uint32, x int64, y int64) {
	syscall.Syscall(gpVertexAttribL2i64NV, 3, uintptr(index), uintptr(x), uintptr(y))
}
func VertexAttribL2i64vNV(index uint32, v *int64) {
	syscall.Syscall(gpVertexAttribL2i64vNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL2ui64NV(index uint32, x uint64, y uint64) {
	syscall.Syscall(gpVertexAttribL2ui64NV, 3, uintptr(index), uintptr(x), uintptr(y))
}
func VertexAttribL2ui64vNV(index uint32, v *uint64) {
	syscall.Syscall(gpVertexAttribL2ui64vNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL3d(index uint32, x float64, y float64, z float64) {
	syscall.Syscall6(gpVertexAttribL3d, 4, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func VertexAttribL3dEXT(index uint32, x float64, y float64, z float64) {
	syscall.Syscall6(gpVertexAttribL3dEXT, 4, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func VertexAttribL3dv(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttribL3dv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL3dvEXT(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttribL3dvEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL3i64NV(index uint32, x int64, y int64, z int64) {
	syscall.Syscall6(gpVertexAttribL3i64NV, 4, uintptr(index), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexAttribL3i64vNV(index uint32, v *int64) {
	syscall.Syscall(gpVertexAttribL3i64vNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL3ui64NV(index uint32, x uint64, y uint64, z uint64) {
	syscall.Syscall6(gpVertexAttribL3ui64NV, 4, uintptr(index), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexAttribL3ui64vNV(index uint32, v *uint64) {
	syscall.Syscall(gpVertexAttribL3ui64vNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL4d(index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpVertexAttribL4d, 5, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0)
}
func VertexAttribL4dEXT(index uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpVertexAttribL4dEXT, 5, uintptr(index), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0)
}
func VertexAttribL4dv(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttribL4dv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL4dvEXT(index uint32, v *float64) {
	syscall.Syscall(gpVertexAttribL4dvEXT, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL4i64NV(index uint32, x int64, y int64, z int64, w int64) {
	syscall.Syscall6(gpVertexAttribL4i64NV, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttribL4i64vNV(index uint32, v *int64) {
	syscall.Syscall(gpVertexAttribL4i64vNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribL4ui64NV(index uint32, x uint64, y uint64, z uint64, w uint64) {
	syscall.Syscall6(gpVertexAttribL4ui64NV, 5, uintptr(index), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexAttribL4ui64vNV(index uint32, v *uint64) {
	syscall.Syscall(gpVertexAttribL4ui64vNV, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func VertexAttribLFormat(attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	syscall.Syscall6(gpVertexAttribLFormat, 4, uintptr(attribindex), uintptr(size), uintptr(xtype), uintptr(relativeoffset), 0, 0)
}
func VertexAttribLFormatNV(index uint32, size int32, xtype uint32, stride int32) {
	syscall.Syscall6(gpVertexAttribLFormatNV, 4, uintptr(index), uintptr(size), uintptr(xtype), uintptr(stride), 0, 0)
}
func VertexAttribLPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpVertexAttribLPointer, 5, uintptr(index), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0)
}
func VertexAttribLPointerEXT(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpVertexAttribLPointerEXT, 5, uintptr(index), uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0)
}
func VertexAttribP1ui(index uint32, xtype uint32, normalized bool, value uint32) {
	syscall.Syscall6(gpVertexAttribP1ui, 4, uintptr(index), uintptr(xtype), boolToUintptr(normalized), uintptr(value), 0, 0)
}
func VertexAttribP1uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	syscall.Syscall6(gpVertexAttribP1uiv, 4, uintptr(index), uintptr(xtype), boolToUintptr(normalized), uintptr(unsafe.Pointer(value)), 0, 0)
}
func VertexAttribP2ui(index uint32, xtype uint32, normalized bool, value uint32) {
	syscall.Syscall6(gpVertexAttribP2ui, 4, uintptr(index), uintptr(xtype), boolToUintptr(normalized), uintptr(value), 0, 0)
}
func VertexAttribP2uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	syscall.Syscall6(gpVertexAttribP2uiv, 4, uintptr(index), uintptr(xtype), boolToUintptr(normalized), uintptr(unsafe.Pointer(value)), 0, 0)
}
func VertexAttribP3ui(index uint32, xtype uint32, normalized bool, value uint32) {
	syscall.Syscall6(gpVertexAttribP3ui, 4, uintptr(index), uintptr(xtype), boolToUintptr(normalized), uintptr(value), 0, 0)
}
func VertexAttribP3uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	syscall.Syscall6(gpVertexAttribP3uiv, 4, uintptr(index), uintptr(xtype), boolToUintptr(normalized), uintptr(unsafe.Pointer(value)), 0, 0)
}
func VertexAttribP4ui(index uint32, xtype uint32, normalized bool, value uint32) {
	syscall.Syscall6(gpVertexAttribP4ui, 4, uintptr(index), uintptr(xtype), boolToUintptr(normalized), uintptr(value), 0, 0)
}
func VertexAttribP4uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	syscall.Syscall6(gpVertexAttribP4uiv, 4, uintptr(index), uintptr(xtype), boolToUintptr(normalized), uintptr(unsafe.Pointer(value)), 0, 0)
}
func VertexAttribParameteriAMD(index uint32, pname uint32, param int32) {
	syscall.Syscall(gpVertexAttribParameteriAMD, 3, uintptr(index), uintptr(pname), uintptr(param))
}

// define an array of generic vertex attribute data
func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer uintptr) {
	syscall.Syscall6(gpVertexAttribPointer, 6, uintptr(index), uintptr(size), uintptr(xtype), boolToUintptr(normalized), uintptr(stride), uintptr(pointer))
}
func VertexAttribPointerARB(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpVertexAttribPointerARB, 6, uintptr(index), uintptr(size), uintptr(xtype), boolToUintptr(normalized), uintptr(stride), uintptr(pointer))
}
func VertexAttribPointerNV(index uint32, fsize int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpVertexAttribPointerNV, 5, uintptr(index), uintptr(fsize), uintptr(xtype), uintptr(stride), uintptr(pointer), 0)
}
func VertexAttribs1dvNV(index uint32, count int32, v *float64) {
	syscall.Syscall(gpVertexAttribs1dvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs1fvNV(index uint32, count int32, v *float32) {
	syscall.Syscall(gpVertexAttribs1fvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs1hvNV(index uint32, n int32, v *uint16) {
	syscall.Syscall(gpVertexAttribs1hvNV, 3, uintptr(index), uintptr(n), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs1svNV(index uint32, count int32, v *int16) {
	syscall.Syscall(gpVertexAttribs1svNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs2dvNV(index uint32, count int32, v *float64) {
	syscall.Syscall(gpVertexAttribs2dvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs2fvNV(index uint32, count int32, v *float32) {
	syscall.Syscall(gpVertexAttribs2fvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs2hvNV(index uint32, n int32, v *uint16) {
	syscall.Syscall(gpVertexAttribs2hvNV, 3, uintptr(index), uintptr(n), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs2svNV(index uint32, count int32, v *int16) {
	syscall.Syscall(gpVertexAttribs2svNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs3dvNV(index uint32, count int32, v *float64) {
	syscall.Syscall(gpVertexAttribs3dvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs3fvNV(index uint32, count int32, v *float32) {
	syscall.Syscall(gpVertexAttribs3fvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs3hvNV(index uint32, n int32, v *uint16) {
	syscall.Syscall(gpVertexAttribs3hvNV, 3, uintptr(index), uintptr(n), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs3svNV(index uint32, count int32, v *int16) {
	syscall.Syscall(gpVertexAttribs3svNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs4dvNV(index uint32, count int32, v *float64) {
	syscall.Syscall(gpVertexAttribs4dvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs4fvNV(index uint32, count int32, v *float32) {
	syscall.Syscall(gpVertexAttribs4fvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs4hvNV(index uint32, n int32, v *uint16) {
	syscall.Syscall(gpVertexAttribs4hvNV, 3, uintptr(index), uintptr(n), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs4svNV(index uint32, count int32, v *int16) {
	syscall.Syscall(gpVertexAttribs4svNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func VertexAttribs4ubvNV(index uint32, count int32, v *uint8) {
	syscall.Syscall(gpVertexAttribs4ubvNV, 3, uintptr(index), uintptr(count), uintptr(unsafe.Pointer(v)))
}

// modify the rate at which generic vertex attributes     advance
func VertexBindingDivisor(bindingindex uint32, divisor uint32) {
	syscall.Syscall(gpVertexBindingDivisor, 2, uintptr(bindingindex), uintptr(divisor), 0)
}
func VertexBlendARB(count int32) {
	syscall.Syscall(gpVertexBlendARB, 1, uintptr(count), 0, 0)
}
func VertexBlendEnvfATI(pname uint32, param float32) {
	syscall.Syscall(gpVertexBlendEnvfATI, 2, uintptr(pname), uintptr(math.Float32bits(param)), 0)
}
func VertexBlendEnviATI(pname uint32, param int32) {
	syscall.Syscall(gpVertexBlendEnviATI, 2, uintptr(pname), uintptr(param), 0)
}
func VertexFormatNV(size int32, xtype uint32, stride int32) {
	syscall.Syscall(gpVertexFormatNV, 3, uintptr(size), uintptr(xtype), uintptr(stride))
}

// define an array of vertex data
func VertexPointer(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpVertexPointer, 4, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0, 0)
}
func VertexPointerEXT(size int32, xtype uint32, stride int32, count int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpVertexPointerEXT, 5, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(count), uintptr(pointer), 0)
}
func VertexPointerListIBM(size int32, xtype uint32, stride int32, pointer *unsafe.Pointer, ptrstride int32) {
	syscall.Syscall6(gpVertexPointerListIBM, 5, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(unsafe.Pointer(pointer)), uintptr(ptrstride), 0)
}
func VertexPointervINTEL(size int32, xtype uint32, pointer *unsafe.Pointer) {
	syscall.Syscall(gpVertexPointervINTEL, 3, uintptr(size), uintptr(xtype), uintptr(unsafe.Pointer(pointer)))
}
func VertexStream1dATI(stream uint32, x float64) {
	syscall.Syscall(gpVertexStream1dATI, 2, uintptr(stream), uintptr(math.Float64bits(x)), 0)
}
func VertexStream1dvATI(stream uint32, coords *float64) {
	syscall.Syscall(gpVertexStream1dvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream1fATI(stream uint32, x float32) {
	syscall.Syscall(gpVertexStream1fATI, 2, uintptr(stream), uintptr(math.Float32bits(x)), 0)
}
func VertexStream1fvATI(stream uint32, coords *float32) {
	syscall.Syscall(gpVertexStream1fvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream1iATI(stream uint32, x int32) {
	syscall.Syscall(gpVertexStream1iATI, 2, uintptr(stream), uintptr(x), 0)
}
func VertexStream1ivATI(stream uint32, coords *int32) {
	syscall.Syscall(gpVertexStream1ivATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream1sATI(stream uint32, x int16) {
	syscall.Syscall(gpVertexStream1sATI, 2, uintptr(stream), uintptr(x), 0)
}
func VertexStream1svATI(stream uint32, coords *int16) {
	syscall.Syscall(gpVertexStream1svATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream2dATI(stream uint32, x float64, y float64) {
	syscall.Syscall(gpVertexStream2dATI, 3, uintptr(stream), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
}
func VertexStream2dvATI(stream uint32, coords *float64) {
	syscall.Syscall(gpVertexStream2dvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream2fATI(stream uint32, x float32, y float32) {
	syscall.Syscall(gpVertexStream2fATI, 3, uintptr(stream), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)))
}
func VertexStream2fvATI(stream uint32, coords *float32) {
	syscall.Syscall(gpVertexStream2fvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream2iATI(stream uint32, x int32, y int32) {
	syscall.Syscall(gpVertexStream2iATI, 3, uintptr(stream), uintptr(x), uintptr(y))
}
func VertexStream2ivATI(stream uint32, coords *int32) {
	syscall.Syscall(gpVertexStream2ivATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream2sATI(stream uint32, x int16, y int16) {
	syscall.Syscall(gpVertexStream2sATI, 3, uintptr(stream), uintptr(x), uintptr(y))
}
func VertexStream2svATI(stream uint32, coords *int16) {
	syscall.Syscall(gpVertexStream2svATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream3dATI(stream uint32, x float64, y float64, z float64) {
	syscall.Syscall6(gpVertexStream3dATI, 4, uintptr(stream), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), 0, 0)
}
func VertexStream3dvATI(stream uint32, coords *float64) {
	syscall.Syscall(gpVertexStream3dvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream3fATI(stream uint32, x float32, y float32, z float32) {
	syscall.Syscall6(gpVertexStream3fATI, 4, uintptr(stream), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), 0, 0)
}
func VertexStream3fvATI(stream uint32, coords *float32) {
	syscall.Syscall(gpVertexStream3fvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream3iATI(stream uint32, x int32, y int32, z int32) {
	syscall.Syscall6(gpVertexStream3iATI, 4, uintptr(stream), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexStream3ivATI(stream uint32, coords *int32) {
	syscall.Syscall(gpVertexStream3ivATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream3sATI(stream uint32, x int16, y int16, z int16) {
	syscall.Syscall6(gpVertexStream3sATI, 4, uintptr(stream), uintptr(x), uintptr(y), uintptr(z), 0, 0)
}
func VertexStream3svATI(stream uint32, coords *int16) {
	syscall.Syscall(gpVertexStream3svATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream4dATI(stream uint32, x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpVertexStream4dATI, 5, uintptr(stream), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0)
}
func VertexStream4dvATI(stream uint32, coords *float64) {
	syscall.Syscall(gpVertexStream4dvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream4fATI(stream uint32, x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpVertexStream4fATI, 5, uintptr(stream), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0)
}
func VertexStream4fvATI(stream uint32, coords *float32) {
	syscall.Syscall(gpVertexStream4fvATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream4iATI(stream uint32, x int32, y int32, z int32, w int32) {
	syscall.Syscall6(gpVertexStream4iATI, 5, uintptr(stream), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexStream4ivATI(stream uint32, coords *int32) {
	syscall.Syscall(gpVertexStream4ivATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexStream4sATI(stream uint32, x int16, y int16, z int16, w int16) {
	syscall.Syscall6(gpVertexStream4sATI, 5, uintptr(stream), uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0)
}
func VertexStream4svATI(stream uint32, coords *int16) {
	syscall.Syscall(gpVertexStream4svATI, 2, uintptr(stream), uintptr(unsafe.Pointer(coords)), 0)
}
func VertexWeightPointerEXT(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpVertexWeightPointerEXT, 4, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0, 0)
}
func VertexWeightfEXT(weight float32) {
	syscall.Syscall(gpVertexWeightfEXT, 1, uintptr(math.Float32bits(weight)), 0, 0)
}
func VertexWeightfvEXT(weight *float32) {
	syscall.Syscall(gpVertexWeightfvEXT, 1, uintptr(unsafe.Pointer(weight)), 0, 0)
}
func VertexWeighthNV(weight uint16) {
	syscall.Syscall(gpVertexWeighthNV, 1, uintptr(weight), 0, 0)
}
func VertexWeighthvNV(weight *uint16) {
	syscall.Syscall(gpVertexWeighthvNV, 1, uintptr(unsafe.Pointer(weight)), 0, 0)
}
func VideoCaptureNV(video_capture_slot uint32, sequence_num *uint32, capture_time *uint64) uint32 {
	ret, _, _ := syscall.Syscall(gpVideoCaptureNV, 3, uintptr(video_capture_slot), uintptr(unsafe.Pointer(sequence_num)), uintptr(unsafe.Pointer(capture_time)))
	return (uint32)(ret)
}
func VideoCaptureStreamParameterdvNV(video_capture_slot uint32, stream uint32, pname uint32, params *float64) {
	syscall.Syscall6(gpVideoCaptureStreamParameterdvNV, 4, uintptr(video_capture_slot), uintptr(stream), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func VideoCaptureStreamParameterfvNV(video_capture_slot uint32, stream uint32, pname uint32, params *float32) {
	syscall.Syscall6(gpVideoCaptureStreamParameterfvNV, 4, uintptr(video_capture_slot), uintptr(stream), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}
func VideoCaptureStreamParameterivNV(video_capture_slot uint32, stream uint32, pname uint32, params *int32) {
	syscall.Syscall6(gpVideoCaptureStreamParameterivNV, 4, uintptr(video_capture_slot), uintptr(stream), uintptr(pname), uintptr(unsafe.Pointer(params)), 0, 0)
}

// set the viewport
func Viewport(x int32, y int32, width int32, height int32) {
	syscall.Syscall6(gpViewport, 4, uintptr(x), uintptr(y), uintptr(width), uintptr(height), 0, 0)
}
func ViewportArrayv(first uint32, count int32, v *float32) {
	syscall.Syscall(gpViewportArrayv, 3, uintptr(first), uintptr(count), uintptr(unsafe.Pointer(v)))
}
func ViewportIndexedf(index uint32, x float32, y float32, w float32, h float32) {
	syscall.Syscall6(gpViewportIndexedf, 5, uintptr(index), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(w)), uintptr(math.Float32bits(h)), 0)
}
func ViewportIndexedfv(index uint32, v *float32) {
	syscall.Syscall(gpViewportIndexedfv, 2, uintptr(index), uintptr(unsafe.Pointer(v)), 0)
}
func ViewportPositionWScaleNV(index uint32, xcoeff float32, ycoeff float32) {
	syscall.Syscall(gpViewportPositionWScaleNV, 3, uintptr(index), uintptr(math.Float32bits(xcoeff)), uintptr(math.Float32bits(ycoeff)))
}
func ViewportSwizzleNV(index uint32, swizzlex uint32, swizzley uint32, swizzlez uint32, swizzlew uint32) {
	syscall.Syscall6(gpViewportSwizzleNV, 5, uintptr(index), uintptr(swizzlex), uintptr(swizzley), uintptr(swizzlez), uintptr(swizzlew), 0)
}
func WaitSemaphoreEXT(semaphore uint32, numBufferBarriers uint32, buffers *uint32, numTextureBarriers uint32, textures *uint32, srcLayouts *uint32) {
	syscall.Syscall6(gpWaitSemaphoreEXT, 6, uintptr(semaphore), uintptr(numBufferBarriers), uintptr(unsafe.Pointer(buffers)), uintptr(numTextureBarriers), uintptr(unsafe.Pointer(textures)), uintptr(unsafe.Pointer(srcLayouts)))
}

// instruct the GL server to block until the specified sync object becomes signaled
func WaitSync(sync uintptr, flags uint32, timeout uint64) {
	syscall.Syscall(gpWaitSync, 3, uintptr(sync), uintptr(flags), uintptr(timeout))
}
func WaitVkSemaphoreNV(vkSemaphore uint64) {
	syscall.Syscall(gpWaitVkSemaphoreNV, 1, uintptr(vkSemaphore), 0, 0)
}
func WeightPathsNV(resultPath uint32, numPaths int32, paths *uint32, weights *float32) {
	syscall.Syscall6(gpWeightPathsNV, 4, uintptr(resultPath), uintptr(numPaths), uintptr(unsafe.Pointer(paths)), uintptr(unsafe.Pointer(weights)), 0, 0)
}
func WeightPointerARB(size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	syscall.Syscall6(gpWeightPointerARB, 4, uintptr(size), uintptr(xtype), uintptr(stride), uintptr(pointer), 0, 0)
}
func WeightbvARB(size int32, weights *int8) {
	syscall.Syscall(gpWeightbvARB, 2, uintptr(size), uintptr(unsafe.Pointer(weights)), 0)
}
func WeightdvARB(size int32, weights *float64) {
	syscall.Syscall(gpWeightdvARB, 2, uintptr(size), uintptr(unsafe.Pointer(weights)), 0)
}
func WeightfvARB(size int32, weights *float32) {
	syscall.Syscall(gpWeightfvARB, 2, uintptr(size), uintptr(unsafe.Pointer(weights)), 0)
}
func WeightivARB(size int32, weights *int32) {
	syscall.Syscall(gpWeightivARB, 2, uintptr(size), uintptr(unsafe.Pointer(weights)), 0)
}
func WeightsvARB(size int32, weights *int16) {
	syscall.Syscall(gpWeightsvARB, 2, uintptr(size), uintptr(unsafe.Pointer(weights)), 0)
}
func WeightubvARB(size int32, weights *uint8) {
	syscall.Syscall(gpWeightubvARB, 2, uintptr(size), uintptr(unsafe.Pointer(weights)), 0)
}
func WeightuivARB(size int32, weights *uint32) {
	syscall.Syscall(gpWeightuivARB, 2, uintptr(size), uintptr(unsafe.Pointer(weights)), 0)
}
func WeightusvARB(size int32, weights *uint16) {
	syscall.Syscall(gpWeightusvARB, 2, uintptr(size), uintptr(unsafe.Pointer(weights)), 0)
}
func WindowPos2d(x float64, y float64) {
	syscall.Syscall(gpWindowPos2d, 2, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), 0)
}
func WindowPos2dARB(x float64, y float64) {
	syscall.Syscall(gpWindowPos2dARB, 2, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), 0)
}
func WindowPos2dMESA(x float64, y float64) {
	syscall.Syscall(gpWindowPos2dMESA, 2, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), 0)
}
func WindowPos2dv(v *float64) {
	syscall.Syscall(gpWindowPos2dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2dvARB(v *float64) {
	syscall.Syscall(gpWindowPos2dvARB, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2dvMESA(v *float64) {
	syscall.Syscall(gpWindowPos2dvMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2f(x float32, y float32) {
	syscall.Syscall(gpWindowPos2f, 2, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), 0)
}
func WindowPos2fARB(x float32, y float32) {
	syscall.Syscall(gpWindowPos2fARB, 2, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), 0)
}
func WindowPos2fMESA(x float32, y float32) {
	syscall.Syscall(gpWindowPos2fMESA, 2, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), 0)
}
func WindowPos2fv(v *float32) {
	syscall.Syscall(gpWindowPos2fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2fvARB(v *float32) {
	syscall.Syscall(gpWindowPos2fvARB, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2fvMESA(v *float32) {
	syscall.Syscall(gpWindowPos2fvMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2i(x int32, y int32) {
	syscall.Syscall(gpWindowPos2i, 2, uintptr(x), uintptr(y), 0)
}
func WindowPos2iARB(x int32, y int32) {
	syscall.Syscall(gpWindowPos2iARB, 2, uintptr(x), uintptr(y), 0)
}
func WindowPos2iMESA(x int32, y int32) {
	syscall.Syscall(gpWindowPos2iMESA, 2, uintptr(x), uintptr(y), 0)
}
func WindowPos2iv(v *int32) {
	syscall.Syscall(gpWindowPos2iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2ivARB(v *int32) {
	syscall.Syscall(gpWindowPos2ivARB, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2ivMESA(v *int32) {
	syscall.Syscall(gpWindowPos2ivMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2s(x int16, y int16) {
	syscall.Syscall(gpWindowPos2s, 2, uintptr(x), uintptr(y), 0)
}
func WindowPos2sARB(x int16, y int16) {
	syscall.Syscall(gpWindowPos2sARB, 2, uintptr(x), uintptr(y), 0)
}
func WindowPos2sMESA(x int16, y int16) {
	syscall.Syscall(gpWindowPos2sMESA, 2, uintptr(x), uintptr(y), 0)
}
func WindowPos2sv(v *int16) {
	syscall.Syscall(gpWindowPos2sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2svARB(v *int16) {
	syscall.Syscall(gpWindowPos2svARB, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos2svMESA(v *int16) {
	syscall.Syscall(gpWindowPos2svMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3d(x float64, y float64, z float64) {
	syscall.Syscall(gpWindowPos3d, 3, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)))
}
func WindowPos3dARB(x float64, y float64, z float64) {
	syscall.Syscall(gpWindowPos3dARB, 3, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)))
}
func WindowPos3dMESA(x float64, y float64, z float64) {
	syscall.Syscall(gpWindowPos3dMESA, 3, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)))
}
func WindowPos3dv(v *float64) {
	syscall.Syscall(gpWindowPos3dv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3dvARB(v *float64) {
	syscall.Syscall(gpWindowPos3dvARB, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3dvMESA(v *float64) {
	syscall.Syscall(gpWindowPos3dvMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3f(x float32, y float32, z float32) {
	syscall.Syscall(gpWindowPos3f, 3, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func WindowPos3fARB(x float32, y float32, z float32) {
	syscall.Syscall(gpWindowPos3fARB, 3, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func WindowPos3fMESA(x float32, y float32, z float32) {
	syscall.Syscall(gpWindowPos3fMESA, 3, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
}
func WindowPos3fv(v *float32) {
	syscall.Syscall(gpWindowPos3fv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3fvARB(v *float32) {
	syscall.Syscall(gpWindowPos3fvARB, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3fvMESA(v *float32) {
	syscall.Syscall(gpWindowPos3fvMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3i(x int32, y int32, z int32) {
	syscall.Syscall(gpWindowPos3i, 3, uintptr(x), uintptr(y), uintptr(z))
}
func WindowPos3iARB(x int32, y int32, z int32) {
	syscall.Syscall(gpWindowPos3iARB, 3, uintptr(x), uintptr(y), uintptr(z))
}
func WindowPos3iMESA(x int32, y int32, z int32) {
	syscall.Syscall(gpWindowPos3iMESA, 3, uintptr(x), uintptr(y), uintptr(z))
}
func WindowPos3iv(v *int32) {
	syscall.Syscall(gpWindowPos3iv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3ivARB(v *int32) {
	syscall.Syscall(gpWindowPos3ivARB, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3ivMESA(v *int32) {
	syscall.Syscall(gpWindowPos3ivMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3s(x int16, y int16, z int16) {
	syscall.Syscall(gpWindowPos3s, 3, uintptr(x), uintptr(y), uintptr(z))
}
func WindowPos3sARB(x int16, y int16, z int16) {
	syscall.Syscall(gpWindowPos3sARB, 3, uintptr(x), uintptr(y), uintptr(z))
}
func WindowPos3sMESA(x int16, y int16, z int16) {
	syscall.Syscall(gpWindowPos3sMESA, 3, uintptr(x), uintptr(y), uintptr(z))
}
func WindowPos3sv(v *int16) {
	syscall.Syscall(gpWindowPos3sv, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3svARB(v *int16) {
	syscall.Syscall(gpWindowPos3svARB, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos3svMESA(v *int16) {
	syscall.Syscall(gpWindowPos3svMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos4dMESA(x float64, y float64, z float64, w float64) {
	syscall.Syscall6(gpWindowPos4dMESA, 4, uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)), uintptr(math.Float64bits(z)), uintptr(math.Float64bits(w)), 0, 0)
}
func WindowPos4dvMESA(v *float64) {
	syscall.Syscall(gpWindowPos4dvMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos4fMESA(x float32, y float32, z float32, w float32) {
	syscall.Syscall6(gpWindowPos4fMESA, 4, uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)), uintptr(math.Float32bits(w)), 0, 0)
}
func WindowPos4fvMESA(v *float32) {
	syscall.Syscall(gpWindowPos4fvMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos4iMESA(x int32, y int32, z int32, w int32) {
	syscall.Syscall6(gpWindowPos4iMESA, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func WindowPos4ivMESA(v *int32) {
	syscall.Syscall(gpWindowPos4ivMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowPos4sMESA(x int16, y int16, z int16, w int16) {
	syscall.Syscall6(gpWindowPos4sMESA, 4, uintptr(x), uintptr(y), uintptr(z), uintptr(w), 0, 0)
}
func WindowPos4svMESA(v *int16) {
	syscall.Syscall(gpWindowPos4svMESA, 1, uintptr(unsafe.Pointer(v)), 0, 0)
}
func WindowRectanglesEXT(mode uint32, count int32, box *int32) {
	syscall.Syscall(gpWindowRectanglesEXT, 3, uintptr(mode), uintptr(count), uintptr(unsafe.Pointer(box)))
}
func WriteMaskEXT(res uint32, in uint32, outX uint32, outY uint32, outZ uint32, outW uint32) {
	syscall.Syscall6(gpWriteMaskEXT, 6, uintptr(res), uintptr(in), uintptr(outX), uintptr(outY), uintptr(outZ), uintptr(outW))
}

// InitWithProcAddrFunc intializes the package using the specified OpenGL
// function pointer loading function. For more cases Init should be used
// instead.
func InitWithProcAddrFunc(getProcAddr func(name string) unsafe.Pointer) error {
	gpAccum = uintptr(getProcAddr("glAccum"))
	if gpAccum == 0 {
		return errors.New("glAccum")
	}
	gpAccumxOES = uintptr(getProcAddr("glAccumxOES"))
	gpAcquireKeyedMutexWin32EXT = uintptr(getProcAddr("glAcquireKeyedMutexWin32EXT"))
	gpActiveProgramEXT = uintptr(getProcAddr("glActiveProgramEXT"))
	gpActiveShaderProgram = uintptr(getProcAddr("glActiveShaderProgram"))
	gpActiveShaderProgramEXT = uintptr(getProcAddr("glActiveShaderProgramEXT"))
	gpActiveStencilFaceEXT = uintptr(getProcAddr("glActiveStencilFaceEXT"))
	gpActiveTexture = uintptr(getProcAddr("glActiveTexture"))
	if gpActiveTexture == 0 {
		return errors.New("glActiveTexture")
	}
	gpActiveTextureARB = uintptr(getProcAddr("glActiveTextureARB"))
	gpActiveVaryingNV = uintptr(getProcAddr("glActiveVaryingNV"))
	gpAlphaFragmentOp1ATI = uintptr(getProcAddr("glAlphaFragmentOp1ATI"))
	gpAlphaFragmentOp2ATI = uintptr(getProcAddr("glAlphaFragmentOp2ATI"))
	gpAlphaFragmentOp3ATI = uintptr(getProcAddr("glAlphaFragmentOp3ATI"))
	gpAlphaFunc = uintptr(getProcAddr("glAlphaFunc"))
	if gpAlphaFunc == 0 {
		return errors.New("glAlphaFunc")
	}
	gpAlphaFuncxOES = uintptr(getProcAddr("glAlphaFuncxOES"))
	gpAlphaToCoverageDitherControlNV = uintptr(getProcAddr("glAlphaToCoverageDitherControlNV"))
	gpApplyFramebufferAttachmentCMAAINTEL = uintptr(getProcAddr("glApplyFramebufferAttachmentCMAAINTEL"))
	gpApplyTextureEXT = uintptr(getProcAddr("glApplyTextureEXT"))
	gpAreProgramsResidentNV = uintptr(getProcAddr("glAreProgramsResidentNV"))
	gpAreTexturesResident = uintptr(getProcAddr("glAreTexturesResident"))
	if gpAreTexturesResident == 0 {
		return errors.New("glAreTexturesResident")
	}
	gpAreTexturesResidentEXT = uintptr(getProcAddr("glAreTexturesResidentEXT"))
	gpArrayElement = uintptr(getProcAddr("glArrayElement"))
	if gpArrayElement == 0 {
		return errors.New("glArrayElement")
	}
	gpArrayElementEXT = uintptr(getProcAddr("glArrayElementEXT"))
	gpArrayObjectATI = uintptr(getProcAddr("glArrayObjectATI"))
	gpAsyncMarkerSGIX = uintptr(getProcAddr("glAsyncMarkerSGIX"))
	gpAttachObjectARB = uintptr(getProcAddr("glAttachObjectARB"))
	gpAttachShader = uintptr(getProcAddr("glAttachShader"))
	if gpAttachShader == 0 {
		return errors.New("glAttachShader")
	}
	gpBegin = uintptr(getProcAddr("glBegin"))
	if gpBegin == 0 {
		return errors.New("glBegin")
	}
	gpBeginConditionalRenderNV = uintptr(getProcAddr("glBeginConditionalRenderNV"))
	gpBeginConditionalRenderNVX = uintptr(getProcAddr("glBeginConditionalRenderNVX"))
	gpBeginFragmentShaderATI = uintptr(getProcAddr("glBeginFragmentShaderATI"))
	gpBeginOcclusionQueryNV = uintptr(getProcAddr("glBeginOcclusionQueryNV"))
	gpBeginPerfMonitorAMD = uintptr(getProcAddr("glBeginPerfMonitorAMD"))
	gpBeginPerfQueryINTEL = uintptr(getProcAddr("glBeginPerfQueryINTEL"))
	gpBeginQuery = uintptr(getProcAddr("glBeginQuery"))
	if gpBeginQuery == 0 {
		return errors.New("glBeginQuery")
	}
	gpBeginQueryARB = uintptr(getProcAddr("glBeginQueryARB"))
	gpBeginQueryIndexed = uintptr(getProcAddr("glBeginQueryIndexed"))
	gpBeginTransformFeedbackEXT = uintptr(getProcAddr("glBeginTransformFeedbackEXT"))
	gpBeginTransformFeedbackNV = uintptr(getProcAddr("glBeginTransformFeedbackNV"))
	gpBeginVertexShaderEXT = uintptr(getProcAddr("glBeginVertexShaderEXT"))
	gpBeginVideoCaptureNV = uintptr(getProcAddr("glBeginVideoCaptureNV"))
	gpBindAttribLocation = uintptr(getProcAddr("glBindAttribLocation"))
	if gpBindAttribLocation == 0 {
		return errors.New("glBindAttribLocation")
	}
	gpBindAttribLocationARB = uintptr(getProcAddr("glBindAttribLocationARB"))
	gpBindBuffer = uintptr(getProcAddr("glBindBuffer"))
	if gpBindBuffer == 0 {
		return errors.New("glBindBuffer")
	}
	gpBindBufferARB = uintptr(getProcAddr("glBindBufferARB"))
	gpBindBufferBase = uintptr(getProcAddr("glBindBufferBase"))
	gpBindBufferBaseEXT = uintptr(getProcAddr("glBindBufferBaseEXT"))
	gpBindBufferBaseNV = uintptr(getProcAddr("glBindBufferBaseNV"))
	gpBindBufferOffsetEXT = uintptr(getProcAddr("glBindBufferOffsetEXT"))
	gpBindBufferOffsetNV = uintptr(getProcAddr("glBindBufferOffsetNV"))
	gpBindBufferRange = uintptr(getProcAddr("glBindBufferRange"))
	gpBindBufferRangeEXT = uintptr(getProcAddr("glBindBufferRangeEXT"))
	gpBindBufferRangeNV = uintptr(getProcAddr("glBindBufferRangeNV"))
	gpBindBuffersBase = uintptr(getProcAddr("glBindBuffersBase"))
	gpBindBuffersRange = uintptr(getProcAddr("glBindBuffersRange"))
	gpBindFragDataLocationEXT = uintptr(getProcAddr("glBindFragDataLocationEXT"))
	gpBindFragDataLocationIndexed = uintptr(getProcAddr("glBindFragDataLocationIndexed"))
	gpBindFragmentShaderATI = uintptr(getProcAddr("glBindFragmentShaderATI"))
	gpBindFramebuffer = uintptr(getProcAddr("glBindFramebuffer"))
	gpBindFramebufferEXT = uintptr(getProcAddr("glBindFramebufferEXT"))
	gpBindImageTexture = uintptr(getProcAddr("glBindImageTexture"))
	gpBindImageTextureEXT = uintptr(getProcAddr("glBindImageTextureEXT"))
	gpBindImageTextures = uintptr(getProcAddr("glBindImageTextures"))
	gpBindLightParameterEXT = uintptr(getProcAddr("glBindLightParameterEXT"))
	gpBindMaterialParameterEXT = uintptr(getProcAddr("glBindMaterialParameterEXT"))
	gpBindMultiTextureEXT = uintptr(getProcAddr("glBindMultiTextureEXT"))
	gpBindParameterEXT = uintptr(getProcAddr("glBindParameterEXT"))
	gpBindProgramARB = uintptr(getProcAddr("glBindProgramARB"))
	gpBindProgramNV = uintptr(getProcAddr("glBindProgramNV"))
	gpBindProgramPipeline = uintptr(getProcAddr("glBindProgramPipeline"))
	gpBindProgramPipelineEXT = uintptr(getProcAddr("glBindProgramPipelineEXT"))
	gpBindRenderbuffer = uintptr(getProcAddr("glBindRenderbuffer"))
	gpBindRenderbufferEXT = uintptr(getProcAddr("glBindRenderbufferEXT"))
	gpBindSampler = uintptr(getProcAddr("glBindSampler"))
	gpBindSamplers = uintptr(getProcAddr("glBindSamplers"))
	gpBindTexGenParameterEXT = uintptr(getProcAddr("glBindTexGenParameterEXT"))
	gpBindTexture = uintptr(getProcAddr("glBindTexture"))
	if gpBindTexture == 0 {
		return errors.New("glBindTexture")
	}
	gpBindTextureEXT = uintptr(getProcAddr("glBindTextureEXT"))
	gpBindTextureUnit = uintptr(getProcAddr("glBindTextureUnit"))
	gpBindTextureUnitParameterEXT = uintptr(getProcAddr("glBindTextureUnitParameterEXT"))
	gpBindTextures = uintptr(getProcAddr("glBindTextures"))
	gpBindTransformFeedback = uintptr(getProcAddr("glBindTransformFeedback"))
	gpBindTransformFeedbackNV = uintptr(getProcAddr("glBindTransformFeedbackNV"))
	gpBindVertexArray = uintptr(getProcAddr("glBindVertexArray"))
	gpBindVertexArrayAPPLE = uintptr(getProcAddr("glBindVertexArrayAPPLE"))
	gpBindVertexBuffer = uintptr(getProcAddr("glBindVertexBuffer"))
	gpBindVertexBuffers = uintptr(getProcAddr("glBindVertexBuffers"))
	gpBindVertexShaderEXT = uintptr(getProcAddr("glBindVertexShaderEXT"))
	gpBindVideoCaptureStreamBufferNV = uintptr(getProcAddr("glBindVideoCaptureStreamBufferNV"))
	gpBindVideoCaptureStreamTextureNV = uintptr(getProcAddr("glBindVideoCaptureStreamTextureNV"))
	gpBinormal3bEXT = uintptr(getProcAddr("glBinormal3bEXT"))
	gpBinormal3bvEXT = uintptr(getProcAddr("glBinormal3bvEXT"))
	gpBinormal3dEXT = uintptr(getProcAddr("glBinormal3dEXT"))
	gpBinormal3dvEXT = uintptr(getProcAddr("glBinormal3dvEXT"))
	gpBinormal3fEXT = uintptr(getProcAddr("glBinormal3fEXT"))
	gpBinormal3fvEXT = uintptr(getProcAddr("glBinormal3fvEXT"))
	gpBinormal3iEXT = uintptr(getProcAddr("glBinormal3iEXT"))
	gpBinormal3ivEXT = uintptr(getProcAddr("glBinormal3ivEXT"))
	gpBinormal3sEXT = uintptr(getProcAddr("glBinormal3sEXT"))
	gpBinormal3svEXT = uintptr(getProcAddr("glBinormal3svEXT"))
	gpBinormalPointerEXT = uintptr(getProcAddr("glBinormalPointerEXT"))
	gpBitmap = uintptr(getProcAddr("glBitmap"))
	if gpBitmap == 0 {
		return errors.New("glBitmap")
	}
	gpBitmapxOES = uintptr(getProcAddr("glBitmapxOES"))
	gpBlendBarrierKHR = uintptr(getProcAddr("glBlendBarrierKHR"))
	gpBlendBarrierNV = uintptr(getProcAddr("glBlendBarrierNV"))
	gpBlendColor = uintptr(getProcAddr("glBlendColor"))
	if gpBlendColor == 0 {
		return errors.New("glBlendColor")
	}
	gpBlendColorEXT = uintptr(getProcAddr("glBlendColorEXT"))
	gpBlendColorxOES = uintptr(getProcAddr("glBlendColorxOES"))
	gpBlendEquation = uintptr(getProcAddr("glBlendEquation"))
	if gpBlendEquation == 0 {
		return errors.New("glBlendEquation")
	}
	gpBlendEquationEXT = uintptr(getProcAddr("glBlendEquationEXT"))
	gpBlendEquationIndexedAMD = uintptr(getProcAddr("glBlendEquationIndexedAMD"))
	gpBlendEquationSeparate = uintptr(getProcAddr("glBlendEquationSeparate"))
	if gpBlendEquationSeparate == 0 {
		return errors.New("glBlendEquationSeparate")
	}
	gpBlendEquationSeparateEXT = uintptr(getProcAddr("glBlendEquationSeparateEXT"))
	gpBlendEquationSeparateIndexedAMD = uintptr(getProcAddr("glBlendEquationSeparateIndexedAMD"))
	gpBlendEquationSeparateiARB = uintptr(getProcAddr("glBlendEquationSeparateiARB"))
	gpBlendEquationiARB = uintptr(getProcAddr("glBlendEquationiARB"))
	gpBlendFunc = uintptr(getProcAddr("glBlendFunc"))
	if gpBlendFunc == 0 {
		return errors.New("glBlendFunc")
	}
	gpBlendFuncIndexedAMD = uintptr(getProcAddr("glBlendFuncIndexedAMD"))
	gpBlendFuncSeparate = uintptr(getProcAddr("glBlendFuncSeparate"))
	if gpBlendFuncSeparate == 0 {
		return errors.New("glBlendFuncSeparate")
	}
	gpBlendFuncSeparateEXT = uintptr(getProcAddr("glBlendFuncSeparateEXT"))
	gpBlendFuncSeparateINGR = uintptr(getProcAddr("glBlendFuncSeparateINGR"))
	gpBlendFuncSeparateIndexedAMD = uintptr(getProcAddr("glBlendFuncSeparateIndexedAMD"))
	gpBlendFuncSeparateiARB = uintptr(getProcAddr("glBlendFuncSeparateiARB"))
	gpBlendFunciARB = uintptr(getProcAddr("glBlendFunciARB"))
	gpBlendParameteriNV = uintptr(getProcAddr("glBlendParameteriNV"))
	gpBlitFramebuffer = uintptr(getProcAddr("glBlitFramebuffer"))
	gpBlitFramebufferEXT = uintptr(getProcAddr("glBlitFramebufferEXT"))
	gpBlitNamedFramebuffer = uintptr(getProcAddr("glBlitNamedFramebuffer"))
	gpBufferAddressRangeNV = uintptr(getProcAddr("glBufferAddressRangeNV"))
	gpBufferData = uintptr(getProcAddr("glBufferData"))
	if gpBufferData == 0 {
		return errors.New("glBufferData")
	}
	gpBufferDataARB = uintptr(getProcAddr("glBufferDataARB"))
	gpBufferPageCommitmentARB = uintptr(getProcAddr("glBufferPageCommitmentARB"))
	gpBufferParameteriAPPLE = uintptr(getProcAddr("glBufferParameteriAPPLE"))
	gpBufferStorage = uintptr(getProcAddr("glBufferStorage"))
	gpBufferStorageExternalEXT = uintptr(getProcAddr("glBufferStorageExternalEXT"))
	gpBufferStorageMemEXT = uintptr(getProcAddr("glBufferStorageMemEXT"))
	gpBufferSubData = uintptr(getProcAddr("glBufferSubData"))
	if gpBufferSubData == 0 {
		return errors.New("glBufferSubData")
	}
	gpBufferSubDataARB = uintptr(getProcAddr("glBufferSubDataARB"))
	gpCallCommandListNV = uintptr(getProcAddr("glCallCommandListNV"))
	gpCallList = uintptr(getProcAddr("glCallList"))
	if gpCallList == 0 {
		return errors.New("glCallList")
	}
	gpCallLists = uintptr(getProcAddr("glCallLists"))
	if gpCallLists == 0 {
		return errors.New("glCallLists")
	}
	gpCheckFramebufferStatus = uintptr(getProcAddr("glCheckFramebufferStatus"))
	gpCheckFramebufferStatusEXT = uintptr(getProcAddr("glCheckFramebufferStatusEXT"))
	gpCheckNamedFramebufferStatus = uintptr(getProcAddr("glCheckNamedFramebufferStatus"))
	gpCheckNamedFramebufferStatusEXT = uintptr(getProcAddr("glCheckNamedFramebufferStatusEXT"))
	gpClampColorARB = uintptr(getProcAddr("glClampColorARB"))
	gpClear = uintptr(getProcAddr("glClear"))
	if gpClear == 0 {
		return errors.New("glClear")
	}
	gpClearAccum = uintptr(getProcAddr("glClearAccum"))
	if gpClearAccum == 0 {
		return errors.New("glClearAccum")
	}
	gpClearAccumxOES = uintptr(getProcAddr("glClearAccumxOES"))
	gpClearBufferData = uintptr(getProcAddr("glClearBufferData"))
	gpClearBufferSubData = uintptr(getProcAddr("glClearBufferSubData"))
	gpClearColor = uintptr(getProcAddr("glClearColor"))
	if gpClearColor == 0 {
		return errors.New("glClearColor")
	}
	gpClearColorIiEXT = uintptr(getProcAddr("glClearColorIiEXT"))
	gpClearColorIuiEXT = uintptr(getProcAddr("glClearColorIuiEXT"))
	gpClearColorxOES = uintptr(getProcAddr("glClearColorxOES"))
	gpClearDepth = uintptr(getProcAddr("glClearDepth"))
	if gpClearDepth == 0 {
		return errors.New("glClearDepth")
	}
	gpClearDepthdNV = uintptr(getProcAddr("glClearDepthdNV"))
	gpClearDepthf = uintptr(getProcAddr("glClearDepthf"))
	gpClearDepthfOES = uintptr(getProcAddr("glClearDepthfOES"))
	gpClearDepthxOES = uintptr(getProcAddr("glClearDepthxOES"))
	gpClearIndex = uintptr(getProcAddr("glClearIndex"))
	if gpClearIndex == 0 {
		return errors.New("glClearIndex")
	}
	gpClearNamedBufferData = uintptr(getProcAddr("glClearNamedBufferData"))
	gpClearNamedBufferDataEXT = uintptr(getProcAddr("glClearNamedBufferDataEXT"))
	gpClearNamedBufferSubData = uintptr(getProcAddr("glClearNamedBufferSubData"))
	gpClearNamedBufferSubDataEXT = uintptr(getProcAddr("glClearNamedBufferSubDataEXT"))
	gpClearNamedFramebufferfi = uintptr(getProcAddr("glClearNamedFramebufferfi"))
	gpClearNamedFramebufferfv = uintptr(getProcAddr("glClearNamedFramebufferfv"))
	gpClearNamedFramebufferiv = uintptr(getProcAddr("glClearNamedFramebufferiv"))
	gpClearNamedFramebufferuiv = uintptr(getProcAddr("glClearNamedFramebufferuiv"))
	gpClearStencil = uintptr(getProcAddr("glClearStencil"))
	if gpClearStencil == 0 {
		return errors.New("glClearStencil")
	}
	gpClearTexImage = uintptr(getProcAddr("glClearTexImage"))
	gpClearTexSubImage = uintptr(getProcAddr("glClearTexSubImage"))
	gpClientActiveTexture = uintptr(getProcAddr("glClientActiveTexture"))
	if gpClientActiveTexture == 0 {
		return errors.New("glClientActiveTexture")
	}
	gpClientActiveTextureARB = uintptr(getProcAddr("glClientActiveTextureARB"))
	gpClientActiveVertexStreamATI = uintptr(getProcAddr("glClientActiveVertexStreamATI"))
	gpClientAttribDefaultEXT = uintptr(getProcAddr("glClientAttribDefaultEXT"))
	gpClientWaitSync = uintptr(getProcAddr("glClientWaitSync"))
	gpClipControl = uintptr(getProcAddr("glClipControl"))
	gpClipPlane = uintptr(getProcAddr("glClipPlane"))
	if gpClipPlane == 0 {
		return errors.New("glClipPlane")
	}
	gpClipPlanefOES = uintptr(getProcAddr("glClipPlanefOES"))
	gpClipPlanexOES = uintptr(getProcAddr("glClipPlanexOES"))
	gpColor3b = uintptr(getProcAddr("glColor3b"))
	if gpColor3b == 0 {
		return errors.New("glColor3b")
	}
	gpColor3bv = uintptr(getProcAddr("glColor3bv"))
	if gpColor3bv == 0 {
		return errors.New("glColor3bv")
	}
	gpColor3d = uintptr(getProcAddr("glColor3d"))
	if gpColor3d == 0 {
		return errors.New("glColor3d")
	}
	gpColor3dv = uintptr(getProcAddr("glColor3dv"))
	if gpColor3dv == 0 {
		return errors.New("glColor3dv")
	}
	gpColor3f = uintptr(getProcAddr("glColor3f"))
	if gpColor3f == 0 {
		return errors.New("glColor3f")
	}
	gpColor3fVertex3fSUN = uintptr(getProcAddr("glColor3fVertex3fSUN"))
	gpColor3fVertex3fvSUN = uintptr(getProcAddr("glColor3fVertex3fvSUN"))
	gpColor3fv = uintptr(getProcAddr("glColor3fv"))
	if gpColor3fv == 0 {
		return errors.New("glColor3fv")
	}
	gpColor3hNV = uintptr(getProcAddr("glColor3hNV"))
	gpColor3hvNV = uintptr(getProcAddr("glColor3hvNV"))
	gpColor3i = uintptr(getProcAddr("glColor3i"))
	if gpColor3i == 0 {
		return errors.New("glColor3i")
	}
	gpColor3iv = uintptr(getProcAddr("glColor3iv"))
	if gpColor3iv == 0 {
		return errors.New("glColor3iv")
	}
	gpColor3s = uintptr(getProcAddr("glColor3s"))
	if gpColor3s == 0 {
		return errors.New("glColor3s")
	}
	gpColor3sv = uintptr(getProcAddr("glColor3sv"))
	if gpColor3sv == 0 {
		return errors.New("glColor3sv")
	}
	gpColor3ub = uintptr(getProcAddr("glColor3ub"))
	if gpColor3ub == 0 {
		return errors.New("glColor3ub")
	}
	gpColor3ubv = uintptr(getProcAddr("glColor3ubv"))
	if gpColor3ubv == 0 {
		return errors.New("glColor3ubv")
	}
	gpColor3ui = uintptr(getProcAddr("glColor3ui"))
	if gpColor3ui == 0 {
		return errors.New("glColor3ui")
	}
	gpColor3uiv = uintptr(getProcAddr("glColor3uiv"))
	if gpColor3uiv == 0 {
		return errors.New("glColor3uiv")
	}
	gpColor3us = uintptr(getProcAddr("glColor3us"))
	if gpColor3us == 0 {
		return errors.New("glColor3us")
	}
	gpColor3usv = uintptr(getProcAddr("glColor3usv"))
	if gpColor3usv == 0 {
		return errors.New("glColor3usv")
	}
	gpColor3xOES = uintptr(getProcAddr("glColor3xOES"))
	gpColor3xvOES = uintptr(getProcAddr("glColor3xvOES"))
	gpColor4b = uintptr(getProcAddr("glColor4b"))
	if gpColor4b == 0 {
		return errors.New("glColor4b")
	}
	gpColor4bv = uintptr(getProcAddr("glColor4bv"))
	if gpColor4bv == 0 {
		return errors.New("glColor4bv")
	}
	gpColor4d = uintptr(getProcAddr("glColor4d"))
	if gpColor4d == 0 {
		return errors.New("glColor4d")
	}
	gpColor4dv = uintptr(getProcAddr("glColor4dv"))
	if gpColor4dv == 0 {
		return errors.New("glColor4dv")
	}
	gpColor4f = uintptr(getProcAddr("glColor4f"))
	if gpColor4f == 0 {
		return errors.New("glColor4f")
	}
	gpColor4fNormal3fVertex3fSUN = uintptr(getProcAddr("glColor4fNormal3fVertex3fSUN"))
	gpColor4fNormal3fVertex3fvSUN = uintptr(getProcAddr("glColor4fNormal3fVertex3fvSUN"))
	gpColor4fv = uintptr(getProcAddr("glColor4fv"))
	if gpColor4fv == 0 {
		return errors.New("glColor4fv")
	}
	gpColor4hNV = uintptr(getProcAddr("glColor4hNV"))
	gpColor4hvNV = uintptr(getProcAddr("glColor4hvNV"))
	gpColor4i = uintptr(getProcAddr("glColor4i"))
	if gpColor4i == 0 {
		return errors.New("glColor4i")
	}
	gpColor4iv = uintptr(getProcAddr("glColor4iv"))
	if gpColor4iv == 0 {
		return errors.New("glColor4iv")
	}
	gpColor4s = uintptr(getProcAddr("glColor4s"))
	if gpColor4s == 0 {
		return errors.New("glColor4s")
	}
	gpColor4sv = uintptr(getProcAddr("glColor4sv"))
	if gpColor4sv == 0 {
		return errors.New("glColor4sv")
	}
	gpColor4ub = uintptr(getProcAddr("glColor4ub"))
	if gpColor4ub == 0 {
		return errors.New("glColor4ub")
	}
	gpColor4ubVertex2fSUN = uintptr(getProcAddr("glColor4ubVertex2fSUN"))
	gpColor4ubVertex2fvSUN = uintptr(getProcAddr("glColor4ubVertex2fvSUN"))
	gpColor4ubVertex3fSUN = uintptr(getProcAddr("glColor4ubVertex3fSUN"))
	gpColor4ubVertex3fvSUN = uintptr(getProcAddr("glColor4ubVertex3fvSUN"))
	gpColor4ubv = uintptr(getProcAddr("glColor4ubv"))
	if gpColor4ubv == 0 {
		return errors.New("glColor4ubv")
	}
	gpColor4ui = uintptr(getProcAddr("glColor4ui"))
	if gpColor4ui == 0 {
		return errors.New("glColor4ui")
	}
	gpColor4uiv = uintptr(getProcAddr("glColor4uiv"))
	if gpColor4uiv == 0 {
		return errors.New("glColor4uiv")
	}
	gpColor4us = uintptr(getProcAddr("glColor4us"))
	if gpColor4us == 0 {
		return errors.New("glColor4us")
	}
	gpColor4usv = uintptr(getProcAddr("glColor4usv"))
	if gpColor4usv == 0 {
		return errors.New("glColor4usv")
	}
	gpColor4xOES = uintptr(getProcAddr("glColor4xOES"))
	gpColor4xvOES = uintptr(getProcAddr("glColor4xvOES"))
	gpColorFormatNV = uintptr(getProcAddr("glColorFormatNV"))
	gpColorFragmentOp1ATI = uintptr(getProcAddr("glColorFragmentOp1ATI"))
	gpColorFragmentOp2ATI = uintptr(getProcAddr("glColorFragmentOp2ATI"))
	gpColorFragmentOp3ATI = uintptr(getProcAddr("glColorFragmentOp3ATI"))
	gpColorMask = uintptr(getProcAddr("glColorMask"))
	if gpColorMask == 0 {
		return errors.New("glColorMask")
	}
	gpColorMaskIndexedEXT = uintptr(getProcAddr("glColorMaskIndexedEXT"))
	gpColorMaterial = uintptr(getProcAddr("glColorMaterial"))
	if gpColorMaterial == 0 {
		return errors.New("glColorMaterial")
	}
	gpColorPointer = uintptr(getProcAddr("glColorPointer"))
	if gpColorPointer == 0 {
		return errors.New("glColorPointer")
	}
	gpColorPointerEXT = uintptr(getProcAddr("glColorPointerEXT"))
	gpColorPointerListIBM = uintptr(getProcAddr("glColorPointerListIBM"))
	gpColorPointervINTEL = uintptr(getProcAddr("glColorPointervINTEL"))
	gpColorSubTableEXT = uintptr(getProcAddr("glColorSubTableEXT"))
	gpColorTableEXT = uintptr(getProcAddr("glColorTableEXT"))
	gpColorTableParameterfvSGI = uintptr(getProcAddr("glColorTableParameterfvSGI"))
	gpColorTableParameterivSGI = uintptr(getProcAddr("glColorTableParameterivSGI"))
	gpColorTableSGI = uintptr(getProcAddr("glColorTableSGI"))
	gpCombinerInputNV = uintptr(getProcAddr("glCombinerInputNV"))
	gpCombinerOutputNV = uintptr(getProcAddr("glCombinerOutputNV"))
	gpCombinerParameterfNV = uintptr(getProcAddr("glCombinerParameterfNV"))
	gpCombinerParameterfvNV = uintptr(getProcAddr("glCombinerParameterfvNV"))
	gpCombinerParameteriNV = uintptr(getProcAddr("glCombinerParameteriNV"))
	gpCombinerParameterivNV = uintptr(getProcAddr("glCombinerParameterivNV"))
	gpCombinerStageParameterfvNV = uintptr(getProcAddr("glCombinerStageParameterfvNV"))
	gpCommandListSegmentsNV = uintptr(getProcAddr("glCommandListSegmentsNV"))
	gpCompileCommandListNV = uintptr(getProcAddr("glCompileCommandListNV"))
	gpCompileShader = uintptr(getProcAddr("glCompileShader"))
	if gpCompileShader == 0 {
		return errors.New("glCompileShader")
	}
	gpCompileShaderARB = uintptr(getProcAddr("glCompileShaderARB"))
	gpCompileShaderIncludeARB = uintptr(getProcAddr("glCompileShaderIncludeARB"))
	gpCompressedMultiTexImage1DEXT = uintptr(getProcAddr("glCompressedMultiTexImage1DEXT"))
	gpCompressedMultiTexImage2DEXT = uintptr(getProcAddr("glCompressedMultiTexImage2DEXT"))
	gpCompressedMultiTexImage3DEXT = uintptr(getProcAddr("glCompressedMultiTexImage3DEXT"))
	gpCompressedMultiTexSubImage1DEXT = uintptr(getProcAddr("glCompressedMultiTexSubImage1DEXT"))
	gpCompressedMultiTexSubImage2DEXT = uintptr(getProcAddr("glCompressedMultiTexSubImage2DEXT"))
	gpCompressedMultiTexSubImage3DEXT = uintptr(getProcAddr("glCompressedMultiTexSubImage3DEXT"))
	gpCompressedTexImage1D = uintptr(getProcAddr("glCompressedTexImage1D"))
	if gpCompressedTexImage1D == 0 {
		return errors.New("glCompressedTexImage1D")
	}
	gpCompressedTexImage1DARB = uintptr(getProcAddr("glCompressedTexImage1DARB"))
	gpCompressedTexImage2D = uintptr(getProcAddr("glCompressedTexImage2D"))
	if gpCompressedTexImage2D == 0 {
		return errors.New("glCompressedTexImage2D")
	}
	gpCompressedTexImage2DARB = uintptr(getProcAddr("glCompressedTexImage2DARB"))
	gpCompressedTexImage3D = uintptr(getProcAddr("glCompressedTexImage3D"))
	if gpCompressedTexImage3D == 0 {
		return errors.New("glCompressedTexImage3D")
	}
	gpCompressedTexImage3DARB = uintptr(getProcAddr("glCompressedTexImage3DARB"))
	gpCompressedTexSubImage1D = uintptr(getProcAddr("glCompressedTexSubImage1D"))
	if gpCompressedTexSubImage1D == 0 {
		return errors.New("glCompressedTexSubImage1D")
	}
	gpCompressedTexSubImage1DARB = uintptr(getProcAddr("glCompressedTexSubImage1DARB"))
	gpCompressedTexSubImage2D = uintptr(getProcAddr("glCompressedTexSubImage2D"))
	if gpCompressedTexSubImage2D == 0 {
		return errors.New("glCompressedTexSubImage2D")
	}
	gpCompressedTexSubImage2DARB = uintptr(getProcAddr("glCompressedTexSubImage2DARB"))
	gpCompressedTexSubImage3D = uintptr(getProcAddr("glCompressedTexSubImage3D"))
	if gpCompressedTexSubImage3D == 0 {
		return errors.New("glCompressedTexSubImage3D")
	}
	gpCompressedTexSubImage3DARB = uintptr(getProcAddr("glCompressedTexSubImage3DARB"))
	gpCompressedTextureImage1DEXT = uintptr(getProcAddr("glCompressedTextureImage1DEXT"))
	gpCompressedTextureImage2DEXT = uintptr(getProcAddr("glCompressedTextureImage2DEXT"))
	gpCompressedTextureImage3DEXT = uintptr(getProcAddr("glCompressedTextureImage3DEXT"))
	gpCompressedTextureSubImage1D = uintptr(getProcAddr("glCompressedTextureSubImage1D"))
	gpCompressedTextureSubImage1DEXT = uintptr(getProcAddr("glCompressedTextureSubImage1DEXT"))
	gpCompressedTextureSubImage2D = uintptr(getProcAddr("glCompressedTextureSubImage2D"))
	gpCompressedTextureSubImage2DEXT = uintptr(getProcAddr("glCompressedTextureSubImage2DEXT"))
	gpCompressedTextureSubImage3D = uintptr(getProcAddr("glCompressedTextureSubImage3D"))
	gpCompressedTextureSubImage3DEXT = uintptr(getProcAddr("glCompressedTextureSubImage3DEXT"))
	gpConservativeRasterParameterfNV = uintptr(getProcAddr("glConservativeRasterParameterfNV"))
	gpConservativeRasterParameteriNV = uintptr(getProcAddr("glConservativeRasterParameteriNV"))
	gpConvolutionFilter1DEXT = uintptr(getProcAddr("glConvolutionFilter1DEXT"))
	gpConvolutionFilter2DEXT = uintptr(getProcAddr("glConvolutionFilter2DEXT"))
	gpConvolutionParameterfEXT = uintptr(getProcAddr("glConvolutionParameterfEXT"))
	gpConvolutionParameterfvEXT = uintptr(getProcAddr("glConvolutionParameterfvEXT"))
	gpConvolutionParameteriEXT = uintptr(getProcAddr("glConvolutionParameteriEXT"))
	gpConvolutionParameterivEXT = uintptr(getProcAddr("glConvolutionParameterivEXT"))
	gpConvolutionParameterxOES = uintptr(getProcAddr("glConvolutionParameterxOES"))
	gpConvolutionParameterxvOES = uintptr(getProcAddr("glConvolutionParameterxvOES"))
	gpCopyBufferSubData = uintptr(getProcAddr("glCopyBufferSubData"))
	gpCopyColorSubTableEXT = uintptr(getProcAddr("glCopyColorSubTableEXT"))
	gpCopyColorTableSGI = uintptr(getProcAddr("glCopyColorTableSGI"))
	gpCopyConvolutionFilter1DEXT = uintptr(getProcAddr("glCopyConvolutionFilter1DEXT"))
	gpCopyConvolutionFilter2DEXT = uintptr(getProcAddr("glCopyConvolutionFilter2DEXT"))
	gpCopyImageSubData = uintptr(getProcAddr("glCopyImageSubData"))
	gpCopyImageSubDataNV = uintptr(getProcAddr("glCopyImageSubDataNV"))
	gpCopyMultiTexImage1DEXT = uintptr(getProcAddr("glCopyMultiTexImage1DEXT"))
	gpCopyMultiTexImage2DEXT = uintptr(getProcAddr("glCopyMultiTexImage2DEXT"))
	gpCopyMultiTexSubImage1DEXT = uintptr(getProcAddr("glCopyMultiTexSubImage1DEXT"))
	gpCopyMultiTexSubImage2DEXT = uintptr(getProcAddr("glCopyMultiTexSubImage2DEXT"))
	gpCopyMultiTexSubImage3DEXT = uintptr(getProcAddr("glCopyMultiTexSubImage3DEXT"))
	gpCopyNamedBufferSubData = uintptr(getProcAddr("glCopyNamedBufferSubData"))
	gpCopyPathNV = uintptr(getProcAddr("glCopyPathNV"))
	gpCopyPixels = uintptr(getProcAddr("glCopyPixels"))
	if gpCopyPixels == 0 {
		return errors.New("glCopyPixels")
	}
	gpCopyTexImage1D = uintptr(getProcAddr("glCopyTexImage1D"))
	if gpCopyTexImage1D == 0 {
		return errors.New("glCopyTexImage1D")
	}
	gpCopyTexImage1DEXT = uintptr(getProcAddr("glCopyTexImage1DEXT"))
	gpCopyTexImage2D = uintptr(getProcAddr("glCopyTexImage2D"))
	if gpCopyTexImage2D == 0 {
		return errors.New("glCopyTexImage2D")
	}
	gpCopyTexImage2DEXT = uintptr(getProcAddr("glCopyTexImage2DEXT"))
	gpCopyTexSubImage1D = uintptr(getProcAddr("glCopyTexSubImage1D"))
	if gpCopyTexSubImage1D == 0 {
		return errors.New("glCopyTexSubImage1D")
	}
	gpCopyTexSubImage1DEXT = uintptr(getProcAddr("glCopyTexSubImage1DEXT"))
	gpCopyTexSubImage2D = uintptr(getProcAddr("glCopyTexSubImage2D"))
	if gpCopyTexSubImage2D == 0 {
		return errors.New("glCopyTexSubImage2D")
	}
	gpCopyTexSubImage2DEXT = uintptr(getProcAddr("glCopyTexSubImage2DEXT"))
	gpCopyTexSubImage3D = uintptr(getProcAddr("glCopyTexSubImage3D"))
	if gpCopyTexSubImage3D == 0 {
		return errors.New("glCopyTexSubImage3D")
	}
	gpCopyTexSubImage3DEXT = uintptr(getProcAddr("glCopyTexSubImage3DEXT"))
	gpCopyTextureImage1DEXT = uintptr(getProcAddr("glCopyTextureImage1DEXT"))
	gpCopyTextureImage2DEXT = uintptr(getProcAddr("glCopyTextureImage2DEXT"))
	gpCopyTextureSubImage1D = uintptr(getProcAddr("glCopyTextureSubImage1D"))
	gpCopyTextureSubImage1DEXT = uintptr(getProcAddr("glCopyTextureSubImage1DEXT"))
	gpCopyTextureSubImage2D = uintptr(getProcAddr("glCopyTextureSubImage2D"))
	gpCopyTextureSubImage2DEXT = uintptr(getProcAddr("glCopyTextureSubImage2DEXT"))
	gpCopyTextureSubImage3D = uintptr(getProcAddr("glCopyTextureSubImage3D"))
	gpCopyTextureSubImage3DEXT = uintptr(getProcAddr("glCopyTextureSubImage3DEXT"))
	gpCoverFillPathInstancedNV = uintptr(getProcAddr("glCoverFillPathInstancedNV"))
	gpCoverFillPathNV = uintptr(getProcAddr("glCoverFillPathNV"))
	gpCoverStrokePathInstancedNV = uintptr(getProcAddr("glCoverStrokePathInstancedNV"))
	gpCoverStrokePathNV = uintptr(getProcAddr("glCoverStrokePathNV"))
	gpCoverageModulationNV = uintptr(getProcAddr("glCoverageModulationNV"))
	gpCoverageModulationTableNV = uintptr(getProcAddr("glCoverageModulationTableNV"))
	gpCreateBuffers = uintptr(getProcAddr("glCreateBuffers"))
	gpCreateCommandListsNV = uintptr(getProcAddr("glCreateCommandListsNV"))
	gpCreateFramebuffers = uintptr(getProcAddr("glCreateFramebuffers"))
	gpCreateMemoryObjectsEXT = uintptr(getProcAddr("glCreateMemoryObjectsEXT"))
	gpCreatePerfQueryINTEL = uintptr(getProcAddr("glCreatePerfQueryINTEL"))
	gpCreateProgram = uintptr(getProcAddr("glCreateProgram"))
	if gpCreateProgram == 0 {
		return errors.New("glCreateProgram")
	}
	gpCreateProgramObjectARB = uintptr(getProcAddr("glCreateProgramObjectARB"))
	gpCreateProgramPipelines = uintptr(getProcAddr("glCreateProgramPipelines"))
	gpCreateQueries = uintptr(getProcAddr("glCreateQueries"))
	gpCreateRenderbuffers = uintptr(getProcAddr("glCreateRenderbuffers"))
	gpCreateSamplers = uintptr(getProcAddr("glCreateSamplers"))
	gpCreateShader = uintptr(getProcAddr("glCreateShader"))
	if gpCreateShader == 0 {
		return errors.New("glCreateShader")
	}
	gpCreateShaderObjectARB = uintptr(getProcAddr("glCreateShaderObjectARB"))
	gpCreateShaderProgramEXT = uintptr(getProcAddr("glCreateShaderProgramEXT"))
	gpCreateShaderProgramv = uintptr(getProcAddr("glCreateShaderProgramv"))
	gpCreateShaderProgramvEXT = uintptr(getProcAddr("glCreateShaderProgramvEXT"))
	gpCreateStatesNV = uintptr(getProcAddr("glCreateStatesNV"))
	gpCreateSyncFromCLeventARB = uintptr(getProcAddr("glCreateSyncFromCLeventARB"))
	gpCreateTextures = uintptr(getProcAddr("glCreateTextures"))
	gpCreateTransformFeedbacks = uintptr(getProcAddr("glCreateTransformFeedbacks"))
	gpCreateVertexArrays = uintptr(getProcAddr("glCreateVertexArrays"))
	gpCullFace = uintptr(getProcAddr("glCullFace"))
	if gpCullFace == 0 {
		return errors.New("glCullFace")
	}
	gpCullParameterdvEXT = uintptr(getProcAddr("glCullParameterdvEXT"))
	gpCullParameterfvEXT = uintptr(getProcAddr("glCullParameterfvEXT"))
	gpCurrentPaletteMatrixARB = uintptr(getProcAddr("glCurrentPaletteMatrixARB"))
	gpDebugMessageCallback = uintptr(getProcAddr("glDebugMessageCallback"))
	gpDebugMessageCallbackAMD = uintptr(getProcAddr("glDebugMessageCallbackAMD"))
	gpDebugMessageCallbackARB = uintptr(getProcAddr("glDebugMessageCallbackARB"))
	gpDebugMessageCallbackKHR = uintptr(getProcAddr("glDebugMessageCallbackKHR"))
	gpDebugMessageControl = uintptr(getProcAddr("glDebugMessageControl"))
	gpDebugMessageControlARB = uintptr(getProcAddr("glDebugMessageControlARB"))
	gpDebugMessageControlKHR = uintptr(getProcAddr("glDebugMessageControlKHR"))
	gpDebugMessageEnableAMD = uintptr(getProcAddr("glDebugMessageEnableAMD"))
	gpDebugMessageInsert = uintptr(getProcAddr("glDebugMessageInsert"))
	gpDebugMessageInsertAMD = uintptr(getProcAddr("glDebugMessageInsertAMD"))
	gpDebugMessageInsertARB = uintptr(getProcAddr("glDebugMessageInsertARB"))
	gpDebugMessageInsertKHR = uintptr(getProcAddr("glDebugMessageInsertKHR"))
	gpDeformSGIX = uintptr(getProcAddr("glDeformSGIX"))
	gpDeformationMap3dSGIX = uintptr(getProcAddr("glDeformationMap3dSGIX"))
	gpDeformationMap3fSGIX = uintptr(getProcAddr("glDeformationMap3fSGIX"))
	gpDeleteAsyncMarkersSGIX = uintptr(getProcAddr("glDeleteAsyncMarkersSGIX"))
	gpDeleteBuffers = uintptr(getProcAddr("glDeleteBuffers"))
	if gpDeleteBuffers == 0 {
		return errors.New("glDeleteBuffers")
	}
	gpDeleteBuffersARB = uintptr(getProcAddr("glDeleteBuffersARB"))
	gpDeleteCommandListsNV = uintptr(getProcAddr("glDeleteCommandListsNV"))
	gpDeleteFencesAPPLE = uintptr(getProcAddr("glDeleteFencesAPPLE"))
	gpDeleteFencesNV = uintptr(getProcAddr("glDeleteFencesNV"))
	gpDeleteFragmentShaderATI = uintptr(getProcAddr("glDeleteFragmentShaderATI"))
	gpDeleteFramebuffers = uintptr(getProcAddr("glDeleteFramebuffers"))
	gpDeleteFramebuffersEXT = uintptr(getProcAddr("glDeleteFramebuffersEXT"))
	gpDeleteLists = uintptr(getProcAddr("glDeleteLists"))
	if gpDeleteLists == 0 {
		return errors.New("glDeleteLists")
	}
	gpDeleteMemoryObjectsEXT = uintptr(getProcAddr("glDeleteMemoryObjectsEXT"))
	gpDeleteNamedStringARB = uintptr(getProcAddr("glDeleteNamedStringARB"))
	gpDeleteNamesAMD = uintptr(getProcAddr("glDeleteNamesAMD"))
	gpDeleteObjectARB = uintptr(getProcAddr("glDeleteObjectARB"))
	gpDeleteOcclusionQueriesNV = uintptr(getProcAddr("glDeleteOcclusionQueriesNV"))
	gpDeletePathsNV = uintptr(getProcAddr("glDeletePathsNV"))
	gpDeletePerfMonitorsAMD = uintptr(getProcAddr("glDeletePerfMonitorsAMD"))
	gpDeletePerfQueryINTEL = uintptr(getProcAddr("glDeletePerfQueryINTEL"))
	gpDeleteProgram = uintptr(getProcAddr("glDeleteProgram"))
	if gpDeleteProgram == 0 {
		return errors.New("glDeleteProgram")
	}
	gpDeleteProgramPipelines = uintptr(getProcAddr("glDeleteProgramPipelines"))
	gpDeleteProgramPipelinesEXT = uintptr(getProcAddr("glDeleteProgramPipelinesEXT"))
	gpDeleteProgramsARB = uintptr(getProcAddr("glDeleteProgramsARB"))
	gpDeleteProgramsNV = uintptr(getProcAddr("glDeleteProgramsNV"))
	gpDeleteQueries = uintptr(getProcAddr("glDeleteQueries"))
	if gpDeleteQueries == 0 {
		return errors.New("glDeleteQueries")
	}
	gpDeleteQueriesARB = uintptr(getProcAddr("glDeleteQueriesARB"))
	gpDeleteQueryResourceTagNV = uintptr(getProcAddr("glDeleteQueryResourceTagNV"))
	gpDeleteRenderbuffers = uintptr(getProcAddr("glDeleteRenderbuffers"))
	gpDeleteRenderbuffersEXT = uintptr(getProcAddr("glDeleteRenderbuffersEXT"))
	gpDeleteSamplers = uintptr(getProcAddr("glDeleteSamplers"))
	gpDeleteSemaphoresEXT = uintptr(getProcAddr("glDeleteSemaphoresEXT"))
	gpDeleteShader = uintptr(getProcAddr("glDeleteShader"))
	if gpDeleteShader == 0 {
		return errors.New("glDeleteShader")
	}
	gpDeleteStatesNV = uintptr(getProcAddr("glDeleteStatesNV"))
	gpDeleteSync = uintptr(getProcAddr("glDeleteSync"))
	gpDeleteTextures = uintptr(getProcAddr("glDeleteTextures"))
	if gpDeleteTextures == 0 {
		return errors.New("glDeleteTextures")
	}
	gpDeleteTexturesEXT = uintptr(getProcAddr("glDeleteTexturesEXT"))
	gpDeleteTransformFeedbacks = uintptr(getProcAddr("glDeleteTransformFeedbacks"))
	gpDeleteTransformFeedbacksNV = uintptr(getProcAddr("glDeleteTransformFeedbacksNV"))
	gpDeleteVertexArrays = uintptr(getProcAddr("glDeleteVertexArrays"))
	gpDeleteVertexArraysAPPLE = uintptr(getProcAddr("glDeleteVertexArraysAPPLE"))
	gpDeleteVertexShaderEXT = uintptr(getProcAddr("glDeleteVertexShaderEXT"))
	gpDepthBoundsEXT = uintptr(getProcAddr("glDepthBoundsEXT"))
	gpDepthBoundsdNV = uintptr(getProcAddr("glDepthBoundsdNV"))
	gpDepthFunc = uintptr(getProcAddr("glDepthFunc"))
	if gpDepthFunc == 0 {
		return errors.New("glDepthFunc")
	}
	gpDepthMask = uintptr(getProcAddr("glDepthMask"))
	if gpDepthMask == 0 {
		return errors.New("glDepthMask")
	}
	gpDepthRange = uintptr(getProcAddr("glDepthRange"))
	if gpDepthRange == 0 {
		return errors.New("glDepthRange")
	}
	gpDepthRangeArrayv = uintptr(getProcAddr("glDepthRangeArrayv"))
	gpDepthRangeIndexed = uintptr(getProcAddr("glDepthRangeIndexed"))
	gpDepthRangedNV = uintptr(getProcAddr("glDepthRangedNV"))
	gpDepthRangef = uintptr(getProcAddr("glDepthRangef"))
	gpDepthRangefOES = uintptr(getProcAddr("glDepthRangefOES"))
	gpDepthRangexOES = uintptr(getProcAddr("glDepthRangexOES"))
	gpDetachObjectARB = uintptr(getProcAddr("glDetachObjectARB"))
	gpDetachShader = uintptr(getProcAddr("glDetachShader"))
	if gpDetachShader == 0 {
		return errors.New("glDetachShader")
	}
	gpDetailTexFuncSGIS = uintptr(getProcAddr("glDetailTexFuncSGIS"))
	gpDisable = uintptr(getProcAddr("glDisable"))
	if gpDisable == 0 {
		return errors.New("glDisable")
	}
	gpDisableClientState = uintptr(getProcAddr("glDisableClientState"))
	if gpDisableClientState == 0 {
		return errors.New("glDisableClientState")
	}
	gpDisableClientStateIndexedEXT = uintptr(getProcAddr("glDisableClientStateIndexedEXT"))
	gpDisableClientStateiEXT = uintptr(getProcAddr("glDisableClientStateiEXT"))
	gpDisableIndexedEXT = uintptr(getProcAddr("glDisableIndexedEXT"))
	gpDisableVariantClientStateEXT = uintptr(getProcAddr("glDisableVariantClientStateEXT"))
	gpDisableVertexArrayAttrib = uintptr(getProcAddr("glDisableVertexArrayAttrib"))
	gpDisableVertexArrayAttribEXT = uintptr(getProcAddr("glDisableVertexArrayAttribEXT"))
	gpDisableVertexArrayEXT = uintptr(getProcAddr("glDisableVertexArrayEXT"))
	gpDisableVertexAttribAPPLE = uintptr(getProcAddr("glDisableVertexAttribAPPLE"))
	gpDisableVertexAttribArray = uintptr(getProcAddr("glDisableVertexAttribArray"))
	if gpDisableVertexAttribArray == 0 {
		return errors.New("glDisableVertexAttribArray")
	}
	gpDisableVertexAttribArrayARB = uintptr(getProcAddr("glDisableVertexAttribArrayARB"))
	gpDispatchCompute = uintptr(getProcAddr("glDispatchCompute"))
	gpDispatchComputeGroupSizeARB = uintptr(getProcAddr("glDispatchComputeGroupSizeARB"))
	gpDispatchComputeIndirect = uintptr(getProcAddr("glDispatchComputeIndirect"))
	gpDrawArrays = uintptr(getProcAddr("glDrawArrays"))
	if gpDrawArrays == 0 {
		return errors.New("glDrawArrays")
	}
	gpDrawArraysEXT = uintptr(getProcAddr("glDrawArraysEXT"))
	gpDrawArraysIndirect = uintptr(getProcAddr("glDrawArraysIndirect"))
	gpDrawArraysInstancedARB = uintptr(getProcAddr("glDrawArraysInstancedARB"))
	gpDrawArraysInstancedBaseInstance = uintptr(getProcAddr("glDrawArraysInstancedBaseInstance"))
	gpDrawArraysInstancedEXT = uintptr(getProcAddr("glDrawArraysInstancedEXT"))
	gpDrawBuffer = uintptr(getProcAddr("glDrawBuffer"))
	if gpDrawBuffer == 0 {
		return errors.New("glDrawBuffer")
	}
	gpDrawBuffers = uintptr(getProcAddr("glDrawBuffers"))
	if gpDrawBuffers == 0 {
		return errors.New("glDrawBuffers")
	}
	gpDrawBuffersARB = uintptr(getProcAddr("glDrawBuffersARB"))
	gpDrawBuffersATI = uintptr(getProcAddr("glDrawBuffersATI"))
	gpDrawCommandsAddressNV = uintptr(getProcAddr("glDrawCommandsAddressNV"))
	gpDrawCommandsNV = uintptr(getProcAddr("glDrawCommandsNV"))
	gpDrawCommandsStatesAddressNV = uintptr(getProcAddr("glDrawCommandsStatesAddressNV"))
	gpDrawCommandsStatesNV = uintptr(getProcAddr("glDrawCommandsStatesNV"))
	gpDrawElementArrayAPPLE = uintptr(getProcAddr("glDrawElementArrayAPPLE"))
	gpDrawElementArrayATI = uintptr(getProcAddr("glDrawElementArrayATI"))
	gpDrawElements = uintptr(getProcAddr("glDrawElements"))
	if gpDrawElements == 0 {
		return errors.New("glDrawElements")
	}
	gpDrawElementsBaseVertex = uintptr(getProcAddr("glDrawElementsBaseVertex"))
	gpDrawElementsIndirect = uintptr(getProcAddr("glDrawElementsIndirect"))
	gpDrawElementsInstancedARB = uintptr(getProcAddr("glDrawElementsInstancedARB"))
	gpDrawElementsInstancedBaseInstance = uintptr(getProcAddr("glDrawElementsInstancedBaseInstance"))
	gpDrawElementsInstancedBaseVertex = uintptr(getProcAddr("glDrawElementsInstancedBaseVertex"))
	gpDrawElementsInstancedBaseVertexBaseInstance = uintptr(getProcAddr("glDrawElementsInstancedBaseVertexBaseInstance"))
	gpDrawElementsInstancedEXT = uintptr(getProcAddr("glDrawElementsInstancedEXT"))
	gpDrawMeshArraysSUN = uintptr(getProcAddr("glDrawMeshArraysSUN"))
	gpDrawPixels = uintptr(getProcAddr("glDrawPixels"))
	if gpDrawPixels == 0 {
		return errors.New("glDrawPixels")
	}
	gpDrawRangeElementArrayAPPLE = uintptr(getProcAddr("glDrawRangeElementArrayAPPLE"))
	gpDrawRangeElementArrayATI = uintptr(getProcAddr("glDrawRangeElementArrayATI"))
	gpDrawRangeElements = uintptr(getProcAddr("glDrawRangeElements"))
	if gpDrawRangeElements == 0 {
		return errors.New("glDrawRangeElements")
	}
	gpDrawRangeElementsBaseVertex = uintptr(getProcAddr("glDrawRangeElementsBaseVertex"))
	gpDrawRangeElementsEXT = uintptr(getProcAddr("glDrawRangeElementsEXT"))
	gpDrawTextureNV = uintptr(getProcAddr("glDrawTextureNV"))
	gpDrawTransformFeedback = uintptr(getProcAddr("glDrawTransformFeedback"))
	gpDrawTransformFeedbackInstanced = uintptr(getProcAddr("glDrawTransformFeedbackInstanced"))
	gpDrawTransformFeedbackNV = uintptr(getProcAddr("glDrawTransformFeedbackNV"))
	gpDrawTransformFeedbackStream = uintptr(getProcAddr("glDrawTransformFeedbackStream"))
	gpDrawTransformFeedbackStreamInstanced = uintptr(getProcAddr("glDrawTransformFeedbackStreamInstanced"))
	gpDrawVkImageNV = uintptr(getProcAddr("glDrawVkImageNV"))
	gpEGLImageTargetTexStorageEXT = uintptr(getProcAddr("glEGLImageTargetTexStorageEXT"))
	gpEGLImageTargetTextureStorageEXT = uintptr(getProcAddr("glEGLImageTargetTextureStorageEXT"))
	gpEdgeFlag = uintptr(getProcAddr("glEdgeFlag"))
	if gpEdgeFlag == 0 {
		return errors.New("glEdgeFlag")
	}
	gpEdgeFlagFormatNV = uintptr(getProcAddr("glEdgeFlagFormatNV"))
	gpEdgeFlagPointer = uintptr(getProcAddr("glEdgeFlagPointer"))
	if gpEdgeFlagPointer == 0 {
		return errors.New("glEdgeFlagPointer")
	}
	gpEdgeFlagPointerEXT = uintptr(getProcAddr("glEdgeFlagPointerEXT"))
	gpEdgeFlagPointerListIBM = uintptr(getProcAddr("glEdgeFlagPointerListIBM"))
	gpEdgeFlagv = uintptr(getProcAddr("glEdgeFlagv"))
	if gpEdgeFlagv == 0 {
		return errors.New("glEdgeFlagv")
	}
	gpElementPointerAPPLE = uintptr(getProcAddr("glElementPointerAPPLE"))
	gpElementPointerATI = uintptr(getProcAddr("glElementPointerATI"))
	gpEnable = uintptr(getProcAddr("glEnable"))
	if gpEnable == 0 {
		return errors.New("glEnable")
	}
	gpEnableClientState = uintptr(getProcAddr("glEnableClientState"))
	if gpEnableClientState == 0 {
		return errors.New("glEnableClientState")
	}
	gpEnableClientStateIndexedEXT = uintptr(getProcAddr("glEnableClientStateIndexedEXT"))
	gpEnableClientStateiEXT = uintptr(getProcAddr("glEnableClientStateiEXT"))
	gpEnableIndexedEXT = uintptr(getProcAddr("glEnableIndexedEXT"))
	gpEnableVariantClientStateEXT = uintptr(getProcAddr("glEnableVariantClientStateEXT"))
	gpEnableVertexArrayAttrib = uintptr(getProcAddr("glEnableVertexArrayAttrib"))
	gpEnableVertexArrayAttribEXT = uintptr(getProcAddr("glEnableVertexArrayAttribEXT"))
	gpEnableVertexArrayEXT = uintptr(getProcAddr("glEnableVertexArrayEXT"))
	gpEnableVertexAttribAPPLE = uintptr(getProcAddr("glEnableVertexAttribAPPLE"))
	gpEnableVertexAttribArray = uintptr(getProcAddr("glEnableVertexAttribArray"))
	if gpEnableVertexAttribArray == 0 {
		return errors.New("glEnableVertexAttribArray")
	}
	gpEnableVertexAttribArrayARB = uintptr(getProcAddr("glEnableVertexAttribArrayARB"))
	gpEnd = uintptr(getProcAddr("glEnd"))
	if gpEnd == 0 {
		return errors.New("glEnd")
	}
	gpEndConditionalRenderNV = uintptr(getProcAddr("glEndConditionalRenderNV"))
	gpEndConditionalRenderNVX = uintptr(getProcAddr("glEndConditionalRenderNVX"))
	gpEndFragmentShaderATI = uintptr(getProcAddr("glEndFragmentShaderATI"))
	gpEndList = uintptr(getProcAddr("glEndList"))
	if gpEndList == 0 {
		return errors.New("glEndList")
	}
	gpEndOcclusionQueryNV = uintptr(getProcAddr("glEndOcclusionQueryNV"))
	gpEndPerfMonitorAMD = uintptr(getProcAddr("glEndPerfMonitorAMD"))
	gpEndPerfQueryINTEL = uintptr(getProcAddr("glEndPerfQueryINTEL"))
	gpEndQuery = uintptr(getProcAddr("glEndQuery"))
	if gpEndQuery == 0 {
		return errors.New("glEndQuery")
	}
	gpEndQueryARB = uintptr(getProcAddr("glEndQueryARB"))
	gpEndQueryIndexed = uintptr(getProcAddr("glEndQueryIndexed"))
	gpEndTransformFeedbackEXT = uintptr(getProcAddr("glEndTransformFeedbackEXT"))
	gpEndTransformFeedbackNV = uintptr(getProcAddr("glEndTransformFeedbackNV"))
	gpEndVertexShaderEXT = uintptr(getProcAddr("glEndVertexShaderEXT"))
	gpEndVideoCaptureNV = uintptr(getProcAddr("glEndVideoCaptureNV"))
	gpEvalCoord1d = uintptr(getProcAddr("glEvalCoord1d"))
	if gpEvalCoord1d == 0 {
		return errors.New("glEvalCoord1d")
	}
	gpEvalCoord1dv = uintptr(getProcAddr("glEvalCoord1dv"))
	if gpEvalCoord1dv == 0 {
		return errors.New("glEvalCoord1dv")
	}
	gpEvalCoord1f = uintptr(getProcAddr("glEvalCoord1f"))
	if gpEvalCoord1f == 0 {
		return errors.New("glEvalCoord1f")
	}
	gpEvalCoord1fv = uintptr(getProcAddr("glEvalCoord1fv"))
	if gpEvalCoord1fv == 0 {
		return errors.New("glEvalCoord1fv")
	}
	gpEvalCoord1xOES = uintptr(getProcAddr("glEvalCoord1xOES"))
	gpEvalCoord1xvOES = uintptr(getProcAddr("glEvalCoord1xvOES"))
	gpEvalCoord2d = uintptr(getProcAddr("glEvalCoord2d"))
	if gpEvalCoord2d == 0 {
		return errors.New("glEvalCoord2d")
	}
	gpEvalCoord2dv = uintptr(getProcAddr("glEvalCoord2dv"))
	if gpEvalCoord2dv == 0 {
		return errors.New("glEvalCoord2dv")
	}
	gpEvalCoord2f = uintptr(getProcAddr("glEvalCoord2f"))
	if gpEvalCoord2f == 0 {
		return errors.New("glEvalCoord2f")
	}
	gpEvalCoord2fv = uintptr(getProcAddr("glEvalCoord2fv"))
	if gpEvalCoord2fv == 0 {
		return errors.New("glEvalCoord2fv")
	}
	gpEvalCoord2xOES = uintptr(getProcAddr("glEvalCoord2xOES"))
	gpEvalCoord2xvOES = uintptr(getProcAddr("glEvalCoord2xvOES"))
	gpEvalMapsNV = uintptr(getProcAddr("glEvalMapsNV"))
	gpEvalMesh1 = uintptr(getProcAddr("glEvalMesh1"))
	if gpEvalMesh1 == 0 {
		return errors.New("glEvalMesh1")
	}
	gpEvalMesh2 = uintptr(getProcAddr("glEvalMesh2"))
	if gpEvalMesh2 == 0 {
		return errors.New("glEvalMesh2")
	}
	gpEvalPoint1 = uintptr(getProcAddr("glEvalPoint1"))
	if gpEvalPoint1 == 0 {
		return errors.New("glEvalPoint1")
	}
	gpEvalPoint2 = uintptr(getProcAddr("glEvalPoint2"))
	if gpEvalPoint2 == 0 {
		return errors.New("glEvalPoint2")
	}
	gpEvaluateDepthValuesARB = uintptr(getProcAddr("glEvaluateDepthValuesARB"))
	gpExecuteProgramNV = uintptr(getProcAddr("glExecuteProgramNV"))
	gpExtractComponentEXT = uintptr(getProcAddr("glExtractComponentEXT"))
	gpFeedbackBuffer = uintptr(getProcAddr("glFeedbackBuffer"))
	if gpFeedbackBuffer == 0 {
		return errors.New("glFeedbackBuffer")
	}
	gpFeedbackBufferxOES = uintptr(getProcAddr("glFeedbackBufferxOES"))
	gpFenceSync = uintptr(getProcAddr("glFenceSync"))
	gpFinalCombinerInputNV = uintptr(getProcAddr("glFinalCombinerInputNV"))
	gpFinish = uintptr(getProcAddr("glFinish"))
	if gpFinish == 0 {
		return errors.New("glFinish")
	}
	gpFinishAsyncSGIX = uintptr(getProcAddr("glFinishAsyncSGIX"))
	gpFinishFenceAPPLE = uintptr(getProcAddr("glFinishFenceAPPLE"))
	gpFinishFenceNV = uintptr(getProcAddr("glFinishFenceNV"))
	gpFinishObjectAPPLE = uintptr(getProcAddr("glFinishObjectAPPLE"))
	gpFinishTextureSUNX = uintptr(getProcAddr("glFinishTextureSUNX"))
	gpFlush = uintptr(getProcAddr("glFlush"))
	if gpFlush == 0 {
		return errors.New("glFlush")
	}
	gpFlushMappedBufferRange = uintptr(getProcAddr("glFlushMappedBufferRange"))
	gpFlushMappedBufferRangeAPPLE = uintptr(getProcAddr("glFlushMappedBufferRangeAPPLE"))
	gpFlushMappedNamedBufferRange = uintptr(getProcAddr("glFlushMappedNamedBufferRange"))
	gpFlushMappedNamedBufferRangeEXT = uintptr(getProcAddr("glFlushMappedNamedBufferRangeEXT"))
	gpFlushPixelDataRangeNV = uintptr(getProcAddr("glFlushPixelDataRangeNV"))
	gpFlushRasterSGIX = uintptr(getProcAddr("glFlushRasterSGIX"))
	gpFlushStaticDataIBM = uintptr(getProcAddr("glFlushStaticDataIBM"))
	gpFlushVertexArrayRangeAPPLE = uintptr(getProcAddr("glFlushVertexArrayRangeAPPLE"))
	gpFlushVertexArrayRangeNV = uintptr(getProcAddr("glFlushVertexArrayRangeNV"))
	gpFogCoordFormatNV = uintptr(getProcAddr("glFogCoordFormatNV"))
	gpFogCoordPointer = uintptr(getProcAddr("glFogCoordPointer"))
	if gpFogCoordPointer == 0 {
		return errors.New("glFogCoordPointer")
	}
	gpFogCoordPointerEXT = uintptr(getProcAddr("glFogCoordPointerEXT"))
	gpFogCoordPointerListIBM = uintptr(getProcAddr("glFogCoordPointerListIBM"))
	gpFogCoordd = uintptr(getProcAddr("glFogCoordd"))
	if gpFogCoordd == 0 {
		return errors.New("glFogCoordd")
	}
	gpFogCoorddEXT = uintptr(getProcAddr("glFogCoorddEXT"))
	gpFogCoorddv = uintptr(getProcAddr("glFogCoorddv"))
	if gpFogCoorddv == 0 {
		return errors.New("glFogCoorddv")
	}
	gpFogCoorddvEXT = uintptr(getProcAddr("glFogCoorddvEXT"))
	gpFogCoordf = uintptr(getProcAddr("glFogCoordf"))
	if gpFogCoordf == 0 {
		return errors.New("glFogCoordf")
	}
	gpFogCoordfEXT = uintptr(getProcAddr("glFogCoordfEXT"))
	gpFogCoordfv = uintptr(getProcAddr("glFogCoordfv"))
	if gpFogCoordfv == 0 {
		return errors.New("glFogCoordfv")
	}
	gpFogCoordfvEXT = uintptr(getProcAddr("glFogCoordfvEXT"))
	gpFogCoordhNV = uintptr(getProcAddr("glFogCoordhNV"))
	gpFogCoordhvNV = uintptr(getProcAddr("glFogCoordhvNV"))
	gpFogFuncSGIS = uintptr(getProcAddr("glFogFuncSGIS"))
	gpFogf = uintptr(getProcAddr("glFogf"))
	if gpFogf == 0 {
		return errors.New("glFogf")
	}
	gpFogfv = uintptr(getProcAddr("glFogfv"))
	if gpFogfv == 0 {
		return errors.New("glFogfv")
	}
	gpFogi = uintptr(getProcAddr("glFogi"))
	if gpFogi == 0 {
		return errors.New("glFogi")
	}
	gpFogiv = uintptr(getProcAddr("glFogiv"))
	if gpFogiv == 0 {
		return errors.New("glFogiv")
	}
	gpFogxOES = uintptr(getProcAddr("glFogxOES"))
	gpFogxvOES = uintptr(getProcAddr("glFogxvOES"))
	gpFragmentColorMaterialSGIX = uintptr(getProcAddr("glFragmentColorMaterialSGIX"))
	gpFragmentCoverageColorNV = uintptr(getProcAddr("glFragmentCoverageColorNV"))
	gpFragmentLightModelfSGIX = uintptr(getProcAddr("glFragmentLightModelfSGIX"))
	gpFragmentLightModelfvSGIX = uintptr(getProcAddr("glFragmentLightModelfvSGIX"))
	gpFragmentLightModeliSGIX = uintptr(getProcAddr("glFragmentLightModeliSGIX"))
	gpFragmentLightModelivSGIX = uintptr(getProcAddr("glFragmentLightModelivSGIX"))
	gpFragmentLightfSGIX = uintptr(getProcAddr("glFragmentLightfSGIX"))
	gpFragmentLightfvSGIX = uintptr(getProcAddr("glFragmentLightfvSGIX"))
	gpFragmentLightiSGIX = uintptr(getProcAddr("glFragmentLightiSGIX"))
	gpFragmentLightivSGIX = uintptr(getProcAddr("glFragmentLightivSGIX"))
	gpFragmentMaterialfSGIX = uintptr(getProcAddr("glFragmentMaterialfSGIX"))
	gpFragmentMaterialfvSGIX = uintptr(getProcAddr("glFragmentMaterialfvSGIX"))
	gpFragmentMaterialiSGIX = uintptr(getProcAddr("glFragmentMaterialiSGIX"))
	gpFragmentMaterialivSGIX = uintptr(getProcAddr("glFragmentMaterialivSGIX"))
	gpFrameTerminatorGREMEDY = uintptr(getProcAddr("glFrameTerminatorGREMEDY"))
	gpFrameZoomSGIX = uintptr(getProcAddr("glFrameZoomSGIX"))
	gpFramebufferDrawBufferEXT = uintptr(getProcAddr("glFramebufferDrawBufferEXT"))
	gpFramebufferDrawBuffersEXT = uintptr(getProcAddr("glFramebufferDrawBuffersEXT"))
	gpFramebufferFetchBarrierEXT = uintptr(getProcAddr("glFramebufferFetchBarrierEXT"))
	gpFramebufferParameteri = uintptr(getProcAddr("glFramebufferParameteri"))
	gpFramebufferReadBufferEXT = uintptr(getProcAddr("glFramebufferReadBufferEXT"))
	gpFramebufferRenderbuffer = uintptr(getProcAddr("glFramebufferRenderbuffer"))
	gpFramebufferRenderbufferEXT = uintptr(getProcAddr("glFramebufferRenderbufferEXT"))
	gpFramebufferSampleLocationsfvARB = uintptr(getProcAddr("glFramebufferSampleLocationsfvARB"))
	gpFramebufferSampleLocationsfvNV = uintptr(getProcAddr("glFramebufferSampleLocationsfvNV"))
	gpFramebufferSamplePositionsfvAMD = uintptr(getProcAddr("glFramebufferSamplePositionsfvAMD"))
	gpFramebufferTexture1D = uintptr(getProcAddr("glFramebufferTexture1D"))
	gpFramebufferTexture1DEXT = uintptr(getProcAddr("glFramebufferTexture1DEXT"))
	gpFramebufferTexture2D = uintptr(getProcAddr("glFramebufferTexture2D"))
	gpFramebufferTexture2DEXT = uintptr(getProcAddr("glFramebufferTexture2DEXT"))
	gpFramebufferTexture3D = uintptr(getProcAddr("glFramebufferTexture3D"))
	gpFramebufferTexture3DEXT = uintptr(getProcAddr("glFramebufferTexture3DEXT"))
	gpFramebufferTextureARB = uintptr(getProcAddr("glFramebufferTextureARB"))
	gpFramebufferTextureEXT = uintptr(getProcAddr("glFramebufferTextureEXT"))
	gpFramebufferTextureFaceARB = uintptr(getProcAddr("glFramebufferTextureFaceARB"))
	gpFramebufferTextureFaceEXT = uintptr(getProcAddr("glFramebufferTextureFaceEXT"))
	gpFramebufferTextureLayer = uintptr(getProcAddr("glFramebufferTextureLayer"))
	gpFramebufferTextureLayerARB = uintptr(getProcAddr("glFramebufferTextureLayerARB"))
	gpFramebufferTextureLayerEXT = uintptr(getProcAddr("glFramebufferTextureLayerEXT"))
	gpFramebufferTextureMultiviewOVR = uintptr(getProcAddr("glFramebufferTextureMultiviewOVR"))
	gpFreeObjectBufferATI = uintptr(getProcAddr("glFreeObjectBufferATI"))
	gpFrontFace = uintptr(getProcAddr("glFrontFace"))
	if gpFrontFace == 0 {
		return errors.New("glFrontFace")
	}
	gpFrustum = uintptr(getProcAddr("glFrustum"))
	if gpFrustum == 0 {
		return errors.New("glFrustum")
	}
	gpFrustumfOES = uintptr(getProcAddr("glFrustumfOES"))
	gpFrustumxOES = uintptr(getProcAddr("glFrustumxOES"))
	gpGenAsyncMarkersSGIX = uintptr(getProcAddr("glGenAsyncMarkersSGIX"))
	gpGenBuffers = uintptr(getProcAddr("glGenBuffers"))
	if gpGenBuffers == 0 {
		return errors.New("glGenBuffers")
	}
	gpGenBuffersARB = uintptr(getProcAddr("glGenBuffersARB"))
	gpGenFencesAPPLE = uintptr(getProcAddr("glGenFencesAPPLE"))
	gpGenFencesNV = uintptr(getProcAddr("glGenFencesNV"))
	gpGenFragmentShadersATI = uintptr(getProcAddr("glGenFragmentShadersATI"))
	gpGenFramebuffers = uintptr(getProcAddr("glGenFramebuffers"))
	gpGenFramebuffersEXT = uintptr(getProcAddr("glGenFramebuffersEXT"))
	gpGenLists = uintptr(getProcAddr("glGenLists"))
	if gpGenLists == 0 {
		return errors.New("glGenLists")
	}
	gpGenNamesAMD = uintptr(getProcAddr("glGenNamesAMD"))
	gpGenOcclusionQueriesNV = uintptr(getProcAddr("glGenOcclusionQueriesNV"))
	gpGenPathsNV = uintptr(getProcAddr("glGenPathsNV"))
	gpGenPerfMonitorsAMD = uintptr(getProcAddr("glGenPerfMonitorsAMD"))
	gpGenProgramPipelines = uintptr(getProcAddr("glGenProgramPipelines"))
	gpGenProgramPipelinesEXT = uintptr(getProcAddr("glGenProgramPipelinesEXT"))
	gpGenProgramsARB = uintptr(getProcAddr("glGenProgramsARB"))
	gpGenProgramsNV = uintptr(getProcAddr("glGenProgramsNV"))
	gpGenQueries = uintptr(getProcAddr("glGenQueries"))
	if gpGenQueries == 0 {
		return errors.New("glGenQueries")
	}
	gpGenQueriesARB = uintptr(getProcAddr("glGenQueriesARB"))
	gpGenQueryResourceTagNV = uintptr(getProcAddr("glGenQueryResourceTagNV"))
	gpGenRenderbuffers = uintptr(getProcAddr("glGenRenderbuffers"))
	gpGenRenderbuffersEXT = uintptr(getProcAddr("glGenRenderbuffersEXT"))
	gpGenSamplers = uintptr(getProcAddr("glGenSamplers"))
	gpGenSemaphoresEXT = uintptr(getProcAddr("glGenSemaphoresEXT"))
	gpGenSymbolsEXT = uintptr(getProcAddr("glGenSymbolsEXT"))
	gpGenTextures = uintptr(getProcAddr("glGenTextures"))
	if gpGenTextures == 0 {
		return errors.New("glGenTextures")
	}
	gpGenTexturesEXT = uintptr(getProcAddr("glGenTexturesEXT"))
	gpGenTransformFeedbacks = uintptr(getProcAddr("glGenTransformFeedbacks"))
	gpGenTransformFeedbacksNV = uintptr(getProcAddr("glGenTransformFeedbacksNV"))
	gpGenVertexArrays = uintptr(getProcAddr("glGenVertexArrays"))
	gpGenVertexArraysAPPLE = uintptr(getProcAddr("glGenVertexArraysAPPLE"))
	gpGenVertexShadersEXT = uintptr(getProcAddr("glGenVertexShadersEXT"))
	gpGenerateMipmap = uintptr(getProcAddr("glGenerateMipmap"))
	gpGenerateMipmapEXT = uintptr(getProcAddr("glGenerateMipmapEXT"))
	gpGenerateMultiTexMipmapEXT = uintptr(getProcAddr("glGenerateMultiTexMipmapEXT"))
	gpGenerateTextureMipmap = uintptr(getProcAddr("glGenerateTextureMipmap"))
	gpGenerateTextureMipmapEXT = uintptr(getProcAddr("glGenerateTextureMipmapEXT"))
	gpGetActiveAtomicCounterBufferiv = uintptr(getProcAddr("glGetActiveAtomicCounterBufferiv"))
	gpGetActiveAttrib = uintptr(getProcAddr("glGetActiveAttrib"))
	if gpGetActiveAttrib == 0 {
		return errors.New("glGetActiveAttrib")
	}
	gpGetActiveAttribARB = uintptr(getProcAddr("glGetActiveAttribARB"))
	gpGetActiveSubroutineName = uintptr(getProcAddr("glGetActiveSubroutineName"))
	gpGetActiveSubroutineUniformName = uintptr(getProcAddr("glGetActiveSubroutineUniformName"))
	gpGetActiveSubroutineUniformiv = uintptr(getProcAddr("glGetActiveSubroutineUniformiv"))
	gpGetActiveUniform = uintptr(getProcAddr("glGetActiveUniform"))
	if gpGetActiveUniform == 0 {
		return errors.New("glGetActiveUniform")
	}
	gpGetActiveUniformARB = uintptr(getProcAddr("glGetActiveUniformARB"))
	gpGetActiveUniformBlockName = uintptr(getProcAddr("glGetActiveUniformBlockName"))
	gpGetActiveUniformBlockiv = uintptr(getProcAddr("glGetActiveUniformBlockiv"))
	gpGetActiveUniformName = uintptr(getProcAddr("glGetActiveUniformName"))
	gpGetActiveUniformsiv = uintptr(getProcAddr("glGetActiveUniformsiv"))
	gpGetActiveVaryingNV = uintptr(getProcAddr("glGetActiveVaryingNV"))
	gpGetArrayObjectfvATI = uintptr(getProcAddr("glGetArrayObjectfvATI"))
	gpGetArrayObjectivATI = uintptr(getProcAddr("glGetArrayObjectivATI"))
	gpGetAttachedObjectsARB = uintptr(getProcAddr("glGetAttachedObjectsARB"))
	gpGetAttachedShaders = uintptr(getProcAddr("glGetAttachedShaders"))
	if gpGetAttachedShaders == 0 {
		return errors.New("glGetAttachedShaders")
	}
	gpGetAttribLocation = uintptr(getProcAddr("glGetAttribLocation"))
	if gpGetAttribLocation == 0 {
		return errors.New("glGetAttribLocation")
	}
	gpGetAttribLocationARB = uintptr(getProcAddr("glGetAttribLocationARB"))
	gpGetBooleanIndexedvEXT = uintptr(getProcAddr("glGetBooleanIndexedvEXT"))
	gpGetBooleanv = uintptr(getProcAddr("glGetBooleanv"))
	if gpGetBooleanv == 0 {
		return errors.New("glGetBooleanv")
	}
	gpGetBufferParameteriv = uintptr(getProcAddr("glGetBufferParameteriv"))
	if gpGetBufferParameteriv == 0 {
		return errors.New("glGetBufferParameteriv")
	}
	gpGetBufferParameterivARB = uintptr(getProcAddr("glGetBufferParameterivARB"))
	gpGetBufferParameterui64vNV = uintptr(getProcAddr("glGetBufferParameterui64vNV"))
	gpGetBufferPointerv = uintptr(getProcAddr("glGetBufferPointerv"))
	if gpGetBufferPointerv == 0 {
		return errors.New("glGetBufferPointerv")
	}
	gpGetBufferPointervARB = uintptr(getProcAddr("glGetBufferPointervARB"))
	gpGetBufferSubData = uintptr(getProcAddr("glGetBufferSubData"))
	if gpGetBufferSubData == 0 {
		return errors.New("glGetBufferSubData")
	}
	gpGetBufferSubDataARB = uintptr(getProcAddr("glGetBufferSubDataARB"))
	gpGetClipPlane = uintptr(getProcAddr("glGetClipPlane"))
	if gpGetClipPlane == 0 {
		return errors.New("glGetClipPlane")
	}
	gpGetClipPlanefOES = uintptr(getProcAddr("glGetClipPlanefOES"))
	gpGetClipPlanexOES = uintptr(getProcAddr("glGetClipPlanexOES"))
	gpGetColorTableEXT = uintptr(getProcAddr("glGetColorTableEXT"))
	gpGetColorTableParameterfvEXT = uintptr(getProcAddr("glGetColorTableParameterfvEXT"))
	gpGetColorTableParameterfvSGI = uintptr(getProcAddr("glGetColorTableParameterfvSGI"))
	gpGetColorTableParameterivEXT = uintptr(getProcAddr("glGetColorTableParameterivEXT"))
	gpGetColorTableParameterivSGI = uintptr(getProcAddr("glGetColorTableParameterivSGI"))
	gpGetColorTableSGI = uintptr(getProcAddr("glGetColorTableSGI"))
	gpGetCombinerInputParameterfvNV = uintptr(getProcAddr("glGetCombinerInputParameterfvNV"))
	gpGetCombinerInputParameterivNV = uintptr(getProcAddr("glGetCombinerInputParameterivNV"))
	gpGetCombinerOutputParameterfvNV = uintptr(getProcAddr("glGetCombinerOutputParameterfvNV"))
	gpGetCombinerOutputParameterivNV = uintptr(getProcAddr("glGetCombinerOutputParameterivNV"))
	gpGetCombinerStageParameterfvNV = uintptr(getProcAddr("glGetCombinerStageParameterfvNV"))
	gpGetCommandHeaderNV = uintptr(getProcAddr("glGetCommandHeaderNV"))
	gpGetCompressedMultiTexImageEXT = uintptr(getProcAddr("glGetCompressedMultiTexImageEXT"))
	gpGetCompressedTexImage = uintptr(getProcAddr("glGetCompressedTexImage"))
	if gpGetCompressedTexImage == 0 {
		return errors.New("glGetCompressedTexImage")
	}
	gpGetCompressedTexImageARB = uintptr(getProcAddr("glGetCompressedTexImageARB"))
	gpGetCompressedTextureImage = uintptr(getProcAddr("glGetCompressedTextureImage"))
	gpGetCompressedTextureImageEXT = uintptr(getProcAddr("glGetCompressedTextureImageEXT"))
	gpGetCompressedTextureSubImage = uintptr(getProcAddr("glGetCompressedTextureSubImage"))
	gpGetConvolutionFilterEXT = uintptr(getProcAddr("glGetConvolutionFilterEXT"))
	gpGetConvolutionParameterfvEXT = uintptr(getProcAddr("glGetConvolutionParameterfvEXT"))
	gpGetConvolutionParameterivEXT = uintptr(getProcAddr("glGetConvolutionParameterivEXT"))
	gpGetConvolutionParameterxvOES = uintptr(getProcAddr("glGetConvolutionParameterxvOES"))
	gpGetCoverageModulationTableNV = uintptr(getProcAddr("glGetCoverageModulationTableNV"))
	gpGetDebugMessageLog = uintptr(getProcAddr("glGetDebugMessageLog"))
	gpGetDebugMessageLogAMD = uintptr(getProcAddr("glGetDebugMessageLogAMD"))
	gpGetDebugMessageLogARB = uintptr(getProcAddr("glGetDebugMessageLogARB"))
	gpGetDebugMessageLogKHR = uintptr(getProcAddr("glGetDebugMessageLogKHR"))
	gpGetDetailTexFuncSGIS = uintptr(getProcAddr("glGetDetailTexFuncSGIS"))
	gpGetDoubleIndexedvEXT = uintptr(getProcAddr("glGetDoubleIndexedvEXT"))
	gpGetDoublei_v = uintptr(getProcAddr("glGetDoublei_v"))
	gpGetDoublei_vEXT = uintptr(getProcAddr("glGetDoublei_vEXT"))
	gpGetDoublev = uintptr(getProcAddr("glGetDoublev"))
	if gpGetDoublev == 0 {
		return errors.New("glGetDoublev")
	}
	gpGetError = uintptr(getProcAddr("glGetError"))
	if gpGetError == 0 {
		return errors.New("glGetError")
	}
	gpGetFenceivNV = uintptr(getProcAddr("glGetFenceivNV"))
	gpGetFinalCombinerInputParameterfvNV = uintptr(getProcAddr("glGetFinalCombinerInputParameterfvNV"))
	gpGetFinalCombinerInputParameterivNV = uintptr(getProcAddr("glGetFinalCombinerInputParameterivNV"))
	gpGetFirstPerfQueryIdINTEL = uintptr(getProcAddr("glGetFirstPerfQueryIdINTEL"))
	gpGetFixedvOES = uintptr(getProcAddr("glGetFixedvOES"))
	gpGetFloatIndexedvEXT = uintptr(getProcAddr("glGetFloatIndexedvEXT"))
	gpGetFloati_v = uintptr(getProcAddr("glGetFloati_v"))
	gpGetFloati_vEXT = uintptr(getProcAddr("glGetFloati_vEXT"))
	gpGetFloatv = uintptr(getProcAddr("glGetFloatv"))
	if gpGetFloatv == 0 {
		return errors.New("glGetFloatv")
	}
	gpGetFogFuncSGIS = uintptr(getProcAddr("glGetFogFuncSGIS"))
	gpGetFragDataIndex = uintptr(getProcAddr("glGetFragDataIndex"))
	gpGetFragDataLocationEXT = uintptr(getProcAddr("glGetFragDataLocationEXT"))
	gpGetFragmentLightfvSGIX = uintptr(getProcAddr("glGetFragmentLightfvSGIX"))
	gpGetFragmentLightivSGIX = uintptr(getProcAddr("glGetFragmentLightivSGIX"))
	gpGetFragmentMaterialfvSGIX = uintptr(getProcAddr("glGetFragmentMaterialfvSGIX"))
	gpGetFragmentMaterialivSGIX = uintptr(getProcAddr("glGetFragmentMaterialivSGIX"))
	gpGetFramebufferAttachmentParameteriv = uintptr(getProcAddr("glGetFramebufferAttachmentParameteriv"))
	gpGetFramebufferAttachmentParameterivEXT = uintptr(getProcAddr("glGetFramebufferAttachmentParameterivEXT"))
	gpGetFramebufferParameterfvAMD = uintptr(getProcAddr("glGetFramebufferParameterfvAMD"))
	gpGetFramebufferParameteriv = uintptr(getProcAddr("glGetFramebufferParameteriv"))
	gpGetFramebufferParameterivEXT = uintptr(getProcAddr("glGetFramebufferParameterivEXT"))
	gpGetGraphicsResetStatus = uintptr(getProcAddr("glGetGraphicsResetStatus"))
	gpGetGraphicsResetStatusARB = uintptr(getProcAddr("glGetGraphicsResetStatusARB"))
	gpGetGraphicsResetStatusKHR = uintptr(getProcAddr("glGetGraphicsResetStatusKHR"))
	gpGetHandleARB = uintptr(getProcAddr("glGetHandleARB"))
	gpGetHistogramEXT = uintptr(getProcAddr("glGetHistogramEXT"))
	gpGetHistogramParameterfvEXT = uintptr(getProcAddr("glGetHistogramParameterfvEXT"))
	gpGetHistogramParameterivEXT = uintptr(getProcAddr("glGetHistogramParameterivEXT"))
	gpGetHistogramParameterxvOES = uintptr(getProcAddr("glGetHistogramParameterxvOES"))
	gpGetImageHandleARB = uintptr(getProcAddr("glGetImageHandleARB"))
	gpGetImageHandleNV = uintptr(getProcAddr("glGetImageHandleNV"))
	gpGetImageTransformParameterfvHP = uintptr(getProcAddr("glGetImageTransformParameterfvHP"))
	gpGetImageTransformParameterivHP = uintptr(getProcAddr("glGetImageTransformParameterivHP"))
	gpGetInfoLogARB = uintptr(getProcAddr("glGetInfoLogARB"))
	gpGetInstrumentsSGIX = uintptr(getProcAddr("glGetInstrumentsSGIX"))
	gpGetInteger64v = uintptr(getProcAddr("glGetInteger64v"))
	gpGetIntegerIndexedvEXT = uintptr(getProcAddr("glGetIntegerIndexedvEXT"))
	gpGetIntegeri_v = uintptr(getProcAddr("glGetIntegeri_v"))
	gpGetIntegerui64i_vNV = uintptr(getProcAddr("glGetIntegerui64i_vNV"))
	gpGetIntegerui64vNV = uintptr(getProcAddr("glGetIntegerui64vNV"))
	gpGetIntegerv = uintptr(getProcAddr("glGetIntegerv"))
	if gpGetIntegerv == 0 {
		return errors.New("glGetIntegerv")
	}
	gpGetInternalformatSampleivNV = uintptr(getProcAddr("glGetInternalformatSampleivNV"))
	gpGetInternalformati64v = uintptr(getProcAddr("glGetInternalformati64v"))
	gpGetInternalformativ = uintptr(getProcAddr("glGetInternalformativ"))
	gpGetInvariantBooleanvEXT = uintptr(getProcAddr("glGetInvariantBooleanvEXT"))
	gpGetInvariantFloatvEXT = uintptr(getProcAddr("glGetInvariantFloatvEXT"))
	gpGetInvariantIntegervEXT = uintptr(getProcAddr("glGetInvariantIntegervEXT"))
	gpGetLightfv = uintptr(getProcAddr("glGetLightfv"))
	if gpGetLightfv == 0 {
		return errors.New("glGetLightfv")
	}
	gpGetLightiv = uintptr(getProcAddr("glGetLightiv"))
	if gpGetLightiv == 0 {
		return errors.New("glGetLightiv")
	}
	gpGetLightxOES = uintptr(getProcAddr("glGetLightxOES"))
	gpGetLightxvOES = uintptr(getProcAddr("glGetLightxvOES"))
	gpGetListParameterfvSGIX = uintptr(getProcAddr("glGetListParameterfvSGIX"))
	gpGetListParameterivSGIX = uintptr(getProcAddr("glGetListParameterivSGIX"))
	gpGetLocalConstantBooleanvEXT = uintptr(getProcAddr("glGetLocalConstantBooleanvEXT"))
	gpGetLocalConstantFloatvEXT = uintptr(getProcAddr("glGetLocalConstantFloatvEXT"))
	gpGetLocalConstantIntegervEXT = uintptr(getProcAddr("glGetLocalConstantIntegervEXT"))
	gpGetMapAttribParameterfvNV = uintptr(getProcAddr("glGetMapAttribParameterfvNV"))
	gpGetMapAttribParameterivNV = uintptr(getProcAddr("glGetMapAttribParameterivNV"))
	gpGetMapControlPointsNV = uintptr(getProcAddr("glGetMapControlPointsNV"))
	gpGetMapParameterfvNV = uintptr(getProcAddr("glGetMapParameterfvNV"))
	gpGetMapParameterivNV = uintptr(getProcAddr("glGetMapParameterivNV"))
	gpGetMapdv = uintptr(getProcAddr("glGetMapdv"))
	if gpGetMapdv == 0 {
		return errors.New("glGetMapdv")
	}
	gpGetMapfv = uintptr(getProcAddr("glGetMapfv"))
	if gpGetMapfv == 0 {
		return errors.New("glGetMapfv")
	}
	gpGetMapiv = uintptr(getProcAddr("glGetMapiv"))
	if gpGetMapiv == 0 {
		return errors.New("glGetMapiv")
	}
	gpGetMapxvOES = uintptr(getProcAddr("glGetMapxvOES"))
	gpGetMaterialfv = uintptr(getProcAddr("glGetMaterialfv"))
	if gpGetMaterialfv == 0 {
		return errors.New("glGetMaterialfv")
	}
	gpGetMaterialiv = uintptr(getProcAddr("glGetMaterialiv"))
	if gpGetMaterialiv == 0 {
		return errors.New("glGetMaterialiv")
	}
	gpGetMaterialxOES = uintptr(getProcAddr("glGetMaterialxOES"))
	gpGetMaterialxvOES = uintptr(getProcAddr("glGetMaterialxvOES"))
	gpGetMemoryObjectParameterivEXT = uintptr(getProcAddr("glGetMemoryObjectParameterivEXT"))
	gpGetMinmaxEXT = uintptr(getProcAddr("glGetMinmaxEXT"))
	gpGetMinmaxParameterfvEXT = uintptr(getProcAddr("glGetMinmaxParameterfvEXT"))
	gpGetMinmaxParameterivEXT = uintptr(getProcAddr("glGetMinmaxParameterivEXT"))
	gpGetMultiTexEnvfvEXT = uintptr(getProcAddr("glGetMultiTexEnvfvEXT"))
	gpGetMultiTexEnvivEXT = uintptr(getProcAddr("glGetMultiTexEnvivEXT"))
	gpGetMultiTexGendvEXT = uintptr(getProcAddr("glGetMultiTexGendvEXT"))
	gpGetMultiTexGenfvEXT = uintptr(getProcAddr("glGetMultiTexGenfvEXT"))
	gpGetMultiTexGenivEXT = uintptr(getProcAddr("glGetMultiTexGenivEXT"))
	gpGetMultiTexImageEXT = uintptr(getProcAddr("glGetMultiTexImageEXT"))
	gpGetMultiTexLevelParameterfvEXT = uintptr(getProcAddr("glGetMultiTexLevelParameterfvEXT"))
	gpGetMultiTexLevelParameterivEXT = uintptr(getProcAddr("glGetMultiTexLevelParameterivEXT"))
	gpGetMultiTexParameterIivEXT = uintptr(getProcAddr("glGetMultiTexParameterIivEXT"))
	gpGetMultiTexParameterIuivEXT = uintptr(getProcAddr("glGetMultiTexParameterIuivEXT"))
	gpGetMultiTexParameterfvEXT = uintptr(getProcAddr("glGetMultiTexParameterfvEXT"))
	gpGetMultiTexParameterivEXT = uintptr(getProcAddr("glGetMultiTexParameterivEXT"))
	gpGetMultisamplefv = uintptr(getProcAddr("glGetMultisamplefv"))
	gpGetMultisamplefvNV = uintptr(getProcAddr("glGetMultisamplefvNV"))
	gpGetNamedBufferParameteri64v = uintptr(getProcAddr("glGetNamedBufferParameteri64v"))
	gpGetNamedBufferParameteriv = uintptr(getProcAddr("glGetNamedBufferParameteriv"))
	gpGetNamedBufferParameterivEXT = uintptr(getProcAddr("glGetNamedBufferParameterivEXT"))
	gpGetNamedBufferParameterui64vNV = uintptr(getProcAddr("glGetNamedBufferParameterui64vNV"))
	gpGetNamedBufferPointerv = uintptr(getProcAddr("glGetNamedBufferPointerv"))
	gpGetNamedBufferPointervEXT = uintptr(getProcAddr("glGetNamedBufferPointervEXT"))
	gpGetNamedBufferSubData = uintptr(getProcAddr("glGetNamedBufferSubData"))
	gpGetNamedBufferSubDataEXT = uintptr(getProcAddr("glGetNamedBufferSubDataEXT"))
	gpGetNamedFramebufferAttachmentParameteriv = uintptr(getProcAddr("glGetNamedFramebufferAttachmentParameteriv"))
	gpGetNamedFramebufferAttachmentParameterivEXT = uintptr(getProcAddr("glGetNamedFramebufferAttachmentParameterivEXT"))
	gpGetNamedFramebufferParameterfvAMD = uintptr(getProcAddr("glGetNamedFramebufferParameterfvAMD"))
	gpGetNamedFramebufferParameteriv = uintptr(getProcAddr("glGetNamedFramebufferParameteriv"))
	gpGetNamedFramebufferParameterivEXT = uintptr(getProcAddr("glGetNamedFramebufferParameterivEXT"))
	gpGetNamedProgramLocalParameterIivEXT = uintptr(getProcAddr("glGetNamedProgramLocalParameterIivEXT"))
	gpGetNamedProgramLocalParameterIuivEXT = uintptr(getProcAddr("glGetNamedProgramLocalParameterIuivEXT"))
	gpGetNamedProgramLocalParameterdvEXT = uintptr(getProcAddr("glGetNamedProgramLocalParameterdvEXT"))
	gpGetNamedProgramLocalParameterfvEXT = uintptr(getProcAddr("glGetNamedProgramLocalParameterfvEXT"))
	gpGetNamedProgramStringEXT = uintptr(getProcAddr("glGetNamedProgramStringEXT"))
	gpGetNamedProgramivEXT = uintptr(getProcAddr("glGetNamedProgramivEXT"))
	gpGetNamedRenderbufferParameteriv = uintptr(getProcAddr("glGetNamedRenderbufferParameteriv"))
	gpGetNamedRenderbufferParameterivEXT = uintptr(getProcAddr("glGetNamedRenderbufferParameterivEXT"))
	gpGetNamedStringARB = uintptr(getProcAddr("glGetNamedStringARB"))
	gpGetNamedStringivARB = uintptr(getProcAddr("glGetNamedStringivARB"))
	gpGetNextPerfQueryIdINTEL = uintptr(getProcAddr("glGetNextPerfQueryIdINTEL"))
	gpGetObjectBufferfvATI = uintptr(getProcAddr("glGetObjectBufferfvATI"))
	gpGetObjectBufferivATI = uintptr(getProcAddr("glGetObjectBufferivATI"))
	gpGetObjectLabel = uintptr(getProcAddr("glGetObjectLabel"))
	gpGetObjectLabelEXT = uintptr(getProcAddr("glGetObjectLabelEXT"))
	gpGetObjectLabelKHR = uintptr(getProcAddr("glGetObjectLabelKHR"))
	gpGetObjectParameterfvARB = uintptr(getProcAddr("glGetObjectParameterfvARB"))
	gpGetObjectParameterivAPPLE = uintptr(getProcAddr("glGetObjectParameterivAPPLE"))
	gpGetObjectParameterivARB = uintptr(getProcAddr("glGetObjectParameterivARB"))
	gpGetObjectPtrLabel = uintptr(getProcAddr("glGetObjectPtrLabel"))
	gpGetObjectPtrLabelKHR = uintptr(getProcAddr("glGetObjectPtrLabelKHR"))
	gpGetOcclusionQueryivNV = uintptr(getProcAddr("glGetOcclusionQueryivNV"))
	gpGetOcclusionQueryuivNV = uintptr(getProcAddr("glGetOcclusionQueryuivNV"))
	gpGetPathCommandsNV = uintptr(getProcAddr("glGetPathCommandsNV"))
	gpGetPathCoordsNV = uintptr(getProcAddr("glGetPathCoordsNV"))
	gpGetPathDashArrayNV = uintptr(getProcAddr("glGetPathDashArrayNV"))
	gpGetPathLengthNV = uintptr(getProcAddr("glGetPathLengthNV"))
	gpGetPathMetricRangeNV = uintptr(getProcAddr("glGetPathMetricRangeNV"))
	gpGetPathMetricsNV = uintptr(getProcAddr("glGetPathMetricsNV"))
	gpGetPathParameterfvNV = uintptr(getProcAddr("glGetPathParameterfvNV"))
	gpGetPathParameterivNV = uintptr(getProcAddr("glGetPathParameterivNV"))
	gpGetPathSpacingNV = uintptr(getProcAddr("glGetPathSpacingNV"))
	gpGetPerfCounterInfoINTEL = uintptr(getProcAddr("glGetPerfCounterInfoINTEL"))
	gpGetPerfMonitorCounterDataAMD = uintptr(getProcAddr("glGetPerfMonitorCounterDataAMD"))
	gpGetPerfMonitorCounterInfoAMD = uintptr(getProcAddr("glGetPerfMonitorCounterInfoAMD"))
	gpGetPerfMonitorCounterStringAMD = uintptr(getProcAddr("glGetPerfMonitorCounterStringAMD"))
	gpGetPerfMonitorCountersAMD = uintptr(getProcAddr("glGetPerfMonitorCountersAMD"))
	gpGetPerfMonitorGroupStringAMD = uintptr(getProcAddr("glGetPerfMonitorGroupStringAMD"))
	gpGetPerfMonitorGroupsAMD = uintptr(getProcAddr("glGetPerfMonitorGroupsAMD"))
	gpGetPerfQueryDataINTEL = uintptr(getProcAddr("glGetPerfQueryDataINTEL"))
	gpGetPerfQueryIdByNameINTEL = uintptr(getProcAddr("glGetPerfQueryIdByNameINTEL"))
	gpGetPerfQueryInfoINTEL = uintptr(getProcAddr("glGetPerfQueryInfoINTEL"))
	gpGetPixelMapfv = uintptr(getProcAddr("glGetPixelMapfv"))
	if gpGetPixelMapfv == 0 {
		return errors.New("glGetPixelMapfv")
	}
	gpGetPixelMapuiv = uintptr(getProcAddr("glGetPixelMapuiv"))
	if gpGetPixelMapuiv == 0 {
		return errors.New("glGetPixelMapuiv")
	}
	gpGetPixelMapusv = uintptr(getProcAddr("glGetPixelMapusv"))
	if gpGetPixelMapusv == 0 {
		return errors.New("glGetPixelMapusv")
	}
	gpGetPixelMapxv = uintptr(getProcAddr("glGetPixelMapxv"))
	gpGetPixelTexGenParameterfvSGIS = uintptr(getProcAddr("glGetPixelTexGenParameterfvSGIS"))
	gpGetPixelTexGenParameterivSGIS = uintptr(getProcAddr("glGetPixelTexGenParameterivSGIS"))
	gpGetPixelTransformParameterfvEXT = uintptr(getProcAddr("glGetPixelTransformParameterfvEXT"))
	gpGetPixelTransformParameterivEXT = uintptr(getProcAddr("glGetPixelTransformParameterivEXT"))
	gpGetPointerIndexedvEXT = uintptr(getProcAddr("glGetPointerIndexedvEXT"))
	gpGetPointeri_vEXT = uintptr(getProcAddr("glGetPointeri_vEXT"))
	gpGetPointerv = uintptr(getProcAddr("glGetPointerv"))
	if gpGetPointerv == 0 {
		return errors.New("glGetPointerv")
	}
	gpGetPointervEXT = uintptr(getProcAddr("glGetPointervEXT"))
	gpGetPointervKHR = uintptr(getProcAddr("glGetPointervKHR"))
	gpGetPolygonStipple = uintptr(getProcAddr("glGetPolygonStipple"))
	if gpGetPolygonStipple == 0 {
		return errors.New("glGetPolygonStipple")
	}
	gpGetProgramBinary = uintptr(getProcAddr("glGetProgramBinary"))
	gpGetProgramEnvParameterIivNV = uintptr(getProcAddr("glGetProgramEnvParameterIivNV"))
	gpGetProgramEnvParameterIuivNV = uintptr(getProcAddr("glGetProgramEnvParameterIuivNV"))
	gpGetProgramEnvParameterdvARB = uintptr(getProcAddr("glGetProgramEnvParameterdvARB"))
	gpGetProgramEnvParameterfvARB = uintptr(getProcAddr("glGetProgramEnvParameterfvARB"))
	gpGetProgramInfoLog = uintptr(getProcAddr("glGetProgramInfoLog"))
	if gpGetProgramInfoLog == 0 {
		return errors.New("glGetProgramInfoLog")
	}
	gpGetProgramInterfaceiv = uintptr(getProcAddr("glGetProgramInterfaceiv"))
	gpGetProgramLocalParameterIivNV = uintptr(getProcAddr("glGetProgramLocalParameterIivNV"))
	gpGetProgramLocalParameterIuivNV = uintptr(getProcAddr("glGetProgramLocalParameterIuivNV"))
	gpGetProgramLocalParameterdvARB = uintptr(getProcAddr("glGetProgramLocalParameterdvARB"))
	gpGetProgramLocalParameterfvARB = uintptr(getProcAddr("glGetProgramLocalParameterfvARB"))
	gpGetProgramNamedParameterdvNV = uintptr(getProcAddr("glGetProgramNamedParameterdvNV"))
	gpGetProgramNamedParameterfvNV = uintptr(getProcAddr("glGetProgramNamedParameterfvNV"))
	gpGetProgramParameterdvNV = uintptr(getProcAddr("glGetProgramParameterdvNV"))
	gpGetProgramParameterfvNV = uintptr(getProcAddr("glGetProgramParameterfvNV"))
	gpGetProgramPipelineInfoLog = uintptr(getProcAddr("glGetProgramPipelineInfoLog"))
	gpGetProgramPipelineInfoLogEXT = uintptr(getProcAddr("glGetProgramPipelineInfoLogEXT"))
	gpGetProgramPipelineiv = uintptr(getProcAddr("glGetProgramPipelineiv"))
	gpGetProgramPipelineivEXT = uintptr(getProcAddr("glGetProgramPipelineivEXT"))
	gpGetProgramResourceIndex = uintptr(getProcAddr("glGetProgramResourceIndex"))
	gpGetProgramResourceLocation = uintptr(getProcAddr("glGetProgramResourceLocation"))
	gpGetProgramResourceLocationIndex = uintptr(getProcAddr("glGetProgramResourceLocationIndex"))
	gpGetProgramResourceName = uintptr(getProcAddr("glGetProgramResourceName"))
	gpGetProgramResourcefvNV = uintptr(getProcAddr("glGetProgramResourcefvNV"))
	gpGetProgramResourceiv = uintptr(getProcAddr("glGetProgramResourceiv"))
	gpGetProgramStageiv = uintptr(getProcAddr("glGetProgramStageiv"))
	gpGetProgramStringARB = uintptr(getProcAddr("glGetProgramStringARB"))
	gpGetProgramStringNV = uintptr(getProcAddr("glGetProgramStringNV"))
	gpGetProgramSubroutineParameteruivNV = uintptr(getProcAddr("glGetProgramSubroutineParameteruivNV"))
	gpGetProgramiv = uintptr(getProcAddr("glGetProgramiv"))
	if gpGetProgramiv == 0 {
		return errors.New("glGetProgramiv")
	}
	gpGetProgramivARB = uintptr(getProcAddr("glGetProgramivARB"))
	gpGetProgramivNV = uintptr(getProcAddr("glGetProgramivNV"))
	gpGetQueryBufferObjecti64v = uintptr(getProcAddr("glGetQueryBufferObjecti64v"))
	gpGetQueryBufferObjectiv = uintptr(getProcAddr("glGetQueryBufferObjectiv"))
	gpGetQueryBufferObjectui64v = uintptr(getProcAddr("glGetQueryBufferObjectui64v"))
	gpGetQueryBufferObjectuiv = uintptr(getProcAddr("glGetQueryBufferObjectuiv"))
	gpGetQueryIndexediv = uintptr(getProcAddr("glGetQueryIndexediv"))
	gpGetQueryObjecti64v = uintptr(getProcAddr("glGetQueryObjecti64v"))
	gpGetQueryObjecti64vEXT = uintptr(getProcAddr("glGetQueryObjecti64vEXT"))
	gpGetQueryObjectiv = uintptr(getProcAddr("glGetQueryObjectiv"))
	if gpGetQueryObjectiv == 0 {
		return errors.New("glGetQueryObjectiv")
	}
	gpGetQueryObjectivARB = uintptr(getProcAddr("glGetQueryObjectivARB"))
	gpGetQueryObjectui64v = uintptr(getProcAddr("glGetQueryObjectui64v"))
	gpGetQueryObjectui64vEXT = uintptr(getProcAddr("glGetQueryObjectui64vEXT"))
	gpGetQueryObjectuiv = uintptr(getProcAddr("glGetQueryObjectuiv"))
	if gpGetQueryObjectuiv == 0 {
		return errors.New("glGetQueryObjectuiv")
	}
	gpGetQueryObjectuivARB = uintptr(getProcAddr("glGetQueryObjectuivARB"))
	gpGetQueryiv = uintptr(getProcAddr("glGetQueryiv"))
	if gpGetQueryiv == 0 {
		return errors.New("glGetQueryiv")
	}
	gpGetQueryivARB = uintptr(getProcAddr("glGetQueryivARB"))
	gpGetRenderbufferParameteriv = uintptr(getProcAddr("glGetRenderbufferParameteriv"))
	gpGetRenderbufferParameterivEXT = uintptr(getProcAddr("glGetRenderbufferParameterivEXT"))
	gpGetSamplerParameterIiv = uintptr(getProcAddr("glGetSamplerParameterIiv"))
	gpGetSamplerParameterIuiv = uintptr(getProcAddr("glGetSamplerParameterIuiv"))
	gpGetSamplerParameterfv = uintptr(getProcAddr("glGetSamplerParameterfv"))
	gpGetSamplerParameteriv = uintptr(getProcAddr("glGetSamplerParameteriv"))
	gpGetSemaphoreParameterui64vEXT = uintptr(getProcAddr("glGetSemaphoreParameterui64vEXT"))
	gpGetSeparableFilterEXT = uintptr(getProcAddr("glGetSeparableFilterEXT"))
	gpGetShaderInfoLog = uintptr(getProcAddr("glGetShaderInfoLog"))
	if gpGetShaderInfoLog == 0 {
		return errors.New("glGetShaderInfoLog")
	}
	gpGetShaderPrecisionFormat = uintptr(getProcAddr("glGetShaderPrecisionFormat"))
	gpGetShaderSource = uintptr(getProcAddr("glGetShaderSource"))
	if gpGetShaderSource == 0 {
		return errors.New("glGetShaderSource")
	}
	gpGetShaderSourceARB = uintptr(getProcAddr("glGetShaderSourceARB"))
	gpGetShaderiv = uintptr(getProcAddr("glGetShaderiv"))
	if gpGetShaderiv == 0 {
		return errors.New("glGetShaderiv")
	}
	gpGetSharpenTexFuncSGIS = uintptr(getProcAddr("glGetSharpenTexFuncSGIS"))
	gpGetStageIndexNV = uintptr(getProcAddr("glGetStageIndexNV"))
	gpGetString = uintptr(getProcAddr("glGetString"))
	if gpGetString == 0 {
		return errors.New("glGetString")
	}
	gpGetSubroutineIndex = uintptr(getProcAddr("glGetSubroutineIndex"))
	gpGetSubroutineUniformLocation = uintptr(getProcAddr("glGetSubroutineUniformLocation"))
	gpGetSynciv = uintptr(getProcAddr("glGetSynciv"))
	gpGetTexBumpParameterfvATI = uintptr(getProcAddr("glGetTexBumpParameterfvATI"))
	gpGetTexBumpParameterivATI = uintptr(getProcAddr("glGetTexBumpParameterivATI"))
	gpGetTexEnvfv = uintptr(getProcAddr("glGetTexEnvfv"))
	if gpGetTexEnvfv == 0 {
		return errors.New("glGetTexEnvfv")
	}
	gpGetTexEnviv = uintptr(getProcAddr("glGetTexEnviv"))
	if gpGetTexEnviv == 0 {
		return errors.New("glGetTexEnviv")
	}
	gpGetTexEnvxvOES = uintptr(getProcAddr("glGetTexEnvxvOES"))
	gpGetTexFilterFuncSGIS = uintptr(getProcAddr("glGetTexFilterFuncSGIS"))
	gpGetTexGendv = uintptr(getProcAddr("glGetTexGendv"))
	if gpGetTexGendv == 0 {
		return errors.New("glGetTexGendv")
	}
	gpGetTexGenfv = uintptr(getProcAddr("glGetTexGenfv"))
	if gpGetTexGenfv == 0 {
		return errors.New("glGetTexGenfv")
	}
	gpGetTexGeniv = uintptr(getProcAddr("glGetTexGeniv"))
	if gpGetTexGeniv == 0 {
		return errors.New("glGetTexGeniv")
	}
	gpGetTexGenxvOES = uintptr(getProcAddr("glGetTexGenxvOES"))
	gpGetTexImage = uintptr(getProcAddr("glGetTexImage"))
	if gpGetTexImage == 0 {
		return errors.New("glGetTexImage")
	}
	gpGetTexLevelParameterfv = uintptr(getProcAddr("glGetTexLevelParameterfv"))
	if gpGetTexLevelParameterfv == 0 {
		return errors.New("glGetTexLevelParameterfv")
	}
	gpGetTexLevelParameteriv = uintptr(getProcAddr("glGetTexLevelParameteriv"))
	if gpGetTexLevelParameteriv == 0 {
		return errors.New("glGetTexLevelParameteriv")
	}
	gpGetTexLevelParameterxvOES = uintptr(getProcAddr("glGetTexLevelParameterxvOES"))
	gpGetTexParameterIivEXT = uintptr(getProcAddr("glGetTexParameterIivEXT"))
	gpGetTexParameterIuivEXT = uintptr(getProcAddr("glGetTexParameterIuivEXT"))
	gpGetTexParameterPointervAPPLE = uintptr(getProcAddr("glGetTexParameterPointervAPPLE"))
	gpGetTexParameterfv = uintptr(getProcAddr("glGetTexParameterfv"))
	if gpGetTexParameterfv == 0 {
		return errors.New("glGetTexParameterfv")
	}
	gpGetTexParameteriv = uintptr(getProcAddr("glGetTexParameteriv"))
	if gpGetTexParameteriv == 0 {
		return errors.New("glGetTexParameteriv")
	}
	gpGetTexParameterxvOES = uintptr(getProcAddr("glGetTexParameterxvOES"))
	gpGetTextureHandleARB = uintptr(getProcAddr("glGetTextureHandleARB"))
	gpGetTextureHandleNV = uintptr(getProcAddr("glGetTextureHandleNV"))
	gpGetTextureImage = uintptr(getProcAddr("glGetTextureImage"))
	gpGetTextureImageEXT = uintptr(getProcAddr("glGetTextureImageEXT"))
	gpGetTextureLevelParameterfv = uintptr(getProcAddr("glGetTextureLevelParameterfv"))
	gpGetTextureLevelParameterfvEXT = uintptr(getProcAddr("glGetTextureLevelParameterfvEXT"))
	gpGetTextureLevelParameteriv = uintptr(getProcAddr("glGetTextureLevelParameteriv"))
	gpGetTextureLevelParameterivEXT = uintptr(getProcAddr("glGetTextureLevelParameterivEXT"))
	gpGetTextureParameterIiv = uintptr(getProcAddr("glGetTextureParameterIiv"))
	gpGetTextureParameterIivEXT = uintptr(getProcAddr("glGetTextureParameterIivEXT"))
	gpGetTextureParameterIuiv = uintptr(getProcAddr("glGetTextureParameterIuiv"))
	gpGetTextureParameterIuivEXT = uintptr(getProcAddr("glGetTextureParameterIuivEXT"))
	gpGetTextureParameterfv = uintptr(getProcAddr("glGetTextureParameterfv"))
	gpGetTextureParameterfvEXT = uintptr(getProcAddr("glGetTextureParameterfvEXT"))
	gpGetTextureParameteriv = uintptr(getProcAddr("glGetTextureParameteriv"))
	gpGetTextureParameterivEXT = uintptr(getProcAddr("glGetTextureParameterivEXT"))
	gpGetTextureSamplerHandleARB = uintptr(getProcAddr("glGetTextureSamplerHandleARB"))
	gpGetTextureSamplerHandleNV = uintptr(getProcAddr("glGetTextureSamplerHandleNV"))
	gpGetTextureSubImage = uintptr(getProcAddr("glGetTextureSubImage"))
	gpGetTrackMatrixivNV = uintptr(getProcAddr("glGetTrackMatrixivNV"))
	gpGetTransformFeedbackVaryingEXT = uintptr(getProcAddr("glGetTransformFeedbackVaryingEXT"))
	gpGetTransformFeedbackVaryingNV = uintptr(getProcAddr("glGetTransformFeedbackVaryingNV"))
	gpGetTransformFeedbacki64_v = uintptr(getProcAddr("glGetTransformFeedbacki64_v"))
	gpGetTransformFeedbacki_v = uintptr(getProcAddr("glGetTransformFeedbacki_v"))
	gpGetTransformFeedbackiv = uintptr(getProcAddr("glGetTransformFeedbackiv"))
	gpGetUniformBlockIndex = uintptr(getProcAddr("glGetUniformBlockIndex"))
	gpGetUniformBufferSizeEXT = uintptr(getProcAddr("glGetUniformBufferSizeEXT"))
	gpGetUniformIndices = uintptr(getProcAddr("glGetUniformIndices"))
	gpGetUniformLocation = uintptr(getProcAddr("glGetUniformLocation"))
	if gpGetUniformLocation == 0 {
		return errors.New("glGetUniformLocation")
	}
	gpGetUniformLocationARB = uintptr(getProcAddr("glGetUniformLocationARB"))
	gpGetUniformOffsetEXT = uintptr(getProcAddr("glGetUniformOffsetEXT"))
	gpGetUniformSubroutineuiv = uintptr(getProcAddr("glGetUniformSubroutineuiv"))
	gpGetUniformdv = uintptr(getProcAddr("glGetUniformdv"))
	gpGetUniformfv = uintptr(getProcAddr("glGetUniformfv"))
	if gpGetUniformfv == 0 {
		return errors.New("glGetUniformfv")
	}
	gpGetUniformfvARB = uintptr(getProcAddr("glGetUniformfvARB"))
	gpGetUniformi64vARB = uintptr(getProcAddr("glGetUniformi64vARB"))
	gpGetUniformi64vNV = uintptr(getProcAddr("glGetUniformi64vNV"))
	gpGetUniformiv = uintptr(getProcAddr("glGetUniformiv"))
	if gpGetUniformiv == 0 {
		return errors.New("glGetUniformiv")
	}
	gpGetUniformivARB = uintptr(getProcAddr("glGetUniformivARB"))
	gpGetUniformui64vARB = uintptr(getProcAddr("glGetUniformui64vARB"))
	gpGetUniformui64vNV = uintptr(getProcAddr("glGetUniformui64vNV"))
	gpGetUniformuivEXT = uintptr(getProcAddr("glGetUniformuivEXT"))
	gpGetUnsignedBytei_vEXT = uintptr(getProcAddr("glGetUnsignedBytei_vEXT"))
	gpGetUnsignedBytevEXT = uintptr(getProcAddr("glGetUnsignedBytevEXT"))
	gpGetVariantArrayObjectfvATI = uintptr(getProcAddr("glGetVariantArrayObjectfvATI"))
	gpGetVariantArrayObjectivATI = uintptr(getProcAddr("glGetVariantArrayObjectivATI"))
	gpGetVariantBooleanvEXT = uintptr(getProcAddr("glGetVariantBooleanvEXT"))
	gpGetVariantFloatvEXT = uintptr(getProcAddr("glGetVariantFloatvEXT"))
	gpGetVariantIntegervEXT = uintptr(getProcAddr("glGetVariantIntegervEXT"))
	gpGetVariantPointervEXT = uintptr(getProcAddr("glGetVariantPointervEXT"))
	gpGetVaryingLocationNV = uintptr(getProcAddr("glGetVaryingLocationNV"))
	gpGetVertexArrayIndexed64iv = uintptr(getProcAddr("glGetVertexArrayIndexed64iv"))
	gpGetVertexArrayIndexediv = uintptr(getProcAddr("glGetVertexArrayIndexediv"))
	gpGetVertexArrayIntegeri_vEXT = uintptr(getProcAddr("glGetVertexArrayIntegeri_vEXT"))
	gpGetVertexArrayIntegervEXT = uintptr(getProcAddr("glGetVertexArrayIntegervEXT"))
	gpGetVertexArrayPointeri_vEXT = uintptr(getProcAddr("glGetVertexArrayPointeri_vEXT"))
	gpGetVertexArrayPointervEXT = uintptr(getProcAddr("glGetVertexArrayPointervEXT"))
	gpGetVertexArrayiv = uintptr(getProcAddr("glGetVertexArrayiv"))
	gpGetVertexAttribArrayObjectfvATI = uintptr(getProcAddr("glGetVertexAttribArrayObjectfvATI"))
	gpGetVertexAttribArrayObjectivATI = uintptr(getProcAddr("glGetVertexAttribArrayObjectivATI"))
	gpGetVertexAttribIivEXT = uintptr(getProcAddr("glGetVertexAttribIivEXT"))
	gpGetVertexAttribIuivEXT = uintptr(getProcAddr("glGetVertexAttribIuivEXT"))
	gpGetVertexAttribLdv = uintptr(getProcAddr("glGetVertexAttribLdv"))
	gpGetVertexAttribLdvEXT = uintptr(getProcAddr("glGetVertexAttribLdvEXT"))
	gpGetVertexAttribLi64vNV = uintptr(getProcAddr("glGetVertexAttribLi64vNV"))
	gpGetVertexAttribLui64vARB = uintptr(getProcAddr("glGetVertexAttribLui64vARB"))
	gpGetVertexAttribLui64vNV = uintptr(getProcAddr("glGetVertexAttribLui64vNV"))
	gpGetVertexAttribPointerv = uintptr(getProcAddr("glGetVertexAttribPointerv"))
	if gpGetVertexAttribPointerv == 0 {
		return errors.New("glGetVertexAttribPointerv")
	}
	gpGetVertexAttribPointervARB = uintptr(getProcAddr("glGetVertexAttribPointervARB"))
	gpGetVertexAttribPointervNV = uintptr(getProcAddr("glGetVertexAttribPointervNV"))
	gpGetVertexAttribdv = uintptr(getProcAddr("glGetVertexAttribdv"))
	if gpGetVertexAttribdv == 0 {
		return errors.New("glGetVertexAttribdv")
	}
	gpGetVertexAttribdvARB = uintptr(getProcAddr("glGetVertexAttribdvARB"))
	gpGetVertexAttribdvNV = uintptr(getProcAddr("glGetVertexAttribdvNV"))
	gpGetVertexAttribfv = uintptr(getProcAddr("glGetVertexAttribfv"))
	if gpGetVertexAttribfv == 0 {
		return errors.New("glGetVertexAttribfv")
	}
	gpGetVertexAttribfvARB = uintptr(getProcAddr("glGetVertexAttribfvARB"))
	gpGetVertexAttribfvNV = uintptr(getProcAddr("glGetVertexAttribfvNV"))
	gpGetVertexAttribiv = uintptr(getProcAddr("glGetVertexAttribiv"))
	if gpGetVertexAttribiv == 0 {
		return errors.New("glGetVertexAttribiv")
	}
	gpGetVertexAttribivARB = uintptr(getProcAddr("glGetVertexAttribivARB"))
	gpGetVertexAttribivNV = uintptr(getProcAddr("glGetVertexAttribivNV"))
	gpGetVideoCaptureStreamdvNV = uintptr(getProcAddr("glGetVideoCaptureStreamdvNV"))
	gpGetVideoCaptureStreamfvNV = uintptr(getProcAddr("glGetVideoCaptureStreamfvNV"))
	gpGetVideoCaptureStreamivNV = uintptr(getProcAddr("glGetVideoCaptureStreamivNV"))
	gpGetVideoCaptureivNV = uintptr(getProcAddr("glGetVideoCaptureivNV"))
	gpGetVideoi64vNV = uintptr(getProcAddr("glGetVideoi64vNV"))
	gpGetVideoivNV = uintptr(getProcAddr("glGetVideoivNV"))
	gpGetVideoui64vNV = uintptr(getProcAddr("glGetVideoui64vNV"))
	gpGetVideouivNV = uintptr(getProcAddr("glGetVideouivNV"))
	gpGetVkProcAddrNV = uintptr(getProcAddr("glGetVkProcAddrNV"))
	gpGetnCompressedTexImageARB = uintptr(getProcAddr("glGetnCompressedTexImageARB"))
	gpGetnTexImageARB = uintptr(getProcAddr("glGetnTexImageARB"))
	gpGetnUniformdvARB = uintptr(getProcAddr("glGetnUniformdvARB"))
	gpGetnUniformfv = uintptr(getProcAddr("glGetnUniformfv"))
	gpGetnUniformfvARB = uintptr(getProcAddr("glGetnUniformfvARB"))
	gpGetnUniformfvKHR = uintptr(getProcAddr("glGetnUniformfvKHR"))
	gpGetnUniformi64vARB = uintptr(getProcAddr("glGetnUniformi64vARB"))
	gpGetnUniformiv = uintptr(getProcAddr("glGetnUniformiv"))
	gpGetnUniformivARB = uintptr(getProcAddr("glGetnUniformivARB"))
	gpGetnUniformivKHR = uintptr(getProcAddr("glGetnUniformivKHR"))
	gpGetnUniformui64vARB = uintptr(getProcAddr("glGetnUniformui64vARB"))
	gpGetnUniformuiv = uintptr(getProcAddr("glGetnUniformuiv"))
	gpGetnUniformuivARB = uintptr(getProcAddr("glGetnUniformuivARB"))
	gpGetnUniformuivKHR = uintptr(getProcAddr("glGetnUniformuivKHR"))
	gpGlobalAlphaFactorbSUN = uintptr(getProcAddr("glGlobalAlphaFactorbSUN"))
	gpGlobalAlphaFactordSUN = uintptr(getProcAddr("glGlobalAlphaFactordSUN"))
	gpGlobalAlphaFactorfSUN = uintptr(getProcAddr("glGlobalAlphaFactorfSUN"))
	gpGlobalAlphaFactoriSUN = uintptr(getProcAddr("glGlobalAlphaFactoriSUN"))
	gpGlobalAlphaFactorsSUN = uintptr(getProcAddr("glGlobalAlphaFactorsSUN"))
	gpGlobalAlphaFactorubSUN = uintptr(getProcAddr("glGlobalAlphaFactorubSUN"))
	gpGlobalAlphaFactoruiSUN = uintptr(getProcAddr("glGlobalAlphaFactoruiSUN"))
	gpGlobalAlphaFactorusSUN = uintptr(getProcAddr("glGlobalAlphaFactorusSUN"))
	gpHint = uintptr(getProcAddr("glHint"))
	if gpHint == 0 {
		return errors.New("glHint")
	}
	gpHintPGI = uintptr(getProcAddr("glHintPGI"))
	gpHistogramEXT = uintptr(getProcAddr("glHistogramEXT"))
	gpIglooInterfaceSGIX = uintptr(getProcAddr("glIglooInterfaceSGIX"))
	gpImageTransformParameterfHP = uintptr(getProcAddr("glImageTransformParameterfHP"))
	gpImageTransformParameterfvHP = uintptr(getProcAddr("glImageTransformParameterfvHP"))
	gpImageTransformParameteriHP = uintptr(getProcAddr("glImageTransformParameteriHP"))
	gpImageTransformParameterivHP = uintptr(getProcAddr("glImageTransformParameterivHP"))
	gpImportMemoryFdEXT = uintptr(getProcAddr("glImportMemoryFdEXT"))
	gpImportMemoryWin32HandleEXT = uintptr(getProcAddr("glImportMemoryWin32HandleEXT"))
	gpImportMemoryWin32NameEXT = uintptr(getProcAddr("glImportMemoryWin32NameEXT"))
	gpImportSemaphoreFdEXT = uintptr(getProcAddr("glImportSemaphoreFdEXT"))
	gpImportSemaphoreWin32HandleEXT = uintptr(getProcAddr("glImportSemaphoreWin32HandleEXT"))
	gpImportSemaphoreWin32NameEXT = uintptr(getProcAddr("glImportSemaphoreWin32NameEXT"))
	gpImportSyncEXT = uintptr(getProcAddr("glImportSyncEXT"))
	gpIndexFormatNV = uintptr(getProcAddr("glIndexFormatNV"))
	gpIndexFuncEXT = uintptr(getProcAddr("glIndexFuncEXT"))
	gpIndexMask = uintptr(getProcAddr("glIndexMask"))
	if gpIndexMask == 0 {
		return errors.New("glIndexMask")
	}
	gpIndexMaterialEXT = uintptr(getProcAddr("glIndexMaterialEXT"))
	gpIndexPointer = uintptr(getProcAddr("glIndexPointer"))
	if gpIndexPointer == 0 {
		return errors.New("glIndexPointer")
	}
	gpIndexPointerEXT = uintptr(getProcAddr("glIndexPointerEXT"))
	gpIndexPointerListIBM = uintptr(getProcAddr("glIndexPointerListIBM"))
	gpIndexd = uintptr(getProcAddr("glIndexd"))
	if gpIndexd == 0 {
		return errors.New("glIndexd")
	}
	gpIndexdv = uintptr(getProcAddr("glIndexdv"))
	if gpIndexdv == 0 {
		return errors.New("glIndexdv")
	}
	gpIndexf = uintptr(getProcAddr("glIndexf"))
	if gpIndexf == 0 {
		return errors.New("glIndexf")
	}
	gpIndexfv = uintptr(getProcAddr("glIndexfv"))
	if gpIndexfv == 0 {
		return errors.New("glIndexfv")
	}
	gpIndexi = uintptr(getProcAddr("glIndexi"))
	if gpIndexi == 0 {
		return errors.New("glIndexi")
	}
	gpIndexiv = uintptr(getProcAddr("glIndexiv"))
	if gpIndexiv == 0 {
		return errors.New("glIndexiv")
	}
	gpIndexs = uintptr(getProcAddr("glIndexs"))
	if gpIndexs == 0 {
		return errors.New("glIndexs")
	}
	gpIndexsv = uintptr(getProcAddr("glIndexsv"))
	if gpIndexsv == 0 {
		return errors.New("glIndexsv")
	}
	gpIndexub = uintptr(getProcAddr("glIndexub"))
	if gpIndexub == 0 {
		return errors.New("glIndexub")
	}
	gpIndexubv = uintptr(getProcAddr("glIndexubv"))
	if gpIndexubv == 0 {
		return errors.New("glIndexubv")
	}
	gpIndexxOES = uintptr(getProcAddr("glIndexxOES"))
	gpIndexxvOES = uintptr(getProcAddr("glIndexxvOES"))
	gpInitNames = uintptr(getProcAddr("glInitNames"))
	if gpInitNames == 0 {
		return errors.New("glInitNames")
	}
	gpInsertComponentEXT = uintptr(getProcAddr("glInsertComponentEXT"))
	gpInsertEventMarkerEXT = uintptr(getProcAddr("glInsertEventMarkerEXT"))
	gpInstrumentsBufferSGIX = uintptr(getProcAddr("glInstrumentsBufferSGIX"))
	gpInterleavedArrays = uintptr(getProcAddr("glInterleavedArrays"))
	if gpInterleavedArrays == 0 {
		return errors.New("glInterleavedArrays")
	}
	gpInterpolatePathsNV = uintptr(getProcAddr("glInterpolatePathsNV"))
	gpInvalidateBufferData = uintptr(getProcAddr("glInvalidateBufferData"))
	gpInvalidateBufferSubData = uintptr(getProcAddr("glInvalidateBufferSubData"))
	gpInvalidateFramebuffer = uintptr(getProcAddr("glInvalidateFramebuffer"))
	gpInvalidateNamedFramebufferData = uintptr(getProcAddr("glInvalidateNamedFramebufferData"))
	gpInvalidateNamedFramebufferSubData = uintptr(getProcAddr("glInvalidateNamedFramebufferSubData"))
	gpInvalidateSubFramebuffer = uintptr(getProcAddr("glInvalidateSubFramebuffer"))
	gpInvalidateTexImage = uintptr(getProcAddr("glInvalidateTexImage"))
	gpInvalidateTexSubImage = uintptr(getProcAddr("glInvalidateTexSubImage"))
	gpIsAsyncMarkerSGIX = uintptr(getProcAddr("glIsAsyncMarkerSGIX"))
	gpIsBuffer = uintptr(getProcAddr("glIsBuffer"))
	if gpIsBuffer == 0 {
		return errors.New("glIsBuffer")
	}
	gpIsBufferARB = uintptr(getProcAddr("glIsBufferARB"))
	gpIsBufferResidentNV = uintptr(getProcAddr("glIsBufferResidentNV"))
	gpIsCommandListNV = uintptr(getProcAddr("glIsCommandListNV"))
	gpIsEnabled = uintptr(getProcAddr("glIsEnabled"))
	if gpIsEnabled == 0 {
		return errors.New("glIsEnabled")
	}
	gpIsEnabledIndexedEXT = uintptr(getProcAddr("glIsEnabledIndexedEXT"))
	gpIsFenceAPPLE = uintptr(getProcAddr("glIsFenceAPPLE"))
	gpIsFenceNV = uintptr(getProcAddr("glIsFenceNV"))
	gpIsFramebuffer = uintptr(getProcAddr("glIsFramebuffer"))
	gpIsFramebufferEXT = uintptr(getProcAddr("glIsFramebufferEXT"))
	gpIsImageHandleResidentARB = uintptr(getProcAddr("glIsImageHandleResidentARB"))
	gpIsImageHandleResidentNV = uintptr(getProcAddr("glIsImageHandleResidentNV"))
	gpIsList = uintptr(getProcAddr("glIsList"))
	if gpIsList == 0 {
		return errors.New("glIsList")
	}
	gpIsMemoryObjectEXT = uintptr(getProcAddr("glIsMemoryObjectEXT"))
	gpIsNameAMD = uintptr(getProcAddr("glIsNameAMD"))
	gpIsNamedBufferResidentNV = uintptr(getProcAddr("glIsNamedBufferResidentNV"))
	gpIsNamedStringARB = uintptr(getProcAddr("glIsNamedStringARB"))
	gpIsObjectBufferATI = uintptr(getProcAddr("glIsObjectBufferATI"))
	gpIsOcclusionQueryNV = uintptr(getProcAddr("glIsOcclusionQueryNV"))
	gpIsPathNV = uintptr(getProcAddr("glIsPathNV"))
	gpIsPointInFillPathNV = uintptr(getProcAddr("glIsPointInFillPathNV"))
	gpIsPointInStrokePathNV = uintptr(getProcAddr("glIsPointInStrokePathNV"))
	gpIsProgram = uintptr(getProcAddr("glIsProgram"))
	if gpIsProgram == 0 {
		return errors.New("glIsProgram")
	}
	gpIsProgramARB = uintptr(getProcAddr("glIsProgramARB"))
	gpIsProgramNV = uintptr(getProcAddr("glIsProgramNV"))
	gpIsProgramPipeline = uintptr(getProcAddr("glIsProgramPipeline"))
	gpIsProgramPipelineEXT = uintptr(getProcAddr("glIsProgramPipelineEXT"))
	gpIsQuery = uintptr(getProcAddr("glIsQuery"))
	if gpIsQuery == 0 {
		return errors.New("glIsQuery")
	}
	gpIsQueryARB = uintptr(getProcAddr("glIsQueryARB"))
	gpIsRenderbuffer = uintptr(getProcAddr("glIsRenderbuffer"))
	gpIsRenderbufferEXT = uintptr(getProcAddr("glIsRenderbufferEXT"))
	gpIsSampler = uintptr(getProcAddr("glIsSampler"))
	gpIsSemaphoreEXT = uintptr(getProcAddr("glIsSemaphoreEXT"))
	gpIsShader = uintptr(getProcAddr("glIsShader"))
	if gpIsShader == 0 {
		return errors.New("glIsShader")
	}
	gpIsStateNV = uintptr(getProcAddr("glIsStateNV"))
	gpIsSync = uintptr(getProcAddr("glIsSync"))
	gpIsTexture = uintptr(getProcAddr("glIsTexture"))
	if gpIsTexture == 0 {
		return errors.New("glIsTexture")
	}
	gpIsTextureEXT = uintptr(getProcAddr("glIsTextureEXT"))
	gpIsTextureHandleResidentARB = uintptr(getProcAddr("glIsTextureHandleResidentARB"))
	gpIsTextureHandleResidentNV = uintptr(getProcAddr("glIsTextureHandleResidentNV"))
	gpIsTransformFeedback = uintptr(getProcAddr("glIsTransformFeedback"))
	gpIsTransformFeedbackNV = uintptr(getProcAddr("glIsTransformFeedbackNV"))
	gpIsVariantEnabledEXT = uintptr(getProcAddr("glIsVariantEnabledEXT"))
	gpIsVertexArray = uintptr(getProcAddr("glIsVertexArray"))
	gpIsVertexArrayAPPLE = uintptr(getProcAddr("glIsVertexArrayAPPLE"))
	gpIsVertexAttribEnabledAPPLE = uintptr(getProcAddr("glIsVertexAttribEnabledAPPLE"))
	gpLGPUCopyImageSubDataNVX = uintptr(getProcAddr("glLGPUCopyImageSubDataNVX"))
	gpLGPUInterlockNVX = uintptr(getProcAddr("glLGPUInterlockNVX"))
	gpLGPUNamedBufferSubDataNVX = uintptr(getProcAddr("glLGPUNamedBufferSubDataNVX"))
	gpLabelObjectEXT = uintptr(getProcAddr("glLabelObjectEXT"))
	gpLightEnviSGIX = uintptr(getProcAddr("glLightEnviSGIX"))
	gpLightModelf = uintptr(getProcAddr("glLightModelf"))
	if gpLightModelf == 0 {
		return errors.New("glLightModelf")
	}
	gpLightModelfv = uintptr(getProcAddr("glLightModelfv"))
	if gpLightModelfv == 0 {
		return errors.New("glLightModelfv")
	}
	gpLightModeli = uintptr(getProcAddr("glLightModeli"))
	if gpLightModeli == 0 {
		return errors.New("glLightModeli")
	}
	gpLightModeliv = uintptr(getProcAddr("glLightModeliv"))
	if gpLightModeliv == 0 {
		return errors.New("glLightModeliv")
	}
	gpLightModelxOES = uintptr(getProcAddr("glLightModelxOES"))
	gpLightModelxvOES = uintptr(getProcAddr("glLightModelxvOES"))
	gpLightf = uintptr(getProcAddr("glLightf"))
	if gpLightf == 0 {
		return errors.New("glLightf")
	}
	gpLightfv = uintptr(getProcAddr("glLightfv"))
	if gpLightfv == 0 {
		return errors.New("glLightfv")
	}
	gpLighti = uintptr(getProcAddr("glLighti"))
	if gpLighti == 0 {
		return errors.New("glLighti")
	}
	gpLightiv = uintptr(getProcAddr("glLightiv"))
	if gpLightiv == 0 {
		return errors.New("glLightiv")
	}
	gpLightxOES = uintptr(getProcAddr("glLightxOES"))
	gpLightxvOES = uintptr(getProcAddr("glLightxvOES"))
	gpLineStipple = uintptr(getProcAddr("glLineStipple"))
	if gpLineStipple == 0 {
		return errors.New("glLineStipple")
	}
	gpLineWidth = uintptr(getProcAddr("glLineWidth"))
	if gpLineWidth == 0 {
		return errors.New("glLineWidth")
	}
	gpLineWidthxOES = uintptr(getProcAddr("glLineWidthxOES"))
	gpLinkProgram = uintptr(getProcAddr("glLinkProgram"))
	if gpLinkProgram == 0 {
		return errors.New("glLinkProgram")
	}
	gpLinkProgramARB = uintptr(getProcAddr("glLinkProgramARB"))
	gpListBase = uintptr(getProcAddr("glListBase"))
	if gpListBase == 0 {
		return errors.New("glListBase")
	}
	gpListDrawCommandsStatesClientNV = uintptr(getProcAddr("glListDrawCommandsStatesClientNV"))
	gpListParameterfSGIX = uintptr(getProcAddr("glListParameterfSGIX"))
	gpListParameterfvSGIX = uintptr(getProcAddr("glListParameterfvSGIX"))
	gpListParameteriSGIX = uintptr(getProcAddr("glListParameteriSGIX"))
	gpListParameterivSGIX = uintptr(getProcAddr("glListParameterivSGIX"))
	gpLoadIdentity = uintptr(getProcAddr("glLoadIdentity"))
	if gpLoadIdentity == 0 {
		return errors.New("glLoadIdentity")
	}
	gpLoadIdentityDeformationMapSGIX = uintptr(getProcAddr("glLoadIdentityDeformationMapSGIX"))
	gpLoadMatrixd = uintptr(getProcAddr("glLoadMatrixd"))
	if gpLoadMatrixd == 0 {
		return errors.New("glLoadMatrixd")
	}
	gpLoadMatrixf = uintptr(getProcAddr("glLoadMatrixf"))
	if gpLoadMatrixf == 0 {
		return errors.New("glLoadMatrixf")
	}
	gpLoadMatrixxOES = uintptr(getProcAddr("glLoadMatrixxOES"))
	gpLoadName = uintptr(getProcAddr("glLoadName"))
	if gpLoadName == 0 {
		return errors.New("glLoadName")
	}
	gpLoadProgramNV = uintptr(getProcAddr("glLoadProgramNV"))
	gpLoadTransposeMatrixd = uintptr(getProcAddr("glLoadTransposeMatrixd"))
	if gpLoadTransposeMatrixd == 0 {
		return errors.New("glLoadTransposeMatrixd")
	}
	gpLoadTransposeMatrixdARB = uintptr(getProcAddr("glLoadTransposeMatrixdARB"))
	gpLoadTransposeMatrixf = uintptr(getProcAddr("glLoadTransposeMatrixf"))
	if gpLoadTransposeMatrixf == 0 {
		return errors.New("glLoadTransposeMatrixf")
	}
	gpLoadTransposeMatrixfARB = uintptr(getProcAddr("glLoadTransposeMatrixfARB"))
	gpLoadTransposeMatrixxOES = uintptr(getProcAddr("glLoadTransposeMatrixxOES"))
	gpLockArraysEXT = uintptr(getProcAddr("glLockArraysEXT"))
	gpLogicOp = uintptr(getProcAddr("glLogicOp"))
	if gpLogicOp == 0 {
		return errors.New("glLogicOp")
	}
	gpMakeBufferNonResidentNV = uintptr(getProcAddr("glMakeBufferNonResidentNV"))
	gpMakeBufferResidentNV = uintptr(getProcAddr("glMakeBufferResidentNV"))
	gpMakeImageHandleNonResidentARB = uintptr(getProcAddr("glMakeImageHandleNonResidentARB"))
	gpMakeImageHandleNonResidentNV = uintptr(getProcAddr("glMakeImageHandleNonResidentNV"))
	gpMakeImageHandleResidentARB = uintptr(getProcAddr("glMakeImageHandleResidentARB"))
	gpMakeImageHandleResidentNV = uintptr(getProcAddr("glMakeImageHandleResidentNV"))
	gpMakeNamedBufferNonResidentNV = uintptr(getProcAddr("glMakeNamedBufferNonResidentNV"))
	gpMakeNamedBufferResidentNV = uintptr(getProcAddr("glMakeNamedBufferResidentNV"))
	gpMakeTextureHandleNonResidentARB = uintptr(getProcAddr("glMakeTextureHandleNonResidentARB"))
	gpMakeTextureHandleNonResidentNV = uintptr(getProcAddr("glMakeTextureHandleNonResidentNV"))
	gpMakeTextureHandleResidentARB = uintptr(getProcAddr("glMakeTextureHandleResidentARB"))
	gpMakeTextureHandleResidentNV = uintptr(getProcAddr("glMakeTextureHandleResidentNV"))
	gpMap1d = uintptr(getProcAddr("glMap1d"))
	if gpMap1d == 0 {
		return errors.New("glMap1d")
	}
	gpMap1f = uintptr(getProcAddr("glMap1f"))
	if gpMap1f == 0 {
		return errors.New("glMap1f")
	}
	gpMap1xOES = uintptr(getProcAddr("glMap1xOES"))
	gpMap2d = uintptr(getProcAddr("glMap2d"))
	if gpMap2d == 0 {
		return errors.New("glMap2d")
	}
	gpMap2f = uintptr(getProcAddr("glMap2f"))
	if gpMap2f == 0 {
		return errors.New("glMap2f")
	}
	gpMap2xOES = uintptr(getProcAddr("glMap2xOES"))
	gpMapBuffer = uintptr(getProcAddr("glMapBuffer"))
	if gpMapBuffer == 0 {
		return errors.New("glMapBuffer")
	}
	gpMapBufferARB = uintptr(getProcAddr("glMapBufferARB"))
	gpMapBufferRange = uintptr(getProcAddr("glMapBufferRange"))
	gpMapControlPointsNV = uintptr(getProcAddr("glMapControlPointsNV"))
	gpMapGrid1d = uintptr(getProcAddr("glMapGrid1d"))
	if gpMapGrid1d == 0 {
		return errors.New("glMapGrid1d")
	}
	gpMapGrid1f = uintptr(getProcAddr("glMapGrid1f"))
	if gpMapGrid1f == 0 {
		return errors.New("glMapGrid1f")
	}
	gpMapGrid1xOES = uintptr(getProcAddr("glMapGrid1xOES"))
	gpMapGrid2d = uintptr(getProcAddr("glMapGrid2d"))
	if gpMapGrid2d == 0 {
		return errors.New("glMapGrid2d")
	}
	gpMapGrid2f = uintptr(getProcAddr("glMapGrid2f"))
	if gpMapGrid2f == 0 {
		return errors.New("glMapGrid2f")
	}
	gpMapGrid2xOES = uintptr(getProcAddr("glMapGrid2xOES"))
	gpMapNamedBuffer = uintptr(getProcAddr("glMapNamedBuffer"))
	gpMapNamedBufferEXT = uintptr(getProcAddr("glMapNamedBufferEXT"))
	gpMapNamedBufferRange = uintptr(getProcAddr("glMapNamedBufferRange"))
	gpMapNamedBufferRangeEXT = uintptr(getProcAddr("glMapNamedBufferRangeEXT"))
	gpMapObjectBufferATI = uintptr(getProcAddr("glMapObjectBufferATI"))
	gpMapParameterfvNV = uintptr(getProcAddr("glMapParameterfvNV"))
	gpMapParameterivNV = uintptr(getProcAddr("glMapParameterivNV"))
	gpMapTexture2DINTEL = uintptr(getProcAddr("glMapTexture2DINTEL"))
	gpMapVertexAttrib1dAPPLE = uintptr(getProcAddr("glMapVertexAttrib1dAPPLE"))
	gpMapVertexAttrib1fAPPLE = uintptr(getProcAddr("glMapVertexAttrib1fAPPLE"))
	gpMapVertexAttrib2dAPPLE = uintptr(getProcAddr("glMapVertexAttrib2dAPPLE"))
	gpMapVertexAttrib2fAPPLE = uintptr(getProcAddr("glMapVertexAttrib2fAPPLE"))
	gpMaterialf = uintptr(getProcAddr("glMaterialf"))
	if gpMaterialf == 0 {
		return errors.New("glMaterialf")
	}
	gpMaterialfv = uintptr(getProcAddr("glMaterialfv"))
	if gpMaterialfv == 0 {
		return errors.New("glMaterialfv")
	}
	gpMateriali = uintptr(getProcAddr("glMateriali"))
	if gpMateriali == 0 {
		return errors.New("glMateriali")
	}
	gpMaterialiv = uintptr(getProcAddr("glMaterialiv"))
	if gpMaterialiv == 0 {
		return errors.New("glMaterialiv")
	}
	gpMaterialxOES = uintptr(getProcAddr("glMaterialxOES"))
	gpMaterialxvOES = uintptr(getProcAddr("glMaterialxvOES"))
	gpMatrixFrustumEXT = uintptr(getProcAddr("glMatrixFrustumEXT"))
	gpMatrixIndexPointerARB = uintptr(getProcAddr("glMatrixIndexPointerARB"))
	gpMatrixIndexubvARB = uintptr(getProcAddr("glMatrixIndexubvARB"))
	gpMatrixIndexuivARB = uintptr(getProcAddr("glMatrixIndexuivARB"))
	gpMatrixIndexusvARB = uintptr(getProcAddr("glMatrixIndexusvARB"))
	gpMatrixLoad3x2fNV = uintptr(getProcAddr("glMatrixLoad3x2fNV"))
	gpMatrixLoad3x3fNV = uintptr(getProcAddr("glMatrixLoad3x3fNV"))
	gpMatrixLoadIdentityEXT = uintptr(getProcAddr("glMatrixLoadIdentityEXT"))
	gpMatrixLoadTranspose3x3fNV = uintptr(getProcAddr("glMatrixLoadTranspose3x3fNV"))
	gpMatrixLoadTransposedEXT = uintptr(getProcAddr("glMatrixLoadTransposedEXT"))
	gpMatrixLoadTransposefEXT = uintptr(getProcAddr("glMatrixLoadTransposefEXT"))
	gpMatrixLoaddEXT = uintptr(getProcAddr("glMatrixLoaddEXT"))
	gpMatrixLoadfEXT = uintptr(getProcAddr("glMatrixLoadfEXT"))
	gpMatrixMode = uintptr(getProcAddr("glMatrixMode"))
	if gpMatrixMode == 0 {
		return errors.New("glMatrixMode")
	}
	gpMatrixMult3x2fNV = uintptr(getProcAddr("glMatrixMult3x2fNV"))
	gpMatrixMult3x3fNV = uintptr(getProcAddr("glMatrixMult3x3fNV"))
	gpMatrixMultTranspose3x3fNV = uintptr(getProcAddr("glMatrixMultTranspose3x3fNV"))
	gpMatrixMultTransposedEXT = uintptr(getProcAddr("glMatrixMultTransposedEXT"))
	gpMatrixMultTransposefEXT = uintptr(getProcAddr("glMatrixMultTransposefEXT"))
	gpMatrixMultdEXT = uintptr(getProcAddr("glMatrixMultdEXT"))
	gpMatrixMultfEXT = uintptr(getProcAddr("glMatrixMultfEXT"))
	gpMatrixOrthoEXT = uintptr(getProcAddr("glMatrixOrthoEXT"))
	gpMatrixPopEXT = uintptr(getProcAddr("glMatrixPopEXT"))
	gpMatrixPushEXT = uintptr(getProcAddr("glMatrixPushEXT"))
	gpMatrixRotatedEXT = uintptr(getProcAddr("glMatrixRotatedEXT"))
	gpMatrixRotatefEXT = uintptr(getProcAddr("glMatrixRotatefEXT"))
	gpMatrixScaledEXT = uintptr(getProcAddr("glMatrixScaledEXT"))
	gpMatrixScalefEXT = uintptr(getProcAddr("glMatrixScalefEXT"))
	gpMatrixTranslatedEXT = uintptr(getProcAddr("glMatrixTranslatedEXT"))
	gpMatrixTranslatefEXT = uintptr(getProcAddr("glMatrixTranslatefEXT"))
	gpMaxShaderCompilerThreadsARB = uintptr(getProcAddr("glMaxShaderCompilerThreadsARB"))
	gpMaxShaderCompilerThreadsKHR = uintptr(getProcAddr("glMaxShaderCompilerThreadsKHR"))
	gpMemoryBarrier = uintptr(getProcAddr("glMemoryBarrier"))
	gpMemoryBarrierByRegion = uintptr(getProcAddr("glMemoryBarrierByRegion"))
	gpMemoryBarrierEXT = uintptr(getProcAddr("glMemoryBarrierEXT"))
	gpMemoryObjectParameterivEXT = uintptr(getProcAddr("glMemoryObjectParameterivEXT"))
	gpMinSampleShadingARB = uintptr(getProcAddr("glMinSampleShadingARB"))
	gpMinmaxEXT = uintptr(getProcAddr("glMinmaxEXT"))
	gpMultMatrixd = uintptr(getProcAddr("glMultMatrixd"))
	if gpMultMatrixd == 0 {
		return errors.New("glMultMatrixd")
	}
	gpMultMatrixf = uintptr(getProcAddr("glMultMatrixf"))
	if gpMultMatrixf == 0 {
		return errors.New("glMultMatrixf")
	}
	gpMultMatrixxOES = uintptr(getProcAddr("glMultMatrixxOES"))
	gpMultTransposeMatrixd = uintptr(getProcAddr("glMultTransposeMatrixd"))
	if gpMultTransposeMatrixd == 0 {
		return errors.New("glMultTransposeMatrixd")
	}
	gpMultTransposeMatrixdARB = uintptr(getProcAddr("glMultTransposeMatrixdARB"))
	gpMultTransposeMatrixf = uintptr(getProcAddr("glMultTransposeMatrixf"))
	if gpMultTransposeMatrixf == 0 {
		return errors.New("glMultTransposeMatrixf")
	}
	gpMultTransposeMatrixfARB = uintptr(getProcAddr("glMultTransposeMatrixfARB"))
	gpMultTransposeMatrixxOES = uintptr(getProcAddr("glMultTransposeMatrixxOES"))
	gpMultiDrawArrays = uintptr(getProcAddr("glMultiDrawArrays"))
	if gpMultiDrawArrays == 0 {
		return errors.New("glMultiDrawArrays")
	}
	gpMultiDrawArraysEXT = uintptr(getProcAddr("glMultiDrawArraysEXT"))
	gpMultiDrawArraysIndirect = uintptr(getProcAddr("glMultiDrawArraysIndirect"))
	gpMultiDrawArraysIndirectAMD = uintptr(getProcAddr("glMultiDrawArraysIndirectAMD"))
	gpMultiDrawArraysIndirectBindlessCountNV = uintptr(getProcAddr("glMultiDrawArraysIndirectBindlessCountNV"))
	gpMultiDrawArraysIndirectBindlessNV = uintptr(getProcAddr("glMultiDrawArraysIndirectBindlessNV"))
	gpMultiDrawArraysIndirectCountARB = uintptr(getProcAddr("glMultiDrawArraysIndirectCountARB"))
	gpMultiDrawElementArrayAPPLE = uintptr(getProcAddr("glMultiDrawElementArrayAPPLE"))
	gpMultiDrawElements = uintptr(getProcAddr("glMultiDrawElements"))
	if gpMultiDrawElements == 0 {
		return errors.New("glMultiDrawElements")
	}
	gpMultiDrawElementsBaseVertex = uintptr(getProcAddr("glMultiDrawElementsBaseVertex"))
	gpMultiDrawElementsEXT = uintptr(getProcAddr("glMultiDrawElementsEXT"))
	gpMultiDrawElementsIndirect = uintptr(getProcAddr("glMultiDrawElementsIndirect"))
	gpMultiDrawElementsIndirectAMD = uintptr(getProcAddr("glMultiDrawElementsIndirectAMD"))
	gpMultiDrawElementsIndirectBindlessCountNV = uintptr(getProcAddr("glMultiDrawElementsIndirectBindlessCountNV"))
	gpMultiDrawElementsIndirectBindlessNV = uintptr(getProcAddr("glMultiDrawElementsIndirectBindlessNV"))
	gpMultiDrawElementsIndirectCountARB = uintptr(getProcAddr("glMultiDrawElementsIndirectCountARB"))
	gpMultiDrawRangeElementArrayAPPLE = uintptr(getProcAddr("glMultiDrawRangeElementArrayAPPLE"))
	gpMultiModeDrawArraysIBM = uintptr(getProcAddr("glMultiModeDrawArraysIBM"))
	gpMultiModeDrawElementsIBM = uintptr(getProcAddr("glMultiModeDrawElementsIBM"))
	gpMultiTexBufferEXT = uintptr(getProcAddr("glMultiTexBufferEXT"))
	gpMultiTexCoord1bOES = uintptr(getProcAddr("glMultiTexCoord1bOES"))
	gpMultiTexCoord1bvOES = uintptr(getProcAddr("glMultiTexCoord1bvOES"))
	gpMultiTexCoord1d = uintptr(getProcAddr("glMultiTexCoord1d"))
	if gpMultiTexCoord1d == 0 {
		return errors.New("glMultiTexCoord1d")
	}
	gpMultiTexCoord1dARB = uintptr(getProcAddr("glMultiTexCoord1dARB"))
	gpMultiTexCoord1dv = uintptr(getProcAddr("glMultiTexCoord1dv"))
	if gpMultiTexCoord1dv == 0 {
		return errors.New("glMultiTexCoord1dv")
	}
	gpMultiTexCoord1dvARB = uintptr(getProcAddr("glMultiTexCoord1dvARB"))
	gpMultiTexCoord1f = uintptr(getProcAddr("glMultiTexCoord1f"))
	if gpMultiTexCoord1f == 0 {
		return errors.New("glMultiTexCoord1f")
	}
	gpMultiTexCoord1fARB = uintptr(getProcAddr("glMultiTexCoord1fARB"))
	gpMultiTexCoord1fv = uintptr(getProcAddr("glMultiTexCoord1fv"))
	if gpMultiTexCoord1fv == 0 {
		return errors.New("glMultiTexCoord1fv")
	}
	gpMultiTexCoord1fvARB = uintptr(getProcAddr("glMultiTexCoord1fvARB"))
	gpMultiTexCoord1hNV = uintptr(getProcAddr("glMultiTexCoord1hNV"))
	gpMultiTexCoord1hvNV = uintptr(getProcAddr("glMultiTexCoord1hvNV"))
	gpMultiTexCoord1i = uintptr(getProcAddr("glMultiTexCoord1i"))
	if gpMultiTexCoord1i == 0 {
		return errors.New("glMultiTexCoord1i")
	}
	gpMultiTexCoord1iARB = uintptr(getProcAddr("glMultiTexCoord1iARB"))
	gpMultiTexCoord1iv = uintptr(getProcAddr("glMultiTexCoord1iv"))
	if gpMultiTexCoord1iv == 0 {
		return errors.New("glMultiTexCoord1iv")
	}
	gpMultiTexCoord1ivARB = uintptr(getProcAddr("glMultiTexCoord1ivARB"))
	gpMultiTexCoord1s = uintptr(getProcAddr("glMultiTexCoord1s"))
	if gpMultiTexCoord1s == 0 {
		return errors.New("glMultiTexCoord1s")
	}
	gpMultiTexCoord1sARB = uintptr(getProcAddr("glMultiTexCoord1sARB"))
	gpMultiTexCoord1sv = uintptr(getProcAddr("glMultiTexCoord1sv"))
	if gpMultiTexCoord1sv == 0 {
		return errors.New("glMultiTexCoord1sv")
	}
	gpMultiTexCoord1svARB = uintptr(getProcAddr("glMultiTexCoord1svARB"))
	gpMultiTexCoord1xOES = uintptr(getProcAddr("glMultiTexCoord1xOES"))
	gpMultiTexCoord1xvOES = uintptr(getProcAddr("glMultiTexCoord1xvOES"))
	gpMultiTexCoord2bOES = uintptr(getProcAddr("glMultiTexCoord2bOES"))
	gpMultiTexCoord2bvOES = uintptr(getProcAddr("glMultiTexCoord2bvOES"))
	gpMultiTexCoord2d = uintptr(getProcAddr("glMultiTexCoord2d"))
	if gpMultiTexCoord2d == 0 {
		return errors.New("glMultiTexCoord2d")
	}
	gpMultiTexCoord2dARB = uintptr(getProcAddr("glMultiTexCoord2dARB"))
	gpMultiTexCoord2dv = uintptr(getProcAddr("glMultiTexCoord2dv"))
	if gpMultiTexCoord2dv == 0 {
		return errors.New("glMultiTexCoord2dv")
	}
	gpMultiTexCoord2dvARB = uintptr(getProcAddr("glMultiTexCoord2dvARB"))
	gpMultiTexCoord2f = uintptr(getProcAddr("glMultiTexCoord2f"))
	if gpMultiTexCoord2f == 0 {
		return errors.New("glMultiTexCoord2f")
	}
	gpMultiTexCoord2fARB = uintptr(getProcAddr("glMultiTexCoord2fARB"))
	gpMultiTexCoord2fv = uintptr(getProcAddr("glMultiTexCoord2fv"))
	if gpMultiTexCoord2fv == 0 {
		return errors.New("glMultiTexCoord2fv")
	}
	gpMultiTexCoord2fvARB = uintptr(getProcAddr("glMultiTexCoord2fvARB"))
	gpMultiTexCoord2hNV = uintptr(getProcAddr("glMultiTexCoord2hNV"))
	gpMultiTexCoord2hvNV = uintptr(getProcAddr("glMultiTexCoord2hvNV"))
	gpMultiTexCoord2i = uintptr(getProcAddr("glMultiTexCoord2i"))
	if gpMultiTexCoord2i == 0 {
		return errors.New("glMultiTexCoord2i")
	}
	gpMultiTexCoord2iARB = uintptr(getProcAddr("glMultiTexCoord2iARB"))
	gpMultiTexCoord2iv = uintptr(getProcAddr("glMultiTexCoord2iv"))
	if gpMultiTexCoord2iv == 0 {
		return errors.New("glMultiTexCoord2iv")
	}
	gpMultiTexCoord2ivARB = uintptr(getProcAddr("glMultiTexCoord2ivARB"))
	gpMultiTexCoord2s = uintptr(getProcAddr("glMultiTexCoord2s"))
	if gpMultiTexCoord2s == 0 {
		return errors.New("glMultiTexCoord2s")
	}
	gpMultiTexCoord2sARB = uintptr(getProcAddr("glMultiTexCoord2sARB"))
	gpMultiTexCoord2sv = uintptr(getProcAddr("glMultiTexCoord2sv"))
	if gpMultiTexCoord2sv == 0 {
		return errors.New("glMultiTexCoord2sv")
	}
	gpMultiTexCoord2svARB = uintptr(getProcAddr("glMultiTexCoord2svARB"))
	gpMultiTexCoord2xOES = uintptr(getProcAddr("glMultiTexCoord2xOES"))
	gpMultiTexCoord2xvOES = uintptr(getProcAddr("glMultiTexCoord2xvOES"))
	gpMultiTexCoord3bOES = uintptr(getProcAddr("glMultiTexCoord3bOES"))
	gpMultiTexCoord3bvOES = uintptr(getProcAddr("glMultiTexCoord3bvOES"))
	gpMultiTexCoord3d = uintptr(getProcAddr("glMultiTexCoord3d"))
	if gpMultiTexCoord3d == 0 {
		return errors.New("glMultiTexCoord3d")
	}
	gpMultiTexCoord3dARB = uintptr(getProcAddr("glMultiTexCoord3dARB"))
	gpMultiTexCoord3dv = uintptr(getProcAddr("glMultiTexCoord3dv"))
	if gpMultiTexCoord3dv == 0 {
		return errors.New("glMultiTexCoord3dv")
	}
	gpMultiTexCoord3dvARB = uintptr(getProcAddr("glMultiTexCoord3dvARB"))
	gpMultiTexCoord3f = uintptr(getProcAddr("glMultiTexCoord3f"))
	if gpMultiTexCoord3f == 0 {
		return errors.New("glMultiTexCoord3f")
	}
	gpMultiTexCoord3fARB = uintptr(getProcAddr("glMultiTexCoord3fARB"))
	gpMultiTexCoord3fv = uintptr(getProcAddr("glMultiTexCoord3fv"))
	if gpMultiTexCoord3fv == 0 {
		return errors.New("glMultiTexCoord3fv")
	}
	gpMultiTexCoord3fvARB = uintptr(getProcAddr("glMultiTexCoord3fvARB"))
	gpMultiTexCoord3hNV = uintptr(getProcAddr("glMultiTexCoord3hNV"))
	gpMultiTexCoord3hvNV = uintptr(getProcAddr("glMultiTexCoord3hvNV"))
	gpMultiTexCoord3i = uintptr(getProcAddr("glMultiTexCoord3i"))
	if gpMultiTexCoord3i == 0 {
		return errors.New("glMultiTexCoord3i")
	}
	gpMultiTexCoord3iARB = uintptr(getProcAddr("glMultiTexCoord3iARB"))
	gpMultiTexCoord3iv = uintptr(getProcAddr("glMultiTexCoord3iv"))
	if gpMultiTexCoord3iv == 0 {
		return errors.New("glMultiTexCoord3iv")
	}
	gpMultiTexCoord3ivARB = uintptr(getProcAddr("glMultiTexCoord3ivARB"))
	gpMultiTexCoord3s = uintptr(getProcAddr("glMultiTexCoord3s"))
	if gpMultiTexCoord3s == 0 {
		return errors.New("glMultiTexCoord3s")
	}
	gpMultiTexCoord3sARB = uintptr(getProcAddr("glMultiTexCoord3sARB"))
	gpMultiTexCoord3sv = uintptr(getProcAddr("glMultiTexCoord3sv"))
	if gpMultiTexCoord3sv == 0 {
		return errors.New("glMultiTexCoord3sv")
	}
	gpMultiTexCoord3svARB = uintptr(getProcAddr("glMultiTexCoord3svARB"))
	gpMultiTexCoord3xOES = uintptr(getProcAddr("glMultiTexCoord3xOES"))
	gpMultiTexCoord3xvOES = uintptr(getProcAddr("glMultiTexCoord3xvOES"))
	gpMultiTexCoord4bOES = uintptr(getProcAddr("glMultiTexCoord4bOES"))
	gpMultiTexCoord4bvOES = uintptr(getProcAddr("glMultiTexCoord4bvOES"))
	gpMultiTexCoord4d = uintptr(getProcAddr("glMultiTexCoord4d"))
	if gpMultiTexCoord4d == 0 {
		return errors.New("glMultiTexCoord4d")
	}
	gpMultiTexCoord4dARB = uintptr(getProcAddr("glMultiTexCoord4dARB"))
	gpMultiTexCoord4dv = uintptr(getProcAddr("glMultiTexCoord4dv"))
	if gpMultiTexCoord4dv == 0 {
		return errors.New("glMultiTexCoord4dv")
	}
	gpMultiTexCoord4dvARB = uintptr(getProcAddr("glMultiTexCoord4dvARB"))
	gpMultiTexCoord4f = uintptr(getProcAddr("glMultiTexCoord4f"))
	if gpMultiTexCoord4f == 0 {
		return errors.New("glMultiTexCoord4f")
	}
	gpMultiTexCoord4fARB = uintptr(getProcAddr("glMultiTexCoord4fARB"))
	gpMultiTexCoord4fv = uintptr(getProcAddr("glMultiTexCoord4fv"))
	if gpMultiTexCoord4fv == 0 {
		return errors.New("glMultiTexCoord4fv")
	}
	gpMultiTexCoord4fvARB = uintptr(getProcAddr("glMultiTexCoord4fvARB"))
	gpMultiTexCoord4hNV = uintptr(getProcAddr("glMultiTexCoord4hNV"))
	gpMultiTexCoord4hvNV = uintptr(getProcAddr("glMultiTexCoord4hvNV"))
	gpMultiTexCoord4i = uintptr(getProcAddr("glMultiTexCoord4i"))
	if gpMultiTexCoord4i == 0 {
		return errors.New("glMultiTexCoord4i")
	}
	gpMultiTexCoord4iARB = uintptr(getProcAddr("glMultiTexCoord4iARB"))
	gpMultiTexCoord4iv = uintptr(getProcAddr("glMultiTexCoord4iv"))
	if gpMultiTexCoord4iv == 0 {
		return errors.New("glMultiTexCoord4iv")
	}
	gpMultiTexCoord4ivARB = uintptr(getProcAddr("glMultiTexCoord4ivARB"))
	gpMultiTexCoord4s = uintptr(getProcAddr("glMultiTexCoord4s"))
	if gpMultiTexCoord4s == 0 {
		return errors.New("glMultiTexCoord4s")
	}
	gpMultiTexCoord4sARB = uintptr(getProcAddr("glMultiTexCoord4sARB"))
	gpMultiTexCoord4sv = uintptr(getProcAddr("glMultiTexCoord4sv"))
	if gpMultiTexCoord4sv == 0 {
		return errors.New("glMultiTexCoord4sv")
	}
	gpMultiTexCoord4svARB = uintptr(getProcAddr("glMultiTexCoord4svARB"))
	gpMultiTexCoord4xOES = uintptr(getProcAddr("glMultiTexCoord4xOES"))
	gpMultiTexCoord4xvOES = uintptr(getProcAddr("glMultiTexCoord4xvOES"))
	gpMultiTexCoordPointerEXT = uintptr(getProcAddr("glMultiTexCoordPointerEXT"))
	gpMultiTexEnvfEXT = uintptr(getProcAddr("glMultiTexEnvfEXT"))
	gpMultiTexEnvfvEXT = uintptr(getProcAddr("glMultiTexEnvfvEXT"))
	gpMultiTexEnviEXT = uintptr(getProcAddr("glMultiTexEnviEXT"))
	gpMultiTexEnvivEXT = uintptr(getProcAddr("glMultiTexEnvivEXT"))
	gpMultiTexGendEXT = uintptr(getProcAddr("glMultiTexGendEXT"))
	gpMultiTexGendvEXT = uintptr(getProcAddr("glMultiTexGendvEXT"))
	gpMultiTexGenfEXT = uintptr(getProcAddr("glMultiTexGenfEXT"))
	gpMultiTexGenfvEXT = uintptr(getProcAddr("glMultiTexGenfvEXT"))
	gpMultiTexGeniEXT = uintptr(getProcAddr("glMultiTexGeniEXT"))
	gpMultiTexGenivEXT = uintptr(getProcAddr("glMultiTexGenivEXT"))
	gpMultiTexImage1DEXT = uintptr(getProcAddr("glMultiTexImage1DEXT"))
	gpMultiTexImage2DEXT = uintptr(getProcAddr("glMultiTexImage2DEXT"))
	gpMultiTexImage3DEXT = uintptr(getProcAddr("glMultiTexImage3DEXT"))
	gpMultiTexParameterIivEXT = uintptr(getProcAddr("glMultiTexParameterIivEXT"))
	gpMultiTexParameterIuivEXT = uintptr(getProcAddr("glMultiTexParameterIuivEXT"))
	gpMultiTexParameterfEXT = uintptr(getProcAddr("glMultiTexParameterfEXT"))
	gpMultiTexParameterfvEXT = uintptr(getProcAddr("glMultiTexParameterfvEXT"))
	gpMultiTexParameteriEXT = uintptr(getProcAddr("glMultiTexParameteriEXT"))
	gpMultiTexParameterivEXT = uintptr(getProcAddr("glMultiTexParameterivEXT"))
	gpMultiTexRenderbufferEXT = uintptr(getProcAddr("glMultiTexRenderbufferEXT"))
	gpMultiTexSubImage1DEXT = uintptr(getProcAddr("glMultiTexSubImage1DEXT"))
	gpMultiTexSubImage2DEXT = uintptr(getProcAddr("glMultiTexSubImage2DEXT"))
	gpMultiTexSubImage3DEXT = uintptr(getProcAddr("glMultiTexSubImage3DEXT"))
	gpMulticastBarrierNV = uintptr(getProcAddr("glMulticastBarrierNV"))
	gpMulticastBlitFramebufferNV = uintptr(getProcAddr("glMulticastBlitFramebufferNV"))
	gpMulticastBufferSubDataNV = uintptr(getProcAddr("glMulticastBufferSubDataNV"))
	gpMulticastCopyBufferSubDataNV = uintptr(getProcAddr("glMulticastCopyBufferSubDataNV"))
	gpMulticastCopyImageSubDataNV = uintptr(getProcAddr("glMulticastCopyImageSubDataNV"))
	gpMulticastFramebufferSampleLocationsfvNV = uintptr(getProcAddr("glMulticastFramebufferSampleLocationsfvNV"))
	gpMulticastGetQueryObjecti64vNV = uintptr(getProcAddr("glMulticastGetQueryObjecti64vNV"))
	gpMulticastGetQueryObjectivNV = uintptr(getProcAddr("glMulticastGetQueryObjectivNV"))
	gpMulticastGetQueryObjectui64vNV = uintptr(getProcAddr("glMulticastGetQueryObjectui64vNV"))
	gpMulticastGetQueryObjectuivNV = uintptr(getProcAddr("glMulticastGetQueryObjectuivNV"))
	gpMulticastWaitSyncNV = uintptr(getProcAddr("glMulticastWaitSyncNV"))
	gpNamedBufferData = uintptr(getProcAddr("glNamedBufferData"))
	gpNamedBufferDataEXT = uintptr(getProcAddr("glNamedBufferDataEXT"))
	gpNamedBufferPageCommitmentARB = uintptr(getProcAddr("glNamedBufferPageCommitmentARB"))
	gpNamedBufferPageCommitmentEXT = uintptr(getProcAddr("glNamedBufferPageCommitmentEXT"))
	gpNamedBufferStorage = uintptr(getProcAddr("glNamedBufferStorage"))
	gpNamedBufferStorageEXT = uintptr(getProcAddr("glNamedBufferStorageEXT"))
	gpNamedBufferStorageExternalEXT = uintptr(getProcAddr("glNamedBufferStorageExternalEXT"))
	gpNamedBufferStorageMemEXT = uintptr(getProcAddr("glNamedBufferStorageMemEXT"))
	gpNamedBufferSubData = uintptr(getProcAddr("glNamedBufferSubData"))
	gpNamedBufferSubDataEXT = uintptr(getProcAddr("glNamedBufferSubDataEXT"))
	gpNamedCopyBufferSubDataEXT = uintptr(getProcAddr("glNamedCopyBufferSubDataEXT"))
	gpNamedFramebufferDrawBuffer = uintptr(getProcAddr("glNamedFramebufferDrawBuffer"))
	gpNamedFramebufferDrawBuffers = uintptr(getProcAddr("glNamedFramebufferDrawBuffers"))
	gpNamedFramebufferParameteri = uintptr(getProcAddr("glNamedFramebufferParameteri"))
	gpNamedFramebufferParameteriEXT = uintptr(getProcAddr("glNamedFramebufferParameteriEXT"))
	gpNamedFramebufferReadBuffer = uintptr(getProcAddr("glNamedFramebufferReadBuffer"))
	gpNamedFramebufferRenderbuffer = uintptr(getProcAddr("glNamedFramebufferRenderbuffer"))
	gpNamedFramebufferRenderbufferEXT = uintptr(getProcAddr("glNamedFramebufferRenderbufferEXT"))
	gpNamedFramebufferSampleLocationsfvARB = uintptr(getProcAddr("glNamedFramebufferSampleLocationsfvARB"))
	gpNamedFramebufferSampleLocationsfvNV = uintptr(getProcAddr("glNamedFramebufferSampleLocationsfvNV"))
	gpNamedFramebufferSamplePositionsfvAMD = uintptr(getProcAddr("glNamedFramebufferSamplePositionsfvAMD"))
	gpNamedFramebufferTexture = uintptr(getProcAddr("glNamedFramebufferTexture"))
	gpNamedFramebufferTexture1DEXT = uintptr(getProcAddr("glNamedFramebufferTexture1DEXT"))
	gpNamedFramebufferTexture2DEXT = uintptr(getProcAddr("glNamedFramebufferTexture2DEXT"))
	gpNamedFramebufferTexture3DEXT = uintptr(getProcAddr("glNamedFramebufferTexture3DEXT"))
	gpNamedFramebufferTextureEXT = uintptr(getProcAddr("glNamedFramebufferTextureEXT"))
	gpNamedFramebufferTextureFaceEXT = uintptr(getProcAddr("glNamedFramebufferTextureFaceEXT"))
	gpNamedFramebufferTextureLayer = uintptr(getProcAddr("glNamedFramebufferTextureLayer"))
	gpNamedFramebufferTextureLayerEXT = uintptr(getProcAddr("glNamedFramebufferTextureLayerEXT"))
	gpNamedProgramLocalParameter4dEXT = uintptr(getProcAddr("glNamedProgramLocalParameter4dEXT"))
	gpNamedProgramLocalParameter4dvEXT = uintptr(getProcAddr("glNamedProgramLocalParameter4dvEXT"))
	gpNamedProgramLocalParameter4fEXT = uintptr(getProcAddr("glNamedProgramLocalParameter4fEXT"))
	gpNamedProgramLocalParameter4fvEXT = uintptr(getProcAddr("glNamedProgramLocalParameter4fvEXT"))
	gpNamedProgramLocalParameterI4iEXT = uintptr(getProcAddr("glNamedProgramLocalParameterI4iEXT"))
	gpNamedProgramLocalParameterI4ivEXT = uintptr(getProcAddr("glNamedProgramLocalParameterI4ivEXT"))
	gpNamedProgramLocalParameterI4uiEXT = uintptr(getProcAddr("glNamedProgramLocalParameterI4uiEXT"))
	gpNamedProgramLocalParameterI4uivEXT = uintptr(getProcAddr("glNamedProgramLocalParameterI4uivEXT"))
	gpNamedProgramLocalParameters4fvEXT = uintptr(getProcAddr("glNamedProgramLocalParameters4fvEXT"))
	gpNamedProgramLocalParametersI4ivEXT = uintptr(getProcAddr("glNamedProgramLocalParametersI4ivEXT"))
	gpNamedProgramLocalParametersI4uivEXT = uintptr(getProcAddr("glNamedProgramLocalParametersI4uivEXT"))
	gpNamedProgramStringEXT = uintptr(getProcAddr("glNamedProgramStringEXT"))
	gpNamedRenderbufferStorage = uintptr(getProcAddr("glNamedRenderbufferStorage"))
	gpNamedRenderbufferStorageEXT = uintptr(getProcAddr("glNamedRenderbufferStorageEXT"))
	gpNamedRenderbufferStorageMultisample = uintptr(getProcAddr("glNamedRenderbufferStorageMultisample"))
	gpNamedRenderbufferStorageMultisampleCoverageEXT = uintptr(getProcAddr("glNamedRenderbufferStorageMultisampleCoverageEXT"))
	gpNamedRenderbufferStorageMultisampleEXT = uintptr(getProcAddr("glNamedRenderbufferStorageMultisampleEXT"))
	gpNamedStringARB = uintptr(getProcAddr("glNamedStringARB"))
	gpNewList = uintptr(getProcAddr("glNewList"))
	if gpNewList == 0 {
		return errors.New("glNewList")
	}
	gpNewObjectBufferATI = uintptr(getProcAddr("glNewObjectBufferATI"))
	gpNormal3b = uintptr(getProcAddr("glNormal3b"))
	if gpNormal3b == 0 {
		return errors.New("glNormal3b")
	}
	gpNormal3bv = uintptr(getProcAddr("glNormal3bv"))
	if gpNormal3bv == 0 {
		return errors.New("glNormal3bv")
	}
	gpNormal3d = uintptr(getProcAddr("glNormal3d"))
	if gpNormal3d == 0 {
		return errors.New("glNormal3d")
	}
	gpNormal3dv = uintptr(getProcAddr("glNormal3dv"))
	if gpNormal3dv == 0 {
		return errors.New("glNormal3dv")
	}
	gpNormal3f = uintptr(getProcAddr("glNormal3f"))
	if gpNormal3f == 0 {
		return errors.New("glNormal3f")
	}
	gpNormal3fVertex3fSUN = uintptr(getProcAddr("glNormal3fVertex3fSUN"))
	gpNormal3fVertex3fvSUN = uintptr(getProcAddr("glNormal3fVertex3fvSUN"))
	gpNormal3fv = uintptr(getProcAddr("glNormal3fv"))
	if gpNormal3fv == 0 {
		return errors.New("glNormal3fv")
	}
	gpNormal3hNV = uintptr(getProcAddr("glNormal3hNV"))
	gpNormal3hvNV = uintptr(getProcAddr("glNormal3hvNV"))
	gpNormal3i = uintptr(getProcAddr("glNormal3i"))
	if gpNormal3i == 0 {
		return errors.New("glNormal3i")
	}
	gpNormal3iv = uintptr(getProcAddr("glNormal3iv"))
	if gpNormal3iv == 0 {
		return errors.New("glNormal3iv")
	}
	gpNormal3s = uintptr(getProcAddr("glNormal3s"))
	if gpNormal3s == 0 {
		return errors.New("glNormal3s")
	}
	gpNormal3sv = uintptr(getProcAddr("glNormal3sv"))
	if gpNormal3sv == 0 {
		return errors.New("glNormal3sv")
	}
	gpNormal3xOES = uintptr(getProcAddr("glNormal3xOES"))
	gpNormal3xvOES = uintptr(getProcAddr("glNormal3xvOES"))
	gpNormalFormatNV = uintptr(getProcAddr("glNormalFormatNV"))
	gpNormalPointer = uintptr(getProcAddr("glNormalPointer"))
	if gpNormalPointer == 0 {
		return errors.New("glNormalPointer")
	}
	gpNormalPointerEXT = uintptr(getProcAddr("glNormalPointerEXT"))
	gpNormalPointerListIBM = uintptr(getProcAddr("glNormalPointerListIBM"))
	gpNormalPointervINTEL = uintptr(getProcAddr("glNormalPointervINTEL"))
	gpNormalStream3bATI = uintptr(getProcAddr("glNormalStream3bATI"))
	gpNormalStream3bvATI = uintptr(getProcAddr("glNormalStream3bvATI"))
	gpNormalStream3dATI = uintptr(getProcAddr("glNormalStream3dATI"))
	gpNormalStream3dvATI = uintptr(getProcAddr("glNormalStream3dvATI"))
	gpNormalStream3fATI = uintptr(getProcAddr("glNormalStream3fATI"))
	gpNormalStream3fvATI = uintptr(getProcAddr("glNormalStream3fvATI"))
	gpNormalStream3iATI = uintptr(getProcAddr("glNormalStream3iATI"))
	gpNormalStream3ivATI = uintptr(getProcAddr("glNormalStream3ivATI"))
	gpNormalStream3sATI = uintptr(getProcAddr("glNormalStream3sATI"))
	gpNormalStream3svATI = uintptr(getProcAddr("glNormalStream3svATI"))
	gpObjectLabel = uintptr(getProcAddr("glObjectLabel"))
	gpObjectLabelKHR = uintptr(getProcAddr("glObjectLabelKHR"))
	gpObjectPtrLabel = uintptr(getProcAddr("glObjectPtrLabel"))
	gpObjectPtrLabelKHR = uintptr(getProcAddr("glObjectPtrLabelKHR"))
	gpObjectPurgeableAPPLE = uintptr(getProcAddr("glObjectPurgeableAPPLE"))
	gpObjectUnpurgeableAPPLE = uintptr(getProcAddr("glObjectUnpurgeableAPPLE"))
	gpOrtho = uintptr(getProcAddr("glOrtho"))
	if gpOrtho == 0 {
		return errors.New("glOrtho")
	}
	gpOrthofOES = uintptr(getProcAddr("glOrthofOES"))
	gpOrthoxOES = uintptr(getProcAddr("glOrthoxOES"))
	gpPNTrianglesfATI = uintptr(getProcAddr("glPNTrianglesfATI"))
	gpPNTrianglesiATI = uintptr(getProcAddr("glPNTrianglesiATI"))
	gpPassTexCoordATI = uintptr(getProcAddr("glPassTexCoordATI"))
	gpPassThrough = uintptr(getProcAddr("glPassThrough"))
	if gpPassThrough == 0 {
		return errors.New("glPassThrough")
	}
	gpPassThroughxOES = uintptr(getProcAddr("glPassThroughxOES"))
	gpPatchParameterfv = uintptr(getProcAddr("glPatchParameterfv"))
	gpPatchParameteri = uintptr(getProcAddr("glPatchParameteri"))
	gpPathCommandsNV = uintptr(getProcAddr("glPathCommandsNV"))
	gpPathCoordsNV = uintptr(getProcAddr("glPathCoordsNV"))
	gpPathCoverDepthFuncNV = uintptr(getProcAddr("glPathCoverDepthFuncNV"))
	gpPathDashArrayNV = uintptr(getProcAddr("glPathDashArrayNV"))
	gpPathGlyphIndexArrayNV = uintptr(getProcAddr("glPathGlyphIndexArrayNV"))
	gpPathGlyphIndexRangeNV = uintptr(getProcAddr("glPathGlyphIndexRangeNV"))
	gpPathGlyphRangeNV = uintptr(getProcAddr("glPathGlyphRangeNV"))
	gpPathGlyphsNV = uintptr(getProcAddr("glPathGlyphsNV"))
	gpPathMemoryGlyphIndexArrayNV = uintptr(getProcAddr("glPathMemoryGlyphIndexArrayNV"))
	gpPathParameterfNV = uintptr(getProcAddr("glPathParameterfNV"))
	gpPathParameterfvNV = uintptr(getProcAddr("glPathParameterfvNV"))
	gpPathParameteriNV = uintptr(getProcAddr("glPathParameteriNV"))
	gpPathParameterivNV = uintptr(getProcAddr("glPathParameterivNV"))
	gpPathStencilDepthOffsetNV = uintptr(getProcAddr("glPathStencilDepthOffsetNV"))
	gpPathStencilFuncNV = uintptr(getProcAddr("glPathStencilFuncNV"))
	gpPathStringNV = uintptr(getProcAddr("glPathStringNV"))
	gpPathSubCommandsNV = uintptr(getProcAddr("glPathSubCommandsNV"))
	gpPathSubCoordsNV = uintptr(getProcAddr("glPathSubCoordsNV"))
	gpPauseTransformFeedback = uintptr(getProcAddr("glPauseTransformFeedback"))
	gpPauseTransformFeedbackNV = uintptr(getProcAddr("glPauseTransformFeedbackNV"))
	gpPixelDataRangeNV = uintptr(getProcAddr("glPixelDataRangeNV"))
	gpPixelMapfv = uintptr(getProcAddr("glPixelMapfv"))
	if gpPixelMapfv == 0 {
		return errors.New("glPixelMapfv")
	}
	gpPixelMapuiv = uintptr(getProcAddr("glPixelMapuiv"))
	if gpPixelMapuiv == 0 {
		return errors.New("glPixelMapuiv")
	}
	gpPixelMapusv = uintptr(getProcAddr("glPixelMapusv"))
	if gpPixelMapusv == 0 {
		return errors.New("glPixelMapusv")
	}
	gpPixelMapx = uintptr(getProcAddr("glPixelMapx"))
	gpPixelStoref = uintptr(getProcAddr("glPixelStoref"))
	if gpPixelStoref == 0 {
		return errors.New("glPixelStoref")
	}
	gpPixelStorei = uintptr(getProcAddr("glPixelStorei"))
	if gpPixelStorei == 0 {
		return errors.New("glPixelStorei")
	}
	gpPixelStorex = uintptr(getProcAddr("glPixelStorex"))
	gpPixelTexGenParameterfSGIS = uintptr(getProcAddr("glPixelTexGenParameterfSGIS"))
	gpPixelTexGenParameterfvSGIS = uintptr(getProcAddr("glPixelTexGenParameterfvSGIS"))
	gpPixelTexGenParameteriSGIS = uintptr(getProcAddr("glPixelTexGenParameteriSGIS"))
	gpPixelTexGenParameterivSGIS = uintptr(getProcAddr("glPixelTexGenParameterivSGIS"))
	gpPixelTexGenSGIX = uintptr(getProcAddr("glPixelTexGenSGIX"))
	gpPixelTransferf = uintptr(getProcAddr("glPixelTransferf"))
	if gpPixelTransferf == 0 {
		return errors.New("glPixelTransferf")
	}
	gpPixelTransferi = uintptr(getProcAddr("glPixelTransferi"))
	if gpPixelTransferi == 0 {
		return errors.New("glPixelTransferi")
	}
	gpPixelTransferxOES = uintptr(getProcAddr("glPixelTransferxOES"))
	gpPixelTransformParameterfEXT = uintptr(getProcAddr("glPixelTransformParameterfEXT"))
	gpPixelTransformParameterfvEXT = uintptr(getProcAddr("glPixelTransformParameterfvEXT"))
	gpPixelTransformParameteriEXT = uintptr(getProcAddr("glPixelTransformParameteriEXT"))
	gpPixelTransformParameterivEXT = uintptr(getProcAddr("glPixelTransformParameterivEXT"))
	gpPixelZoom = uintptr(getProcAddr("glPixelZoom"))
	if gpPixelZoom == 0 {
		return errors.New("glPixelZoom")
	}
	gpPixelZoomxOES = uintptr(getProcAddr("glPixelZoomxOES"))
	gpPointAlongPathNV = uintptr(getProcAddr("glPointAlongPathNV"))
	gpPointParameterf = uintptr(getProcAddr("glPointParameterf"))
	if gpPointParameterf == 0 {
		return errors.New("glPointParameterf")
	}
	gpPointParameterfARB = uintptr(getProcAddr("glPointParameterfARB"))
	gpPointParameterfEXT = uintptr(getProcAddr("glPointParameterfEXT"))
	gpPointParameterfSGIS = uintptr(getProcAddr("glPointParameterfSGIS"))
	gpPointParameterfv = uintptr(getProcAddr("glPointParameterfv"))
	if gpPointParameterfv == 0 {
		return errors.New("glPointParameterfv")
	}
	gpPointParameterfvARB = uintptr(getProcAddr("glPointParameterfvARB"))
	gpPointParameterfvEXT = uintptr(getProcAddr("glPointParameterfvEXT"))
	gpPointParameterfvSGIS = uintptr(getProcAddr("glPointParameterfvSGIS"))
	gpPointParameteri = uintptr(getProcAddr("glPointParameteri"))
	if gpPointParameteri == 0 {
		return errors.New("glPointParameteri")
	}
	gpPointParameteriNV = uintptr(getProcAddr("glPointParameteriNV"))
	gpPointParameteriv = uintptr(getProcAddr("glPointParameteriv"))
	if gpPointParameteriv == 0 {
		return errors.New("glPointParameteriv")
	}
	gpPointParameterivNV = uintptr(getProcAddr("glPointParameterivNV"))
	gpPointParameterxOES = uintptr(getProcAddr("glPointParameterxOES"))
	gpPointParameterxvOES = uintptr(getProcAddr("glPointParameterxvOES"))
	gpPointSize = uintptr(getProcAddr("glPointSize"))
	if gpPointSize == 0 {
		return errors.New("glPointSize")
	}
	gpPointSizexOES = uintptr(getProcAddr("glPointSizexOES"))
	gpPollAsyncSGIX = uintptr(getProcAddr("glPollAsyncSGIX"))
	gpPollInstrumentsSGIX = uintptr(getProcAddr("glPollInstrumentsSGIX"))
	gpPolygonMode = uintptr(getProcAddr("glPolygonMode"))
	if gpPolygonMode == 0 {
		return errors.New("glPolygonMode")
	}
	gpPolygonOffset = uintptr(getProcAddr("glPolygonOffset"))
	if gpPolygonOffset == 0 {
		return errors.New("glPolygonOffset")
	}
	gpPolygonOffsetClamp = uintptr(getProcAddr("glPolygonOffsetClamp"))
	gpPolygonOffsetClampEXT = uintptr(getProcAddr("glPolygonOffsetClampEXT"))
	gpPolygonOffsetEXT = uintptr(getProcAddr("glPolygonOffsetEXT"))
	gpPolygonOffsetxOES = uintptr(getProcAddr("glPolygonOffsetxOES"))
	gpPolygonStipple = uintptr(getProcAddr("glPolygonStipple"))
	if gpPolygonStipple == 0 {
		return errors.New("glPolygonStipple")
	}
	gpPopAttrib = uintptr(getProcAddr("glPopAttrib"))
	if gpPopAttrib == 0 {
		return errors.New("glPopAttrib")
	}
	gpPopClientAttrib = uintptr(getProcAddr("glPopClientAttrib"))
	if gpPopClientAttrib == 0 {
		return errors.New("glPopClientAttrib")
	}
	gpPopDebugGroup = uintptr(getProcAddr("glPopDebugGroup"))
	gpPopDebugGroupKHR = uintptr(getProcAddr("glPopDebugGroupKHR"))
	gpPopGroupMarkerEXT = uintptr(getProcAddr("glPopGroupMarkerEXT"))
	gpPopMatrix = uintptr(getProcAddr("glPopMatrix"))
	if gpPopMatrix == 0 {
		return errors.New("glPopMatrix")
	}
	gpPopName = uintptr(getProcAddr("glPopName"))
	if gpPopName == 0 {
		return errors.New("glPopName")
	}
	gpPresentFrameDualFillNV = uintptr(getProcAddr("glPresentFrameDualFillNV"))
	gpPresentFrameKeyedNV = uintptr(getProcAddr("glPresentFrameKeyedNV"))
	gpPrimitiveBoundingBoxARB = uintptr(getProcAddr("glPrimitiveBoundingBoxARB"))
	gpPrimitiveRestartIndexNV = uintptr(getProcAddr("glPrimitiveRestartIndexNV"))
	gpPrimitiveRestartNV = uintptr(getProcAddr("glPrimitiveRestartNV"))
	gpPrioritizeTextures = uintptr(getProcAddr("glPrioritizeTextures"))
	if gpPrioritizeTextures == 0 {
		return errors.New("glPrioritizeTextures")
	}
	gpPrioritizeTexturesEXT = uintptr(getProcAddr("glPrioritizeTexturesEXT"))
	gpPrioritizeTexturesxOES = uintptr(getProcAddr("glPrioritizeTexturesxOES"))
	gpProgramBinary = uintptr(getProcAddr("glProgramBinary"))
	gpProgramBufferParametersIivNV = uintptr(getProcAddr("glProgramBufferParametersIivNV"))
	gpProgramBufferParametersIuivNV = uintptr(getProcAddr("glProgramBufferParametersIuivNV"))
	gpProgramBufferParametersfvNV = uintptr(getProcAddr("glProgramBufferParametersfvNV"))
	gpProgramEnvParameter4dARB = uintptr(getProcAddr("glProgramEnvParameter4dARB"))
	gpProgramEnvParameter4dvARB = uintptr(getProcAddr("glProgramEnvParameter4dvARB"))
	gpProgramEnvParameter4fARB = uintptr(getProcAddr("glProgramEnvParameter4fARB"))
	gpProgramEnvParameter4fvARB = uintptr(getProcAddr("glProgramEnvParameter4fvARB"))
	gpProgramEnvParameterI4iNV = uintptr(getProcAddr("glProgramEnvParameterI4iNV"))
	gpProgramEnvParameterI4ivNV = uintptr(getProcAddr("glProgramEnvParameterI4ivNV"))
	gpProgramEnvParameterI4uiNV = uintptr(getProcAddr("glProgramEnvParameterI4uiNV"))
	gpProgramEnvParameterI4uivNV = uintptr(getProcAddr("glProgramEnvParameterI4uivNV"))
	gpProgramEnvParameters4fvEXT = uintptr(getProcAddr("glProgramEnvParameters4fvEXT"))
	gpProgramEnvParametersI4ivNV = uintptr(getProcAddr("glProgramEnvParametersI4ivNV"))
	gpProgramEnvParametersI4uivNV = uintptr(getProcAddr("glProgramEnvParametersI4uivNV"))
	gpProgramLocalParameter4dARB = uintptr(getProcAddr("glProgramLocalParameter4dARB"))
	gpProgramLocalParameter4dvARB = uintptr(getProcAddr("glProgramLocalParameter4dvARB"))
	gpProgramLocalParameter4fARB = uintptr(getProcAddr("glProgramLocalParameter4fARB"))
	gpProgramLocalParameter4fvARB = uintptr(getProcAddr("glProgramLocalParameter4fvARB"))
	gpProgramLocalParameterI4iNV = uintptr(getProcAddr("glProgramLocalParameterI4iNV"))
	gpProgramLocalParameterI4ivNV = uintptr(getProcAddr("glProgramLocalParameterI4ivNV"))
	gpProgramLocalParameterI4uiNV = uintptr(getProcAddr("glProgramLocalParameterI4uiNV"))
	gpProgramLocalParameterI4uivNV = uintptr(getProcAddr("glProgramLocalParameterI4uivNV"))
	gpProgramLocalParameters4fvEXT = uintptr(getProcAddr("glProgramLocalParameters4fvEXT"))
	gpProgramLocalParametersI4ivNV = uintptr(getProcAddr("glProgramLocalParametersI4ivNV"))
	gpProgramLocalParametersI4uivNV = uintptr(getProcAddr("glProgramLocalParametersI4uivNV"))
	gpProgramNamedParameter4dNV = uintptr(getProcAddr("glProgramNamedParameter4dNV"))
	gpProgramNamedParameter4dvNV = uintptr(getProcAddr("glProgramNamedParameter4dvNV"))
	gpProgramNamedParameter4fNV = uintptr(getProcAddr("glProgramNamedParameter4fNV"))
	gpProgramNamedParameter4fvNV = uintptr(getProcAddr("glProgramNamedParameter4fvNV"))
	gpProgramParameter4dNV = uintptr(getProcAddr("glProgramParameter4dNV"))
	gpProgramParameter4dvNV = uintptr(getProcAddr("glProgramParameter4dvNV"))
	gpProgramParameter4fNV = uintptr(getProcAddr("glProgramParameter4fNV"))
	gpProgramParameter4fvNV = uintptr(getProcAddr("glProgramParameter4fvNV"))
	gpProgramParameteri = uintptr(getProcAddr("glProgramParameteri"))
	gpProgramParameteriARB = uintptr(getProcAddr("glProgramParameteriARB"))
	gpProgramParameteriEXT = uintptr(getProcAddr("glProgramParameteriEXT"))
	gpProgramParameters4dvNV = uintptr(getProcAddr("glProgramParameters4dvNV"))
	gpProgramParameters4fvNV = uintptr(getProcAddr("glProgramParameters4fvNV"))
	gpProgramPathFragmentInputGenNV = uintptr(getProcAddr("glProgramPathFragmentInputGenNV"))
	gpProgramStringARB = uintptr(getProcAddr("glProgramStringARB"))
	gpProgramSubroutineParametersuivNV = uintptr(getProcAddr("glProgramSubroutineParametersuivNV"))
	gpProgramUniform1d = uintptr(getProcAddr("glProgramUniform1d"))
	gpProgramUniform1dEXT = uintptr(getProcAddr("glProgramUniform1dEXT"))
	gpProgramUniform1dv = uintptr(getProcAddr("glProgramUniform1dv"))
	gpProgramUniform1dvEXT = uintptr(getProcAddr("glProgramUniform1dvEXT"))
	gpProgramUniform1f = uintptr(getProcAddr("glProgramUniform1f"))
	gpProgramUniform1fEXT = uintptr(getProcAddr("glProgramUniform1fEXT"))
	gpProgramUniform1fv = uintptr(getProcAddr("glProgramUniform1fv"))
	gpProgramUniform1fvEXT = uintptr(getProcAddr("glProgramUniform1fvEXT"))
	gpProgramUniform1i = uintptr(getProcAddr("glProgramUniform1i"))
	gpProgramUniform1i64ARB = uintptr(getProcAddr("glProgramUniform1i64ARB"))
	gpProgramUniform1i64NV = uintptr(getProcAddr("glProgramUniform1i64NV"))
	gpProgramUniform1i64vARB = uintptr(getProcAddr("glProgramUniform1i64vARB"))
	gpProgramUniform1i64vNV = uintptr(getProcAddr("glProgramUniform1i64vNV"))
	gpProgramUniform1iEXT = uintptr(getProcAddr("glProgramUniform1iEXT"))
	gpProgramUniform1iv = uintptr(getProcAddr("glProgramUniform1iv"))
	gpProgramUniform1ivEXT = uintptr(getProcAddr("glProgramUniform1ivEXT"))
	gpProgramUniform1ui = uintptr(getProcAddr("glProgramUniform1ui"))
	gpProgramUniform1ui64ARB = uintptr(getProcAddr("glProgramUniform1ui64ARB"))
	gpProgramUniform1ui64NV = uintptr(getProcAddr("glProgramUniform1ui64NV"))
	gpProgramUniform1ui64vARB = uintptr(getProcAddr("glProgramUniform1ui64vARB"))
	gpProgramUniform1ui64vNV = uintptr(getProcAddr("glProgramUniform1ui64vNV"))
	gpProgramUniform1uiEXT = uintptr(getProcAddr("glProgramUniform1uiEXT"))
	gpProgramUniform1uiv = uintptr(getProcAddr("glProgramUniform1uiv"))
	gpProgramUniform1uivEXT = uintptr(getProcAddr("glProgramUniform1uivEXT"))
	gpProgramUniform2d = uintptr(getProcAddr("glProgramUniform2d"))
	gpProgramUniform2dEXT = uintptr(getProcAddr("glProgramUniform2dEXT"))
	gpProgramUniform2dv = uintptr(getProcAddr("glProgramUniform2dv"))
	gpProgramUniform2dvEXT = uintptr(getProcAddr("glProgramUniform2dvEXT"))
	gpProgramUniform2f = uintptr(getProcAddr("glProgramUniform2f"))
	gpProgramUniform2fEXT = uintptr(getProcAddr("glProgramUniform2fEXT"))
	gpProgramUniform2fv = uintptr(getProcAddr("glProgramUniform2fv"))
	gpProgramUniform2fvEXT = uintptr(getProcAddr("glProgramUniform2fvEXT"))
	gpProgramUniform2i = uintptr(getProcAddr("glProgramUniform2i"))
	gpProgramUniform2i64ARB = uintptr(getProcAddr("glProgramUniform2i64ARB"))
	gpProgramUniform2i64NV = uintptr(getProcAddr("glProgramUniform2i64NV"))
	gpProgramUniform2i64vARB = uintptr(getProcAddr("glProgramUniform2i64vARB"))
	gpProgramUniform2i64vNV = uintptr(getProcAddr("glProgramUniform2i64vNV"))
	gpProgramUniform2iEXT = uintptr(getProcAddr("glProgramUniform2iEXT"))
	gpProgramUniform2iv = uintptr(getProcAddr("glProgramUniform2iv"))
	gpProgramUniform2ivEXT = uintptr(getProcAddr("glProgramUniform2ivEXT"))
	gpProgramUniform2ui = uintptr(getProcAddr("glProgramUniform2ui"))
	gpProgramUniform2ui64ARB = uintptr(getProcAddr("glProgramUniform2ui64ARB"))
	gpProgramUniform2ui64NV = uintptr(getProcAddr("glProgramUniform2ui64NV"))
	gpProgramUniform2ui64vARB = uintptr(getProcAddr("glProgramUniform2ui64vARB"))
	gpProgramUniform2ui64vNV = uintptr(getProcAddr("glProgramUniform2ui64vNV"))
	gpProgramUniform2uiEXT = uintptr(getProcAddr("glProgramUniform2uiEXT"))
	gpProgramUniform2uiv = uintptr(getProcAddr("glProgramUniform2uiv"))
	gpProgramUniform2uivEXT = uintptr(getProcAddr("glProgramUniform2uivEXT"))
	gpProgramUniform3d = uintptr(getProcAddr("glProgramUniform3d"))
	gpProgramUniform3dEXT = uintptr(getProcAddr("glProgramUniform3dEXT"))
	gpProgramUniform3dv = uintptr(getProcAddr("glProgramUniform3dv"))
	gpProgramUniform3dvEXT = uintptr(getProcAddr("glProgramUniform3dvEXT"))
	gpProgramUniform3f = uintptr(getProcAddr("glProgramUniform3f"))
	gpProgramUniform3fEXT = uintptr(getProcAddr("glProgramUniform3fEXT"))
	gpProgramUniform3fv = uintptr(getProcAddr("glProgramUniform3fv"))
	gpProgramUniform3fvEXT = uintptr(getProcAddr("glProgramUniform3fvEXT"))
	gpProgramUniform3i = uintptr(getProcAddr("glProgramUniform3i"))
	gpProgramUniform3i64ARB = uintptr(getProcAddr("glProgramUniform3i64ARB"))
	gpProgramUniform3i64NV = uintptr(getProcAddr("glProgramUniform3i64NV"))
	gpProgramUniform3i64vARB = uintptr(getProcAddr("glProgramUniform3i64vARB"))
	gpProgramUniform3i64vNV = uintptr(getProcAddr("glProgramUniform3i64vNV"))
	gpProgramUniform3iEXT = uintptr(getProcAddr("glProgramUniform3iEXT"))
	gpProgramUniform3iv = uintptr(getProcAddr("glProgramUniform3iv"))
	gpProgramUniform3ivEXT = uintptr(getProcAddr("glProgramUniform3ivEXT"))
	gpProgramUniform3ui = uintptr(getProcAddr("glProgramUniform3ui"))
	gpProgramUniform3ui64ARB = uintptr(getProcAddr("glProgramUniform3ui64ARB"))
	gpProgramUniform3ui64NV = uintptr(getProcAddr("glProgramUniform3ui64NV"))
	gpProgramUniform3ui64vARB = uintptr(getProcAddr("glProgramUniform3ui64vARB"))
	gpProgramUniform3ui64vNV = uintptr(getProcAddr("glProgramUniform3ui64vNV"))
	gpProgramUniform3uiEXT = uintptr(getProcAddr("glProgramUniform3uiEXT"))
	gpProgramUniform3uiv = uintptr(getProcAddr("glProgramUniform3uiv"))
	gpProgramUniform3uivEXT = uintptr(getProcAddr("glProgramUniform3uivEXT"))
	gpProgramUniform4d = uintptr(getProcAddr("glProgramUniform4d"))
	gpProgramUniform4dEXT = uintptr(getProcAddr("glProgramUniform4dEXT"))
	gpProgramUniform4dv = uintptr(getProcAddr("glProgramUniform4dv"))
	gpProgramUniform4dvEXT = uintptr(getProcAddr("glProgramUniform4dvEXT"))
	gpProgramUniform4f = uintptr(getProcAddr("glProgramUniform4f"))
	gpProgramUniform4fEXT = uintptr(getProcAddr("glProgramUniform4fEXT"))
	gpProgramUniform4fv = uintptr(getProcAddr("glProgramUniform4fv"))
	gpProgramUniform4fvEXT = uintptr(getProcAddr("glProgramUniform4fvEXT"))
	gpProgramUniform4i = uintptr(getProcAddr("glProgramUniform4i"))
	gpProgramUniform4i64ARB = uintptr(getProcAddr("glProgramUniform4i64ARB"))
	gpProgramUniform4i64NV = uintptr(getProcAddr("glProgramUniform4i64NV"))
	gpProgramUniform4i64vARB = uintptr(getProcAddr("glProgramUniform4i64vARB"))
	gpProgramUniform4i64vNV = uintptr(getProcAddr("glProgramUniform4i64vNV"))
	gpProgramUniform4iEXT = uintptr(getProcAddr("glProgramUniform4iEXT"))
	gpProgramUniform4iv = uintptr(getProcAddr("glProgramUniform4iv"))
	gpProgramUniform4ivEXT = uintptr(getProcAddr("glProgramUniform4ivEXT"))
	gpProgramUniform4ui = uintptr(getProcAddr("glProgramUniform4ui"))
	gpProgramUniform4ui64ARB = uintptr(getProcAddr("glProgramUniform4ui64ARB"))
	gpProgramUniform4ui64NV = uintptr(getProcAddr("glProgramUniform4ui64NV"))
	gpProgramUniform4ui64vARB = uintptr(getProcAddr("glProgramUniform4ui64vARB"))
	gpProgramUniform4ui64vNV = uintptr(getProcAddr("glProgramUniform4ui64vNV"))
	gpProgramUniform4uiEXT = uintptr(getProcAddr("glProgramUniform4uiEXT"))
	gpProgramUniform4uiv = uintptr(getProcAddr("glProgramUniform4uiv"))
	gpProgramUniform4uivEXT = uintptr(getProcAddr("glProgramUniform4uivEXT"))
	gpProgramUniformHandleui64ARB = uintptr(getProcAddr("glProgramUniformHandleui64ARB"))
	gpProgramUniformHandleui64NV = uintptr(getProcAddr("glProgramUniformHandleui64NV"))
	gpProgramUniformHandleui64vARB = uintptr(getProcAddr("glProgramUniformHandleui64vARB"))
	gpProgramUniformHandleui64vNV = uintptr(getProcAddr("glProgramUniformHandleui64vNV"))
	gpProgramUniformMatrix2dv = uintptr(getProcAddr("glProgramUniformMatrix2dv"))
	gpProgramUniformMatrix2dvEXT = uintptr(getProcAddr("glProgramUniformMatrix2dvEXT"))
	gpProgramUniformMatrix2fv = uintptr(getProcAddr("glProgramUniformMatrix2fv"))
	gpProgramUniformMatrix2fvEXT = uintptr(getProcAddr("glProgramUniformMatrix2fvEXT"))
	gpProgramUniformMatrix2x3dv = uintptr(getProcAddr("glProgramUniformMatrix2x3dv"))
	gpProgramUniformMatrix2x3dvEXT = uintptr(getProcAddr("glProgramUniformMatrix2x3dvEXT"))
	gpProgramUniformMatrix2x3fv = uintptr(getProcAddr("glProgramUniformMatrix2x3fv"))
	gpProgramUniformMatrix2x3fvEXT = uintptr(getProcAddr("glProgramUniformMatrix2x3fvEXT"))
	gpProgramUniformMatrix2x4dv = uintptr(getProcAddr("glProgramUniformMatrix2x4dv"))
	gpProgramUniformMatrix2x4dvEXT = uintptr(getProcAddr("glProgramUniformMatrix2x4dvEXT"))
	gpProgramUniformMatrix2x4fv = uintptr(getProcAddr("glProgramUniformMatrix2x4fv"))
	gpProgramUniformMatrix2x4fvEXT = uintptr(getProcAddr("glProgramUniformMatrix2x4fvEXT"))
	gpProgramUniformMatrix3dv = uintptr(getProcAddr("glProgramUniformMatrix3dv"))
	gpProgramUniformMatrix3dvEXT = uintptr(getProcAddr("glProgramUniformMatrix3dvEXT"))
	gpProgramUniformMatrix3fv = uintptr(getProcAddr("glProgramUniformMatrix3fv"))
	gpProgramUniformMatrix3fvEXT = uintptr(getProcAddr("glProgramUniformMatrix3fvEXT"))
	gpProgramUniformMatrix3x2dv = uintptr(getProcAddr("glProgramUniformMatrix3x2dv"))
	gpProgramUniformMatrix3x2dvEXT = uintptr(getProcAddr("glProgramUniformMatrix3x2dvEXT"))
	gpProgramUniformMatrix3x2fv = uintptr(getProcAddr("glProgramUniformMatrix3x2fv"))
	gpProgramUniformMatrix3x2fvEXT = uintptr(getProcAddr("glProgramUniformMatrix3x2fvEXT"))
	gpProgramUniformMatrix3x4dv = uintptr(getProcAddr("glProgramUniformMatrix3x4dv"))
	gpProgramUniformMatrix3x4dvEXT = uintptr(getProcAddr("glProgramUniformMatrix3x4dvEXT"))
	gpProgramUniformMatrix3x4fv = uintptr(getProcAddr("glProgramUniformMatrix3x4fv"))
	gpProgramUniformMatrix3x4fvEXT = uintptr(getProcAddr("glProgramUniformMatrix3x4fvEXT"))
	gpProgramUniformMatrix4dv = uintptr(getProcAddr("glProgramUniformMatrix4dv"))
	gpProgramUniformMatrix4dvEXT = uintptr(getProcAddr("glProgramUniformMatrix4dvEXT"))
	gpProgramUniformMatrix4fv = uintptr(getProcAddr("glProgramUniformMatrix4fv"))
	gpProgramUniformMatrix4fvEXT = uintptr(getProcAddr("glProgramUniformMatrix4fvEXT"))
	gpProgramUniformMatrix4x2dv = uintptr(getProcAddr("glProgramUniformMatrix4x2dv"))
	gpProgramUniformMatrix4x2dvEXT = uintptr(getProcAddr("glProgramUniformMatrix4x2dvEXT"))
	gpProgramUniformMatrix4x2fv = uintptr(getProcAddr("glProgramUniformMatrix4x2fv"))
	gpProgramUniformMatrix4x2fvEXT = uintptr(getProcAddr("glProgramUniformMatrix4x2fvEXT"))
	gpProgramUniformMatrix4x3dv = uintptr(getProcAddr("glProgramUniformMatrix4x3dv"))
	gpProgramUniformMatrix4x3dvEXT = uintptr(getProcAddr("glProgramUniformMatrix4x3dvEXT"))
	gpProgramUniformMatrix4x3fv = uintptr(getProcAddr("glProgramUniformMatrix4x3fv"))
	gpProgramUniformMatrix4x3fvEXT = uintptr(getProcAddr("glProgramUniformMatrix4x3fvEXT"))
	gpProgramUniformui64NV = uintptr(getProcAddr("glProgramUniformui64NV"))
	gpProgramUniformui64vNV = uintptr(getProcAddr("glProgramUniformui64vNV"))
	gpProgramVertexLimitNV = uintptr(getProcAddr("glProgramVertexLimitNV"))
	gpProvokingVertex = uintptr(getProcAddr("glProvokingVertex"))
	gpProvokingVertexEXT = uintptr(getProcAddr("glProvokingVertexEXT"))
	gpPushAttrib = uintptr(getProcAddr("glPushAttrib"))
	if gpPushAttrib == 0 {
		return errors.New("glPushAttrib")
	}
	gpPushClientAttrib = uintptr(getProcAddr("glPushClientAttrib"))
	if gpPushClientAttrib == 0 {
		return errors.New("glPushClientAttrib")
	}
	gpPushClientAttribDefaultEXT = uintptr(getProcAddr("glPushClientAttribDefaultEXT"))
	gpPushDebugGroup = uintptr(getProcAddr("glPushDebugGroup"))
	gpPushDebugGroupKHR = uintptr(getProcAddr("glPushDebugGroupKHR"))
	gpPushGroupMarkerEXT = uintptr(getProcAddr("glPushGroupMarkerEXT"))
	gpPushMatrix = uintptr(getProcAddr("glPushMatrix"))
	if gpPushMatrix == 0 {
		return errors.New("glPushMatrix")
	}
	gpPushName = uintptr(getProcAddr("glPushName"))
	if gpPushName == 0 {
		return errors.New("glPushName")
	}
	gpQueryCounter = uintptr(getProcAddr("glQueryCounter"))
	gpQueryMatrixxOES = uintptr(getProcAddr("glQueryMatrixxOES"))
	gpQueryObjectParameteruiAMD = uintptr(getProcAddr("glQueryObjectParameteruiAMD"))
	gpQueryResourceNV = uintptr(getProcAddr("glQueryResourceNV"))
	gpQueryResourceTagNV = uintptr(getProcAddr("glQueryResourceTagNV"))
	gpRasterPos2d = uintptr(getProcAddr("glRasterPos2d"))
	if gpRasterPos2d == 0 {
		return errors.New("glRasterPos2d")
	}
	gpRasterPos2dv = uintptr(getProcAddr("glRasterPos2dv"))
	if gpRasterPos2dv == 0 {
		return errors.New("glRasterPos2dv")
	}
	gpRasterPos2f = uintptr(getProcAddr("glRasterPos2f"))
	if gpRasterPos2f == 0 {
		return errors.New("glRasterPos2f")
	}
	gpRasterPos2fv = uintptr(getProcAddr("glRasterPos2fv"))
	if gpRasterPos2fv == 0 {
		return errors.New("glRasterPos2fv")
	}
	gpRasterPos2i = uintptr(getProcAddr("glRasterPos2i"))
	if gpRasterPos2i == 0 {
		return errors.New("glRasterPos2i")
	}
	gpRasterPos2iv = uintptr(getProcAddr("glRasterPos2iv"))
	if gpRasterPos2iv == 0 {
		return errors.New("glRasterPos2iv")
	}
	gpRasterPos2s = uintptr(getProcAddr("glRasterPos2s"))
	if gpRasterPos2s == 0 {
		return errors.New("glRasterPos2s")
	}
	gpRasterPos2sv = uintptr(getProcAddr("glRasterPos2sv"))
	if gpRasterPos2sv == 0 {
		return errors.New("glRasterPos2sv")
	}
	gpRasterPos2xOES = uintptr(getProcAddr("glRasterPos2xOES"))
	gpRasterPos2xvOES = uintptr(getProcAddr("glRasterPos2xvOES"))
	gpRasterPos3d = uintptr(getProcAddr("glRasterPos3d"))
	if gpRasterPos3d == 0 {
		return errors.New("glRasterPos3d")
	}
	gpRasterPos3dv = uintptr(getProcAddr("glRasterPos3dv"))
	if gpRasterPos3dv == 0 {
		return errors.New("glRasterPos3dv")
	}
	gpRasterPos3f = uintptr(getProcAddr("glRasterPos3f"))
	if gpRasterPos3f == 0 {
		return errors.New("glRasterPos3f")
	}
	gpRasterPos3fv = uintptr(getProcAddr("glRasterPos3fv"))
	if gpRasterPos3fv == 0 {
		return errors.New("glRasterPos3fv")
	}
	gpRasterPos3i = uintptr(getProcAddr("glRasterPos3i"))
	if gpRasterPos3i == 0 {
		return errors.New("glRasterPos3i")
	}
	gpRasterPos3iv = uintptr(getProcAddr("glRasterPos3iv"))
	if gpRasterPos3iv == 0 {
		return errors.New("glRasterPos3iv")
	}
	gpRasterPos3s = uintptr(getProcAddr("glRasterPos3s"))
	if gpRasterPos3s == 0 {
		return errors.New("glRasterPos3s")
	}
	gpRasterPos3sv = uintptr(getProcAddr("glRasterPos3sv"))
	if gpRasterPos3sv == 0 {
		return errors.New("glRasterPos3sv")
	}
	gpRasterPos3xOES = uintptr(getProcAddr("glRasterPos3xOES"))
	gpRasterPos3xvOES = uintptr(getProcAddr("glRasterPos3xvOES"))
	gpRasterPos4d = uintptr(getProcAddr("glRasterPos4d"))
	if gpRasterPos4d == 0 {
		return errors.New("glRasterPos4d")
	}
	gpRasterPos4dv = uintptr(getProcAddr("glRasterPos4dv"))
	if gpRasterPos4dv == 0 {
		return errors.New("glRasterPos4dv")
	}
	gpRasterPos4f = uintptr(getProcAddr("glRasterPos4f"))
	if gpRasterPos4f == 0 {
		return errors.New("glRasterPos4f")
	}
	gpRasterPos4fv = uintptr(getProcAddr("glRasterPos4fv"))
	if gpRasterPos4fv == 0 {
		return errors.New("glRasterPos4fv")
	}
	gpRasterPos4i = uintptr(getProcAddr("glRasterPos4i"))
	if gpRasterPos4i == 0 {
		return errors.New("glRasterPos4i")
	}
	gpRasterPos4iv = uintptr(getProcAddr("glRasterPos4iv"))
	if gpRasterPos4iv == 0 {
		return errors.New("glRasterPos4iv")
	}
	gpRasterPos4s = uintptr(getProcAddr("glRasterPos4s"))
	if gpRasterPos4s == 0 {
		return errors.New("glRasterPos4s")
	}
	gpRasterPos4sv = uintptr(getProcAddr("glRasterPos4sv"))
	if gpRasterPos4sv == 0 {
		return errors.New("glRasterPos4sv")
	}
	gpRasterPos4xOES = uintptr(getProcAddr("glRasterPos4xOES"))
	gpRasterPos4xvOES = uintptr(getProcAddr("glRasterPos4xvOES"))
	gpRasterSamplesEXT = uintptr(getProcAddr("glRasterSamplesEXT"))
	gpReadBuffer = uintptr(getProcAddr("glReadBuffer"))
	if gpReadBuffer == 0 {
		return errors.New("glReadBuffer")
	}
	gpReadInstrumentsSGIX = uintptr(getProcAddr("glReadInstrumentsSGIX"))
	gpReadPixels = uintptr(getProcAddr("glReadPixels"))
	if gpReadPixels == 0 {
		return errors.New("glReadPixels")
	}
	gpReadnPixels = uintptr(getProcAddr("glReadnPixels"))
	gpReadnPixelsARB = uintptr(getProcAddr("glReadnPixelsARB"))
	gpReadnPixelsKHR = uintptr(getProcAddr("glReadnPixelsKHR"))
	gpRectd = uintptr(getProcAddr("glRectd"))
	if gpRectd == 0 {
		return errors.New("glRectd")
	}
	gpRectdv = uintptr(getProcAddr("glRectdv"))
	if gpRectdv == 0 {
		return errors.New("glRectdv")
	}
	gpRectf = uintptr(getProcAddr("glRectf"))
	if gpRectf == 0 {
		return errors.New("glRectf")
	}
	gpRectfv = uintptr(getProcAddr("glRectfv"))
	if gpRectfv == 0 {
		return errors.New("glRectfv")
	}
	gpRecti = uintptr(getProcAddr("glRecti"))
	if gpRecti == 0 {
		return errors.New("glRecti")
	}
	gpRectiv = uintptr(getProcAddr("glRectiv"))
	if gpRectiv == 0 {
		return errors.New("glRectiv")
	}
	gpRects = uintptr(getProcAddr("glRects"))
	if gpRects == 0 {
		return errors.New("glRects")
	}
	gpRectsv = uintptr(getProcAddr("glRectsv"))
	if gpRectsv == 0 {
		return errors.New("glRectsv")
	}
	gpRectxOES = uintptr(getProcAddr("glRectxOES"))
	gpRectxvOES = uintptr(getProcAddr("glRectxvOES"))
	gpReferencePlaneSGIX = uintptr(getProcAddr("glReferencePlaneSGIX"))
	gpReleaseKeyedMutexWin32EXT = uintptr(getProcAddr("glReleaseKeyedMutexWin32EXT"))
	gpReleaseShaderCompiler = uintptr(getProcAddr("glReleaseShaderCompiler"))
	gpRenderGpuMaskNV = uintptr(getProcAddr("glRenderGpuMaskNV"))
	gpRenderMode = uintptr(getProcAddr("glRenderMode"))
	if gpRenderMode == 0 {
		return errors.New("glRenderMode")
	}
	gpRenderbufferStorage = uintptr(getProcAddr("glRenderbufferStorage"))
	gpRenderbufferStorageEXT = uintptr(getProcAddr("glRenderbufferStorageEXT"))
	gpRenderbufferStorageMultisample = uintptr(getProcAddr("glRenderbufferStorageMultisample"))
	gpRenderbufferStorageMultisampleCoverageNV = uintptr(getProcAddr("glRenderbufferStorageMultisampleCoverageNV"))
	gpRenderbufferStorageMultisampleEXT = uintptr(getProcAddr("glRenderbufferStorageMultisampleEXT"))
	gpReplacementCodePointerSUN = uintptr(getProcAddr("glReplacementCodePointerSUN"))
	gpReplacementCodeubSUN = uintptr(getProcAddr("glReplacementCodeubSUN"))
	gpReplacementCodeubvSUN = uintptr(getProcAddr("glReplacementCodeubvSUN"))
	gpReplacementCodeuiColor3fVertex3fSUN = uintptr(getProcAddr("glReplacementCodeuiColor3fVertex3fSUN"))
	gpReplacementCodeuiColor3fVertex3fvSUN = uintptr(getProcAddr("glReplacementCodeuiColor3fVertex3fvSUN"))
	gpReplacementCodeuiColor4fNormal3fVertex3fSUN = uintptr(getProcAddr("glReplacementCodeuiColor4fNormal3fVertex3fSUN"))
	gpReplacementCodeuiColor4fNormal3fVertex3fvSUN = uintptr(getProcAddr("glReplacementCodeuiColor4fNormal3fVertex3fvSUN"))
	gpReplacementCodeuiColor4ubVertex3fSUN = uintptr(getProcAddr("glReplacementCodeuiColor4ubVertex3fSUN"))
	gpReplacementCodeuiColor4ubVertex3fvSUN = uintptr(getProcAddr("glReplacementCodeuiColor4ubVertex3fvSUN"))
	gpReplacementCodeuiNormal3fVertex3fSUN = uintptr(getProcAddr("glReplacementCodeuiNormal3fVertex3fSUN"))
	gpReplacementCodeuiNormal3fVertex3fvSUN = uintptr(getProcAddr("glReplacementCodeuiNormal3fVertex3fvSUN"))
	gpReplacementCodeuiSUN = uintptr(getProcAddr("glReplacementCodeuiSUN"))
	gpReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN = uintptr(getProcAddr("glReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN"))
	gpReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN = uintptr(getProcAddr("glReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN"))
	gpReplacementCodeuiTexCoord2fNormal3fVertex3fSUN = uintptr(getProcAddr("glReplacementCodeuiTexCoord2fNormal3fVertex3fSUN"))
	gpReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN = uintptr(getProcAddr("glReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN"))
	gpReplacementCodeuiTexCoord2fVertex3fSUN = uintptr(getProcAddr("glReplacementCodeuiTexCoord2fVertex3fSUN"))
	gpReplacementCodeuiTexCoord2fVertex3fvSUN = uintptr(getProcAddr("glReplacementCodeuiTexCoord2fVertex3fvSUN"))
	gpReplacementCodeuiVertex3fSUN = uintptr(getProcAddr("glReplacementCodeuiVertex3fSUN"))
	gpReplacementCodeuiVertex3fvSUN = uintptr(getProcAddr("glReplacementCodeuiVertex3fvSUN"))
	gpReplacementCodeuivSUN = uintptr(getProcAddr("glReplacementCodeuivSUN"))
	gpReplacementCodeusSUN = uintptr(getProcAddr("glReplacementCodeusSUN"))
	gpReplacementCodeusvSUN = uintptr(getProcAddr("glReplacementCodeusvSUN"))
	gpRequestResidentProgramsNV = uintptr(getProcAddr("glRequestResidentProgramsNV"))
	gpResetHistogramEXT = uintptr(getProcAddr("glResetHistogramEXT"))
	gpResetMinmaxEXT = uintptr(getProcAddr("glResetMinmaxEXT"))
	gpResizeBuffersMESA = uintptr(getProcAddr("glResizeBuffersMESA"))
	gpResolveDepthValuesNV = uintptr(getProcAddr("glResolveDepthValuesNV"))
	gpResumeTransformFeedback = uintptr(getProcAddr("glResumeTransformFeedback"))
	gpResumeTransformFeedbackNV = uintptr(getProcAddr("glResumeTransformFeedbackNV"))
	gpRotated = uintptr(getProcAddr("glRotated"))
	if gpRotated == 0 {
		return errors.New("glRotated")
	}
	gpRotatef = uintptr(getProcAddr("glRotatef"))
	if gpRotatef == 0 {
		return errors.New("glRotatef")
	}
	gpRotatexOES = uintptr(getProcAddr("glRotatexOES"))
	gpSampleCoverage = uintptr(getProcAddr("glSampleCoverage"))
	if gpSampleCoverage == 0 {
		return errors.New("glSampleCoverage")
	}
	gpSampleCoverageARB = uintptr(getProcAddr("glSampleCoverageARB"))
	gpSampleCoveragexOES = uintptr(getProcAddr("glSampleCoveragexOES"))
	gpSampleMapATI = uintptr(getProcAddr("glSampleMapATI"))
	gpSampleMaskEXT = uintptr(getProcAddr("glSampleMaskEXT"))
	gpSampleMaskIndexedNV = uintptr(getProcAddr("glSampleMaskIndexedNV"))
	gpSampleMaskSGIS = uintptr(getProcAddr("glSampleMaskSGIS"))
	gpSampleMaski = uintptr(getProcAddr("glSampleMaski"))
	gpSamplePatternEXT = uintptr(getProcAddr("glSamplePatternEXT"))
	gpSamplePatternSGIS = uintptr(getProcAddr("glSamplePatternSGIS"))
	gpSamplerParameterIiv = uintptr(getProcAddr("glSamplerParameterIiv"))
	gpSamplerParameterIuiv = uintptr(getProcAddr("glSamplerParameterIuiv"))
	gpSamplerParameterf = uintptr(getProcAddr("glSamplerParameterf"))
	gpSamplerParameterfv = uintptr(getProcAddr("glSamplerParameterfv"))
	gpSamplerParameteri = uintptr(getProcAddr("glSamplerParameteri"))
	gpSamplerParameteriv = uintptr(getProcAddr("glSamplerParameteriv"))
	gpScaled = uintptr(getProcAddr("glScaled"))
	if gpScaled == 0 {
		return errors.New("glScaled")
	}
	gpScalef = uintptr(getProcAddr("glScalef"))
	if gpScalef == 0 {
		return errors.New("glScalef")
	}
	gpScalexOES = uintptr(getProcAddr("glScalexOES"))
	gpScissor = uintptr(getProcAddr("glScissor"))
	if gpScissor == 0 {
		return errors.New("glScissor")
	}
	gpScissorArrayv = uintptr(getProcAddr("glScissorArrayv"))
	gpScissorIndexed = uintptr(getProcAddr("glScissorIndexed"))
	gpScissorIndexedv = uintptr(getProcAddr("glScissorIndexedv"))
	gpSecondaryColor3b = uintptr(getProcAddr("glSecondaryColor3b"))
	if gpSecondaryColor3b == 0 {
		return errors.New("glSecondaryColor3b")
	}
	gpSecondaryColor3bEXT = uintptr(getProcAddr("glSecondaryColor3bEXT"))
	gpSecondaryColor3bv = uintptr(getProcAddr("glSecondaryColor3bv"))
	if gpSecondaryColor3bv == 0 {
		return errors.New("glSecondaryColor3bv")
	}
	gpSecondaryColor3bvEXT = uintptr(getProcAddr("glSecondaryColor3bvEXT"))
	gpSecondaryColor3d = uintptr(getProcAddr("glSecondaryColor3d"))
	if gpSecondaryColor3d == 0 {
		return errors.New("glSecondaryColor3d")
	}
	gpSecondaryColor3dEXT = uintptr(getProcAddr("glSecondaryColor3dEXT"))
	gpSecondaryColor3dv = uintptr(getProcAddr("glSecondaryColor3dv"))
	if gpSecondaryColor3dv == 0 {
		return errors.New("glSecondaryColor3dv")
	}
	gpSecondaryColor3dvEXT = uintptr(getProcAddr("glSecondaryColor3dvEXT"))
	gpSecondaryColor3f = uintptr(getProcAddr("glSecondaryColor3f"))
	if gpSecondaryColor3f == 0 {
		return errors.New("glSecondaryColor3f")
	}
	gpSecondaryColor3fEXT = uintptr(getProcAddr("glSecondaryColor3fEXT"))
	gpSecondaryColor3fv = uintptr(getProcAddr("glSecondaryColor3fv"))
	if gpSecondaryColor3fv == 0 {
		return errors.New("glSecondaryColor3fv")
	}
	gpSecondaryColor3fvEXT = uintptr(getProcAddr("glSecondaryColor3fvEXT"))
	gpSecondaryColor3hNV = uintptr(getProcAddr("glSecondaryColor3hNV"))
	gpSecondaryColor3hvNV = uintptr(getProcAddr("glSecondaryColor3hvNV"))
	gpSecondaryColor3i = uintptr(getProcAddr("glSecondaryColor3i"))
	if gpSecondaryColor3i == 0 {
		return errors.New("glSecondaryColor3i")
	}
	gpSecondaryColor3iEXT = uintptr(getProcAddr("glSecondaryColor3iEXT"))
	gpSecondaryColor3iv = uintptr(getProcAddr("glSecondaryColor3iv"))
	if gpSecondaryColor3iv == 0 {
		return errors.New("glSecondaryColor3iv")
	}
	gpSecondaryColor3ivEXT = uintptr(getProcAddr("glSecondaryColor3ivEXT"))
	gpSecondaryColor3s = uintptr(getProcAddr("glSecondaryColor3s"))
	if gpSecondaryColor3s == 0 {
		return errors.New("glSecondaryColor3s")
	}
	gpSecondaryColor3sEXT = uintptr(getProcAddr("glSecondaryColor3sEXT"))
	gpSecondaryColor3sv = uintptr(getProcAddr("glSecondaryColor3sv"))
	if gpSecondaryColor3sv == 0 {
		return errors.New("glSecondaryColor3sv")
	}
	gpSecondaryColor3svEXT = uintptr(getProcAddr("glSecondaryColor3svEXT"))
	gpSecondaryColor3ub = uintptr(getProcAddr("glSecondaryColor3ub"))
	if gpSecondaryColor3ub == 0 {
		return errors.New("glSecondaryColor3ub")
	}
	gpSecondaryColor3ubEXT = uintptr(getProcAddr("glSecondaryColor3ubEXT"))
	gpSecondaryColor3ubv = uintptr(getProcAddr("glSecondaryColor3ubv"))
	if gpSecondaryColor3ubv == 0 {
		return errors.New("glSecondaryColor3ubv")
	}
	gpSecondaryColor3ubvEXT = uintptr(getProcAddr("glSecondaryColor3ubvEXT"))
	gpSecondaryColor3ui = uintptr(getProcAddr("glSecondaryColor3ui"))
	if gpSecondaryColor3ui == 0 {
		return errors.New("glSecondaryColor3ui")
	}
	gpSecondaryColor3uiEXT = uintptr(getProcAddr("glSecondaryColor3uiEXT"))
	gpSecondaryColor3uiv = uintptr(getProcAddr("glSecondaryColor3uiv"))
	if gpSecondaryColor3uiv == 0 {
		return errors.New("glSecondaryColor3uiv")
	}
	gpSecondaryColor3uivEXT = uintptr(getProcAddr("glSecondaryColor3uivEXT"))
	gpSecondaryColor3us = uintptr(getProcAddr("glSecondaryColor3us"))
	if gpSecondaryColor3us == 0 {
		return errors.New("glSecondaryColor3us")
	}
	gpSecondaryColor3usEXT = uintptr(getProcAddr("glSecondaryColor3usEXT"))
	gpSecondaryColor3usv = uintptr(getProcAddr("glSecondaryColor3usv"))
	if gpSecondaryColor3usv == 0 {
		return errors.New("glSecondaryColor3usv")
	}
	gpSecondaryColor3usvEXT = uintptr(getProcAddr("glSecondaryColor3usvEXT"))
	gpSecondaryColorFormatNV = uintptr(getProcAddr("glSecondaryColorFormatNV"))
	gpSecondaryColorPointer = uintptr(getProcAddr("glSecondaryColorPointer"))
	if gpSecondaryColorPointer == 0 {
		return errors.New("glSecondaryColorPointer")
	}
	gpSecondaryColorPointerEXT = uintptr(getProcAddr("glSecondaryColorPointerEXT"))
	gpSecondaryColorPointerListIBM = uintptr(getProcAddr("glSecondaryColorPointerListIBM"))
	gpSelectBuffer = uintptr(getProcAddr("glSelectBuffer"))
	if gpSelectBuffer == 0 {
		return errors.New("glSelectBuffer")
	}
	gpSelectPerfMonitorCountersAMD = uintptr(getProcAddr("glSelectPerfMonitorCountersAMD"))
	gpSemaphoreParameterui64vEXT = uintptr(getProcAddr("glSemaphoreParameterui64vEXT"))
	gpSeparableFilter2DEXT = uintptr(getProcAddr("glSeparableFilter2DEXT"))
	gpSetFenceAPPLE = uintptr(getProcAddr("glSetFenceAPPLE"))
	gpSetFenceNV = uintptr(getProcAddr("glSetFenceNV"))
	gpSetFragmentShaderConstantATI = uintptr(getProcAddr("glSetFragmentShaderConstantATI"))
	gpSetInvariantEXT = uintptr(getProcAddr("glSetInvariantEXT"))
	gpSetLocalConstantEXT = uintptr(getProcAddr("glSetLocalConstantEXT"))
	gpSetMultisamplefvAMD = uintptr(getProcAddr("glSetMultisamplefvAMD"))
	gpShadeModel = uintptr(getProcAddr("glShadeModel"))
	if gpShadeModel == 0 {
		return errors.New("glShadeModel")
	}
	gpShaderBinary = uintptr(getProcAddr("glShaderBinary"))
	gpShaderOp1EXT = uintptr(getProcAddr("glShaderOp1EXT"))
	gpShaderOp2EXT = uintptr(getProcAddr("glShaderOp2EXT"))
	gpShaderOp3EXT = uintptr(getProcAddr("glShaderOp3EXT"))
	gpShaderSource = uintptr(getProcAddr("glShaderSource"))
	if gpShaderSource == 0 {
		return errors.New("glShaderSource")
	}
	gpShaderSourceARB = uintptr(getProcAddr("glShaderSourceARB"))
	gpShaderStorageBlockBinding = uintptr(getProcAddr("glShaderStorageBlockBinding"))
	gpSharpenTexFuncSGIS = uintptr(getProcAddr("glSharpenTexFuncSGIS"))
	gpSignalSemaphoreEXT = uintptr(getProcAddr("glSignalSemaphoreEXT"))
	gpSignalVkFenceNV = uintptr(getProcAddr("glSignalVkFenceNV"))
	gpSignalVkSemaphoreNV = uintptr(getProcAddr("glSignalVkSemaphoreNV"))
	gpSpecializeShaderARB = uintptr(getProcAddr("glSpecializeShaderARB"))
	gpSpriteParameterfSGIX = uintptr(getProcAddr("glSpriteParameterfSGIX"))
	gpSpriteParameterfvSGIX = uintptr(getProcAddr("glSpriteParameterfvSGIX"))
	gpSpriteParameteriSGIX = uintptr(getProcAddr("glSpriteParameteriSGIX"))
	gpSpriteParameterivSGIX = uintptr(getProcAddr("glSpriteParameterivSGIX"))
	gpStartInstrumentsSGIX = uintptr(getProcAddr("glStartInstrumentsSGIX"))
	gpStateCaptureNV = uintptr(getProcAddr("glStateCaptureNV"))
	gpStencilClearTagEXT = uintptr(getProcAddr("glStencilClearTagEXT"))
	gpStencilFillPathInstancedNV = uintptr(getProcAddr("glStencilFillPathInstancedNV"))
	gpStencilFillPathNV = uintptr(getProcAddr("glStencilFillPathNV"))
	gpStencilFunc = uintptr(getProcAddr("glStencilFunc"))
	if gpStencilFunc == 0 {
		return errors.New("glStencilFunc")
	}
	gpStencilFuncSeparate = uintptr(getProcAddr("glStencilFuncSeparate"))
	if gpStencilFuncSeparate == 0 {
		return errors.New("glStencilFuncSeparate")
	}
	gpStencilFuncSeparateATI = uintptr(getProcAddr("glStencilFuncSeparateATI"))
	gpStencilMask = uintptr(getProcAddr("glStencilMask"))
	if gpStencilMask == 0 {
		return errors.New("glStencilMask")
	}
	gpStencilMaskSeparate = uintptr(getProcAddr("glStencilMaskSeparate"))
	if gpStencilMaskSeparate == 0 {
		return errors.New("glStencilMaskSeparate")
	}
	gpStencilOp = uintptr(getProcAddr("glStencilOp"))
	if gpStencilOp == 0 {
		return errors.New("glStencilOp")
	}
	gpStencilOpSeparate = uintptr(getProcAddr("glStencilOpSeparate"))
	if gpStencilOpSeparate == 0 {
		return errors.New("glStencilOpSeparate")
	}
	gpStencilOpSeparateATI = uintptr(getProcAddr("glStencilOpSeparateATI"))
	gpStencilOpValueAMD = uintptr(getProcAddr("glStencilOpValueAMD"))
	gpStencilStrokePathInstancedNV = uintptr(getProcAddr("glStencilStrokePathInstancedNV"))
	gpStencilStrokePathNV = uintptr(getProcAddr("glStencilStrokePathNV"))
	gpStencilThenCoverFillPathInstancedNV = uintptr(getProcAddr("glStencilThenCoverFillPathInstancedNV"))
	gpStencilThenCoverFillPathNV = uintptr(getProcAddr("glStencilThenCoverFillPathNV"))
	gpStencilThenCoverStrokePathInstancedNV = uintptr(getProcAddr("glStencilThenCoverStrokePathInstancedNV"))
	gpStencilThenCoverStrokePathNV = uintptr(getProcAddr("glStencilThenCoverStrokePathNV"))
	gpStopInstrumentsSGIX = uintptr(getProcAddr("glStopInstrumentsSGIX"))
	gpStringMarkerGREMEDY = uintptr(getProcAddr("glStringMarkerGREMEDY"))
	gpSubpixelPrecisionBiasNV = uintptr(getProcAddr("glSubpixelPrecisionBiasNV"))
	gpSwizzleEXT = uintptr(getProcAddr("glSwizzleEXT"))
	gpSyncTextureINTEL = uintptr(getProcAddr("glSyncTextureINTEL"))
	gpTagSampleBufferSGIX = uintptr(getProcAddr("glTagSampleBufferSGIX"))
	gpTangent3bEXT = uintptr(getProcAddr("glTangent3bEXT"))
	gpTangent3bvEXT = uintptr(getProcAddr("glTangent3bvEXT"))
	gpTangent3dEXT = uintptr(getProcAddr("glTangent3dEXT"))
	gpTangent3dvEXT = uintptr(getProcAddr("glTangent3dvEXT"))
	gpTangent3fEXT = uintptr(getProcAddr("glTangent3fEXT"))
	gpTangent3fvEXT = uintptr(getProcAddr("glTangent3fvEXT"))
	gpTangent3iEXT = uintptr(getProcAddr("glTangent3iEXT"))
	gpTangent3ivEXT = uintptr(getProcAddr("glTangent3ivEXT"))
	gpTangent3sEXT = uintptr(getProcAddr("glTangent3sEXT"))
	gpTangent3svEXT = uintptr(getProcAddr("glTangent3svEXT"))
	gpTangentPointerEXT = uintptr(getProcAddr("glTangentPointerEXT"))
	gpTbufferMask3DFX = uintptr(getProcAddr("glTbufferMask3DFX"))
	gpTessellationFactorAMD = uintptr(getProcAddr("glTessellationFactorAMD"))
	gpTessellationModeAMD = uintptr(getProcAddr("glTessellationModeAMD"))
	gpTestFenceAPPLE = uintptr(getProcAddr("glTestFenceAPPLE"))
	gpTestFenceNV = uintptr(getProcAddr("glTestFenceNV"))
	gpTestObjectAPPLE = uintptr(getProcAddr("glTestObjectAPPLE"))
	gpTexBufferARB = uintptr(getProcAddr("glTexBufferARB"))
	gpTexBufferEXT = uintptr(getProcAddr("glTexBufferEXT"))
	gpTexBufferRange = uintptr(getProcAddr("glTexBufferRange"))
	gpTexBumpParameterfvATI = uintptr(getProcAddr("glTexBumpParameterfvATI"))
	gpTexBumpParameterivATI = uintptr(getProcAddr("glTexBumpParameterivATI"))
	gpTexCoord1bOES = uintptr(getProcAddr("glTexCoord1bOES"))
	gpTexCoord1bvOES = uintptr(getProcAddr("glTexCoord1bvOES"))
	gpTexCoord1d = uintptr(getProcAddr("glTexCoord1d"))
	if gpTexCoord1d == 0 {
		return errors.New("glTexCoord1d")
	}
	gpTexCoord1dv = uintptr(getProcAddr("glTexCoord1dv"))
	if gpTexCoord1dv == 0 {
		return errors.New("glTexCoord1dv")
	}
	gpTexCoord1f = uintptr(getProcAddr("glTexCoord1f"))
	if gpTexCoord1f == 0 {
		return errors.New("glTexCoord1f")
	}
	gpTexCoord1fv = uintptr(getProcAddr("glTexCoord1fv"))
	if gpTexCoord1fv == 0 {
		return errors.New("glTexCoord1fv")
	}
	gpTexCoord1hNV = uintptr(getProcAddr("glTexCoord1hNV"))
	gpTexCoord1hvNV = uintptr(getProcAddr("glTexCoord1hvNV"))
	gpTexCoord1i = uintptr(getProcAddr("glTexCoord1i"))
	if gpTexCoord1i == 0 {
		return errors.New("glTexCoord1i")
	}
	gpTexCoord1iv = uintptr(getProcAddr("glTexCoord1iv"))
	if gpTexCoord1iv == 0 {
		return errors.New("glTexCoord1iv")
	}
	gpTexCoord1s = uintptr(getProcAddr("glTexCoord1s"))
	if gpTexCoord1s == 0 {
		return errors.New("glTexCoord1s")
	}
	gpTexCoord1sv = uintptr(getProcAddr("glTexCoord1sv"))
	if gpTexCoord1sv == 0 {
		return errors.New("glTexCoord1sv")
	}
	gpTexCoord1xOES = uintptr(getProcAddr("glTexCoord1xOES"))
	gpTexCoord1xvOES = uintptr(getProcAddr("glTexCoord1xvOES"))
	gpTexCoord2bOES = uintptr(getProcAddr("glTexCoord2bOES"))
	gpTexCoord2bvOES = uintptr(getProcAddr("glTexCoord2bvOES"))
	gpTexCoord2d = uintptr(getProcAddr("glTexCoord2d"))
	if gpTexCoord2d == 0 {
		return errors.New("glTexCoord2d")
	}
	gpTexCoord2dv = uintptr(getProcAddr("glTexCoord2dv"))
	if gpTexCoord2dv == 0 {
		return errors.New("glTexCoord2dv")
	}
	gpTexCoord2f = uintptr(getProcAddr("glTexCoord2f"))
	if gpTexCoord2f == 0 {
		return errors.New("glTexCoord2f")
	}
	gpTexCoord2fColor3fVertex3fSUN = uintptr(getProcAddr("glTexCoord2fColor3fVertex3fSUN"))
	gpTexCoord2fColor3fVertex3fvSUN = uintptr(getProcAddr("glTexCoord2fColor3fVertex3fvSUN"))
	gpTexCoord2fColor4fNormal3fVertex3fSUN = uintptr(getProcAddr("glTexCoord2fColor4fNormal3fVertex3fSUN"))
	gpTexCoord2fColor4fNormal3fVertex3fvSUN = uintptr(getProcAddr("glTexCoord2fColor4fNormal3fVertex3fvSUN"))
	gpTexCoord2fColor4ubVertex3fSUN = uintptr(getProcAddr("glTexCoord2fColor4ubVertex3fSUN"))
	gpTexCoord2fColor4ubVertex3fvSUN = uintptr(getProcAddr("glTexCoord2fColor4ubVertex3fvSUN"))
	gpTexCoord2fNormal3fVertex3fSUN = uintptr(getProcAddr("glTexCoord2fNormal3fVertex3fSUN"))
	gpTexCoord2fNormal3fVertex3fvSUN = uintptr(getProcAddr("glTexCoord2fNormal3fVertex3fvSUN"))
	gpTexCoord2fVertex3fSUN = uintptr(getProcAddr("glTexCoord2fVertex3fSUN"))
	gpTexCoord2fVertex3fvSUN = uintptr(getProcAddr("glTexCoord2fVertex3fvSUN"))
	gpTexCoord2fv = uintptr(getProcAddr("glTexCoord2fv"))
	if gpTexCoord2fv == 0 {
		return errors.New("glTexCoord2fv")
	}
	gpTexCoord2hNV = uintptr(getProcAddr("glTexCoord2hNV"))
	gpTexCoord2hvNV = uintptr(getProcAddr("glTexCoord2hvNV"))
	gpTexCoord2i = uintptr(getProcAddr("glTexCoord2i"))
	if gpTexCoord2i == 0 {
		return errors.New("glTexCoord2i")
	}
	gpTexCoord2iv = uintptr(getProcAddr("glTexCoord2iv"))
	if gpTexCoord2iv == 0 {
		return errors.New("glTexCoord2iv")
	}
	gpTexCoord2s = uintptr(getProcAddr("glTexCoord2s"))
	if gpTexCoord2s == 0 {
		return errors.New("glTexCoord2s")
	}
	gpTexCoord2sv = uintptr(getProcAddr("glTexCoord2sv"))
	if gpTexCoord2sv == 0 {
		return errors.New("glTexCoord2sv")
	}
	gpTexCoord2xOES = uintptr(getProcAddr("glTexCoord2xOES"))
	gpTexCoord2xvOES = uintptr(getProcAddr("glTexCoord2xvOES"))
	gpTexCoord3bOES = uintptr(getProcAddr("glTexCoord3bOES"))
	gpTexCoord3bvOES = uintptr(getProcAddr("glTexCoord3bvOES"))
	gpTexCoord3d = uintptr(getProcAddr("glTexCoord3d"))
	if gpTexCoord3d == 0 {
		return errors.New("glTexCoord3d")
	}
	gpTexCoord3dv = uintptr(getProcAddr("glTexCoord3dv"))
	if gpTexCoord3dv == 0 {
		return errors.New("glTexCoord3dv")
	}
	gpTexCoord3f = uintptr(getProcAddr("glTexCoord3f"))
	if gpTexCoord3f == 0 {
		return errors.New("glTexCoord3f")
	}
	gpTexCoord3fv = uintptr(getProcAddr("glTexCoord3fv"))
	if gpTexCoord3fv == 0 {
		return errors.New("glTexCoord3fv")
	}
	gpTexCoord3hNV = uintptr(getProcAddr("glTexCoord3hNV"))
	gpTexCoord3hvNV = uintptr(getProcAddr("glTexCoord3hvNV"))
	gpTexCoord3i = uintptr(getProcAddr("glTexCoord3i"))
	if gpTexCoord3i == 0 {
		return errors.New("glTexCoord3i")
	}
	gpTexCoord3iv = uintptr(getProcAddr("glTexCoord3iv"))
	if gpTexCoord3iv == 0 {
		return errors.New("glTexCoord3iv")
	}
	gpTexCoord3s = uintptr(getProcAddr("glTexCoord3s"))
	if gpTexCoord3s == 0 {
		return errors.New("glTexCoord3s")
	}
	gpTexCoord3sv = uintptr(getProcAddr("glTexCoord3sv"))
	if gpTexCoord3sv == 0 {
		return errors.New("glTexCoord3sv")
	}
	gpTexCoord3xOES = uintptr(getProcAddr("glTexCoord3xOES"))
	gpTexCoord3xvOES = uintptr(getProcAddr("glTexCoord3xvOES"))
	gpTexCoord4bOES = uintptr(getProcAddr("glTexCoord4bOES"))
	gpTexCoord4bvOES = uintptr(getProcAddr("glTexCoord4bvOES"))
	gpTexCoord4d = uintptr(getProcAddr("glTexCoord4d"))
	if gpTexCoord4d == 0 {
		return errors.New("glTexCoord4d")
	}
	gpTexCoord4dv = uintptr(getProcAddr("glTexCoord4dv"))
	if gpTexCoord4dv == 0 {
		return errors.New("glTexCoord4dv")
	}
	gpTexCoord4f = uintptr(getProcAddr("glTexCoord4f"))
	if gpTexCoord4f == 0 {
		return errors.New("glTexCoord4f")
	}
	gpTexCoord4fColor4fNormal3fVertex4fSUN = uintptr(getProcAddr("glTexCoord4fColor4fNormal3fVertex4fSUN"))
	gpTexCoord4fColor4fNormal3fVertex4fvSUN = uintptr(getProcAddr("glTexCoord4fColor4fNormal3fVertex4fvSUN"))
	gpTexCoord4fVertex4fSUN = uintptr(getProcAddr("glTexCoord4fVertex4fSUN"))
	gpTexCoord4fVertex4fvSUN = uintptr(getProcAddr("glTexCoord4fVertex4fvSUN"))
	gpTexCoord4fv = uintptr(getProcAddr("glTexCoord4fv"))
	if gpTexCoord4fv == 0 {
		return errors.New("glTexCoord4fv")
	}
	gpTexCoord4hNV = uintptr(getProcAddr("glTexCoord4hNV"))
	gpTexCoord4hvNV = uintptr(getProcAddr("glTexCoord4hvNV"))
	gpTexCoord4i = uintptr(getProcAddr("glTexCoord4i"))
	if gpTexCoord4i == 0 {
		return errors.New("glTexCoord4i")
	}
	gpTexCoord4iv = uintptr(getProcAddr("glTexCoord4iv"))
	if gpTexCoord4iv == 0 {
		return errors.New("glTexCoord4iv")
	}
	gpTexCoord4s = uintptr(getProcAddr("glTexCoord4s"))
	if gpTexCoord4s == 0 {
		return errors.New("glTexCoord4s")
	}
	gpTexCoord4sv = uintptr(getProcAddr("glTexCoord4sv"))
	if gpTexCoord4sv == 0 {
		return errors.New("glTexCoord4sv")
	}
	gpTexCoord4xOES = uintptr(getProcAddr("glTexCoord4xOES"))
	gpTexCoord4xvOES = uintptr(getProcAddr("glTexCoord4xvOES"))
	gpTexCoordFormatNV = uintptr(getProcAddr("glTexCoordFormatNV"))
	gpTexCoordPointer = uintptr(getProcAddr("glTexCoordPointer"))
	if gpTexCoordPointer == 0 {
		return errors.New("glTexCoordPointer")
	}
	gpTexCoordPointerEXT = uintptr(getProcAddr("glTexCoordPointerEXT"))
	gpTexCoordPointerListIBM = uintptr(getProcAddr("glTexCoordPointerListIBM"))
	gpTexCoordPointervINTEL = uintptr(getProcAddr("glTexCoordPointervINTEL"))
	gpTexEnvf = uintptr(getProcAddr("glTexEnvf"))
	if gpTexEnvf == 0 {
		return errors.New("glTexEnvf")
	}
	gpTexEnvfv = uintptr(getProcAddr("glTexEnvfv"))
	if gpTexEnvfv == 0 {
		return errors.New("glTexEnvfv")
	}
	gpTexEnvi = uintptr(getProcAddr("glTexEnvi"))
	if gpTexEnvi == 0 {
		return errors.New("glTexEnvi")
	}
	gpTexEnviv = uintptr(getProcAddr("glTexEnviv"))
	if gpTexEnviv == 0 {
		return errors.New("glTexEnviv")
	}
	gpTexEnvxOES = uintptr(getProcAddr("glTexEnvxOES"))
	gpTexEnvxvOES = uintptr(getProcAddr("glTexEnvxvOES"))
	gpTexFilterFuncSGIS = uintptr(getProcAddr("glTexFilterFuncSGIS"))
	gpTexGend = uintptr(getProcAddr("glTexGend"))
	if gpTexGend == 0 {
		return errors.New("glTexGend")
	}
	gpTexGendv = uintptr(getProcAddr("glTexGendv"))
	if gpTexGendv == 0 {
		return errors.New("glTexGendv")
	}
	gpTexGenf = uintptr(getProcAddr("glTexGenf"))
	if gpTexGenf == 0 {
		return errors.New("glTexGenf")
	}
	gpTexGenfv = uintptr(getProcAddr("glTexGenfv"))
	if gpTexGenfv == 0 {
		return errors.New("glTexGenfv")
	}
	gpTexGeni = uintptr(getProcAddr("glTexGeni"))
	if gpTexGeni == 0 {
		return errors.New("glTexGeni")
	}
	gpTexGeniv = uintptr(getProcAddr("glTexGeniv"))
	if gpTexGeniv == 0 {
		return errors.New("glTexGeniv")
	}
	gpTexGenxOES = uintptr(getProcAddr("glTexGenxOES"))
	gpTexGenxvOES = uintptr(getProcAddr("glTexGenxvOES"))
	gpTexImage1D = uintptr(getProcAddr("glTexImage1D"))
	if gpTexImage1D == 0 {
		return errors.New("glTexImage1D")
	}
	gpTexImage2D = uintptr(getProcAddr("glTexImage2D"))
	if gpTexImage2D == 0 {
		return errors.New("glTexImage2D")
	}
	gpTexImage2DMultisample = uintptr(getProcAddr("glTexImage2DMultisample"))
	gpTexImage2DMultisampleCoverageNV = uintptr(getProcAddr("glTexImage2DMultisampleCoverageNV"))
	gpTexImage3D = uintptr(getProcAddr("glTexImage3D"))
	if gpTexImage3D == 0 {
		return errors.New("glTexImage3D")
	}
	gpTexImage3DEXT = uintptr(getProcAddr("glTexImage3DEXT"))
	gpTexImage3DMultisample = uintptr(getProcAddr("glTexImage3DMultisample"))
	gpTexImage3DMultisampleCoverageNV = uintptr(getProcAddr("glTexImage3DMultisampleCoverageNV"))
	gpTexImage4DSGIS = uintptr(getProcAddr("glTexImage4DSGIS"))
	gpTexPageCommitmentARB = uintptr(getProcAddr("glTexPageCommitmentARB"))
	gpTexParameterIivEXT = uintptr(getProcAddr("glTexParameterIivEXT"))
	gpTexParameterIuivEXT = uintptr(getProcAddr("glTexParameterIuivEXT"))
	gpTexParameterf = uintptr(getProcAddr("glTexParameterf"))
	if gpTexParameterf == 0 {
		return errors.New("glTexParameterf")
	}
	gpTexParameterfv = uintptr(getProcAddr("glTexParameterfv"))
	if gpTexParameterfv == 0 {
		return errors.New("glTexParameterfv")
	}
	gpTexParameteri = uintptr(getProcAddr("glTexParameteri"))
	if gpTexParameteri == 0 {
		return errors.New("glTexParameteri")
	}
	gpTexParameteriv = uintptr(getProcAddr("glTexParameteriv"))
	if gpTexParameteriv == 0 {
		return errors.New("glTexParameteriv")
	}
	gpTexParameterxOES = uintptr(getProcAddr("glTexParameterxOES"))
	gpTexParameterxvOES = uintptr(getProcAddr("glTexParameterxvOES"))
	gpTexRenderbufferNV = uintptr(getProcAddr("glTexRenderbufferNV"))
	gpTexStorage1D = uintptr(getProcAddr("glTexStorage1D"))
	gpTexStorage2D = uintptr(getProcAddr("glTexStorage2D"))
	gpTexStorage2DMultisample = uintptr(getProcAddr("glTexStorage2DMultisample"))
	gpTexStorage3D = uintptr(getProcAddr("glTexStorage3D"))
	gpTexStorage3DMultisample = uintptr(getProcAddr("glTexStorage3DMultisample"))
	gpTexStorageMem1DEXT = uintptr(getProcAddr("glTexStorageMem1DEXT"))
	gpTexStorageMem2DEXT = uintptr(getProcAddr("glTexStorageMem2DEXT"))
	gpTexStorageMem2DMultisampleEXT = uintptr(getProcAddr("glTexStorageMem2DMultisampleEXT"))
	gpTexStorageMem3DEXT = uintptr(getProcAddr("glTexStorageMem3DEXT"))
	gpTexStorageMem3DMultisampleEXT = uintptr(getProcAddr("glTexStorageMem3DMultisampleEXT"))
	gpTexStorageSparseAMD = uintptr(getProcAddr("glTexStorageSparseAMD"))
	gpTexSubImage1D = uintptr(getProcAddr("glTexSubImage1D"))
	if gpTexSubImage1D == 0 {
		return errors.New("glTexSubImage1D")
	}
	gpTexSubImage1DEXT = uintptr(getProcAddr("glTexSubImage1DEXT"))
	gpTexSubImage2D = uintptr(getProcAddr("glTexSubImage2D"))
	if gpTexSubImage2D == 0 {
		return errors.New("glTexSubImage2D")
	}
	gpTexSubImage2DEXT = uintptr(getProcAddr("glTexSubImage2DEXT"))
	gpTexSubImage3D = uintptr(getProcAddr("glTexSubImage3D"))
	if gpTexSubImage3D == 0 {
		return errors.New("glTexSubImage3D")
	}
	gpTexSubImage3DEXT = uintptr(getProcAddr("glTexSubImage3DEXT"))
	gpTexSubImage4DSGIS = uintptr(getProcAddr("glTexSubImage4DSGIS"))
	gpTextureBarrier = uintptr(getProcAddr("glTextureBarrier"))
	gpTextureBarrierNV = uintptr(getProcAddr("glTextureBarrierNV"))
	gpTextureBuffer = uintptr(getProcAddr("glTextureBuffer"))
	gpTextureBufferEXT = uintptr(getProcAddr("glTextureBufferEXT"))
	gpTextureBufferRange = uintptr(getProcAddr("glTextureBufferRange"))
	gpTextureBufferRangeEXT = uintptr(getProcAddr("glTextureBufferRangeEXT"))
	gpTextureColorMaskSGIS = uintptr(getProcAddr("glTextureColorMaskSGIS"))
	gpTextureImage1DEXT = uintptr(getProcAddr("glTextureImage1DEXT"))
	gpTextureImage2DEXT = uintptr(getProcAddr("glTextureImage2DEXT"))
	gpTextureImage2DMultisampleCoverageNV = uintptr(getProcAddr("glTextureImage2DMultisampleCoverageNV"))
	gpTextureImage2DMultisampleNV = uintptr(getProcAddr("glTextureImage2DMultisampleNV"))
	gpTextureImage3DEXT = uintptr(getProcAddr("glTextureImage3DEXT"))
	gpTextureImage3DMultisampleCoverageNV = uintptr(getProcAddr("glTextureImage3DMultisampleCoverageNV"))
	gpTextureImage3DMultisampleNV = uintptr(getProcAddr("glTextureImage3DMultisampleNV"))
	gpTextureLightEXT = uintptr(getProcAddr("glTextureLightEXT"))
	gpTextureMaterialEXT = uintptr(getProcAddr("glTextureMaterialEXT"))
	gpTextureNormalEXT = uintptr(getProcAddr("glTextureNormalEXT"))
	gpTexturePageCommitmentEXT = uintptr(getProcAddr("glTexturePageCommitmentEXT"))
	gpTextureParameterIiv = uintptr(getProcAddr("glTextureParameterIiv"))
	gpTextureParameterIivEXT = uintptr(getProcAddr("glTextureParameterIivEXT"))
	gpTextureParameterIuiv = uintptr(getProcAddr("glTextureParameterIuiv"))
	gpTextureParameterIuivEXT = uintptr(getProcAddr("glTextureParameterIuivEXT"))
	gpTextureParameterf = uintptr(getProcAddr("glTextureParameterf"))
	gpTextureParameterfEXT = uintptr(getProcAddr("glTextureParameterfEXT"))
	gpTextureParameterfv = uintptr(getProcAddr("glTextureParameterfv"))
	gpTextureParameterfvEXT = uintptr(getProcAddr("glTextureParameterfvEXT"))
	gpTextureParameteri = uintptr(getProcAddr("glTextureParameteri"))
	gpTextureParameteriEXT = uintptr(getProcAddr("glTextureParameteriEXT"))
	gpTextureParameteriv = uintptr(getProcAddr("glTextureParameteriv"))
	gpTextureParameterivEXT = uintptr(getProcAddr("glTextureParameterivEXT"))
	gpTextureRangeAPPLE = uintptr(getProcAddr("glTextureRangeAPPLE"))
	gpTextureRenderbufferEXT = uintptr(getProcAddr("glTextureRenderbufferEXT"))
	gpTextureStorage1D = uintptr(getProcAddr("glTextureStorage1D"))
	gpTextureStorage1DEXT = uintptr(getProcAddr("glTextureStorage1DEXT"))
	gpTextureStorage2D = uintptr(getProcAddr("glTextureStorage2D"))
	gpTextureStorage2DEXT = uintptr(getProcAddr("glTextureStorage2DEXT"))
	gpTextureStorage2DMultisample = uintptr(getProcAddr("glTextureStorage2DMultisample"))
	gpTextureStorage2DMultisampleEXT = uintptr(getProcAddr("glTextureStorage2DMultisampleEXT"))
	gpTextureStorage3D = uintptr(getProcAddr("glTextureStorage3D"))
	gpTextureStorage3DEXT = uintptr(getProcAddr("glTextureStorage3DEXT"))
	gpTextureStorage3DMultisample = uintptr(getProcAddr("glTextureStorage3DMultisample"))
	gpTextureStorage3DMultisampleEXT = uintptr(getProcAddr("glTextureStorage3DMultisampleEXT"))
	gpTextureStorageMem1DEXT = uintptr(getProcAddr("glTextureStorageMem1DEXT"))
	gpTextureStorageMem2DEXT = uintptr(getProcAddr("glTextureStorageMem2DEXT"))
	gpTextureStorageMem2DMultisampleEXT = uintptr(getProcAddr("glTextureStorageMem2DMultisampleEXT"))
	gpTextureStorageMem3DEXT = uintptr(getProcAddr("glTextureStorageMem3DEXT"))
	gpTextureStorageMem3DMultisampleEXT = uintptr(getProcAddr("glTextureStorageMem3DMultisampleEXT"))
	gpTextureStorageSparseAMD = uintptr(getProcAddr("glTextureStorageSparseAMD"))
	gpTextureSubImage1D = uintptr(getProcAddr("glTextureSubImage1D"))
	gpTextureSubImage1DEXT = uintptr(getProcAddr("glTextureSubImage1DEXT"))
	gpTextureSubImage2D = uintptr(getProcAddr("glTextureSubImage2D"))
	gpTextureSubImage2DEXT = uintptr(getProcAddr("glTextureSubImage2DEXT"))
	gpTextureSubImage3D = uintptr(getProcAddr("glTextureSubImage3D"))
	gpTextureSubImage3DEXT = uintptr(getProcAddr("glTextureSubImage3DEXT"))
	gpTextureView = uintptr(getProcAddr("glTextureView"))
	gpTrackMatrixNV = uintptr(getProcAddr("glTrackMatrixNV"))
	gpTransformFeedbackAttribsNV = uintptr(getProcAddr("glTransformFeedbackAttribsNV"))
	gpTransformFeedbackBufferBase = uintptr(getProcAddr("glTransformFeedbackBufferBase"))
	gpTransformFeedbackBufferRange = uintptr(getProcAddr("glTransformFeedbackBufferRange"))
	gpTransformFeedbackStreamAttribsNV = uintptr(getProcAddr("glTransformFeedbackStreamAttribsNV"))
	gpTransformFeedbackVaryingsEXT = uintptr(getProcAddr("glTransformFeedbackVaryingsEXT"))
	gpTransformFeedbackVaryingsNV = uintptr(getProcAddr("glTransformFeedbackVaryingsNV"))
	gpTransformPathNV = uintptr(getProcAddr("glTransformPathNV"))
	gpTranslated = uintptr(getProcAddr("glTranslated"))
	if gpTranslated == 0 {
		return errors.New("glTranslated")
	}
	gpTranslatef = uintptr(getProcAddr("glTranslatef"))
	if gpTranslatef == 0 {
		return errors.New("glTranslatef")
	}
	gpTranslatexOES = uintptr(getProcAddr("glTranslatexOES"))
	gpUniform1d = uintptr(getProcAddr("glUniform1d"))
	gpUniform1dv = uintptr(getProcAddr("glUniform1dv"))
	gpUniform1f = uintptr(getProcAddr("glUniform1f"))
	if gpUniform1f == 0 {
		return errors.New("glUniform1f")
	}
	gpUniform1fARB = uintptr(getProcAddr("glUniform1fARB"))
	gpUniform1fv = uintptr(getProcAddr("glUniform1fv"))
	if gpUniform1fv == 0 {
		return errors.New("glUniform1fv")
	}
	gpUniform1fvARB = uintptr(getProcAddr("glUniform1fvARB"))
	gpUniform1i = uintptr(getProcAddr("glUniform1i"))
	if gpUniform1i == 0 {
		return errors.New("glUniform1i")
	}
	gpUniform1i64ARB = uintptr(getProcAddr("glUniform1i64ARB"))
	gpUniform1i64NV = uintptr(getProcAddr("glUniform1i64NV"))
	gpUniform1i64vARB = uintptr(getProcAddr("glUniform1i64vARB"))
	gpUniform1i64vNV = uintptr(getProcAddr("glUniform1i64vNV"))
	gpUniform1iARB = uintptr(getProcAddr("glUniform1iARB"))
	gpUniform1iv = uintptr(getProcAddr("glUniform1iv"))
	if gpUniform1iv == 0 {
		return errors.New("glUniform1iv")
	}
	gpUniform1ivARB = uintptr(getProcAddr("glUniform1ivARB"))
	gpUniform1ui64ARB = uintptr(getProcAddr("glUniform1ui64ARB"))
	gpUniform1ui64NV = uintptr(getProcAddr("glUniform1ui64NV"))
	gpUniform1ui64vARB = uintptr(getProcAddr("glUniform1ui64vARB"))
	gpUniform1ui64vNV = uintptr(getProcAddr("glUniform1ui64vNV"))
	gpUniform1uiEXT = uintptr(getProcAddr("glUniform1uiEXT"))
	gpUniform1uivEXT = uintptr(getProcAddr("glUniform1uivEXT"))
	gpUniform2d = uintptr(getProcAddr("glUniform2d"))
	gpUniform2dv = uintptr(getProcAddr("glUniform2dv"))
	gpUniform2f = uintptr(getProcAddr("glUniform2f"))
	if gpUniform2f == 0 {
		return errors.New("glUniform2f")
	}
	gpUniform2fARB = uintptr(getProcAddr("glUniform2fARB"))
	gpUniform2fv = uintptr(getProcAddr("glUniform2fv"))
	if gpUniform2fv == 0 {
		return errors.New("glUniform2fv")
	}
	gpUniform2fvARB = uintptr(getProcAddr("glUniform2fvARB"))
	gpUniform2i = uintptr(getProcAddr("glUniform2i"))
	if gpUniform2i == 0 {
		return errors.New("glUniform2i")
	}
	gpUniform2i64ARB = uintptr(getProcAddr("glUniform2i64ARB"))
	gpUniform2i64NV = uintptr(getProcAddr("glUniform2i64NV"))
	gpUniform2i64vARB = uintptr(getProcAddr("glUniform2i64vARB"))
	gpUniform2i64vNV = uintptr(getProcAddr("glUniform2i64vNV"))
	gpUniform2iARB = uintptr(getProcAddr("glUniform2iARB"))
	gpUniform2iv = uintptr(getProcAddr("glUniform2iv"))
	if gpUniform2iv == 0 {
		return errors.New("glUniform2iv")
	}
	gpUniform2ivARB = uintptr(getProcAddr("glUniform2ivARB"))
	gpUniform2ui64ARB = uintptr(getProcAddr("glUniform2ui64ARB"))
	gpUniform2ui64NV = uintptr(getProcAddr("glUniform2ui64NV"))
	gpUniform2ui64vARB = uintptr(getProcAddr("glUniform2ui64vARB"))
	gpUniform2ui64vNV = uintptr(getProcAddr("glUniform2ui64vNV"))
	gpUniform2uiEXT = uintptr(getProcAddr("glUniform2uiEXT"))
	gpUniform2uivEXT = uintptr(getProcAddr("glUniform2uivEXT"))
	gpUniform3d = uintptr(getProcAddr("glUniform3d"))
	gpUniform3dv = uintptr(getProcAddr("glUniform3dv"))
	gpUniform3f = uintptr(getProcAddr("glUniform3f"))
	if gpUniform3f == 0 {
		return errors.New("glUniform3f")
	}
	gpUniform3fARB = uintptr(getProcAddr("glUniform3fARB"))
	gpUniform3fv = uintptr(getProcAddr("glUniform3fv"))
	if gpUniform3fv == 0 {
		return errors.New("glUniform3fv")
	}
	gpUniform3fvARB = uintptr(getProcAddr("glUniform3fvARB"))
	gpUniform3i = uintptr(getProcAddr("glUniform3i"))
	if gpUniform3i == 0 {
		return errors.New("glUniform3i")
	}
	gpUniform3i64ARB = uintptr(getProcAddr("glUniform3i64ARB"))
	gpUniform3i64NV = uintptr(getProcAddr("glUniform3i64NV"))
	gpUniform3i64vARB = uintptr(getProcAddr("glUniform3i64vARB"))
	gpUniform3i64vNV = uintptr(getProcAddr("glUniform3i64vNV"))
	gpUniform3iARB = uintptr(getProcAddr("glUniform3iARB"))
	gpUniform3iv = uintptr(getProcAddr("glUniform3iv"))
	if gpUniform3iv == 0 {
		return errors.New("glUniform3iv")
	}
	gpUniform3ivARB = uintptr(getProcAddr("glUniform3ivARB"))
	gpUniform3ui64ARB = uintptr(getProcAddr("glUniform3ui64ARB"))
	gpUniform3ui64NV = uintptr(getProcAddr("glUniform3ui64NV"))
	gpUniform3ui64vARB = uintptr(getProcAddr("glUniform3ui64vARB"))
	gpUniform3ui64vNV = uintptr(getProcAddr("glUniform3ui64vNV"))
	gpUniform3uiEXT = uintptr(getProcAddr("glUniform3uiEXT"))
	gpUniform3uivEXT = uintptr(getProcAddr("glUniform3uivEXT"))
	gpUniform4d = uintptr(getProcAddr("glUniform4d"))
	gpUniform4dv = uintptr(getProcAddr("glUniform4dv"))
	gpUniform4f = uintptr(getProcAddr("glUniform4f"))
	if gpUniform4f == 0 {
		return errors.New("glUniform4f")
	}
	gpUniform4fARB = uintptr(getProcAddr("glUniform4fARB"))
	gpUniform4fv = uintptr(getProcAddr("glUniform4fv"))
	if gpUniform4fv == 0 {
		return errors.New("glUniform4fv")
	}
	gpUniform4fvARB = uintptr(getProcAddr("glUniform4fvARB"))
	gpUniform4i = uintptr(getProcAddr("glUniform4i"))
	if gpUniform4i == 0 {
		return errors.New("glUniform4i")
	}
	gpUniform4i64ARB = uintptr(getProcAddr("glUniform4i64ARB"))
	gpUniform4i64NV = uintptr(getProcAddr("glUniform4i64NV"))
	gpUniform4i64vARB = uintptr(getProcAddr("glUniform4i64vARB"))
	gpUniform4i64vNV = uintptr(getProcAddr("glUniform4i64vNV"))
	gpUniform4iARB = uintptr(getProcAddr("glUniform4iARB"))
	gpUniform4iv = uintptr(getProcAddr("glUniform4iv"))
	if gpUniform4iv == 0 {
		return errors.New("glUniform4iv")
	}
	gpUniform4ivARB = uintptr(getProcAddr("glUniform4ivARB"))
	gpUniform4ui64ARB = uintptr(getProcAddr("glUniform4ui64ARB"))
	gpUniform4ui64NV = uintptr(getProcAddr("glUniform4ui64NV"))
	gpUniform4ui64vARB = uintptr(getProcAddr("glUniform4ui64vARB"))
	gpUniform4ui64vNV = uintptr(getProcAddr("glUniform4ui64vNV"))
	gpUniform4uiEXT = uintptr(getProcAddr("glUniform4uiEXT"))
	gpUniform4uivEXT = uintptr(getProcAddr("glUniform4uivEXT"))
	gpUniformBlockBinding = uintptr(getProcAddr("glUniformBlockBinding"))
	gpUniformBufferEXT = uintptr(getProcAddr("glUniformBufferEXT"))
	gpUniformHandleui64ARB = uintptr(getProcAddr("glUniformHandleui64ARB"))
	gpUniformHandleui64NV = uintptr(getProcAddr("glUniformHandleui64NV"))
	gpUniformHandleui64vARB = uintptr(getProcAddr("glUniformHandleui64vARB"))
	gpUniformHandleui64vNV = uintptr(getProcAddr("glUniformHandleui64vNV"))
	gpUniformMatrix2dv = uintptr(getProcAddr("glUniformMatrix2dv"))
	gpUniformMatrix2fv = uintptr(getProcAddr("glUniformMatrix2fv"))
	if gpUniformMatrix2fv == 0 {
		return errors.New("glUniformMatrix2fv")
	}
	gpUniformMatrix2fvARB = uintptr(getProcAddr("glUniformMatrix2fvARB"))
	gpUniformMatrix2x3dv = uintptr(getProcAddr("glUniformMatrix2x3dv"))
	gpUniformMatrix2x3fv = uintptr(getProcAddr("glUniformMatrix2x3fv"))
	if gpUniformMatrix2x3fv == 0 {
		return errors.New("glUniformMatrix2x3fv")
	}
	gpUniformMatrix2x4dv = uintptr(getProcAddr("glUniformMatrix2x4dv"))
	gpUniformMatrix2x4fv = uintptr(getProcAddr("glUniformMatrix2x4fv"))
	if gpUniformMatrix2x4fv == 0 {
		return errors.New("glUniformMatrix2x4fv")
	}
	gpUniformMatrix3dv = uintptr(getProcAddr("glUniformMatrix3dv"))
	gpUniformMatrix3fv = uintptr(getProcAddr("glUniformMatrix3fv"))
	if gpUniformMatrix3fv == 0 {
		return errors.New("glUniformMatrix3fv")
	}
	gpUniformMatrix3fvARB = uintptr(getProcAddr("glUniformMatrix3fvARB"))
	gpUniformMatrix3x2dv = uintptr(getProcAddr("glUniformMatrix3x2dv"))
	gpUniformMatrix3x2fv = uintptr(getProcAddr("glUniformMatrix3x2fv"))
	if gpUniformMatrix3x2fv == 0 {
		return errors.New("glUniformMatrix3x2fv")
	}
	gpUniformMatrix3x4dv = uintptr(getProcAddr("glUniformMatrix3x4dv"))
	gpUniformMatrix3x4fv = uintptr(getProcAddr("glUniformMatrix3x4fv"))
	if gpUniformMatrix3x4fv == 0 {
		return errors.New("glUniformMatrix3x4fv")
	}
	gpUniformMatrix4dv = uintptr(getProcAddr("glUniformMatrix4dv"))
	gpUniformMatrix4fv = uintptr(getProcAddr("glUniformMatrix4fv"))
	if gpUniformMatrix4fv == 0 {
		return errors.New("glUniformMatrix4fv")
	}
	gpUniformMatrix4fvARB = uintptr(getProcAddr("glUniformMatrix4fvARB"))
	gpUniformMatrix4x2dv = uintptr(getProcAddr("glUniformMatrix4x2dv"))
	gpUniformMatrix4x2fv = uintptr(getProcAddr("glUniformMatrix4x2fv"))
	if gpUniformMatrix4x2fv == 0 {
		return errors.New("glUniformMatrix4x2fv")
	}
	gpUniformMatrix4x3dv = uintptr(getProcAddr("glUniformMatrix4x3dv"))
	gpUniformMatrix4x3fv = uintptr(getProcAddr("glUniformMatrix4x3fv"))
	if gpUniformMatrix4x3fv == 0 {
		return errors.New("glUniformMatrix4x3fv")
	}
	gpUniformSubroutinesuiv = uintptr(getProcAddr("glUniformSubroutinesuiv"))
	gpUniformui64NV = uintptr(getProcAddr("glUniformui64NV"))
	gpUniformui64vNV = uintptr(getProcAddr("glUniformui64vNV"))
	gpUnlockArraysEXT = uintptr(getProcAddr("glUnlockArraysEXT"))
	gpUnmapBuffer = uintptr(getProcAddr("glUnmapBuffer"))
	if gpUnmapBuffer == 0 {
		return errors.New("glUnmapBuffer")
	}
	gpUnmapBufferARB = uintptr(getProcAddr("glUnmapBufferARB"))
	gpUnmapNamedBuffer = uintptr(getProcAddr("glUnmapNamedBuffer"))
	gpUnmapNamedBufferEXT = uintptr(getProcAddr("glUnmapNamedBufferEXT"))
	gpUnmapObjectBufferATI = uintptr(getProcAddr("glUnmapObjectBufferATI"))
	gpUnmapTexture2DINTEL = uintptr(getProcAddr("glUnmapTexture2DINTEL"))
	gpUpdateObjectBufferATI = uintptr(getProcAddr("glUpdateObjectBufferATI"))
	gpUseProgram = uintptr(getProcAddr("glUseProgram"))
	if gpUseProgram == 0 {
		return errors.New("glUseProgram")
	}
	gpUseProgramObjectARB = uintptr(getProcAddr("glUseProgramObjectARB"))
	gpUseProgramStages = uintptr(getProcAddr("glUseProgramStages"))
	gpUseProgramStagesEXT = uintptr(getProcAddr("glUseProgramStagesEXT"))
	gpUseShaderProgramEXT = uintptr(getProcAddr("glUseShaderProgramEXT"))
	gpVDPAUFiniNV = uintptr(getProcAddr("glVDPAUFiniNV"))
	gpVDPAUGetSurfaceivNV = uintptr(getProcAddr("glVDPAUGetSurfaceivNV"))
	gpVDPAUInitNV = uintptr(getProcAddr("glVDPAUInitNV"))
	gpVDPAUIsSurfaceNV = uintptr(getProcAddr("glVDPAUIsSurfaceNV"))
	gpVDPAUMapSurfacesNV = uintptr(getProcAddr("glVDPAUMapSurfacesNV"))
	gpVDPAURegisterOutputSurfaceNV = uintptr(getProcAddr("glVDPAURegisterOutputSurfaceNV"))
	gpVDPAURegisterVideoSurfaceNV = uintptr(getProcAddr("glVDPAURegisterVideoSurfaceNV"))
	gpVDPAUSurfaceAccessNV = uintptr(getProcAddr("glVDPAUSurfaceAccessNV"))
	gpVDPAUUnmapSurfacesNV = uintptr(getProcAddr("glVDPAUUnmapSurfacesNV"))
	gpVDPAUUnregisterSurfaceNV = uintptr(getProcAddr("glVDPAUUnregisterSurfaceNV"))
	gpValidateProgram = uintptr(getProcAddr("glValidateProgram"))
	if gpValidateProgram == 0 {
		return errors.New("glValidateProgram")
	}
	gpValidateProgramARB = uintptr(getProcAddr("glValidateProgramARB"))
	gpValidateProgramPipeline = uintptr(getProcAddr("glValidateProgramPipeline"))
	gpValidateProgramPipelineEXT = uintptr(getProcAddr("glValidateProgramPipelineEXT"))
	gpVariantArrayObjectATI = uintptr(getProcAddr("glVariantArrayObjectATI"))
	gpVariantPointerEXT = uintptr(getProcAddr("glVariantPointerEXT"))
	gpVariantbvEXT = uintptr(getProcAddr("glVariantbvEXT"))
	gpVariantdvEXT = uintptr(getProcAddr("glVariantdvEXT"))
	gpVariantfvEXT = uintptr(getProcAddr("glVariantfvEXT"))
	gpVariantivEXT = uintptr(getProcAddr("glVariantivEXT"))
	gpVariantsvEXT = uintptr(getProcAddr("glVariantsvEXT"))
	gpVariantubvEXT = uintptr(getProcAddr("glVariantubvEXT"))
	gpVariantuivEXT = uintptr(getProcAddr("glVariantuivEXT"))
	gpVariantusvEXT = uintptr(getProcAddr("glVariantusvEXT"))
	gpVertex2bOES = uintptr(getProcAddr("glVertex2bOES"))
	gpVertex2bvOES = uintptr(getProcAddr("glVertex2bvOES"))
	gpVertex2d = uintptr(getProcAddr("glVertex2d"))
	if gpVertex2d == 0 {
		return errors.New("glVertex2d")
	}
	gpVertex2dv = uintptr(getProcAddr("glVertex2dv"))
	if gpVertex2dv == 0 {
		return errors.New("glVertex2dv")
	}
	gpVertex2f = uintptr(getProcAddr("glVertex2f"))
	if gpVertex2f == 0 {
		return errors.New("glVertex2f")
	}
	gpVertex2fv = uintptr(getProcAddr("glVertex2fv"))
	if gpVertex2fv == 0 {
		return errors.New("glVertex2fv")
	}
	gpVertex2hNV = uintptr(getProcAddr("glVertex2hNV"))
	gpVertex2hvNV = uintptr(getProcAddr("glVertex2hvNV"))
	gpVertex2i = uintptr(getProcAddr("glVertex2i"))
	if gpVertex2i == 0 {
		return errors.New("glVertex2i")
	}
	gpVertex2iv = uintptr(getProcAddr("glVertex2iv"))
	if gpVertex2iv == 0 {
		return errors.New("glVertex2iv")
	}
	gpVertex2s = uintptr(getProcAddr("glVertex2s"))
	if gpVertex2s == 0 {
		return errors.New("glVertex2s")
	}
	gpVertex2sv = uintptr(getProcAddr("glVertex2sv"))
	if gpVertex2sv == 0 {
		return errors.New("glVertex2sv")
	}
	gpVertex2xOES = uintptr(getProcAddr("glVertex2xOES"))
	gpVertex2xvOES = uintptr(getProcAddr("glVertex2xvOES"))
	gpVertex3bOES = uintptr(getProcAddr("glVertex3bOES"))
	gpVertex3bvOES = uintptr(getProcAddr("glVertex3bvOES"))
	gpVertex3d = uintptr(getProcAddr("glVertex3d"))
	if gpVertex3d == 0 {
		return errors.New("glVertex3d")
	}
	gpVertex3dv = uintptr(getProcAddr("glVertex3dv"))
	if gpVertex3dv == 0 {
		return errors.New("glVertex3dv")
	}
	gpVertex3f = uintptr(getProcAddr("glVertex3f"))
	if gpVertex3f == 0 {
		return errors.New("glVertex3f")
	}
	gpVertex3fv = uintptr(getProcAddr("glVertex3fv"))
	if gpVertex3fv == 0 {
		return errors.New("glVertex3fv")
	}
	gpVertex3hNV = uintptr(getProcAddr("glVertex3hNV"))
	gpVertex3hvNV = uintptr(getProcAddr("glVertex3hvNV"))
	gpVertex3i = uintptr(getProcAddr("glVertex3i"))
	if gpVertex3i == 0 {
		return errors.New("glVertex3i")
	}
	gpVertex3iv = uintptr(getProcAddr("glVertex3iv"))
	if gpVertex3iv == 0 {
		return errors.New("glVertex3iv")
	}
	gpVertex3s = uintptr(getProcAddr("glVertex3s"))
	if gpVertex3s == 0 {
		return errors.New("glVertex3s")
	}
	gpVertex3sv = uintptr(getProcAddr("glVertex3sv"))
	if gpVertex3sv == 0 {
		return errors.New("glVertex3sv")
	}
	gpVertex3xOES = uintptr(getProcAddr("glVertex3xOES"))
	gpVertex3xvOES = uintptr(getProcAddr("glVertex3xvOES"))
	gpVertex4bOES = uintptr(getProcAddr("glVertex4bOES"))
	gpVertex4bvOES = uintptr(getProcAddr("glVertex4bvOES"))
	gpVertex4d = uintptr(getProcAddr("glVertex4d"))
	if gpVertex4d == 0 {
		return errors.New("glVertex4d")
	}
	gpVertex4dv = uintptr(getProcAddr("glVertex4dv"))
	if gpVertex4dv == 0 {
		return errors.New("glVertex4dv")
	}
	gpVertex4f = uintptr(getProcAddr("glVertex4f"))
	if gpVertex4f == 0 {
		return errors.New("glVertex4f")
	}
	gpVertex4fv = uintptr(getProcAddr("glVertex4fv"))
	if gpVertex4fv == 0 {
		return errors.New("glVertex4fv")
	}
	gpVertex4hNV = uintptr(getProcAddr("glVertex4hNV"))
	gpVertex4hvNV = uintptr(getProcAddr("glVertex4hvNV"))
	gpVertex4i = uintptr(getProcAddr("glVertex4i"))
	if gpVertex4i == 0 {
		return errors.New("glVertex4i")
	}
	gpVertex4iv = uintptr(getProcAddr("glVertex4iv"))
	if gpVertex4iv == 0 {
		return errors.New("glVertex4iv")
	}
	gpVertex4s = uintptr(getProcAddr("glVertex4s"))
	if gpVertex4s == 0 {
		return errors.New("glVertex4s")
	}
	gpVertex4sv = uintptr(getProcAddr("glVertex4sv"))
	if gpVertex4sv == 0 {
		return errors.New("glVertex4sv")
	}
	gpVertex4xOES = uintptr(getProcAddr("glVertex4xOES"))
	gpVertex4xvOES = uintptr(getProcAddr("glVertex4xvOES"))
	gpVertexArrayAttribBinding = uintptr(getProcAddr("glVertexArrayAttribBinding"))
	gpVertexArrayAttribFormat = uintptr(getProcAddr("glVertexArrayAttribFormat"))
	gpVertexArrayAttribIFormat = uintptr(getProcAddr("glVertexArrayAttribIFormat"))
	gpVertexArrayAttribLFormat = uintptr(getProcAddr("glVertexArrayAttribLFormat"))
	gpVertexArrayBindVertexBufferEXT = uintptr(getProcAddr("glVertexArrayBindVertexBufferEXT"))
	gpVertexArrayBindingDivisor = uintptr(getProcAddr("glVertexArrayBindingDivisor"))
	gpVertexArrayColorOffsetEXT = uintptr(getProcAddr("glVertexArrayColorOffsetEXT"))
	gpVertexArrayEdgeFlagOffsetEXT = uintptr(getProcAddr("glVertexArrayEdgeFlagOffsetEXT"))
	gpVertexArrayElementBuffer = uintptr(getProcAddr("glVertexArrayElementBuffer"))
	gpVertexArrayFogCoordOffsetEXT = uintptr(getProcAddr("glVertexArrayFogCoordOffsetEXT"))
	gpVertexArrayIndexOffsetEXT = uintptr(getProcAddr("glVertexArrayIndexOffsetEXT"))
	gpVertexArrayMultiTexCoordOffsetEXT = uintptr(getProcAddr("glVertexArrayMultiTexCoordOffsetEXT"))
	gpVertexArrayNormalOffsetEXT = uintptr(getProcAddr("glVertexArrayNormalOffsetEXT"))
	gpVertexArrayParameteriAPPLE = uintptr(getProcAddr("glVertexArrayParameteriAPPLE"))
	gpVertexArrayRangeAPPLE = uintptr(getProcAddr("glVertexArrayRangeAPPLE"))
	gpVertexArrayRangeNV = uintptr(getProcAddr("glVertexArrayRangeNV"))
	gpVertexArraySecondaryColorOffsetEXT = uintptr(getProcAddr("glVertexArraySecondaryColorOffsetEXT"))
	gpVertexArrayTexCoordOffsetEXT = uintptr(getProcAddr("glVertexArrayTexCoordOffsetEXT"))
	gpVertexArrayVertexAttribBindingEXT = uintptr(getProcAddr("glVertexArrayVertexAttribBindingEXT"))
	gpVertexArrayVertexAttribDivisorEXT = uintptr(getProcAddr("glVertexArrayVertexAttribDivisorEXT"))
	gpVertexArrayVertexAttribFormatEXT = uintptr(getProcAddr("glVertexArrayVertexAttribFormatEXT"))
	gpVertexArrayVertexAttribIFormatEXT = uintptr(getProcAddr("glVertexArrayVertexAttribIFormatEXT"))
	gpVertexArrayVertexAttribIOffsetEXT = uintptr(getProcAddr("glVertexArrayVertexAttribIOffsetEXT"))
	gpVertexArrayVertexAttribLFormatEXT = uintptr(getProcAddr("glVertexArrayVertexAttribLFormatEXT"))
	gpVertexArrayVertexAttribLOffsetEXT = uintptr(getProcAddr("glVertexArrayVertexAttribLOffsetEXT"))
	gpVertexArrayVertexAttribOffsetEXT = uintptr(getProcAddr("glVertexArrayVertexAttribOffsetEXT"))
	gpVertexArrayVertexBindingDivisorEXT = uintptr(getProcAddr("glVertexArrayVertexBindingDivisorEXT"))
	gpVertexArrayVertexBuffer = uintptr(getProcAddr("glVertexArrayVertexBuffer"))
	gpVertexArrayVertexBuffers = uintptr(getProcAddr("glVertexArrayVertexBuffers"))
	gpVertexArrayVertexOffsetEXT = uintptr(getProcAddr("glVertexArrayVertexOffsetEXT"))
	gpVertexAttrib1d = uintptr(getProcAddr("glVertexAttrib1d"))
	if gpVertexAttrib1d == 0 {
		return errors.New("glVertexAttrib1d")
	}
	gpVertexAttrib1dARB = uintptr(getProcAddr("glVertexAttrib1dARB"))
	gpVertexAttrib1dNV = uintptr(getProcAddr("glVertexAttrib1dNV"))
	gpVertexAttrib1dv = uintptr(getProcAddr("glVertexAttrib1dv"))
	if gpVertexAttrib1dv == 0 {
		return errors.New("glVertexAttrib1dv")
	}
	gpVertexAttrib1dvARB = uintptr(getProcAddr("glVertexAttrib1dvARB"))
	gpVertexAttrib1dvNV = uintptr(getProcAddr("glVertexAttrib1dvNV"))
	gpVertexAttrib1f = uintptr(getProcAddr("glVertexAttrib1f"))
	if gpVertexAttrib1f == 0 {
		return errors.New("glVertexAttrib1f")
	}
	gpVertexAttrib1fARB = uintptr(getProcAddr("glVertexAttrib1fARB"))
	gpVertexAttrib1fNV = uintptr(getProcAddr("glVertexAttrib1fNV"))
	gpVertexAttrib1fv = uintptr(getProcAddr("glVertexAttrib1fv"))
	if gpVertexAttrib1fv == 0 {
		return errors.New("glVertexAttrib1fv")
	}
	gpVertexAttrib1fvARB = uintptr(getProcAddr("glVertexAttrib1fvARB"))
	gpVertexAttrib1fvNV = uintptr(getProcAddr("glVertexAttrib1fvNV"))
	gpVertexAttrib1hNV = uintptr(getProcAddr("glVertexAttrib1hNV"))
	gpVertexAttrib1hvNV = uintptr(getProcAddr("glVertexAttrib1hvNV"))
	gpVertexAttrib1s = uintptr(getProcAddr("glVertexAttrib1s"))
	if gpVertexAttrib1s == 0 {
		return errors.New("glVertexAttrib1s")
	}
	gpVertexAttrib1sARB = uintptr(getProcAddr("glVertexAttrib1sARB"))
	gpVertexAttrib1sNV = uintptr(getProcAddr("glVertexAttrib1sNV"))
	gpVertexAttrib1sv = uintptr(getProcAddr("glVertexAttrib1sv"))
	if gpVertexAttrib1sv == 0 {
		return errors.New("glVertexAttrib1sv")
	}
	gpVertexAttrib1svARB = uintptr(getProcAddr("glVertexAttrib1svARB"))
	gpVertexAttrib1svNV = uintptr(getProcAddr("glVertexAttrib1svNV"))
	gpVertexAttrib2d = uintptr(getProcAddr("glVertexAttrib2d"))
	if gpVertexAttrib2d == 0 {
		return errors.New("glVertexAttrib2d")
	}
	gpVertexAttrib2dARB = uintptr(getProcAddr("glVertexAttrib2dARB"))
	gpVertexAttrib2dNV = uintptr(getProcAddr("glVertexAttrib2dNV"))
	gpVertexAttrib2dv = uintptr(getProcAddr("glVertexAttrib2dv"))
	if gpVertexAttrib2dv == 0 {
		return errors.New("glVertexAttrib2dv")
	}
	gpVertexAttrib2dvARB = uintptr(getProcAddr("glVertexAttrib2dvARB"))
	gpVertexAttrib2dvNV = uintptr(getProcAddr("glVertexAttrib2dvNV"))
	gpVertexAttrib2f = uintptr(getProcAddr("glVertexAttrib2f"))
	if gpVertexAttrib2f == 0 {
		return errors.New("glVertexAttrib2f")
	}
	gpVertexAttrib2fARB = uintptr(getProcAddr("glVertexAttrib2fARB"))
	gpVertexAttrib2fNV = uintptr(getProcAddr("glVertexAttrib2fNV"))
	gpVertexAttrib2fv = uintptr(getProcAddr("glVertexAttrib2fv"))
	if gpVertexAttrib2fv == 0 {
		return errors.New("glVertexAttrib2fv")
	}
	gpVertexAttrib2fvARB = uintptr(getProcAddr("glVertexAttrib2fvARB"))
	gpVertexAttrib2fvNV = uintptr(getProcAddr("glVertexAttrib2fvNV"))
	gpVertexAttrib2hNV = uintptr(getProcAddr("glVertexAttrib2hNV"))
	gpVertexAttrib2hvNV = uintptr(getProcAddr("glVertexAttrib2hvNV"))
	gpVertexAttrib2s = uintptr(getProcAddr("glVertexAttrib2s"))
	if gpVertexAttrib2s == 0 {
		return errors.New("glVertexAttrib2s")
	}
	gpVertexAttrib2sARB = uintptr(getProcAddr("glVertexAttrib2sARB"))
	gpVertexAttrib2sNV = uintptr(getProcAddr("glVertexAttrib2sNV"))
	gpVertexAttrib2sv = uintptr(getProcAddr("glVertexAttrib2sv"))
	if gpVertexAttrib2sv == 0 {
		return errors.New("glVertexAttrib2sv")
	}
	gpVertexAttrib2svARB = uintptr(getProcAddr("glVertexAttrib2svARB"))
	gpVertexAttrib2svNV = uintptr(getProcAddr("glVertexAttrib2svNV"))
	gpVertexAttrib3d = uintptr(getProcAddr("glVertexAttrib3d"))
	if gpVertexAttrib3d == 0 {
		return errors.New("glVertexAttrib3d")
	}
	gpVertexAttrib3dARB = uintptr(getProcAddr("glVertexAttrib3dARB"))
	gpVertexAttrib3dNV = uintptr(getProcAddr("glVertexAttrib3dNV"))
	gpVertexAttrib3dv = uintptr(getProcAddr("glVertexAttrib3dv"))
	if gpVertexAttrib3dv == 0 {
		return errors.New("glVertexAttrib3dv")
	}
	gpVertexAttrib3dvARB = uintptr(getProcAddr("glVertexAttrib3dvARB"))
	gpVertexAttrib3dvNV = uintptr(getProcAddr("glVertexAttrib3dvNV"))
	gpVertexAttrib3f = uintptr(getProcAddr("glVertexAttrib3f"))
	if gpVertexAttrib3f == 0 {
		return errors.New("glVertexAttrib3f")
	}
	gpVertexAttrib3fARB = uintptr(getProcAddr("glVertexAttrib3fARB"))
	gpVertexAttrib3fNV = uintptr(getProcAddr("glVertexAttrib3fNV"))
	gpVertexAttrib3fv = uintptr(getProcAddr("glVertexAttrib3fv"))
	if gpVertexAttrib3fv == 0 {
		return errors.New("glVertexAttrib3fv")
	}
	gpVertexAttrib3fvARB = uintptr(getProcAddr("glVertexAttrib3fvARB"))
	gpVertexAttrib3fvNV = uintptr(getProcAddr("glVertexAttrib3fvNV"))
	gpVertexAttrib3hNV = uintptr(getProcAddr("glVertexAttrib3hNV"))
	gpVertexAttrib3hvNV = uintptr(getProcAddr("glVertexAttrib3hvNV"))
	gpVertexAttrib3s = uintptr(getProcAddr("glVertexAttrib3s"))
	if gpVertexAttrib3s == 0 {
		return errors.New("glVertexAttrib3s")
	}
	gpVertexAttrib3sARB = uintptr(getProcAddr("glVertexAttrib3sARB"))
	gpVertexAttrib3sNV = uintptr(getProcAddr("glVertexAttrib3sNV"))
	gpVertexAttrib3sv = uintptr(getProcAddr("glVertexAttrib3sv"))
	if gpVertexAttrib3sv == 0 {
		return errors.New("glVertexAttrib3sv")
	}
	gpVertexAttrib3svARB = uintptr(getProcAddr("glVertexAttrib3svARB"))
	gpVertexAttrib3svNV = uintptr(getProcAddr("glVertexAttrib3svNV"))
	gpVertexAttrib4Nbv = uintptr(getProcAddr("glVertexAttrib4Nbv"))
	if gpVertexAttrib4Nbv == 0 {
		return errors.New("glVertexAttrib4Nbv")
	}
	gpVertexAttrib4NbvARB = uintptr(getProcAddr("glVertexAttrib4NbvARB"))
	gpVertexAttrib4Niv = uintptr(getProcAddr("glVertexAttrib4Niv"))
	if gpVertexAttrib4Niv == 0 {
		return errors.New("glVertexAttrib4Niv")
	}
	gpVertexAttrib4NivARB = uintptr(getProcAddr("glVertexAttrib4NivARB"))
	gpVertexAttrib4Nsv = uintptr(getProcAddr("glVertexAttrib4Nsv"))
	if gpVertexAttrib4Nsv == 0 {
		return errors.New("glVertexAttrib4Nsv")
	}
	gpVertexAttrib4NsvARB = uintptr(getProcAddr("glVertexAttrib4NsvARB"))
	gpVertexAttrib4Nub = uintptr(getProcAddr("glVertexAttrib4Nub"))
	if gpVertexAttrib4Nub == 0 {
		return errors.New("glVertexAttrib4Nub")
	}
	gpVertexAttrib4NubARB = uintptr(getProcAddr("glVertexAttrib4NubARB"))
	gpVertexAttrib4Nubv = uintptr(getProcAddr("glVertexAttrib4Nubv"))
	if gpVertexAttrib4Nubv == 0 {
		return errors.New("glVertexAttrib4Nubv")
	}
	gpVertexAttrib4NubvARB = uintptr(getProcAddr("glVertexAttrib4NubvARB"))
	gpVertexAttrib4Nuiv = uintptr(getProcAddr("glVertexAttrib4Nuiv"))
	if gpVertexAttrib4Nuiv == 0 {
		return errors.New("glVertexAttrib4Nuiv")
	}
	gpVertexAttrib4NuivARB = uintptr(getProcAddr("glVertexAttrib4NuivARB"))
	gpVertexAttrib4Nusv = uintptr(getProcAddr("glVertexAttrib4Nusv"))
	if gpVertexAttrib4Nusv == 0 {
		return errors.New("glVertexAttrib4Nusv")
	}
	gpVertexAttrib4NusvARB = uintptr(getProcAddr("glVertexAttrib4NusvARB"))
	gpVertexAttrib4bv = uintptr(getProcAddr("glVertexAttrib4bv"))
	if gpVertexAttrib4bv == 0 {
		return errors.New("glVertexAttrib4bv")
	}
	gpVertexAttrib4bvARB = uintptr(getProcAddr("glVertexAttrib4bvARB"))
	gpVertexAttrib4d = uintptr(getProcAddr("glVertexAttrib4d"))
	if gpVertexAttrib4d == 0 {
		return errors.New("glVertexAttrib4d")
	}
	gpVertexAttrib4dARB = uintptr(getProcAddr("glVertexAttrib4dARB"))
	gpVertexAttrib4dNV = uintptr(getProcAddr("glVertexAttrib4dNV"))
	gpVertexAttrib4dv = uintptr(getProcAddr("glVertexAttrib4dv"))
	if gpVertexAttrib4dv == 0 {
		return errors.New("glVertexAttrib4dv")
	}
	gpVertexAttrib4dvARB = uintptr(getProcAddr("glVertexAttrib4dvARB"))
	gpVertexAttrib4dvNV = uintptr(getProcAddr("glVertexAttrib4dvNV"))
	gpVertexAttrib4f = uintptr(getProcAddr("glVertexAttrib4f"))
	if gpVertexAttrib4f == 0 {
		return errors.New("glVertexAttrib4f")
	}
	gpVertexAttrib4fARB = uintptr(getProcAddr("glVertexAttrib4fARB"))
	gpVertexAttrib4fNV = uintptr(getProcAddr("glVertexAttrib4fNV"))
	gpVertexAttrib4fv = uintptr(getProcAddr("glVertexAttrib4fv"))
	if gpVertexAttrib4fv == 0 {
		return errors.New("glVertexAttrib4fv")
	}
	gpVertexAttrib4fvARB = uintptr(getProcAddr("glVertexAttrib4fvARB"))
	gpVertexAttrib4fvNV = uintptr(getProcAddr("glVertexAttrib4fvNV"))
	gpVertexAttrib4hNV = uintptr(getProcAddr("glVertexAttrib4hNV"))
	gpVertexAttrib4hvNV = uintptr(getProcAddr("glVertexAttrib4hvNV"))
	gpVertexAttrib4iv = uintptr(getProcAddr("glVertexAttrib4iv"))
	if gpVertexAttrib4iv == 0 {
		return errors.New("glVertexAttrib4iv")
	}
	gpVertexAttrib4ivARB = uintptr(getProcAddr("glVertexAttrib4ivARB"))
	gpVertexAttrib4s = uintptr(getProcAddr("glVertexAttrib4s"))
	if gpVertexAttrib4s == 0 {
		return errors.New("glVertexAttrib4s")
	}
	gpVertexAttrib4sARB = uintptr(getProcAddr("glVertexAttrib4sARB"))
	gpVertexAttrib4sNV = uintptr(getProcAddr("glVertexAttrib4sNV"))
	gpVertexAttrib4sv = uintptr(getProcAddr("glVertexAttrib4sv"))
	if gpVertexAttrib4sv == 0 {
		return errors.New("glVertexAttrib4sv")
	}
	gpVertexAttrib4svARB = uintptr(getProcAddr("glVertexAttrib4svARB"))
	gpVertexAttrib4svNV = uintptr(getProcAddr("glVertexAttrib4svNV"))
	gpVertexAttrib4ubNV = uintptr(getProcAddr("glVertexAttrib4ubNV"))
	gpVertexAttrib4ubv = uintptr(getProcAddr("glVertexAttrib4ubv"))
	if gpVertexAttrib4ubv == 0 {
		return errors.New("glVertexAttrib4ubv")
	}
	gpVertexAttrib4ubvARB = uintptr(getProcAddr("glVertexAttrib4ubvARB"))
	gpVertexAttrib4ubvNV = uintptr(getProcAddr("glVertexAttrib4ubvNV"))
	gpVertexAttrib4uiv = uintptr(getProcAddr("glVertexAttrib4uiv"))
	if gpVertexAttrib4uiv == 0 {
		return errors.New("glVertexAttrib4uiv")
	}
	gpVertexAttrib4uivARB = uintptr(getProcAddr("glVertexAttrib4uivARB"))
	gpVertexAttrib4usv = uintptr(getProcAddr("glVertexAttrib4usv"))
	if gpVertexAttrib4usv == 0 {
		return errors.New("glVertexAttrib4usv")
	}
	gpVertexAttrib4usvARB = uintptr(getProcAddr("glVertexAttrib4usvARB"))
	gpVertexAttribArrayObjectATI = uintptr(getProcAddr("glVertexAttribArrayObjectATI"))
	gpVertexAttribBinding = uintptr(getProcAddr("glVertexAttribBinding"))
	gpVertexAttribDivisorARB = uintptr(getProcAddr("glVertexAttribDivisorARB"))
	gpVertexAttribFormat = uintptr(getProcAddr("glVertexAttribFormat"))
	gpVertexAttribFormatNV = uintptr(getProcAddr("glVertexAttribFormatNV"))
	gpVertexAttribI1iEXT = uintptr(getProcAddr("glVertexAttribI1iEXT"))
	gpVertexAttribI1ivEXT = uintptr(getProcAddr("glVertexAttribI1ivEXT"))
	gpVertexAttribI1uiEXT = uintptr(getProcAddr("glVertexAttribI1uiEXT"))
	gpVertexAttribI1uivEXT = uintptr(getProcAddr("glVertexAttribI1uivEXT"))
	gpVertexAttribI2iEXT = uintptr(getProcAddr("glVertexAttribI2iEXT"))
	gpVertexAttribI2ivEXT = uintptr(getProcAddr("glVertexAttribI2ivEXT"))
	gpVertexAttribI2uiEXT = uintptr(getProcAddr("glVertexAttribI2uiEXT"))
	gpVertexAttribI2uivEXT = uintptr(getProcAddr("glVertexAttribI2uivEXT"))
	gpVertexAttribI3iEXT = uintptr(getProcAddr("glVertexAttribI3iEXT"))
	gpVertexAttribI3ivEXT = uintptr(getProcAddr("glVertexAttribI3ivEXT"))
	gpVertexAttribI3uiEXT = uintptr(getProcAddr("glVertexAttribI3uiEXT"))
	gpVertexAttribI3uivEXT = uintptr(getProcAddr("glVertexAttribI3uivEXT"))
	gpVertexAttribI4bvEXT = uintptr(getProcAddr("glVertexAttribI4bvEXT"))
	gpVertexAttribI4iEXT = uintptr(getProcAddr("glVertexAttribI4iEXT"))
	gpVertexAttribI4ivEXT = uintptr(getProcAddr("glVertexAttribI4ivEXT"))
	gpVertexAttribI4svEXT = uintptr(getProcAddr("glVertexAttribI4svEXT"))
	gpVertexAttribI4ubvEXT = uintptr(getProcAddr("glVertexAttribI4ubvEXT"))
	gpVertexAttribI4uiEXT = uintptr(getProcAddr("glVertexAttribI4uiEXT"))
	gpVertexAttribI4uivEXT = uintptr(getProcAddr("glVertexAttribI4uivEXT"))
	gpVertexAttribI4usvEXT = uintptr(getProcAddr("glVertexAttribI4usvEXT"))
	gpVertexAttribIFormat = uintptr(getProcAddr("glVertexAttribIFormat"))
	gpVertexAttribIFormatNV = uintptr(getProcAddr("glVertexAttribIFormatNV"))
	gpVertexAttribIPointerEXT = uintptr(getProcAddr("glVertexAttribIPointerEXT"))
	gpVertexAttribL1d = uintptr(getProcAddr("glVertexAttribL1d"))
	gpVertexAttribL1dEXT = uintptr(getProcAddr("glVertexAttribL1dEXT"))
	gpVertexAttribL1dv = uintptr(getProcAddr("glVertexAttribL1dv"))
	gpVertexAttribL1dvEXT = uintptr(getProcAddr("glVertexAttribL1dvEXT"))
	gpVertexAttribL1i64NV = uintptr(getProcAddr("glVertexAttribL1i64NV"))
	gpVertexAttribL1i64vNV = uintptr(getProcAddr("glVertexAttribL1i64vNV"))
	gpVertexAttribL1ui64ARB = uintptr(getProcAddr("glVertexAttribL1ui64ARB"))
	gpVertexAttribL1ui64NV = uintptr(getProcAddr("glVertexAttribL1ui64NV"))
	gpVertexAttribL1ui64vARB = uintptr(getProcAddr("glVertexAttribL1ui64vARB"))
	gpVertexAttribL1ui64vNV = uintptr(getProcAddr("glVertexAttribL1ui64vNV"))
	gpVertexAttribL2d = uintptr(getProcAddr("glVertexAttribL2d"))
	gpVertexAttribL2dEXT = uintptr(getProcAddr("glVertexAttribL2dEXT"))
	gpVertexAttribL2dv = uintptr(getProcAddr("glVertexAttribL2dv"))
	gpVertexAttribL2dvEXT = uintptr(getProcAddr("glVertexAttribL2dvEXT"))
	gpVertexAttribL2i64NV = uintptr(getProcAddr("glVertexAttribL2i64NV"))
	gpVertexAttribL2i64vNV = uintptr(getProcAddr("glVertexAttribL2i64vNV"))
	gpVertexAttribL2ui64NV = uintptr(getProcAddr("glVertexAttribL2ui64NV"))
	gpVertexAttribL2ui64vNV = uintptr(getProcAddr("glVertexAttribL2ui64vNV"))
	gpVertexAttribL3d = uintptr(getProcAddr("glVertexAttribL3d"))
	gpVertexAttribL3dEXT = uintptr(getProcAddr("glVertexAttribL3dEXT"))
	gpVertexAttribL3dv = uintptr(getProcAddr("glVertexAttribL3dv"))
	gpVertexAttribL3dvEXT = uintptr(getProcAddr("glVertexAttribL3dvEXT"))
	gpVertexAttribL3i64NV = uintptr(getProcAddr("glVertexAttribL3i64NV"))
	gpVertexAttribL3i64vNV = uintptr(getProcAddr("glVertexAttribL3i64vNV"))
	gpVertexAttribL3ui64NV = uintptr(getProcAddr("glVertexAttribL3ui64NV"))
	gpVertexAttribL3ui64vNV = uintptr(getProcAddr("glVertexAttribL3ui64vNV"))
	gpVertexAttribL4d = uintptr(getProcAddr("glVertexAttribL4d"))
	gpVertexAttribL4dEXT = uintptr(getProcAddr("glVertexAttribL4dEXT"))
	gpVertexAttribL4dv = uintptr(getProcAddr("glVertexAttribL4dv"))
	gpVertexAttribL4dvEXT = uintptr(getProcAddr("glVertexAttribL4dvEXT"))
	gpVertexAttribL4i64NV = uintptr(getProcAddr("glVertexAttribL4i64NV"))
	gpVertexAttribL4i64vNV = uintptr(getProcAddr("glVertexAttribL4i64vNV"))
	gpVertexAttribL4ui64NV = uintptr(getProcAddr("glVertexAttribL4ui64NV"))
	gpVertexAttribL4ui64vNV = uintptr(getProcAddr("glVertexAttribL4ui64vNV"))
	gpVertexAttribLFormat = uintptr(getProcAddr("glVertexAttribLFormat"))
	gpVertexAttribLFormatNV = uintptr(getProcAddr("glVertexAttribLFormatNV"))
	gpVertexAttribLPointer = uintptr(getProcAddr("glVertexAttribLPointer"))
	gpVertexAttribLPointerEXT = uintptr(getProcAddr("glVertexAttribLPointerEXT"))
	gpVertexAttribP1ui = uintptr(getProcAddr("glVertexAttribP1ui"))
	gpVertexAttribP1uiv = uintptr(getProcAddr("glVertexAttribP1uiv"))
	gpVertexAttribP2ui = uintptr(getProcAddr("glVertexAttribP2ui"))
	gpVertexAttribP2uiv = uintptr(getProcAddr("glVertexAttribP2uiv"))
	gpVertexAttribP3ui = uintptr(getProcAddr("glVertexAttribP3ui"))
	gpVertexAttribP3uiv = uintptr(getProcAddr("glVertexAttribP3uiv"))
	gpVertexAttribP4ui = uintptr(getProcAddr("glVertexAttribP4ui"))
	gpVertexAttribP4uiv = uintptr(getProcAddr("glVertexAttribP4uiv"))
	gpVertexAttribParameteriAMD = uintptr(getProcAddr("glVertexAttribParameteriAMD"))
	gpVertexAttribPointer = uintptr(getProcAddr("glVertexAttribPointer"))
	if gpVertexAttribPointer == 0 {
		return errors.New("glVertexAttribPointer")
	}
	gpVertexAttribPointerARB = uintptr(getProcAddr("glVertexAttribPointerARB"))
	gpVertexAttribPointerNV = uintptr(getProcAddr("glVertexAttribPointerNV"))
	gpVertexAttribs1dvNV = uintptr(getProcAddr("glVertexAttribs1dvNV"))
	gpVertexAttribs1fvNV = uintptr(getProcAddr("glVertexAttribs1fvNV"))
	gpVertexAttribs1hvNV = uintptr(getProcAddr("glVertexAttribs1hvNV"))
	gpVertexAttribs1svNV = uintptr(getProcAddr("glVertexAttribs1svNV"))
	gpVertexAttribs2dvNV = uintptr(getProcAddr("glVertexAttribs2dvNV"))
	gpVertexAttribs2fvNV = uintptr(getProcAddr("glVertexAttribs2fvNV"))
	gpVertexAttribs2hvNV = uintptr(getProcAddr("glVertexAttribs2hvNV"))
	gpVertexAttribs2svNV = uintptr(getProcAddr("glVertexAttribs2svNV"))
	gpVertexAttribs3dvNV = uintptr(getProcAddr("glVertexAttribs3dvNV"))
	gpVertexAttribs3fvNV = uintptr(getProcAddr("glVertexAttribs3fvNV"))
	gpVertexAttribs3hvNV = uintptr(getProcAddr("glVertexAttribs3hvNV"))
	gpVertexAttribs3svNV = uintptr(getProcAddr("glVertexAttribs3svNV"))
	gpVertexAttribs4dvNV = uintptr(getProcAddr("glVertexAttribs4dvNV"))
	gpVertexAttribs4fvNV = uintptr(getProcAddr("glVertexAttribs4fvNV"))
	gpVertexAttribs4hvNV = uintptr(getProcAddr("glVertexAttribs4hvNV"))
	gpVertexAttribs4svNV = uintptr(getProcAddr("glVertexAttribs4svNV"))
	gpVertexAttribs4ubvNV = uintptr(getProcAddr("glVertexAttribs4ubvNV"))
	gpVertexBindingDivisor = uintptr(getProcAddr("glVertexBindingDivisor"))
	gpVertexBlendARB = uintptr(getProcAddr("glVertexBlendARB"))
	gpVertexBlendEnvfATI = uintptr(getProcAddr("glVertexBlendEnvfATI"))
	gpVertexBlendEnviATI = uintptr(getProcAddr("glVertexBlendEnviATI"))
	gpVertexFormatNV = uintptr(getProcAddr("glVertexFormatNV"))
	gpVertexPointer = uintptr(getProcAddr("glVertexPointer"))
	if gpVertexPointer == 0 {
		return errors.New("glVertexPointer")
	}
	gpVertexPointerEXT = uintptr(getProcAddr("glVertexPointerEXT"))
	gpVertexPointerListIBM = uintptr(getProcAddr("glVertexPointerListIBM"))
	gpVertexPointervINTEL = uintptr(getProcAddr("glVertexPointervINTEL"))
	gpVertexStream1dATI = uintptr(getProcAddr("glVertexStream1dATI"))
	gpVertexStream1dvATI = uintptr(getProcAddr("glVertexStream1dvATI"))
	gpVertexStream1fATI = uintptr(getProcAddr("glVertexStream1fATI"))
	gpVertexStream1fvATI = uintptr(getProcAddr("glVertexStream1fvATI"))
	gpVertexStream1iATI = uintptr(getProcAddr("glVertexStream1iATI"))
	gpVertexStream1ivATI = uintptr(getProcAddr("glVertexStream1ivATI"))
	gpVertexStream1sATI = uintptr(getProcAddr("glVertexStream1sATI"))
	gpVertexStream1svATI = uintptr(getProcAddr("glVertexStream1svATI"))
	gpVertexStream2dATI = uintptr(getProcAddr("glVertexStream2dATI"))
	gpVertexStream2dvATI = uintptr(getProcAddr("glVertexStream2dvATI"))
	gpVertexStream2fATI = uintptr(getProcAddr("glVertexStream2fATI"))
	gpVertexStream2fvATI = uintptr(getProcAddr("glVertexStream2fvATI"))
	gpVertexStream2iATI = uintptr(getProcAddr("glVertexStream2iATI"))
	gpVertexStream2ivATI = uintptr(getProcAddr("glVertexStream2ivATI"))
	gpVertexStream2sATI = uintptr(getProcAddr("glVertexStream2sATI"))
	gpVertexStream2svATI = uintptr(getProcAddr("glVertexStream2svATI"))
	gpVertexStream3dATI = uintptr(getProcAddr("glVertexStream3dATI"))
	gpVertexStream3dvATI = uintptr(getProcAddr("glVertexStream3dvATI"))
	gpVertexStream3fATI = uintptr(getProcAddr("glVertexStream3fATI"))
	gpVertexStream3fvATI = uintptr(getProcAddr("glVertexStream3fvATI"))
	gpVertexStream3iATI = uintptr(getProcAddr("glVertexStream3iATI"))
	gpVertexStream3ivATI = uintptr(getProcAddr("glVertexStream3ivATI"))
	gpVertexStream3sATI = uintptr(getProcAddr("glVertexStream3sATI"))
	gpVertexStream3svATI = uintptr(getProcAddr("glVertexStream3svATI"))
	gpVertexStream4dATI = uintptr(getProcAddr("glVertexStream4dATI"))
	gpVertexStream4dvATI = uintptr(getProcAddr("glVertexStream4dvATI"))
	gpVertexStream4fATI = uintptr(getProcAddr("glVertexStream4fATI"))
	gpVertexStream4fvATI = uintptr(getProcAddr("glVertexStream4fvATI"))
	gpVertexStream4iATI = uintptr(getProcAddr("glVertexStream4iATI"))
	gpVertexStream4ivATI = uintptr(getProcAddr("glVertexStream4ivATI"))
	gpVertexStream4sATI = uintptr(getProcAddr("glVertexStream4sATI"))
	gpVertexStream4svATI = uintptr(getProcAddr("glVertexStream4svATI"))
	gpVertexWeightPointerEXT = uintptr(getProcAddr("glVertexWeightPointerEXT"))
	gpVertexWeightfEXT = uintptr(getProcAddr("glVertexWeightfEXT"))
	gpVertexWeightfvEXT = uintptr(getProcAddr("glVertexWeightfvEXT"))
	gpVertexWeighthNV = uintptr(getProcAddr("glVertexWeighthNV"))
	gpVertexWeighthvNV = uintptr(getProcAddr("glVertexWeighthvNV"))
	gpVideoCaptureNV = uintptr(getProcAddr("glVideoCaptureNV"))
	gpVideoCaptureStreamParameterdvNV = uintptr(getProcAddr("glVideoCaptureStreamParameterdvNV"))
	gpVideoCaptureStreamParameterfvNV = uintptr(getProcAddr("glVideoCaptureStreamParameterfvNV"))
	gpVideoCaptureStreamParameterivNV = uintptr(getProcAddr("glVideoCaptureStreamParameterivNV"))
	gpViewport = uintptr(getProcAddr("glViewport"))
	if gpViewport == 0 {
		return errors.New("glViewport")
	}
	gpViewportArrayv = uintptr(getProcAddr("glViewportArrayv"))
	gpViewportIndexedf = uintptr(getProcAddr("glViewportIndexedf"))
	gpViewportIndexedfv = uintptr(getProcAddr("glViewportIndexedfv"))
	gpViewportPositionWScaleNV = uintptr(getProcAddr("glViewportPositionWScaleNV"))
	gpViewportSwizzleNV = uintptr(getProcAddr("glViewportSwizzleNV"))
	gpWaitSemaphoreEXT = uintptr(getProcAddr("glWaitSemaphoreEXT"))
	gpWaitSync = uintptr(getProcAddr("glWaitSync"))
	gpWaitVkSemaphoreNV = uintptr(getProcAddr("glWaitVkSemaphoreNV"))
	gpWeightPathsNV = uintptr(getProcAddr("glWeightPathsNV"))
	gpWeightPointerARB = uintptr(getProcAddr("glWeightPointerARB"))
	gpWeightbvARB = uintptr(getProcAddr("glWeightbvARB"))
	gpWeightdvARB = uintptr(getProcAddr("glWeightdvARB"))
	gpWeightfvARB = uintptr(getProcAddr("glWeightfvARB"))
	gpWeightivARB = uintptr(getProcAddr("glWeightivARB"))
	gpWeightsvARB = uintptr(getProcAddr("glWeightsvARB"))
	gpWeightubvARB = uintptr(getProcAddr("glWeightubvARB"))
	gpWeightuivARB = uintptr(getProcAddr("glWeightuivARB"))
	gpWeightusvARB = uintptr(getProcAddr("glWeightusvARB"))
	gpWindowPos2d = uintptr(getProcAddr("glWindowPos2d"))
	if gpWindowPos2d == 0 {
		return errors.New("glWindowPos2d")
	}
	gpWindowPos2dARB = uintptr(getProcAddr("glWindowPos2dARB"))
	gpWindowPos2dMESA = uintptr(getProcAddr("glWindowPos2dMESA"))
	gpWindowPos2dv = uintptr(getProcAddr("glWindowPos2dv"))
	if gpWindowPos2dv == 0 {
		return errors.New("glWindowPos2dv")
	}
	gpWindowPos2dvARB = uintptr(getProcAddr("glWindowPos2dvARB"))
	gpWindowPos2dvMESA = uintptr(getProcAddr("glWindowPos2dvMESA"))
	gpWindowPos2f = uintptr(getProcAddr("glWindowPos2f"))
	if gpWindowPos2f == 0 {
		return errors.New("glWindowPos2f")
	}
	gpWindowPos2fARB = uintptr(getProcAddr("glWindowPos2fARB"))
	gpWindowPos2fMESA = uintptr(getProcAddr("glWindowPos2fMESA"))
	gpWindowPos2fv = uintptr(getProcAddr("glWindowPos2fv"))
	if gpWindowPos2fv == 0 {
		return errors.New("glWindowPos2fv")
	}
	gpWindowPos2fvARB = uintptr(getProcAddr("glWindowPos2fvARB"))
	gpWindowPos2fvMESA = uintptr(getProcAddr("glWindowPos2fvMESA"))
	gpWindowPos2i = uintptr(getProcAddr("glWindowPos2i"))
	if gpWindowPos2i == 0 {
		return errors.New("glWindowPos2i")
	}
	gpWindowPos2iARB = uintptr(getProcAddr("glWindowPos2iARB"))
	gpWindowPos2iMESA = uintptr(getProcAddr("glWindowPos2iMESA"))
	gpWindowPos2iv = uintptr(getProcAddr("glWindowPos2iv"))
	if gpWindowPos2iv == 0 {
		return errors.New("glWindowPos2iv")
	}
	gpWindowPos2ivARB = uintptr(getProcAddr("glWindowPos2ivARB"))
	gpWindowPos2ivMESA = uintptr(getProcAddr("glWindowPos2ivMESA"))
	gpWindowPos2s = uintptr(getProcAddr("glWindowPos2s"))
	if gpWindowPos2s == 0 {
		return errors.New("glWindowPos2s")
	}
	gpWindowPos2sARB = uintptr(getProcAddr("glWindowPos2sARB"))
	gpWindowPos2sMESA = uintptr(getProcAddr("glWindowPos2sMESA"))
	gpWindowPos2sv = uintptr(getProcAddr("glWindowPos2sv"))
	if gpWindowPos2sv == 0 {
		return errors.New("glWindowPos2sv")
	}
	gpWindowPos2svARB = uintptr(getProcAddr("glWindowPos2svARB"))
	gpWindowPos2svMESA = uintptr(getProcAddr("glWindowPos2svMESA"))
	gpWindowPos3d = uintptr(getProcAddr("glWindowPos3d"))
	if gpWindowPos3d == 0 {
		return errors.New("glWindowPos3d")
	}
	gpWindowPos3dARB = uintptr(getProcAddr("glWindowPos3dARB"))
	gpWindowPos3dMESA = uintptr(getProcAddr("glWindowPos3dMESA"))
	gpWindowPos3dv = uintptr(getProcAddr("glWindowPos3dv"))
	if gpWindowPos3dv == 0 {
		return errors.New("glWindowPos3dv")
	}
	gpWindowPos3dvARB = uintptr(getProcAddr("glWindowPos3dvARB"))
	gpWindowPos3dvMESA = uintptr(getProcAddr("glWindowPos3dvMESA"))
	gpWindowPos3f = uintptr(getProcAddr("glWindowPos3f"))
	if gpWindowPos3f == 0 {
		return errors.New("glWindowPos3f")
	}
	gpWindowPos3fARB = uintptr(getProcAddr("glWindowPos3fARB"))
	gpWindowPos3fMESA = uintptr(getProcAddr("glWindowPos3fMESA"))
	gpWindowPos3fv = uintptr(getProcAddr("glWindowPos3fv"))
	if gpWindowPos3fv == 0 {
		return errors.New("glWindowPos3fv")
	}
	gpWindowPos3fvARB = uintptr(getProcAddr("glWindowPos3fvARB"))
	gpWindowPos3fvMESA = uintptr(getProcAddr("glWindowPos3fvMESA"))
	gpWindowPos3i = uintptr(getProcAddr("glWindowPos3i"))
	if gpWindowPos3i == 0 {
		return errors.New("glWindowPos3i")
	}
	gpWindowPos3iARB = uintptr(getProcAddr("glWindowPos3iARB"))
	gpWindowPos3iMESA = uintptr(getProcAddr("glWindowPos3iMESA"))
	gpWindowPos3iv = uintptr(getProcAddr("glWindowPos3iv"))
	if gpWindowPos3iv == 0 {
		return errors.New("glWindowPos3iv")
	}
	gpWindowPos3ivARB = uintptr(getProcAddr("glWindowPos3ivARB"))
	gpWindowPos3ivMESA = uintptr(getProcAddr("glWindowPos3ivMESA"))
	gpWindowPos3s = uintptr(getProcAddr("glWindowPos3s"))
	if gpWindowPos3s == 0 {
		return errors.New("glWindowPos3s")
	}
	gpWindowPos3sARB = uintptr(getProcAddr("glWindowPos3sARB"))
	gpWindowPos3sMESA = uintptr(getProcAddr("glWindowPos3sMESA"))
	gpWindowPos3sv = uintptr(getProcAddr("glWindowPos3sv"))
	if gpWindowPos3sv == 0 {
		return errors.New("glWindowPos3sv")
	}
	gpWindowPos3svARB = uintptr(getProcAddr("glWindowPos3svARB"))
	gpWindowPos3svMESA = uintptr(getProcAddr("glWindowPos3svMESA"))
	gpWindowPos4dMESA = uintptr(getProcAddr("glWindowPos4dMESA"))
	gpWindowPos4dvMESA = uintptr(getProcAddr("glWindowPos4dvMESA"))
	gpWindowPos4fMESA = uintptr(getProcAddr("glWindowPos4fMESA"))
	gpWindowPos4fvMESA = uintptr(getProcAddr("glWindowPos4fvMESA"))
	gpWindowPos4iMESA = uintptr(getProcAddr("glWindowPos4iMESA"))
	gpWindowPos4ivMESA = uintptr(getProcAddr("glWindowPos4ivMESA"))
	gpWindowPos4sMESA = uintptr(getProcAddr("glWindowPos4sMESA"))
	gpWindowPos4svMESA = uintptr(getProcAddr("glWindowPos4svMESA"))
	gpWindowRectanglesEXT = uintptr(getProcAddr("glWindowRectanglesEXT"))
	gpWriteMaskEXT = uintptr(getProcAddr("glWriteMaskEXT"))
	return nil
}
