import Button from '../../components/Button';
import { FiTrash2 } from 'react-icons/fi';
import { animateModalIn, animateModalOut } from '../../vars/animation';

export interface DeleteModalData {
  id: number | string;
  name: string;
  label: string;
}

interface DeleteModalProps {
  isDeleteLoading: boolean;
  isDeleteModalAnimating: boolean;
  deleteModalData: DeleteModalData | null;
  handleDelete: () => void;
  closeDeleteModal: () => void;
  onAnimationEnd: () => void;
}

const DeleteModal: React.FC<DeleteModalProps> = ({
  isDeleteLoading,
  isDeleteModalAnimating,
  deleteModalData,
  handleDelete,
  closeDeleteModal,
  onAnimationEnd
}) => {
    return (
        <>
        {/* Backdrop */}
        <div className="fixed inset-0 bg-gray-500/75 transition-opacity duration-300 z-40" 
             aria-hidden="true"
             onClick={closeDeleteModal} />

        {/* Modal container - centered with flex */}
        <div className="fixed inset-0 z-50 overflow-y-auto">
          <div className="flex min-h-full items-center justify-center p-4 text-center sm:p-0">

            {/* Modal panel with animation */}
            <div className={`relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg 
                  ${isDeleteModalAnimating ? animateModalOut : animateModalIn}`}
                  onAnimationEnd={onAnimationEnd}
                 >
                <div className="bg-white px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
                  <div className="sm:flex sm:items-start">
                    <div className="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                      <FiTrash2 className="h-6 w-6 text-red-600" />
                    </div>
                    <div className="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                      <h3 className="text-base font-semibold leading-6 text-gray-900">
                        Hapus {deleteModalData?.label}
                      </h3>
                      <div className="mt-2">
                        <p className="text-sm text-gray-500">
                          Apakah Anda yakin ingin menghapus {deleteModalData?.label} <span className="font-semibold">{deleteModalData?.name}</span>? 
                          Tindakan ini tidak dapat dibatalkan.
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                  <Button 
                    type="button"
                    variant="danger"
                    size="md"
                    isLoading={isDeleteLoading}
                    icon={<FiTrash2 className="h-3.5 w-3.5" />}
                    onClick={handleDelete}
                    className="sm:ml-3 sm:w-auto"
                  >
                    {isDeleteLoading ? 'Menghapus...' : 'Hapus'}
                  </Button>
                  <Button 
                    type="button"
                    variant="secondary"
                    size="md"
                    onClick={closeDeleteModal}
                    disabled={isDeleteLoading}
                  >
                    Batal
                  </Button>
                </div>
              </div>
          </div>
        </div>
        </>
    )
}

export default DeleteModal;